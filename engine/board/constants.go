package board

//squareNumber Number of squares in the board representastion
const squareNumber = 120

//maxPositionMoves the max number of moves we can expect from any given position
const maxPositionMoves = 256

//StartPosFEN The fen string for a starting position
const StartPosFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

//MaxDepth The max depth the engine will try to search to
const MaxDepth = 64

//NoMove constant for no move found
const NoMove = 0

const (
	//empty empty square
	empty = iota
	//wP white Pawn
	wP
	//wN white Knight
	wN
	//wB white Bishop
	wB
	//wR white Rook
	wR
	//wQ white Queen
	wQ
	//wK white King
	wK
	//bP black Pawn
	bP
	//black Knights
	bN
	//bb black bishop
	bB
	//bR black Rook
	bR
	//bQ black Queen
	bQ
	//bK black King
	bK
)

const (
	//fileA const
	fileA = iota
	//fileB const
	fileB
	//filec const
	fileC
	//fileD const
	fileD
	//fileE const
	fileE
	//fileF const
	fileF
	//FileG const
	fileG
	//FileH const
	fileH
	//FileNone const for no file
	fileNone
)

const (
	//rank1 const
	rank1 = iota
	//rank2 const
	rank2
	//rank3 const
	rank3
	//rank4 const
	rank4
	//rank5 const
	rank5
	//rank6 const
	rank6
	//rank7 const
	rank7
	//rank8 const
	rank8
	//rankNone const for no rank
	rankNone
)

const (
	//white side const
	white = iota
	//black side const
	black
	//both side const
	both
)

const (
	//a1 coord const
	a1 = iota + 21
	//b1 coord const
	b1
	//c1 coord const
	c1
	//d1 coord const
	d1
	//e1 coord const
	e1
	//f1 coord const
	f1
	//g1 coord const
	g1
	//h1 coord const
	h1
)

const (
	//a2 coord const
	a2 = iota + 31
	//b2 coord const
	b2
	//c2 coord const
	c2
	//d2 coord const
	d2
	//e2 coord const
	e2
	//f2 coord const
	f2
	//g2 coord const
	g2
	//h2 coord const
	h2
)

const (
	//a3 coord const
	a3 = iota + 41
	//b3 coord const
	b3
	//c3 coord const
	c3
	//d3 coord const
	d3
	//e3 coord const
	e3
	//f3 coord const
	f3
	//g3 coord const
	g3
	//h3 coord const
	h3
)

const (
	//a4 coord const
	a4 = iota + 51
	//b4 coord const
	b4
	//c4 coord const
	c4
	//d4 coord const
	d4
	//e4 coord const
	e4
	//f4 coord const
	f4
	//g4 coord const
	g4
	//h4 coord const
	h4
)

const (
	//a5 coord const
	a5 = iota + 61
	//b5 coord const
	b5
	//c5 coord const
	c5
	//d5 coord const
	d5
	//e5 coord const
	e5
	//f5 coord const
	f5
	//g5 coord const
	g5
	//h5 coord const
	h5
)

const (
	//a6 coord const
	a6 = iota + 71
	//b6 coord const
	b6
	//c6 coord const
	c6
	//d6 coord const
	d6
	//e6 coord const
	e6
	//f6 coord const
	f6
	//g6 coord const
	g6
	//h6 coord const
	h6
)

const (
	//a7 coord const
	a7 = iota + 81
	//b7 coord const
	b7
	//c7 coord const
	c7
	//d7 coord const
	d7
	//e7 coord const
	e7
	//f7 coord const
	f7
	//g7 coord const
	g7
	//h7 coord const
	h7
)

const (
	//a8 coord const
	a8 = iota + 91
	//b8 coord const
	b8
	//c8 coord const
	c8
	//d8 coord const
	d8
	//e8 coord const
	e8
	//f8 coord const
	f8
	//g8 coord const
	g8
	//h8 coord const
	h8
	//noSquare coord const
	noSquare
	//offBoard coord const
	offBoard
)

const (
	//wkcastel white King castel Perm
	wkcastel = 1
	//wqcastel white Queen castel Perm
	wqcastel = 2
	//bkcastel Black King castel Perm
	bkcastel = 4
	//bqcastel Black Queen castel Perm
	bqcastel = 8
)
