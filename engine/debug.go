package engine

import (
	"errors"
	"fmt"
)

var DEBUG bool = false

//CheckBoard VERY expensive only use for debugging
func (pos *BoardStruct) CheckBoard() error {

	var tPawns [3]uint64

	tPawns[White] = pos.Pawns[White]
	tPawns[Black] = pos.Pawns[Black]
	tPawns[Both] = pos.Pawns[Both]

	//Check all piece types and positions are correct
	for tPiece := WP; tPiece <= BK; tPiece++ {
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
		sq120 := Sq64ToSq120[sq64]
		tPiece := pos.Pieces[sq120]
		tPieceNum[tPiece]++
		color := getPieceColor(tPiece)

		if isPieceBig(tPiece) && color != Both {
			tBigPiece[color]++
		}

		if isPieceMinor(tPiece) && color != Both {
			tMinPiece[color]++
		}

		if isPieceMajor(tPiece) && color != Both {
			tMajPiece[color]++
		}

		if color != Both {
			tMaterial[color] += getPieceValue(tPiece)
		}
	}

	for tPiece := WP; tPiece < BK; tPiece++ {
		if tPieceNum[tPiece] != pos.PieceNum[tPiece] {
			return errors.New(fmt.Sprintf("Piece count mis match for %c Expected %d got %d", pieceChar[tPiece], tPieceNum[tPiece], pos.PieceNum[tPiece]))
		}
	}

	pcount := CountBits(tPawns[White])
	if pcount != pos.PieceNum[WP] {
		return errors.New(fmt.Sprintf("White Pawn BitBoard count mismatch, expected %d got %d", pos.PieceNum[WP], pcount))
	}

	pcount = CountBits(tPawns[Black])
	if pcount != pos.PieceNum[BP] {
		return errors.New(fmt.Sprintf("Black Pawn BitBoard count mismatch, expected %d got %d", pos.PieceNum[BP], pcount))
	}

	pcount = CountBits(tPawns[Both])
	if pcount != pos.PieceNum[BP]+pos.PieceNum[WP] {
		return errors.New(fmt.Sprintf("Both Pawn BitBoard count mismatch, expected %d got %d", pos.PieceNum[BP]+pos.PieceNum[WP], pcount))
	}

	for tPawns[White] != 0 {
		sq64 := PopBit(&tPawns[White])
		if pos.Pieces[Sq64ToSq120[sq64]] != WP {
			return errors.New(fmt.Sprintf("White Pawn BitBoard position mismatch on position %d", sq64))
		}
	}

	for tPawns[Black] != 0 {
		sq64 := PopBit(&tPawns[Black])
		if pos.Pieces[Sq64ToSq120[sq64]] != BP {
			return errors.New(fmt.Sprintf("Black Pawn BitBoard position mismatch on position %d", sq64))
		}
	}

	for tPawns[Both] != 0 {
		sq64 := PopBit(&tPawns[Both])
		if pos.Pieces[Sq64ToSq120[sq64]] != BP && pos.Pieces[Sq64ToSq120[sq64]] != WP {
			return errors.New(fmt.Sprintf("Both Pawn BitBoard position mismatch on position %d", sq64))
		}
	}

	if tMaterial[White] != pos.Material[White] || tMaterial[Black] != pos.Material[Black] {
		return errors.New("Material value mismatch")
	}

	if tMinPiece[White] != pos.MinorPieces[White] || tMinPiece[Black] != pos.MinorPieces[Black] {
		return errors.New("Minor piece mismatch")
	}

	if tMajPiece[White] != pos.MajorPieces[White] || tMajPiece[Black] != pos.MajorPieces[Black] {
		return errors.New("Major piece mismatch")
	}

	if tBigPiece[White] != pos.BigPieces[White] || tBigPiece[Black] != pos.BigPieces[Black] {
		return errors.New("Big piece mismatch")
	}

	if pos.Side != White && pos.Side != Black {
		return errors.New("No side to move is set")
	}

	if pos.EnPassant != NoSquare && (RanksBoard[pos.EnPassant] != Rank6 && pos.Side == White) ||
		(RanksBoard[pos.EnPassant] != Rank3 && pos.Side == Black) {
		return errors.New(fmt.Sprintf("Invalid EnPassant rank of %d", RanksBoard[pos.EnPassant]))
	}

	if pos.Pieces[pos.KingSquare[White]] != WK {
		return errors.New(fmt.Sprintf("White king square set to invalid position of %d", pos.KingSquare[White]))
	}

	if pos.Pieces[pos.KingSquare[Black]] != BK {
		return errors.New(fmt.Sprintf("Black king square set to invalid position of %d", pos.KingSquare[Black]))
	}

	return nil
}
