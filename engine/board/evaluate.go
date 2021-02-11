package board

import "fmt"

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

//bishopPair bonus for having a bishop pair
const bishopPair = 30

//endGameMaterial If there are no queens on the board or material is less than this, than we are in a end game
var endGameMaterial = GetPieceValue(wR) + 2*GetPieceValue(wN) + 2*GetPieceValue(wP) + GetPieceValue(wK)

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

var kingE = [64]int{
	-50, -10, 0, 0, 0, 0, -10, -50,
	-10, 0, 10, 10, 10, 10, 0, -10,
	0, 10, 15, 15, 15, 15, 10, 0,
	0, 10, 15, 20, 20, 15, 10, 0,
	0, 10, 15, 20, 20, 15, 10, 0,
	0, 10, 15, 15, 15, 15, 10, 0,
	-10, 0, 10, 10, 10, 10, 0, -10,
	-50, -10, 0, 0, 0, 0, -10, -50,
}

var kingO = [64]int{
	0, 5, 5, -10, -10, 0, 10, 5,
	-30, -30, -30, -30, -30, -30, -30, -30,
	-50, -50, -50, -50, -50, -50, -50, -50,
	-70, -70, -70, -70, -70, -70, -70, -70,
	-70, -70, -70, -70, -70, -70, -70, -70,
	-70, -70, -70, -70, -70, -70, -70, -70,
	-70, -70, -70, -70, -70, -70, -70, -70,
	-70, -70, -70, -70, -70, -70, -70, -70,
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

//materialDraw Test if the position is a material draw
//Credit sjeng 11.2
func (pos *PositionStruct) materialDraw() bool {
	if pos.PieceNum[wR] == 0 && pos.PieceNum[bR] == 0 && pos.PieceNum[wQ] == 0 && pos.PieceNum[bQ] == 0 {
		if pos.PieceNum[bB] == 0 && pos.PieceNum[wB] == 0 {
			if pos.PieceNum[wN] < 3 && pos.PieceNum[bN] < 3 {
				return true
			}
		} else if pos.PieceNum[wN] == 0 && pos.PieceNum[bN] == 0 {
			if pos.PieceNum[wB]-pos.PieceNum[bB] < 2 {
				return true
			}
		} else if (pos.PieceNum[wN] < 3 && pos.PieceNum[wB] == 0) || (pos.PieceNum[wB] == 1 && pos.PieceNum[wN] == 0) {
			if (pos.PieceNum[bN] < 3 && pos.PieceNum[bB] == 0) || (pos.PieceNum[bB] == 1 && pos.PieceNum[bN] == 0) {
				return true
			}
		}
	} else if pos.PieceNum[wQ] == 0 && pos.PieceNum[bQ] == 0 {
		if pos.PieceNum[wR] == 1 && pos.PieceNum[bR] == 1 {
			if (pos.PieceNum[wN]+pos.PieceNum[wB]) < 2 && (pos.PieceNum[bN]+pos.PieceNum[bB]) < 2 {
				return true
			}
		} else if pos.PieceNum[wR] == 1 && pos.PieceNum[bR] == 0 {
			if (pos.PieceNum[wN]+pos.PieceNum[wB] == 0) && (((pos.PieceNum[bN] + pos.PieceNum[bB]) == 1) || ((pos.PieceNum[bN] + pos.PieceNum[bB]) == 2)) {
				return true
			}
		} else if pos.PieceNum[bR] == 1 && pos.PieceNum[wR] == 0 {
			if (pos.PieceNum[bN]+pos.PieceNum[bB] == 0) && (((pos.PieceNum[wN] + pos.PieceNum[wB]) == 1) || ((pos.PieceNum[wN] + pos.PieceNum[wB]) == 2)) {
				return true
			}
		}
	}
	return false
}

//Evaluate the currect position and return a score
func (pos *PositionStruct) Evaluate() int {
	score := pos.Material[white] - pos.Material[black]

	if pos.PieceNum[wP] == 0 && pos.PieceNum[bP] == 0 && pos.materialDraw() {
		return 0
	}

	//Pawn square tables and isolated / passed check
	for i := 0; i < pos.PieceNum[wP]; i++ {
		sq := pos.PieceList[wP][i]
		score += pawnTable[sq120ToSq64[sq]]

		if isolatedMasks[sq120ToSq64[sq]]&pos.Pawns[white] == 0 {
			score += pawnIsolated
		}

		if whitePassedMasks[sq120ToSq64[sq]]&pos.Pawns[black] == 0 {
			score += pawnPassed[ranksBoard[sq]]
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
			score -= queenOpenFile
		} else if pos.Pawns[black]&fileMasks[filesBoard[sq]] == 0 {
			score -= queenSemiOpenFile
		}
	}

	//King piece squares
	sq := pos.PieceList[wK][0]
	if pos.Material[black] <= endGameMaterial {
		score += kingE[sq120ToSq64[sq]]
	} else {
		score += kingO[sq120ToSq64[sq]]
	}

	sq = pos.PieceList[bK][0]
	if pos.Material[white] <= endGameMaterial {
		score -= kingE[mirror64[sq120ToSq64[sq]]]
	} else {
		score -= kingO[mirror64[sq120ToSq64[sq]]]
		fmt.Println(pos.Material[white], endGameMaterial)
	}

	if pos.PieceNum[wB] >= 2 {
		score += bishopPair
	}

	if pos.PieceNum[bB] >= 2 {
		score -= bishopPair
	}

	//Return a positive score no matter the side to move
	if pos.Side == white {
		return score
	}
	return -score
}

//GetPieceValue returns the value of a piece
func GetPieceValue(piece int) int {
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
