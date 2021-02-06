package engine

//MoveStruct used to store a move
type MoveStruct struct {
	Move  int
	Score int
}

//MoveListStruct List of moves
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

//moveFlagEP EnPas flag
var moveFlagEP = 0x40000

//moveFlagPS Pawn Start flag
var moveFlagPS = 0x80000

//moveFlagCA Castel flag
var moveFlagCA = 0x1000000

//moveFlagCAP Capture flag
var moveFlagCAP = 0x7C000

//moveFlagPROM Promotion flag
var moveFlagPROM = 0xF00000

//getFrom value from move int
func getFrom(move int) int {
	return move & 0x7F
}

//getTo get TO value from move int
func getTo(move int) int {
	return (move >> 7) & 0x7F
}

//getCapture get capture value from move int
func getCapture(move int) int {
	return (move >> 14) & 0xF
}

//getPromoted Get promote value from move int
func getPromoted(move int) int {
	return (move >> 20) & 0xf
}

//toMove Puts all move info into a single move int
func toMove(from, to, capture, promotion, flag int) int {
	return (from | (to << 7) | (capture << 14) | (promotion << 20) | flag)
}
