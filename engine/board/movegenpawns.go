package board

import (
	"errors"
	"fmt"
)

//addEnPasMove add an EnPass move
func (list *MoveListStruct) addEnPasMove(move int) {
	list.Moves[list.Count].Move = move
	list.Moves[list.Count].Score = 105 + 1000000
	list.Count++
}

//addWhitePawnCaptureMove Add capture move for white pawn
func (list *MoveListStruct) addWhitePawnCaptureMove(from, to, cap int, pos *PositionStruct) error {
	if DEBUG {
		if !squareOnBoard(from) || !squareOnBoard(to) || !pieceValidEmpty(cap) {
			return errors.New("Invalid perams for addWhitePawnCapture")
		}
	}

	if ranksBoard[from] == rank7 {
		list.addCaptureMove(ToMove(from, to, cap, wQ, 0), pos)
		list.addCaptureMove(ToMove(from, to, cap, wR, 0), pos)
		list.addCaptureMove(ToMove(from, to, cap, wB, 0), pos)
		list.addCaptureMove(ToMove(from, to, cap, wN, 0), pos)
	} else {
		list.addCaptureMove(ToMove(from, to, cap, empty, 0), pos)
	}
	return nil
}

//addWhitePawnMove Add normal white pawn move
func (list *MoveListStruct) addWhitePawnMove(from, to int, pos *PositionStruct) error {
	if DEBUG {
		if !squareOnBoard(from) || !squareOnBoard(to) {
			return errors.New("Invalid perams for addWhitePawnCapture")
		}
	}
	if ranksBoard[from] == rank7 {
		list.addQuietMove(ToMove(from, to, empty, wQ, 0), pos)
		list.addQuietMove(ToMove(from, to, empty, wR, 0), pos)
		list.addQuietMove(ToMove(from, to, empty, wB, 0), pos)
		list.addQuietMove(ToMove(from, to, empty, wN, 0), pos)
	} else {
		list.addQuietMove(ToMove(from, to, empty, empty, 0), pos)
	}
	return nil
}

//addBlackPawnCaptureMove Add capture move for black pawn
func (list *MoveListStruct) addBlackPawnCaptureMove(from, to, cap int, pos *PositionStruct) error {
	if DEBUG {
		if !squareOnBoard(from) || !squareOnBoard(to) || !pieceValidEmpty(cap) {
			return errors.New("Invalid perams for addWhitePawnCapture")
		}
	}

	if ranksBoard[from] == rank2 {
		list.addCaptureMove(ToMove(from, to, cap, bQ, 0), pos)
		list.addCaptureMove(ToMove(from, to, cap, bR, 0), pos)
		list.addCaptureMove(ToMove(from, to, cap, bB, 0), pos)
		list.addCaptureMove(ToMove(from, to, cap, bN, 0), pos)
	} else {
		list.addCaptureMove(ToMove(from, to, cap, empty, 0), pos)
	}

	return nil
}

//addBlackPawnMove add normal black pawn move
func (list *MoveListStruct) addBlackPawnMove(from, to int, pos *PositionStruct) error {
	if DEBUG {
		if !squareOnBoard(from) || !squareOnBoard(to) {
			return errors.New("Invalid perams for addWhitePawnCapture")
		}
	}
	if ranksBoard[from] == rank2 {
		list.addQuietMove(ToMove(from, to, empty, bQ, 0), pos)
		list.addQuietMove(ToMove(from, to, empty, bR, 0), pos)
		list.addQuietMove(ToMove(from, to, empty, bB, 0), pos)
		list.addQuietMove(ToMove(from, to, empty, bN, 0), pos)
	} else {
		list.addQuietMove(ToMove(from, to, empty, empty, 0), pos)
	}
	return nil
}

//generateAllPawnMoves Generate all pawn moves
func (pos *PositionStruct) generateAllPawnMoves(list *MoveListStruct) error {
	var err error = nil
	if pos.Side == white {
		//White pawn moves
		for pieceNum := 0; pieceNum < pos.PieceNum[wP]; pieceNum++ {
			sq := pos.PieceList[wP][pieceNum]

			if DEBUG && !squareOnBoard(sq) {
				return fmt.Errorf("Square: %d not on board", sq)
			}

			//Pawn move forward
			if pos.Pieces[sq+10] == empty {
				err = list.addWhitePawnMove(sq, sq+10, pos)
				if err != nil {
					return err
				}
				//Pawn move 2 forward
				if ranksBoard[sq] == rank2 && pos.Pieces[sq+20] == empty {
					list.addQuietMove(ToMove(sq, sq+20, empty, empty, MoveFlagPS), pos)
				}
			}

			//Pawn Captures
			if squareOnBoard(sq+9) && getPieceColor(pos.Pieces[sq+9]) == black {
				err = list.addWhitePawnCaptureMove(sq, sq+9, pos.Pieces[sq+9], pos)
				if err != nil {
					return err
				}
			}

			if squareOnBoard(sq+11) && getPieceColor(pos.Pieces[sq+11]) == black {
				err = list.addWhitePawnCaptureMove(sq, sq+11, pos.Pieces[sq+11], pos)
				if err != nil {
					return err
				}
			}
			if pos.EnPassant != noSquare {
				//Pawn EnPassant Capture
				if sq+9 == pos.EnPassant {
					list.addEnPasMove(ToMove(sq, sq+9, empty, empty, MoveFlagEP))
				}

				if sq+11 == pos.EnPassant {
					list.addEnPasMove(ToMove(sq, sq+11, empty, empty, MoveFlagEP))
				}
			}
		}
	} else {
		//Black pawn moves
		for pieceNum := 0; pieceNum < pos.PieceNum[bP]; pieceNum++ {
			sq := pos.PieceList[bP][pieceNum]

			if DEBUG && !squareOnBoard(sq) {
				return fmt.Errorf("Square: %d not on board", sq)
			}

			//Pawn move forward
			if pos.Pieces[sq-10] == empty {
				err = list.addBlackPawnMove(sq, sq-10, pos)
				if err != nil {
					return err
				}
				//Pawn move 2 forward
				if ranksBoard[sq] == rank7 && pos.Pieces[sq-20] == empty {
					list.addQuietMove(ToMove(sq, sq-20, empty, empty, MoveFlagPS), pos)
				}
			}

			//Pawn Captures
			if squareOnBoard(sq-9) && getPieceColor(pos.Pieces[sq-9]) == white {
				err = list.addBlackPawnCaptureMove(sq, sq-9, pos.Pieces[sq-9], pos)
				if err != nil {
					return err
				}
			}

			if squareOnBoard(sq-11) && getPieceColor(pos.Pieces[sq-11]) == white {
				err = list.addBlackPawnCaptureMove(sq, sq-11, pos.Pieces[sq-11], pos)
				if err != nil {
					return err
				}
			}
			if pos.EnPassant != noSquare {
				//Pawn EnPassant Capture
				if sq-9 == pos.EnPassant {
					list.addEnPasMove(ToMove(sq, sq-9, empty, empty, MoveFlagEP))
				}

				if sq-11 == pos.EnPassant {
					list.addEnPasMove(ToMove(sq, sq-11, empty, empty, MoveFlagEP))
				}
			}
		}
	}

	return err
}
