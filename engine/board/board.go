package board

// PositionStruct the boards struct
type PositionStruct struct {
	Pieces      [SquareNumber]int
	Pawns       [3]uint64
	KingSquare  [2]int
	PieceNum    [13]int
	BigPieces   [2]int
	MajorPieces [2]int
	MinorPieces [2]int
	Material    [2]int

	CastelPerm int
	Side       int
	EnPassant  int
	FiftyMove  int
	Ply        int
	HisPly     int
	PosKey     uint64

	History [2048]UndoStruct

	PieceList [13][10]int

	HashTable HashTableStruct
	PvArray   [MaxDepth]int

	SearchHistory [13][SquareNumber]int
	SearchKillers [2][MaxDepth]int
}

// UndoStruct the undo move struct
type UndoStruct struct {
	Move       int
	CastelPerm int
	EnPassant  int
	FiftyMove  int
	PosKey     uint64
}

// sq120ToSq64 120 Square board to 64 square board index
var sq120ToSq64 [SquareNumber]int

// sq64ToSq120 64 Square board to 64 square board index
var sq64ToSq120 [64]int

// filesBoard Get a positions file
var filesBoard [SquareNumber]int

// ranksBoard Get a positions rank
var ranksBoard [SquareNumber]int

var pieceChar = [13]rune{'.', 'P', 'N', 'B', 'R', 'Q', 'K', 'p', 'n', 'b', 'r', 'q', 'k'}
var sideChar = [3]rune{'w', 'b', '-'}

// fileRankToSquare takes a file and rank and returns a square number
func fileRankToSquare(file int, rank int) int {
	return 21 + file + rank*10
}

// resetBoard Reset the board
func (pos *PositionStruct) resetBoard() {
	for i := 0; i < SquareNumber; i++ {
		pos.Pieces[i] = offBoard
	}

	for i := 0; i < 64; i++ {
		pos.Pieces[sq64ToSq120[i]] = empty
	}

	for i := 0; i < 2; i++ {
		pos.BigPieces[i] = 0
		pos.MajorPieces[i] = 0
		pos.MinorPieces[i] = 0
		pos.Material[i] = 0
	}

	for i := 0; i < 3; i++ {
		pos.Pawns[i] = uint64(0)
	}

	for i := 0; i < 13; i++ {
		pos.PieceNum[i] = 0
	}

	pos.KingSquare[white] = noSquare
	pos.KingSquare[black] = noSquare

	pos.Side = both
	pos.EnPassant = noSquare
	pos.FiftyMove = 0

	pos.Ply = 0
	pos.HisPly = 0

	pos.CastelPerm = 0

	pos.PosKey = uint64(0)
}

// updateMaterialLists Update the material lists for the baord
func (pos *PositionStruct) updateMaterialLists() {
	for i := 0; i < SquareNumber; i++ {
		piece := pos.Pieces[i]
		if piece != offBoard && piece != empty {
			color := getPieceColor(piece)
			if isPieceBig(piece) {
				pos.BigPieces[color]++
			}

			if isPieceMajor(piece) {
				pos.MajorPieces[color]++
			}

			if isPieceMinor(piece) {
				pos.MinorPieces[color]++
			}

			pos.Material[color] += GetPieceValue(piece)
			pos.PieceList[piece][pos.PieceNum[piece]] = i
			pos.PieceNum[piece]++

			if piece == wK || piece == bK {
				pos.KingSquare[color] = i
			}

			if piece == wP || piece == bP {
				setBit(&pos.Pawns[color], sq120ToSq64[i])
				setBit(&pos.Pawns[both], sq120ToSq64[i])
			}
		}
	}
}

// mirrorBoard Mirror the position
func (pos *PositionStruct) MirrorBoard() error {
	var tempPieceArray [64]int
	tempSide := pos.Side ^ 1
	var swapPieces = [13]int{empty, bP, bN, bB, bR, bQ, bK, wP, wN, wB, wR, wQ, wK}
	var tempCastelPerm = 0
	var tempEnpas = noSquare

	if pos.CastelPerm&wkcastel != 0 {
		tempCastelPerm |= bkcastel
	}

	if pos.CastelPerm&wqcastel != 0 {
		tempCastelPerm |= bqcastel
	}

	if pos.CastelPerm&bkcastel != 0 {
		tempCastelPerm |= wkcastel
	}

	if pos.CastelPerm&bqcastel != 0 {
		tempCastelPerm |= wqcastel
	}

	if pos.EnPassant != noSquare {
		tempEnpas = sq64ToSq120[mirror64[sq120ToSq64[pos.EnPassant]]]
	}

	for sq := 0; sq < 64; sq++ {
		tempPieceArray[sq] = pos.Pieces[sq64ToSq120[mirror64[sq]]]
	}

	pos.resetBoard()

	for sq := 0; sq < 64; sq++ {
		tp := swapPieces[tempPieceArray[sq]]
		pos.Pieces[sq64ToSq120[sq]] = tp
	}

	pos.Side = tempSide
	pos.CastelPerm = tempCastelPerm
	pos.EnPassant = tempEnpas

	var err error
	pos.PosKey, err = pos.generatePosKey()
	if err != nil {
		return err
	}

	pos.updateMaterialLists()

	if DEBUG {
		err = pos.CheckBoard()
		if err != nil {
			return err
		}
	}
	return nil
}
