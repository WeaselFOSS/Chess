package board

//pawnIsolated penalty for having an isolated pawn
const pawnIsolated = -10

//rookOpenFile Bonus for having a rook on a open file
const rookOpenFile = 10

//rookSemiOpenFile Bonus for having a rook on a semi open file
const rookSemiOpenFile = 5

//queenOpenFile Bonus for having a queen on a open file
const queenOpenFile = 5

//queenSemiOpenFile Bonus for having a queen on a semi open file
const queenSemiOpenFile = 3

//Bonus for pushing passed pawns
var pawnPassed = [8]int{0, 5, 10, 20, 35, 60, 100, 200}

//These are piece square tables. A simple method to get a better evaluation
var pawnTable = [64]int{
	0, 0, 0, 0, 0, 0, 0, 0,
	10, 10, 0, -10, -10, 0, 10, 10,
	5, 0, 0, 5, 5, 0, 0, 5,
	0, 0, 10, 20, 20, 10, 0, 0,
	5, 5, 5, 10, 10, 5, 5, 5,
	10, 10, 10, 20, 20, 10, 10, 10,
	20, 20, 20, 30, 30, 20, 20, 20,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var knightTable = [64]int{
	0, -10, 0, 0, 0, 0, -10, 0,
	0, 0, 0, 5, 5, 0, 0, 0,
	0, 0, 10, 10, 10, 10, 0, 0,
	0, 0, 10, 20, 20, 10, 5, 0,
	5, 10, 15, 20, 20, 15, 10, 5,
	5, 10, 10, 20, 20, 10, 10, 5,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var bishopTable = [64]int{
	0, 0, -10, 0, 0, -10, 0, 0,
	0, 0, 0, 10, 10, 0, 0, 0,
	0, 0, 10, 15, 15, 10, 0, 0,
	0, 10, 15, 20, 20, 15, 10, 0,
	0, 10, 15, 20, 20, 15, 10, 0,
	0, 0, 10, 15, 15, 10, 0, 0,
	0, 0, 0, 10, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var rookTable = [64]int{
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	25, 25, 25, 25, 25, 25, 25, 25,
	0, 0, 5, 10, 10, 5, 0, 0,
}

//mirror64 Mirror the piece square tables for the other side
var mirror64 = [64]int{
	56, 57, 58, 59, 60, 61, 62, 63,
	48, 49, 50, 51, 52, 53, 54, 55,
	40, 41, 42, 43, 44, 45, 46, 47,
	32, 33, 34, 35, 36, 37, 38, 39,
	24, 25, 26, 27, 28, 29, 30, 31,
	16, 17, 18, 19, 20, 21, 22, 23,
	8, 9, 10, 11, 12, 13, 14, 15,
	0, 1, 2, 3, 4, 5, 6, 7,
}

//Evaluate the currect position and return a score
func (pos *PositionStruct) Evaluate() int {
	score := pos.Material[white] - pos.Material[black]

	//Pawn square tables and isolated / passed check
	for i := 0; i < pos.PieceNum[wP]; i++ {
		sq := pos.PieceList[wP][i]
		score += pawnTable[sq120ToSq64[sq]]

		if isolatedMasks[sq120ToSq64[sq]]&pos.Pawns[white] == 0 {
			score += pawnIsolated
		}

		if whitePassedMasks[sq120ToSq64[sq]]&pos.Pawns[black] == 0 {
			score -= pawnPassed[ranksBoard[sq]]
		}
	}

	for i := 0; i < pos.PieceNum[bP]; i++ {
		sq := pos.PieceList[bP][i]
		score -= pawnTable[mirror64[sq120ToSq64[sq]]]

		if isolatedMasks[sq120ToSq64[sq]]&pos.Pawns[black] == 0 {
			score -= pawnIsolated
		}

		if blackPassedMasks[sq120ToSq64[sq]]&pos.Pawns[white] == 0 {
			score -= pawnPassed[7-ranksBoard[sq]]
		}
	}

	//Knight square tables
	for i := 0; i < pos.PieceNum[wN]; i++ {
		sq := pos.PieceList[wN][i]
		score += knightTable[sq120ToSq64[sq]]
	}

	for i := 0; i < pos.PieceNum[bN]; i++ {
		sq := pos.PieceList[bN][i]
		score -= knightTable[mirror64[sq120ToSq64[sq]]]
	}

	//Bishop square tables
	for i := 0; i < pos.PieceNum[wB]; i++ {
		sq := pos.PieceList[wB][i]
		score += bishopTable[sq120ToSq64[sq]]
	}

	for i := 0; i < pos.PieceNum[bB]; i++ {
		sq := pos.PieceList[bB][i]
		score -= bishopTable[mirror64[sq120ToSq64[sq]]]
	}

	//Rook square tables and open files
	for i := 0; i < pos.PieceNum[wR]; i++ {
		sq := pos.PieceList[wR][i]
		score += rookTable[sq120ToSq64[sq]]

		if pos.Pawns[both]&fileMasks[filesBoard[sq]] == 0 {
			score += rookOpenFile
		} else if pos.Pawns[white]&fileMasks[filesBoard[sq]] == 0 {
			score += rookSemiOpenFile
		}
	}

	for i := 0; i < pos.PieceNum[bR]; i++ {
		sq := pos.PieceList[bR][i]
		score -= rookTable[mirror64[sq120ToSq64[sq]]]

		if pos.Pawns[both]&fileMasks[filesBoard[sq]] == 0 {
			score -= rookOpenFile
		} else if pos.Pawns[black]&fileMasks[filesBoard[sq]] == 0 {
			score -= rookSemiOpenFile
		}
	}

	//Queen open/semi open files
	for i := 0; i < pos.PieceNum[wQ]; i++ {
		sq := pos.PieceList[wQ][i]
		if pos.Pawns[both]&fileMasks[filesBoard[sq]] == 0 {
			score += queenOpenFile
		} else if pos.Pawns[white]&fileMasks[filesBoard[sq]] == 0 {
			score += queenSemiOpenFile
		}
	}

	for i := 0; i < pos.PieceNum[bQ]; i++ {
		sq := pos.PieceList[bQ][i]
		if pos.Pawns[both]&fileMasks[filesBoard[sq]] == 0 {
			score += queenOpenFile
		} else if pos.Pawns[black]&fileMasks[filesBoard[sq]] == 0 {
			score += queenSemiOpenFile
		}
	}

	//Return a positive score no matter the side to move
	if pos.Side == white {
		return score
	}
	return -score
}

//GetPieceValue returns the value of a piece
func getPieceValue(piece int) int {
	switch piece {
	case wP, bP:
		return 100
	case wN, bN, wB, bB:
		return 325
	case wR, bR:
		return 550
	case wQ, bQ:
		return 1000
	case wK, bK:
		return 50000
	}
	return 0
}
