package engine

import "fmt"

//addEnPasMove add an EnPass move
func (list *MoveListStruct) addEnPasMove(move int) {
	list.Moves[list.Count].Move = move
	list.Moves[list.Count].Score = 0
	list.Count++
}

//addWhitePawnCaptureMove Add capture move for white pawn
func (list *MoveListStruct) addWhitePawnCaptureMove(from, to, cap int) {
	if ranksBoard[from] == rank7 {
		list.addCaptureMove(toMove(from, to, cap, wQ, 0))
		list.addCaptureMove(toMove(from, to, cap, wR, 0))
		list.addCaptureMove(toMove(from, to, cap, wB, 0))
		list.addCaptureMove(toMove(from, to, cap, wN, 0))
	} else {
		list.addCaptureMove(toMove(from, to, cap, empty, 0))
	}
}

//addWhitePawnMove Add normal white pawn move
func (list *MoveListStruct) addWhitePawnMove(from, to int) {
	if ranksBoard[from] == rank7 {
		list.addCaptureMove(toMove(from, to, empty, wQ, 0))
		list.addCaptureMove(toMove(from, to, empty, wR, 0))
		list.addCaptureMove(toMove(from, to, empty, wB, 0))
		list.addCaptureMove(toMove(from, to, empty, wN, 0))
	} else {
		list.addCaptureMove(toMove(from, to, empty, empty, 0))
	}
}

//addBlackPawnCaptureMove Add capture move for black pawn
func (list *MoveListStruct) addBlackPawnCaptureMove(from, to, cap int) {
	if ranksBoard[from] == rank2 {
		list.addCaptureMove(toMove(from, to, cap, bQ, 0))
		list.addCaptureMove(toMove(from, to, cap, bR, 0))
		list.addCaptureMove(toMove(from, to, cap, bB, 0))
		list.addCaptureMove(toMove(from, to, cap, bN, 0))
	} else {
		list.addCaptureMove(toMove(from, to, cap, empty, 0))
	}
}

//addBlackPawnMove add normal black pawn move
func (list *MoveListStruct) addBlackPawnMove(from, to int) {
	if ranksBoard[from] == rank2 {
		list.addCaptureMove(toMove(from, to, empty, bQ, 0))
		list.addCaptureMove(toMove(from, to, empty, bR, 0))
		list.addCaptureMove(toMove(from, to, empty, bB, 0))
		list.addCaptureMove(toMove(from, to, empty, bN, 0))
	} else {
		list.addCaptureMove(toMove(from, to, empty, empty, 0))
	}
}

//generateAllPawnMoves Generate all pawn moves
func (pos *BoardStruct) generateAllPawnMoves(list *MoveListStruct) error {
	if pos.Side == white {
		//White pawn moves
		for pieceNum := 0; pieceNum < pos.PieceNum[wP]; pieceNum++ {
			sq := pos.PieceList[wP][pieceNum]

			if DEBUG && !squareOnBoard(sq) {
				return fmt.Errorf("Square: %d not on board", sq)
			}

			//Pawn move forward
			if pos.Pieces[sq+10] == empty {
				list.addWhitePawnMove(sq, sq+10)
				//Pawn move 2 forward
				if ranksBoard[sq] == rank2 && pos.Pieces[sq+20] == empty {
					list.addQuietMove(toMove(sq, sq+20, empty, empty, moveFlagPS))
				}
			}

			//Pawn Captures
			if squareOnBoard(sq+9) && getPieceColor(pos.Pieces[sq+9]) == black {
				list.addWhitePawnCaptureMove(sq, sq+9, pos.Pieces[sq+9])
			}

			if squareOnBoard(sq+11) && getPieceColor(pos.Pieces[sq+11]) == black {
				list.addWhitePawnCaptureMove(sq, sq+11, pos.Pieces[sq+11])
			}

			//Pawn EnPassant Capture
			if sq+9 == pos.EnPassant {
				list.addCaptureMove(toMove(sq, sq+9, empty, empty, moveFlagEP))
			}

			if sq+11 == pos.EnPassant {
				list.addCaptureMove(toMove(sq, sq+11, empty, empty, moveFlagEP))
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
				list.addBlackPawnMove(sq, sq-10)
				//Pawn move 2 forward
				if ranksBoard[sq] == rank7 && pos.Pieces[sq-20] == empty {
					list.addQuietMove(toMove(sq, sq-20, empty, empty, moveFlagPS))
				}
			}

			//Pawn Captures
			if squareOnBoard(sq-9) && getPieceColor(pos.Pieces[sq-9]) == white {
				list.addBlackPawnCaptureMove(sq, sq-9, pos.Pieces[sq-9])
			}

			if squareOnBoard(sq-11) && getPieceColor(pos.Pieces[sq-11]) == white {
				list.addBlackPawnCaptureMove(sq, sq-11, pos.Pieces[sq-11])
			}

			//Pawn EnPassant Capture
			if sq-9 == pos.EnPassant {
				list.addCaptureMove(toMove(sq, sq-9, empty, empty, moveFlagEP))
			}

			if sq-11 == pos.EnPassant {
				list.addCaptureMove(toMove(sq, sq-11, empty, empty, moveFlagEP))
			}
		}
	}

	return nil
}
