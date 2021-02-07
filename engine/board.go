package engine

import (
	"errors"
	"fmt"
)

//BoardStruct the boards struct
type BoardStruct struct {
	Pieces      [squareNumber]int
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
}

//UndoStruct the undo move struct
type UndoStruct struct {
	Move       int
	CastelPerm int
	EnPassant  int
	FiftyMove  int
	PosKey     uint64
}

//sq120ToSq64 120 Square board to 64 square board index
var sq120ToSq64 [squareNumber]int

//sq64ToSq120 64 Square board to 64 square board index
var sq64ToSq120 [64]int

//filesBoard Get a positions file
var filesBoard [squareNumber]int

//ranksBoard Get a positions rank
var ranksBoard [squareNumber]int

var pieceChar = [13]rune{'.', 'P', 'N', 'B', 'R', 'Q', 'K', 'p', 'n', 'b', 'r', 'q', 'k'}
var sideChar = [3]rune{'w', 'b', '-'}

//fileRankToSquare takes a file and rank and returns a square number
func fileRankToSquare(file int, rank int) int {
	return 21 + file + rank*10
}

//resetBoard Reset the board
func (pos *BoardStruct) resetBoard() {
	for i := 0; i < squareNumber; i++ {
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

//LoadFEN loads the engine with a new board position from a FEN string
func (pos *BoardStruct) LoadFEN(fen string) error {
	if fen == "" {
		return errors.New("FEN String is empty")
	}

	rank := rank8
	file := fileA
	piece := 0
	count := 0

	pos.resetBoard()

	for (rank >= rank1) && len(fen) > 0 {
		count = 1

		switch fen[0] {
		case 'p':
			piece = bP
			break
		case 'n':
			piece = bN
			break
		case 'b':
			piece = bB
			break
		case 'r':
			piece = bR
			break
		case 'q':
			piece = bQ
			break
		case 'k':
			piece = bK
			break

		case 'P':
			piece = wP
			break
		case 'N':
			piece = wN
			break
		case 'B':
			piece = wB
			break
		case 'R':
			piece = wR
			break
		case 'Q':
			piece = wQ
			break
		case 'K':
			piece = wK
			break

		case '1', '2', '3', '4', '5', '6', '7', '8':
			piece = empty
			count = int(fen[0] - '0')
			break

		case '/', ' ':
			rank--
			file = fileA
			fen = fen[1:]
			continue

		default:
			return errors.New("Bad FEN string")
		}

		for i := 0; i < count; i++ {
			sq64 := rank*8 + file
			sq120 := sq64ToSq120[sq64]
			if piece != empty {
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
		pos.Side = white
	} else {
		pos.Side = black
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
			pos.CastelPerm |= wkcastel
			break
		case 'Q':
			pos.CastelPerm |= wqcastel
			break
		case 'k':
			pos.CastelPerm |= bkcastel
			break
		case 'q':
			pos.CastelPerm |= bqcastel
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

		if file < fileA || file > fileH {
			return errors.New("Bad FEN EnPas File")
		}

		if rank < rank1 || rank > rank8 {
			return errors.New("Bad FEN EnPas Rank")
		}

		pos.EnPassant = fileRankToSquare(file, rank)
	}
	//TODO: Add supprot for fifty move rule and current ply
	pos.updateMaterialLists()
	var err error
	pos.PosKey, err = pos.generatePosKey()
	return err
}

//Print a representation of the current board state to the console
func (pos *BoardStruct) Print() {
	fmt.Print("\nBoard State:\n\n")
	for rank := rank8; rank >= rank1; rank-- {
		fmt.Printf("%d", rank+1)
		for file := fileA; file <= fileH; file++ {
			sq := fileRankToSquare(file, rank)
			piece := pos.Pieces[sq]
			fmt.Printf("%3c", pieceChar[piece])
		}
		fmt.Print("\n")
	}

	fmt.Print(" ")
	for file := fileA; file <= fileH; file++ {
		fmt.Printf("%3c", 'A'+file)
	}
	fmt.Print("\n")
	fmt.Printf("Side: %c\n", sideChar[pos.Side])
	fmt.Printf("EnPassant: %s\n", SquareToString(pos.EnPassant))
	WK, WQ, BK, BQ := pos.castelPermToChar()
	fmt.Printf("Castel Perms: %c%c%c%c\n", WK, WQ, BK, BQ)
	fmt.Printf("Position Hash: %X\n", pos.PosKey)
}

//updateMaterialLists Update the material lists for the baord
func (pos *BoardStruct) updateMaterialLists() {
	for i := 0; i < squareNumber; i++ {
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

			pos.Material[color] += getPieceValue(piece)
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
