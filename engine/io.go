package engine

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

//StringToMove algebraic notation to move int
func StringToMove(move string) int {
	fromFile := int(rune(move[0]) - 'a')
	fromRank := int(rune(move[1]) - '1')
	toFile := int(rune(move[2]) - 'a')
	toRank := int(rune(move[3]) - '1')

	fromSquare := fileRankToSquare(fromFile, fromRank)
	toSquare := fileRankToSquare(toFile, toRank)

	var promotion int

	if len(move) > 4 {
		promotion = int(rune(move[4]) - 'a') //TODO: Redo with global piece constants
	}

	return toMove(fromSquare, toSquare, 0, promotion, 0)
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
