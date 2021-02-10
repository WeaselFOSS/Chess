package board

import "errors"

//MakeNullMove Make a null move
//
//For more info go to https://www.chessprogramming.org/Null_Move_Pruning
func (pos *PositionStruct) MakeNullMove() error {
	if DEBUG {
		err := pos.CheckBoard()
		if err != nil {
			return err
		}

		var inCheck bool
		inCheck, err = pos.IsAttacked(pos.KingSquare[pos.Side], pos.Side^1)
		if err != nil {
			return err
		}

		if inCheck {
			return errors.New("Null move attempt whilst in check")
		}
	}

	pos.Ply++
	pos.History[pos.HisPly].PosKey = pos.PosKey
	if pos.EnPassant != noSquare {
		pos.hashEnPas()
	}

	pos.History[pos.HisPly].Move = NoMove
	pos.History[pos.HisPly].FiftyMove = pos.FiftyMove
	pos.History[pos.HisPly].EnPassant = pos.EnPassant
	pos.History[pos.HisPly].CastelPerm = pos.CastelPerm
	pos.EnPassant = noSquare

	//WHITE ^ 1 == BLACK & BLACK ^1 == WHITE
	pos.Side ^= 1
	pos.HisPly++
	pos.hashSide()

	if DEBUG {
		err := pos.CheckBoard()
		if err != nil {
			return err
		}
	}
	return nil
}

//TakeNullMove Take back the last null move
func (pos *PositionStruct) TakeNullMove() error {
	if DEBUG {
		err := pos.CheckBoard()
		if err != nil {
			return err
		}
	}

	pos.HisPly--
	pos.Ply--
	if pos.EnPassant != noSquare {
		pos.hashEnPas()
	}

	pos.CastelPerm = pos.History[pos.HisPly].CastelPerm
	pos.FiftyMove = pos.History[pos.HisPly].FiftyMove
	pos.EnPassant = pos.History[pos.HisPly].EnPassant

	if pos.EnPassant != noSquare {
		pos.hashEnPas()
	}

	//WHITE ^ 1 == BLACK & BLACK ^1 == WHITE
	pos.Side ^= 1
	pos.hashSide()

	return nil
}
