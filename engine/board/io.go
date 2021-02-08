package board

import (
	"fmt"
)

//SquareToString square to algebraic notation
func SquareToString(sq int) string {
	if sq == offBoard || sq == noSquare {
		return "None"
	}
	file := filesBoard[sq]
	rank := ranksBoard[sq]
	return fmt.Sprintf("%c%c", ('a' + file), ('1' + rank))
}

//MoveToString move int to algebraic notation
func MoveToString(move int) string {
	ff := filesBoard[getFrom(move)]
	rf := ranksBoard[getFrom(move)]
	ft := filesBoard[getTo(move)]
	rt := ranksBoard[getTo(move)]

	promoted := getPromoted(move)
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

//ParseMove algebraic notation to move int
func (pos *PositionStruct) ParseMove(move string) (int, error) {
	if len(move) < 4 {
		return NoMove, fmt.Errorf("Move %s out of bounds", move)
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

	var moveInt = 0
	var promotionPiece = empty

	for i := 0; i < list.Count; i++ {
		moveInt = list.Moves[i].Move
		if getFrom(moveInt) == from && getTo(moveInt) == to {
			promotionPiece = getPromoted(moveInt)
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

//Print print current board state to console
func (list *MoveListStruct) Print() {
	fmt.Println("MoveList:")

	for i := 0; i < list.Count; i++ {
		move := list.Moves[i].Move
		score := list.Moves[i].Score
		fmt.Printf("Move: %d > %s (score: %d)\n", i+1, MoveToString(move), score)
	}
	fmt.Printf("MoveList Total of %d Moves\n\n", list.Count)
}
