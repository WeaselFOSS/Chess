package engine

func (pos *BoardStruct) castelPermToChar() (rune, rune, rune, rune) {
	var wKt rune = '-'
	var wQt rune = '-'
	var bKt rune = '-'
	var bQt rune = '-'

	if pos.CastelPerm&wkcastel != 0 {
		wKt = 'K'
	}

	if pos.CastelPerm&wqcastel != 0 {
		wQt = 'Q'
	}

	if pos.CastelPerm&bkcastel != 0 {
		bKt = 'k'
	}

	if pos.CastelPerm&bqcastel != 0 {
		bQt = 'q'
	}

	return wKt, wQt, bKt, bQt
}

//IsPieceBig returns true if piece is big
func isPieceBig(piece int) bool {
	return piece != empty && piece != wP && piece != bP
}

//IsPieceBig returns true if piece is major
func isPieceMajor(piece int) bool {
	return piece != empty && piece != wP && piece != wN && piece != wB &&
		piece != bP && piece != bN && piece != bB

}

//IsPieceBig returns true if piece is minor
func isPieceMinor(piece int) bool {
	return piece == wN || piece == wB || piece == bN || piece == bB

}

//isPieceSlider returns true if piece is slider
func isPieceSlider(piece int) bool {
	return piece == wB || piece == wR || piece == wQ ||
		piece == bB || piece == bR || piece == bQ
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

//GetPieceColor returns the color of a piece
func getPieceColor(piece int) int {
	if piece >= wP && piece <= wK {
		return white
	}

	if piece >= bP && piece <= bK {
		return black
	}
	return both
}
