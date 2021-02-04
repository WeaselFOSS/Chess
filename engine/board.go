package engine

import (
	"errors"
	"fmt"
)

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

//PieceChar Piece const to cahr
var PieceChar = [13]rune{'.', 'P', 'N', 'B', 'R', 'Q', 'K', 'p', 'n', 'b', 'r', 'q', 'k'}

//SideChar Side const to char
var SideChar = [3]rune{'w', 'b', '-'}

//RankChar Rank cosnt to char
var RankChar = [8]rune{'1', '2', '3', '4', '5', '6', '7', '8'}

//FileChar File const to char
var FileChar = [8]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'}

//Sq120ToSq64 120 Square board to 64 square board index
var Sq120ToSq64 [SquareNumber]int

//Sq64ToSq120 64 Square board to 64 square board index
var Sq64ToSq120 [64]int

func initBoard() {
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

//FileRankToSquare takes a file and rank and returns a square number
func FileRankToSquare(file int, rank int) int {
	return 21 + file + rank*10
}

//ResetBoard Reset the board
func (pos *BoardStruct) ResetBoard() {
	for i := 0; i < SquareNumber; i++ {
		pos.Pieces[i] = OffBoard
	}

	for i := 0; i < 64; i++ {
		pos.Pieces[Sq64ToSq120[i]] = Empty
	}

	for i := 0; i < 3; i++ {
		pos.BigPieces[i] = 0
		pos.MajorPieces[i] = 0
		pos.MinorPieces[i] = 0
		pos.Pawns[i] = uint64(0)
	}

	for i := 0; i < 13; i++ {
		pos.PieceNum[i] = 0
	}

	pos.KingSquare[White] = NoSquare
	pos.KingSquare[Black] = NoSquare

	pos.Side = Both
	pos.EnPassant = NoSquare
	pos.FiftyMove = 0

	pos.Ply = 0
	pos.HisPly = 0

	pos.CastelPerm = 0

	pos.PosKey = uint64(0)
}

func (pos *BoardStruct) LoadFEN(fen string) error {
	if fen == "" {
		return errors.New("FEN String is empty")
	}

	rank := Rank8
	file := FileA
	piece := 0
	count := 0

	pos.ResetBoard()

	for (rank >= Rank1) && len(fen) > 0 {
		count = 1

		switch fen[0] {
		case 'p':
			piece = BP
			break
		case 'n':
			piece = BN
			break
		case 'b':
			piece = BB
			break
		case 'r':
			piece = BR
			break
		case 'q':
			piece = BQ
			break
		case 'k':
			piece = BK
			break

		case 'P':
			piece = WP
			break
		case 'N':
			piece = WN
			break
		case 'B':
			piece = WB
			break
		case 'R':
			piece = WR
			break
		case 'Q':
			piece = WQ
			break
		case 'K':
			piece = WK
			break

		case '1', '2', '3', '4', '5', '6', '7', '8':
			piece = Empty
			count = int(fen[0] - '0')
			break

		case '/', ' ':
			rank--
			file = FileA
			fen = fen[1:]
			continue

		default:
			return errors.New("Bad FEN string")
		}

		for i := 0; i < count; i++ {
			sq64 := rank*8 + file
			sq120 := Sq64ToSq120[sq64]
			if piece != Empty {
				pos.Pieces[sq120] = piece
			}
			file++
		}
		fen = fen[1:]
	}

	if fen[0] != 'w' && fen[0] != 'b' {
		return errors.New("Bad FEN Side To move")
	}

	if fen[0] == 'w' {
		pos.Side = White
	} else {
		pos.Side = Black
	}

	if len(fen) < 3 {
		return errors.New("Bad FEN Length")
	}

	fen = fen[2:]

	for i := 0; i < 4; i++ {
		if fen[0] == ' ' {
			break
		}
		switch fen[0] {
		case 'K':
			pos.CastelPerm |= Wkcastel
			break
		case 'Q':
			pos.CastelPerm |= Wqcastel
			break
		case 'k':
			pos.CastelPerm |= Bkcastel
			break
		case 'q':
			pos.CastelPerm |= Bqcastel
			break
		default:
			break
		}
		fen = fen[1:]
	}

	if len(fen) < 2 {
		return errors.New("Bad FEN Length")
	}

	fen = fen[1:]

	if fen[0] != '-' {
		file = int(fen[0] - 'a')
		rank = int(fen[1] - '1')

		if file < FileA || file > FileH {
			return errors.New("Bad FEN EnPas File")
		}

		if rank < Rank1 || rank > Rank8 {
			return errors.New("Bad FEN EnPas Rank")
		}

		pos.EnPassant = FileRankToSquare(file, rank)
	}

	return pos.GeneratePosKey()
}

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

//Print a representation of the current board state to the console
func (pos *BoardStruct) Print() {
	fmt.Print("\nBoard State:\n\n")
	for rank := Rank8; rank >= Rank1; rank-- {
		fmt.Printf("%d", rank+1)
		for file := FileA; file <= FileH; file++ {
			sq := FileRankToSquare(file, rank)
			piece := pos.Pieces[sq]
			fmt.Printf("%3c", PieceChar[piece])
		}
		fmt.Print("\n")
	}

	fmt.Print(" ")
	for file := FileA; file <= FileH; file++ {
		fmt.Printf("%3c", 'A'+file)
	}
	fmt.Print("\n")
	fmt.Printf("Side: %c\n", SideChar[pos.Side])
	fmt.Printf("EnPassant: %d\n", pos.EnPassant) //TODO: Create Decimal to algebraic notation function
	WK, WQ, BK, BQ := pos.castelPermToChar()
	fmt.Printf("Castel Perms: %c%c%c%c\n", WK, WQ, BK, BQ)
	fmt.Printf("Position Hash: %X\n", pos.PosKey)
}
