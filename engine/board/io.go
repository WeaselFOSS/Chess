package board

import (
	"fmt"
)

// SquareToString square to algebraic notation
func SquareToString(sq int) string {
	if sq == offBoard || sq == noSquare {
		return "None"
	}
	file := filesBoard[sq]
	rank := ranksBoard[sq]
	return fmt.Sprintf("%c%c", ('a' + file), ('1' + rank))
}

// MoveToString move int to algebraic notation
func MoveToString(move int) string {
	ff := filesBoard[GetFrom(move)]
	rf := ranksBoard[GetFrom(move)]
	ft := filesBoard[GetTo(move)]
	rt := ranksBoard[GetTo(move)]

	promoted := GetPromoted(move)
	if promoted != 0 {
		pchar := 'q'
		if promoted == wN || promoted == bN {
			pchar = 'n'
		} else if promoted == wR || promoted == bR {
			pchar = 'r'
		} else if promoted == wB || promoted == bB {
			pchar = 'b'
		}
		return fmt.Sprintf("%c%c%c%c%c", ('a' + ff), ('1' + rf), ('a' + ft), ('1' + rt), pchar)
	}
	return fmt.Sprintf("%c%c%c%c", ('a' + ff), ('1' + rf), ('a' + ft), ('1' + rt))
}

// ParseMove algebraic notation to move int
func (pos *PositionStruct) ParseMove(move string) (int, error) {
	if len(move) < 4 {
		return NoMove, nil
	}

	if rune(move[1]) > '8' || rune(move[1]) < '1' {
		return NoMove, fmt.Errorf("Move %s out of bounds", move)
	}

	if rune(move[3]) > '8' || rune(move[3]) < '1' {
		return NoMove, fmt.Errorf("Move %s out of bounds", move)
	}

	if rune(move[0]) > 'h' || rune(move[0]) < 'a' {
		return NoMove, fmt.Errorf("Move %s out of bounds", move)
	}

	if rune(move[2]) > 'h' || rune(move[2]) < 'a' {
		return NoMove, fmt.Errorf("Move %s out of bounds", move)
	}

	from := fileRankToSquare(int(rune(move[0])-'a'), int(rune(move[1])-'1'))
	to := fileRankToSquare(int(rune(move[2])-'a'), int(rune(move[3])-'1'))

	if DEBUG {
		if !squareOnBoard(from) || !squareOnBoard(to) {
			return NoMove, fmt.Errorf("Cant parse move %s Square not on board %d to %d", move, from, to)
		}
	}

	var list MoveListStruct
	err := pos.GenerateAllMoves(&list)
	if err != nil {
		return NoMove, err
	}

	var moveInt int
	var promotionPiece int

	for i := 0; i < list.Count; i++ {
		moveInt = list.Moves[i].Move
		if GetFrom(moveInt) == from && GetTo(moveInt) == to {
			promotionPiece = GetPromoted(moveInt)
			if promotionPiece != empty {
				if (promotionPiece == wQ || promotionPiece == bQ) && move[4] == 'q' {
					return moveInt, nil
				} else if (promotionPiece == wR || promotionPiece == bR) && move[4] == 'r' {
					return moveInt, nil
				} else if (promotionPiece == wB || promotionPiece == bB) && move[4] == 'b' {
					return moveInt, nil
				} else if (promotionPiece == wN || promotionPiece == bN) && move[4] == 'n' {
					return moveInt, nil
				}
				continue
			}
			return moveInt, nil
		}
	}
	return NoMove, nil
}

// Print a representation of the current board state to the console
func (pos *PositionStruct) Print() {
	fmt.Print("\nBoard State:\n\n")
	for rank := rank8; rank >= rank1; rank-- {
		fmt.Printf("%d", rank+1)
		for file := fileA; file <= fileH; file++ {
			sq := fileRankToSquare(file, rank)
			piece := pos.Pieces[sq]
			fmt.Printf("%3c", pieceChar[piece])
		}
		fmt.Print("\n")
	}

	fmt.Print(" ")
	for file := fileA; file <= fileH; file++ {
		fmt.Printf("%3c", 'A'+file)
	}
	fmt.Print("\n")
	fmt.Printf("Side: %c\n", sideChar[pos.Side])
	fmt.Printf("EnPassant: %s\n", SquareToString(pos.EnPassant))
	WK, WQ, BK, BQ := pos.castelPermToChar()
	fmt.Printf("Castel Perms: %c%c%c%c\n", WK, WQ, BK, BQ)
	fmt.Printf("Position Hash: %X\n", pos.PosKey)
	fmt.Printf("Is Repition: %v\n", pos.IsRepition())
	fmt.Printf("Evaluation CP: %d\n", pos.Evaluate())
	fmt.Printf("Fifty Move Count: %d\n", pos.FiftyMove)
	fmt.Printf("HisPly: %d\n", pos.HisPly)
}

// Print prints the move list struct to console
func (list *MoveListStruct) Print() {
	fmt.Println("MoveList:")

	for i := 0; i < list.Count; i++ {
		move := list.Moves[i].Move
		score := list.Moves[i].Score
		fmt.Printf("Move: %d > %s (score: %d)\n", i+1, MoveToString(move), score)
	}
	fmt.Printf("MoveList Total of %d Moves\n\n", list.Count)
}
