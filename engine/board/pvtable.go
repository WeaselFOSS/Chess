package board

import (
	"fmt"
)

//PVEnteryStruct The struct for each entry in the PV table
type PVEnteryStruct struct {
	posKey uint64
	move   int
}

//enterytBytes calculated using  unsafe.Sizeof(pos.PVTable.Entry[0])
//
//if more values are added to the struct (PVEnteryStruct) it will need to be recalculated
const enterytBytes = 16

//PVTableStruct Struct to hole the PVTable
type PVTableStruct struct {
	Entries      []PVEnteryStruct
	EntriesCount uint64
}

//initPVTables Initilize our PV Table slice with a exact amount of memory based on hashSizeMB
func (table *PVTableStruct) Init(hashSizeMB uint64) {
	hashBytes := hashSizeMB * 1000000
	table.EntriesCount = hashBytes / enterytBytes
	table.Entries = make([]PVEnteryStruct, table.EntriesCount)
	table.EntriesCount -= 2 //Saftey net for indexing
}

//Clear Clear the PV table
func (table *PVTableStruct) Clear() {
	table.Entries = make([]PVEnteryStruct, table.EntriesCount)
}

//StorePVMove a move in the PV table
func (pos *PositionStruct) StorePVMove(move int) error {
	//Indexing based off of position hash
	index := pos.PosKey % pos.PVTable.EntriesCount

	if DEBUG && (index <= 0 || index >= pos.PVTable.EntriesCount-1) {
		return fmt.Errorf("PV Index out of range with value of %d", index)
	}

	pos.PVTable.Entries[index].move = move
	pos.PVTable.Entries[index].posKey = pos.PosKey

	return nil
}

//ProbePVTable probe the table for a move on the current position
func (pos *PositionStruct) ProbePVTable() (int, error) {
	index := pos.PosKey % pos.PVTable.EntriesCount

	if DEBUG && (index <= 0 || index >= pos.PVTable.EntriesCount-1) {
		return NoMove, fmt.Errorf("PV Index out of range with value of %d", index)
	}

	if pos.PVTable.Entries[index].posKey == pos.PosKey {
		return pos.PVTable.Entries[index].move, nil
	}

	return NoMove, nil
}

//GetPvLine return the PVLine if found for the curren position
func (pos *PositionStruct) GetPvLine(depth int) (int, error) {
	if DEBUG && depth >= MaxDepth {
		return NoMove, fmt.Errorf("Depth of %d greater than or equal to MaxDepth of %d", depth, MaxDepth)
	}

	move, err := pos.ProbePVTable()
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
		move, err = pos.ProbePVTable()
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
