package engine

func (pos *BoardStruct) castelPermToChar() (rune, rune, rune, rune) {
	var WK rune = '-'
	var WQ rune = '-'
	var BK rune = '-'
	var BQ rune = '-'

	if pos.CastelPerm&Wkcastel != 0 {
		WK = 'K'
	}

	if pos.CastelPerm&Wqcastel != 0 {
		WQ = 'Q'
	}

	if pos.CastelPerm&Bkcastel != 0 {
		BK = 'k'
	}

	if pos.CastelPerm&Bqcastel != 0 {
		BQ = 'q'
	}

	return WK, WQ, BK, BQ
}

//IsPieceBig returns true if piece is big
func isPieceBig(piece int) bool {
	if piece != Empty && piece != WP && piece != BP {
		return true
	}
	return false
}

//IsPieceBig returns true if piece is major
func isPieceMajor(piece int) bool {
	if piece != Empty && piece != WP && piece != WN && piece != WB && piece != BP && piece != BN && piece != BB {
		return true
	}
	return false
}

//IsPieceBig returns true if piece is minor
func isPieceMinor(piece int) bool {
	if piece == WN || piece == WB || piece == BN || piece == BB {
		return true
	}
	return false
}

//GetPieceValue returns the value of a piece
func getPieceValue(piece int) int {
	switch piece {
	case WP, BP:
		return 100
	case WN, BN, WB, BB:
		return 325
	case WR, BR:
		return 550
	case WQ, BQ:
		return 1000
	case WK, BK:
		return 50000
	}
	return 0
}

//GetPieceColor returns the color of a piece
func getPieceColor(piece int) int {
	if piece >= WP && piece <= WK {
		return White
	}

	if piece >= BP && piece <= BK {
		return Black
	}
	return Both
}
