package engine

//MoveStruct used to store a move
type MoveStruct struct {
	Move  int
	Score int
}

type MoveListStruct struct {
	Moves [maxPositionMoves]MoveStruct
	Count int
}

/*
0000 0000 0000 0000 0000 0111 1111 -> From 0x7F
0000 0000 0000 0011 1111 1000 0000 -> To >> 7, 0x7F
0000 0000 0011 1100 0000 0000 0000 -> Captured >> 14, 0xF | 0x7C000
0000 0000 0100 0000 0000 0000 0000 -> EP 0x40000
0000 0000 1000 0000 0000 0000 0000 -> Pawn Start 0x80000
0000 1111 0000 0000 0000 0000 0000 -> Promoted Piece >> 20, 0xF | 0xF00000
0001 0000 0000 0000 0000 0000 0000 -> Castle 0x1000000
*/

var moveFlagEP = 0x40000
var moveFlagPS = 0x80000
var moveFlagCA = 0x1000000
var moveFlagCAP = 0x7C000
var moveFlagPROM = 0xF00000

func getFrom(move int) int {
	return move & 0x7F
}

func getTo(move int) int {
	return (move >> 7) & 0x7F
}

func getCapture(move int) int {
	return (move >> 14) & 0xF
}

func getPromoted(move int) int {
	return (move >> 20) & 0xf
}

func toMove(from, to, capture, promotion, flag int) int {
	return (from | (to << 7) | (capture << 14) | (promotion << 20) | flag)
}
