package board

import "fmt"

//MoveStruct used to store a move
type MoveStruct struct {
	Move  int
	Score int
}

//MoveListStruct List of moves
type MoveListStruct struct {
	Moves [maxPositionMoves]MoveStruct
	Count int
}

//castlePerm index of castel perm change per square
var castelPerm = [120]int{
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 13, 15, 15, 15, 12, 15, 15, 14, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 7, 15, 15, 15, 3, 15, 15, 11, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
}

/*
0000 0000 0000 0000 0000 0111 1111 -> From 0x7F
0000 0000 0000 0011 1111 1000 0000 -> To >> 7, 0x7F
0000 0000 0011 1100 0000 0000 0000 -> Captured >> 14, 0xF | 0x7C000
0000 0000 0100 0000 0000 0000 0000 -> EP 0x40000
0000 0000 1000 0000 0000 0000 0000 -> Pawn Start 0x80000
0000 1111 0000 0000 0000 0000 0000 -> Promoted Piece >> 20, 0xF | 0xF00000
0001 0000 0000 0000 0000 0000 0000 -> Castle 0x1000000
*/

//MoveFlagEP EnPas flag
var MoveFlagEP = 0x40000

//MoveFlagPS Pawn Start flag
var MoveFlagPS = 0x80000

//MoveFlagCA Castel flag
var MoveFlagCA = 0x1000000

//MoveFlagCAP Capture flag
var MoveFlagCAP = 0x7C000

//MoveFlagPROM Promotion flag
var MoveFlagPROM = 0xF00000

//GetFrom value from move int
func GetFrom(move int) int {
	return move & 0x7F
}

//GetTo get TO value from move int
func GetTo(move int) int {
	return (move >> 7) & 0x7F
}

//GetCapture get capture value from move int
func GetCapture(move int) int {
	return (move >> 14) & 0xF
}

//GetPromoted Get promote value from move int
func GetPromoted(move int) int {
	return (move >> 20) & 0xf
}

//ToMove Puts all move info into a single move int
func ToMove(from, to, capture, promotion, flag int) int {
	return (from | (to << 7) | (capture << 14) | (promotion << 20) | flag)
}

//clearPiece clear piece from current square
func (pos *PositionStruct) clearPiece(sq int) error {
	if DEBUG {
		if !squareOnBoard(sq) {
			return fmt.Errorf("Square %d not on board", sq)
		}
		if !pieceValid(pos.Pieces[sq]) {
			return fmt.Errorf("Piece on square %d is invalid with %d", sq, pos.Pieces[sq])
		}
	}

	piece := pos.Pieces[sq]
	color := getPieceColor(piece)

	pos.hashPiece(piece, sq)

	pos.Pieces[sq] = empty
	pos.Material[color] -= GetPieceValue(piece)

	if isPieceBig(piece) {
		pos.BigPieces[color]--
		if isPieceMajor(piece) {
			pos.MajorPieces[color]--
		} else {
			pos.MinorPieces[color]--
		}
	} else {
		clearBit(&pos.Pawns[color], sq120ToSq64[sq])
		clearBit(&pos.Pawns[both], sq120ToSq64[sq])
	}

	tPieceNum := -1
	for i := 0; i < pos.PieceNum[piece]; i++ {
		if pos.PieceList[piece][i] == sq {
			tPieceNum = i
			break
		}
	}

	if DEBUG && tPieceNum == -1 {
		return fmt.Errorf("Could not find piece %d on square %d in piece list", piece, sq)
	}

	pos.PieceNum[piece]--
	pos.PieceList[piece][tPieceNum] = pos.PieceList[piece][pos.PieceNum[piece]]
	return nil
}

func (pos *PositionStruct) addPiece(sq, piece int) error {
	if DEBUG {
		if !squareOnBoard(sq) {
			return fmt.Errorf("Square %d not on board", sq)
		}
		if !pieceValid(piece) {
			return fmt.Errorf("Piece value %d invalid", piece)
		}
	}

	color := getPieceColor(piece)

	pos.hashPiece(piece, sq)
	pos.Pieces[sq] = piece

	if isPieceBig(piece) {
		pos.BigPieces[color]++
		if isPieceMajor(piece) {
			pos.MajorPieces[color]++
		} else {
			pos.MinorPieces[color]++
		}
	} else {
		setBit(&pos.Pawns[color], sq120ToSq64[sq])
		setBit(&pos.Pawns[both], sq120ToSq64[sq])
	}

	pos.Material[color] += GetPieceValue(piece)
	pos.PieceList[piece][pos.PieceNum[piece]] = sq
	pos.PieceNum[piece]++

	return nil
}

//movePiece Move a piece
func (pos *PositionStruct) movePiece(from, to int) error {
	if DEBUG {
		if !squareOnBoard(from) {
			return fmt.Errorf("from value Square %d not on board", from)
		}
		if !squareOnBoard(to) {
			return fmt.Errorf("to value Square %d not on board", to)
		}
	}

	piece := pos.Pieces[from]
	color := getPieceColor(piece)
	//Value only used in debug mode
	pieceFound := false

	pos.hashPiece(piece, from)
	pos.Pieces[from] = empty

	pos.hashPiece(piece, to)
	pos.Pieces[to] = piece

	if !isPieceBig(piece) {
		clearBit(&pos.Pawns[color], sq120ToSq64[from])
		clearBit(&pos.Pawns[both], sq120ToSq64[from])

		setBit(&pos.Pawns[color], sq120ToSq64[to])
		setBit(&pos.Pawns[both], sq120ToSq64[to])
	}

	for i := 0; i < pos.PieceNum[piece]; i++ {
		if pos.PieceList[piece][i] == from {
			pos.PieceList[piece][i] = to
			pieceFound = true
			break
		}
	}

	if DEBUG && !pieceFound {
		return fmt.Errorf("Could not find piece %d in PieceList", piece)
	}

	return nil
}

func (pos *PositionStruct) MoveExists(move int) (bool, error) {
	var list MoveListStruct
	err := pos.GenerateAllMoves(&list)
	if err != nil {
		return false, err
	}

	for i := 0; i < list.Count; i++ {
		var makeMove bool
		makeMove, err = pos.MakeMove(list.Moves[i].Move)
		if err != nil {
			return false, err
		}

		if !makeMove {
			continue
		}
		err = pos.TakeMove()
		if err != nil {
			return false, err
		}
		if list.Moves[i].Move == move {
			return true, nil
		}
	}
	return false, nil
}
