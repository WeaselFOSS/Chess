package board

//BoardStruct the boards struct
type BoardStruct struct {
	Pieces      [SquareNumber]int
	Pawns       [3]uint64
	KingSquare  [2]int
	PieceNum    [13]int
	BigPieces   [3]int
	MajorPieces [3]int
	MinorPieces [3]int

	CastelPerm int
	Side       int
	EnPassant  int
	FiftyMove  int
	Ply        int
	HisPly     int
	PosKey     uint64

	History []UndoStruct

	PieceList [13][10]int
}

//UndoStruct the undo move struct
type UndoStruct struct {
	Move      int
	CastelPem int
	EnPassant int
	FiftyMove int
	PosKey    uint64
}

//Sq120ToSq64 120 Square board to 64 square board index
var Sq120ToSq64 [SquareNumber]int

//Sq64ToSq120 64 Square board to 64 square board index
var Sq64ToSq120 [64]int

//SetMask set mask
var SetMask [64]uint64

//ClearMask clear mask
var ClearMask [64]uint64

func init() {
	initSq120To64()
	initBitMasks()
}

func initSq120To64() {
	for i := 0; i < SquareNumber; i++ {
		Sq120ToSq64[i] = 65
	}

	for i := 0; i < 64; i++ {
		Sq64ToSq120[i] = 120
	}

	sq64 := 0
	for rank := Rank1; rank <= Rank8; rank++ {
		for file := FileA; file <= FileH; file++ {
			sq := FileRankToSquare(file, rank)
			Sq64ToSq120[sq64] = sq
			Sq120ToSq64[sq] = sq64
			sq64++
		}
	}
}

func initBitMasks() {
	for i := 0; i < 64; i++ {
		SetMask[i] |= uint64(1) << uint64(i)
		ClearMask[i] = ^SetMask[i]

	}
}

//FileRankToSquare takes a file and rank and returns a square number
func FileRankToSquare(f int, r int) int {
	return 21 + f + r*10
}

func ClearBit(bitboard *uint64, square int) {
	*bitboard &= ClearMask[square]
}

func SetBit(bitboard *uint64, square int) {
	*bitboard |= SetMask[square]
}
