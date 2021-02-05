package engine

import (
	"errors"
	"fmt"
)

var DEBUG bool = true //Default set to true while debugging, should be false normally

//CheckBoard VERY expensive only use for debugging
func (pos *BoardStruct) CheckBoard() error {

	var tPawns [3]uint64

	tPawns[white] = pos.Pawns[white]
	tPawns[black] = pos.Pawns[black]
	tPawns[both] = pos.Pawns[both]

	//Check all piece types and positions are correct
	for tPiece := wP; tPiece <= bK; tPiece++ {
		for tPieceNum := 0; tPieceNum < pos.PieceNum[tPiece]; tPieceNum++ {
			sq120 := pos.PieceList[tPiece][tPieceNum]
			if pos.Pieces[sq120] != tPiece {
				return errors.New(fmt.Sprintf("Position piece at index %d is type %d expected %d", sq120, pos.Pieces[sq120], tPieceNum))
			}
		}
	}

	//Check piece count and other counters
	var tPieceNum [13]int
	var tBigPiece [2]int
	var tMajPiece [2]int
	var tMinPiece [2]int
	var tMaterial [2]int
	for sq64 := 0; sq64 < 64; sq64++ {
		sq120 := sq64ToSq120[sq64]
		tPiece := pos.Pieces[sq120]
		tPieceNum[tPiece]++
		color := getPieceColor(tPiece)

		if isPieceBig(tPiece) && color != both {
			tBigPiece[color]++
		}

		if isPieceMinor(tPiece) && color != both {
			tMinPiece[color]++
		}

		if isPieceMajor(tPiece) && color != both {
			tMajPiece[color]++
		}

		if color != both {
			tMaterial[color] += getPieceValue(tPiece)
		}
	}

	for tPiece := wP; tPiece < bK; tPiece++ {
		if tPieceNum[tPiece] != pos.PieceNum[tPiece] {
			return errors.New(fmt.Sprintf("Piece count mis match for %c Expected %d got %d", pieceChar[tPiece], tPieceNum[tPiece], pos.PieceNum[tPiece]))
		}
	}

	pcount := countBits(tPawns[white])
	if pcount != pos.PieceNum[wP] {
		return errors.New(fmt.Sprintf("White Pawn BitBoard count mismatch, expected %d got %d", pos.PieceNum[wP], pcount))
	}

	pcount = countBits(tPawns[black])
	if pcount != pos.PieceNum[bP] {
		return errors.New(fmt.Sprintf("Black Pawn BitBoard count mismatch, expected %d got %d", pos.PieceNum[bP], pcount))
	}

	pcount = countBits(tPawns[both])
	if pcount != pos.PieceNum[bP]+pos.PieceNum[wP] {
		return errors.New(fmt.Sprintf("Both Pawn BitBoard count mismatch, expected %d got %d", pos.PieceNum[bP]+pos.PieceNum[wP], pcount))
	}

	for tPawns[white] != 0 {
		sq64 := popBit(&tPawns[white])
		if pos.Pieces[sq64ToSq120[sq64]] != wP {
			return errors.New(fmt.Sprintf("White Pawn BitBoard position mismatch on position %d", sq64))
		}
	}

	for tPawns[black] != 0 {
		sq64 := popBit(&tPawns[black])
		if pos.Pieces[sq64ToSq120[sq64]] != bP {
			return errors.New(fmt.Sprintf("Black Pawn BitBoard position mismatch on position %d", sq64))
		}
	}

	for tPawns[both] != 0 {
		sq64 := popBit(&tPawns[both])
		if pos.Pieces[sq64ToSq120[sq64]] != bP && pos.Pieces[sq64ToSq120[sq64]] != wP {
			return errors.New(fmt.Sprintf("Both Pawn BitBoard position mismatch on position %d", sq64))
		}
	}

	if tMaterial[white] != pos.Material[white] || tMaterial[black] != pos.Material[black] {
		return errors.New("Material value mismatch")
	}

	if tMinPiece[white] != pos.MinorPieces[white] || tMinPiece[black] != pos.MinorPieces[black] {
		return errors.New("Minor piece mismatch")
	}

	if tMajPiece[white] != pos.MajorPieces[white] || tMajPiece[black] != pos.MajorPieces[black] {
		return errors.New("Major piece mismatch")
	}

	if tBigPiece[white] != pos.BigPieces[white] || tBigPiece[black] != pos.BigPieces[black] {
		return errors.New("Big piece mismatch")
	}

	if pos.Side != white && pos.Side != black {
		return errors.New("No side to move is set")
	}

	if pos.EnPassant != noSquare && (ranksBoard[pos.EnPassant] != rank6 && pos.Side == white) ||
		(ranksBoard[pos.EnPassant] != rank3 && pos.Side == black) {
		return errors.New(fmt.Sprintf("Invalid EnPassant rank of %d", ranksBoard[pos.EnPassant]))
	}

	if pos.Pieces[pos.KingSquare[white]] != wK {
		return errors.New(fmt.Sprintf("White king square set to invalid position of %d", pos.KingSquare[white]))
	}

	if pos.Pieces[pos.KingSquare[black]] != bK {
		return errors.New(fmt.Sprintf("Black king square set to invalid position of %d", pos.KingSquare[black]))
	}

	poskey, err := pos.generatePosKey()
	if err != nil {
		return err
	}

	if pos.PosKey != poskey {
		return errors.New(fmt.Sprintf("Position Hash mis match expected %d, got %d", poskey, pos.PosKey))
	}

	return nil
}

func squareOnBoard(sq int) bool {
	return filesBoard[sq] != offBoard
}

func sideValid(side int) bool {
	return (side == white || side == black)
}

func fileRankValid(fr int) bool {
	return (fr >= 0 && fr <= 7)
}

func pieceValidEmpty(piece int) bool {
	return (piece >= empty && piece <= bK)
}

func pieceValid(piece int) bool {
	return (piece >= wP && piece <= bK)
}
