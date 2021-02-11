package board

import (
	"fmt"
	"unsafe"
)

//PVEnteryStruct The struct for each entry in the PV table
type HashEnteryStruct struct {
	PosKey uint64
	Move   int
	Score  int
	Depth  int
	Flags  int
	Ply    int
}

//enterytBytes calculated using  unsafe.Sizeof(pos.PVTable.Entry[0])
//
//if more values are added to the struct (PVEnteryStruct) it will need to be recalculated

//PVTableStruct Struct to hole the PVTable
type HashTableStruct struct {
	Entries     []HashEnteryStruct
	EntrieCount uint64
	NewWrite    int
	OverWrite   int
	Hit         int
	Cut         int
}

//initPVTables Initilize our PV Table slice with a exact amount of memory based on hashSizeMB
func (table *HashTableStruct) Init(hashSizeMB uint64) {
	const enterytBytes = uint64(unsafe.Sizeof(HashEnteryStruct{}))
	hashBytes := hashSizeMB * 1000000
	table.EntrieCount = hashBytes / enterytBytes
	table.Entries = make([]HashEnteryStruct, table.EntrieCount)
	table.EntrieCount -= 2 //Saftey net for indexing
}

//Clear Clear the PV table
func (table *HashTableStruct) Clear() {
	table.Entries = make([]HashEnteryStruct, table.EntrieCount)
	table.NewWrite = 0
}

func (pos *PositionStruct) ProbeHashEntry(move *int, score *int, alpha, beta, depth int) (bool, error) {
	index := pos.PosKey % pos.HashTable.EntrieCount

	if DEBUG && (index <= 0 || index >= pos.HashTable.EntrieCount-1) {
		return false, fmt.Errorf("PV Index out of range with value of %d", index)
	}

	if pos.HashTable.Entries[index].PosKey == pos.PosKey && pos.HashTable.Entries[index].Ply == pos.Ply {

		*move = pos.HashTable.Entries[index].Move
		if pos.HashTable.Entries[index].Depth >= depth {
			pos.HashTable.Hit++

			*score = pos.HashTable.Entries[index].Score
			if *score > IsMate {
				*score -= pos.Ply
			} else if *score < -IsMate {
				*score += pos.Ply
			}

			switch pos.HashTable.Entries[index].Flags {
			case HFALPHA:
				if *score <= alpha {
					*score = alpha
					return true, nil
				}
				break
			case HFBETA:
				if *score >= beta {
					*score = beta
					return true, nil
				}
				break
			case HFEXACT:
				return true, nil
			}
		}
	}
	return false, nil
}

//StorePVMove a move in the PV table
func (pos *PositionStruct) StoreHashEntry(move, score, flags, depth int) error {
	//Indexing based off of position hash
	index := pos.PosKey % pos.HashTable.EntrieCount

	if DEBUG && (index <= 0 || index >= pos.HashTable.EntrieCount-1) {
		return fmt.Errorf("PV Index out of range with value of %d", index)
	}

	if pos.HashTable.Entries[index].PosKey == 0 {
		pos.HashTable.NewWrite++
	} else {
		pos.HashTable.OverWrite++
	}

	if score > IsMate {
		score += pos.Ply
	} else if score < -IsMate {
		score -= pos.Ply
	}

	pos.HashTable.Entries[index].Move = move
	pos.HashTable.Entries[index].PosKey = pos.PosKey
	pos.HashTable.Entries[index].Ply = pos.Ply
	pos.HashTable.Entries[index].Flags = flags
	pos.HashTable.Entries[index].Score = score
	pos.HashTable.Entries[index].Depth = depth

	return nil
}

//ProbePVTable probe the table for a move on the current position
func (pos *PositionStruct) ProbePVMove() (int, error) {
	index := pos.PosKey % pos.HashTable.EntrieCount

	if DEBUG && (index <= 0 || index >= pos.HashTable.EntrieCount-1) {
		return NoMove, fmt.Errorf("PV Index out of range with value of %d", index)
	}

	if pos.HashTable.Entries[index].PosKey == pos.PosKey {
		return pos.HashTable.Entries[index].Move, nil
	}

	return NoMove, nil
}

//GetPvLine return the PVLine if found for the curren position
func (pos *PositionStruct) GetPvLine(depth int) (int, error) {
	if DEBUG && depth >= MaxDepth {
		return NoMove, fmt.Errorf("Depth of %d greater than or equal to MaxDepth of %d", depth, MaxDepth)
	}

	move, err := pos.ProbePVMove()
	if err != nil {
		return NoMove, err
	}

	var count int
	for move != NoMove && count < depth {
		var moveExists bool
		moveExists, err = pos.MoveExists(move)
		if moveExists {
			_, err = pos.MakeMove(move)
			if err != nil {
				return NoMove, err
			}
			pos.PvArray[count] = move
			count++
		} else {
			break
		}
		move, err = pos.ProbePVMove()
		if err != nil {
			return NoMove, err
		}
	}

	for pos.Ply > 0 {
		err = pos.TakeMove()
		if err != nil {
			return NoMove, err
		}
	}

	return count, nil
}
