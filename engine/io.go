package engine

import "fmt"

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
func moveToString(move int) string {
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

func (list *MoveListStruct) Print() {
	fmt.Println("MoveList:")

	for i := 0; i < list.Count; i++ {
		move := list.Moves[i].Move
		score := list.Moves[i].Score
		fmt.Printf("Move: %d > %s (score: %d)\n", i+1, moveToString(move), score)
	}
	fmt.Printf("MoveList Total of %d Moves\n\n", list.Count)
}
