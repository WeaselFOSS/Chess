package board

//SquareNumber Number of squares in the board representastion
const SquareNumber = 120

const (
	//Empty empty square
	Empty = iota
	//WP White Pawn
	WP
	//WB White Bishop
	WB
	//WN White Knight
	WN
	//WR White Rook
	WR
	//WQ White Queen
	WQ
	//WK White King
	WK
	//BP Black Pawn
	BP
	//BB Black Bishop
	BB
	//Black Knights
	BN
	//BR Black Rook
	BR
	//BQ Black Queen
	BQ
	//BK Black King
	BK
)

const (
	//FileA const
	FileA = iota
	//FileB const
	FileB
	//FileC const
	FileC
	//FileD const
	FileD
	//FileE const
	FileE
	//FileF const
	FileF
	//FileG const
	FileG
	//FileH const
	FileH
	//FileNone const for no file
	FileNone
)

const (
	//Rank1 const
	Rank1 = iota
	//Rank2 const
	Rank2
	//Rank3 const
	Rank3
	//Rank4 const
	Rank4
	//Rank5 const
	Rank5
	//Rank6 const
	Rank6
	//Rank7 const
	Rank7
	//Rank8 const
	Rank8
	//RankNone const for no rank
	RankNone
)

const (
	//White side const
	White = iota
	//Black side const
	Black
	//Both side const
	Both
)

const (
	//A1 coord const
	A1 = iota + 21
	//B1 coord const
	B1
	//C1 coord const
	C1
	//D1 coord const
	D1
	//E1 coord const
	E1
	//F1 coord const
	F1
	//G1 coord const
	G1
	//H1 coord const
	H1
)

const (
	//A2 coord const
	A2 = iota + 31
	//B2 coord const
	B2
	//C2 coord const
	C2
	//D2 coord const
	D2
	//E2 coord const
	E2
	//F2 coord const
	F2
	//G2 coord const
	G2
	//H2 coord const
	H2
)

const (
	//A3 coord const
	A3 = iota + 41
	//B3 coord const
	B3
	//C3 coord const
	C3
	//D3 coord const
	D3
	//E3 coord const
	E3
	//F3 coord const
	F3
	//G3 coord const
	G3
	//H3 coord const
	H3
)

const (
	//A4 coord const
	A4 = iota + 51
	//B4 coord const
	B4
	//C4 coord const
	C4
	//D4 coord const
	D4
	//E4 coord const
	E4
	//F4 coord const
	F4
	//G4 coord const
	G4
	//H4 coord const
	H4
)

const (
	//A5 coord const
	A5 = iota + 61
	//B5 coord const
	B5
	//C5 coord const
	C5
	//D5 coord const
	D5
	//E5 coord const
	E5
	//F5 coord const
	F5
	//G5 coord const
	G5
	//H5 coord const
	H5
)

const (
	//A6 coord const
	A6 = iota + 71
	//B6 coord const
	B6
	//C6 coord const
	C6
	//D6 coord const
	D6
	//E6 coord const
	E6
	//F6 coord const
	F6
	//G6 coord const
	G6
	//H6 coord const
	H6
)

const (
	//A7 coord const
	A7 = iota + 81
	//B7 coord const
	B7
	//C7 coord const
	C7
	//D7 coord const
	D7
	//E7 coord const
	E7
	//F7 coord const
	F7
	//G7 coord const
	G7
	//H7 coord const
	H7
)

const (
	//A8 coord const
	A8 = iota + 91
	//B8 coord const
	B8
	//C8 coord const
	C8
	//D8 coord const
	D8
	//E8 coord const
	E8
	//F8 coord const
	F8
	//G8 coord const
	G8
	//H8 coord const
	H8
	//NoSquare coord const
	NoSquare
)

const (
	//Wkcastel White King Castel Perm
	Wkcastel = 1
	//Wqcastel White Queen Castel Perm
	Wqcastel = 2
	//Bkcastel Black King Castel Perm
	Bkcastel = 4
	//Bqcastel Black Queen Castel Perm
	Bqcastel = 8
)
