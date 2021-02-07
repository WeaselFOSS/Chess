package engine

import "fmt"

//takeMove Take back the last move
func (pos *BoardStruct) takeMove() error {

	pos.HisPly--
	pos.Ply--

	move := pos.History[pos.HisPly].Move
	from := getFrom(move)
	to := getTo(move)

	if DEBUG {
		err := pos.CheckBoard()
		if err != nil {
			return err
		}
		if !squareOnBoard(from) || !squareOnBoard(to) {
			return fmt.Errorf("From: %d or To: %d square for move is off the board", from, to)
		}
	}

	if pos.EnPassant != noSquare {
		pos.hashEnPas()
	}
	pos.hashCastel()

	pos.CastelPerm = pos.History[pos.HisPly].CastelPerm
	pos.FiftyMove = pos.History[pos.HisPly].FiftyMove
	pos.EnPassant = pos.History[pos.HisPly].EnPassant

	if pos.EnPassant != noSquare {
		pos.hashEnPas()
	}
	pos.hashCastel()

	//Flipping side to move
	pos.Side ^= 1
	pos.hashSide()

	if moveFlagEP&move != 0 {
		if pos.Side == white {
			pos.addPiece(to-10, bP)
		} else {
			pos.addPiece(to+10, wP)
		}
	} else if moveFlagCA&move != 0 {
		var err error
		switch to {
		case c1:
			err = pos.movePiece(d1, a1)
			break
		case c8:
			err = pos.movePiece(d8, a8)
			break
		case g1:
			err = pos.movePiece(f1, h1)
			break
		case g8:
			err = pos.movePiece(f8, h8)
			break
		default:
			return fmt.Errorf("Invalid castel move to %d", to)
		}
		if err != nil {
			return err
		}
	}

	err := pos.movePiece(to, from)
	if err != nil {
		return err
	}

	if pos.Pieces[from] == wK || pos.Pieces[from] == bK {
		pos.KingSquare[pos.Side] = from
	}

	captured := getCapture(move)
	if captured != empty {
		if DEBUG && !pieceValid(captured) {
			return fmt.Errorf("Invalid capture piece %d", captured)
		}
		err := pos.addPiece(to, captured)
		if err != nil {
			return err
		}
	}

	promotedPiece := getPromoted(move)
	if promotedPiece != empty {
		if !pieceValid(promotedPiece) || !isPieceBig(promotedPiece) {
			return fmt.Errorf("Invalid promotion piece of %d", promotedPiece)
		}
		err = pos.clearPiece(from)
		if err != nil {
			return err
		}
		if getPieceColor(promotedPiece) == white {
			err = pos.addPiece(from, wP)
		} else {
			err = pos.addPiece(from, bP)
		}
		if err != nil {
			return err
		}
	}

	if DEBUG {
		err = pos.CheckBoard()
		if err != nil {
			return err
		}
	}

	return nil
}
