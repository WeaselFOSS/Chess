package board

// SquareNumber Number of squares in the board representastion
const SquareNumber = 120

// maxPositionMoves the max number of moves we can expect from any given position
const maxPositionMoves = 256

// StartPosFEN The fen string for a starting position
const StartPosFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 0"

// MaxDepth The max depth the engine will try to search to
const MaxDepth = 64

// NoMove constant for no move found
const NoMove = 0

// Infinite Largest score value
const Infinite = 30000

const IsMate = Infinite - MaxDepth

const (
	// HFNONE NONE flag for hash table
	HFNONE = iota
	// HFALPHA Alpha flag for hash table
	HFALPHA
	// HFBETA Beta flag for hash table
	HFBETA
	// HFEXACT Exact flag for hash table
	HFEXACT
)

const (
	empty = iota
	wP
	wN
	wB
	wR
	wQ
	wK
	bP
	bN
	bB
	bR
	bQ
	bK
)

//nolint
const (
	fileA = iota
	fileB
	fileC
	fileD
	fileE
	fileF
	fileG
	fileH
	fileNone
)

//nolint
const (
	rank1 = iota
	rank2
	rank3
	rank4
	rank5
	rank6
	rank7
	rank8
	rankNone
)

const (
	white = iota
	black
	both
)

//nolint
const (
	a1 = iota + 21
	b1
	c1
	d1
	e1
	f1
	g1
	h1
)

//nolint
const (
	a2 = iota + 31
	b2
	c2
	d2
	e2
	f2
	g2
	h2
)

//nolint
const (
	a3 = iota + 41
	b3
	c3
	d3
	e3
	f3
	g3
	h3
)

//nolint
const (
	a4 = iota + 51
	b4
	c4
	d4
	e4
	f4
	g4
	h4
)

//nolint
const (
	a5 = iota + 61
	b5
	c5
	d5
	e5
	f5
	g5
	h5
)

//nolint
const (
	a6 = iota + 71
	b6
	c6
	d6
	e6
	f6
	g6
	h6
)

//nolint
const (
	a7 = iota + 81
	b7
	c7
	d7
	e7
	f7
	g7
	h7
)

//nolint
const (
	a8 = iota + 91
	b8
	c8
	d8
	e8
	f8
	g8
	h8
	noSquare
	offBoard
)

const (
	wkcastel = 1
	wqcastel = 2
	bkcastel = 4
	bqcastel = 8
)
