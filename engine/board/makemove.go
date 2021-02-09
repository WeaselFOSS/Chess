package board

import (
	"errors"
	"fmt"
)

//MakeMove Make a move if legal and return legality
func (pos *PositionStruct) MakeMove(move int) (bool, error) {

	from := getFrom(move)
	to := getTo(move)
	side := pos.Side

	if DEBUG {
		err := pos.CheckBoard()
		if err != nil {
			return false, err
		}
		if !squareOnBoard(from) || !squareOnBoard(to) {
			return false, fmt.Errorf("From: %d or To: %d square for move is off the board", from, to)
		}
		if !sideValid(side) {
			return false, errors.New("Side is invalid")
		}
		if !pieceValid(pos.Pieces[from]) {
			return false, fmt.Errorf("Piece is invalid with value of %d", pos.Pieces[from])
		}
	}

	pos.History[pos.HisPly].PosKey = pos.PosKey

	//EnPas moves
	if move&moveFlagEP != 0 {
		var err error
		if side == white {
			err = pos.clearPiece(to - 10)
		} else {
			err = pos.clearPiece(to + 10)
		}
		if err != nil {
			return false, err
		}
	} else if move&moveFlagCA != 0 { //castel moves
		var err error
		switch to {
		case c1:
			err = pos.movePiece(a1, d1)
			break
		case c8:
			err = pos.movePiece(a8, d8)
			break
		case g1:
			err = pos.movePiece(h1, f1)
			break
		case g8:
			err = pos.movePiece(h8, f8)
			break
		default:
			return false, fmt.Errorf("Invalid castel move to %d", to)
		}
		if err != nil {
			return false, err
		}
	}

	if pos.EnPassant != noSquare {
		pos.hashEnPas()
	}
	pos.hashCastel()
	pos.History[pos.HisPly].Move = move
	pos.History[pos.HisPly].FiftyMove = pos.FiftyMove
	pos.History[pos.HisPly].EnPassant = pos.EnPassant
	pos.History[pos.HisPly].CastelPerm = pos.CastelPerm

	//update castel perms
	pos.CastelPerm &= castelPerm[from]
	pos.CastelPerm &= castelPerm[to]
	pos.hashCastel()

	pos.EnPassant = noSquare

	captured := getCapture(move)
	pos.FiftyMove++

	if captured != empty {
		if DEBUG && !pieceValid(captured) {
			return false, fmt.Errorf("Invalid capture piece %d", captured)
		}
		err := pos.clearPiece(to)
		if err != nil {
			return false, err
		}
		pos.FiftyMove = 0
	}

	pos.HisPly++
	pos.Ply++

	if !isPieceBig(pos.Pieces[from]) {
		pos.FiftyMove = 0
		if move&moveFlagEP != 0 {
			if side == white {
				pos.EnPassant = from + 10
				if ranksBoard[pos.EnPassant] != rank3 {
					return false, fmt.Errorf("Invalid enPas rank of %d", ranksBoard[pos.EnPassant])
				}
			} else {
				pos.EnPassant = from - 10
				if ranksBoard[pos.EnPassant] != rank6 {
					return false, fmt.Errorf("Invalid enPas rank of %d", ranksBoard[pos.EnPassant])
				}
			}
			pos.hashEnPas()
		}
	}

	err := pos.movePiece(from, to)
	if err != nil {
		return false, err
	}

	promotedPiece := getPromoted(move)
	if promotedPiece != empty {
		if !pieceValid(promotedPiece) || !isPieceBig(promotedPiece) {
			return false, fmt.Errorf("Invalid promotion piece of %d", promotedPiece)
		}
		err = pos.clearPiece(to)
		if err != nil {
			return false, err
		}
		err = pos.addPiece(to, promotedPiece)
		if err != nil {
			return false, err
		}
	}

	if pos.Pieces[to] == wK || pos.Pieces[to] == bK {
		pos.KingSquare[pos.Side] = to
	}

	//flip side to move
	pos.Side ^= 1
	pos.hashSide()

	if DEBUG {
		err = pos.CheckBoard()
		if err != nil {
			return false, err
		}
	}

	var isKAttacked bool
	isKAttacked, err = pos.IsAttacked(pos.KingSquare[side], pos.Side)
	if err != nil {
		return false, err
	}
	if isKAttacked {
		err = pos.TakeMove()
		if err != nil {
			return false, err
		}
		return false, nil
	}

	return true, nil
}
