package engine

import "fmt"

func (pos *BoardStruct) addQuietMove(move int, list *MoveListStruct) {
	list.Moves[list.Count].Move = move
	list.Moves[list.Count].Score = 0
	list.Count++
}

func (pos *BoardStruct) addCaptureMove(move int, list *MoveListStruct) {
	list.Moves[list.Count].Move = move
	list.Moves[list.Count].Score = 0
	list.Count++
}

func (pos *BoardStruct) addEnPasMove(move int, list *MoveListStruct) {
	list.Moves[list.Count].Move = move
	list.Moves[list.Count].Score = 0
	list.Count++
}

func (pos *BoardStruct) addWhitePawnCaptureMove(from, to, cap int, list *MoveListStruct) {
	if ranksBoard[from] == rank7 {
		pos.addCaptureMove(toMove(from, to, cap, wQ, 0), list)
		pos.addCaptureMove(toMove(from, to, cap, wR, 0), list)
		pos.addCaptureMove(toMove(from, to, cap, wB, 0), list)
		pos.addCaptureMove(toMove(from, to, cap, wN, 0), list)
	} else {
		pos.addCaptureMove(toMove(from, to, cap, empty, 0), list)
	}
}

func (pos *BoardStruct) addWhitePawnMove(from, to int, list *MoveListStruct) {
	if ranksBoard[from] == rank7 {
		pos.addCaptureMove(toMove(from, to, empty, wQ, 0), list)
		pos.addCaptureMove(toMove(from, to, empty, wR, 0), list)
		pos.addCaptureMove(toMove(from, to, empty, wB, 0), list)
		pos.addCaptureMove(toMove(from, to, empty, wN, 0), list)
	} else {
		pos.addCaptureMove(toMove(from, to, empty, empty, 0), list)
	}
}

func (pos *BoardStruct) GenerateAllMoves(list *MoveListStruct) error {
	if DEBUG {
		err := pos.CheckBoard()
		if err != nil {
			return err
		}
	}

	list.Count = 0

	if pos.Side == white {
		//White pawn moves
		for pieceNum := 0; pieceNum < pos.PieceNum[wP]; pieceNum++ {
			sq := pos.PieceList[wP][pieceNum]

			if DEBUG && !squareOnBoard(sq) {
				return fmt.Errorf("Square: %d not on board", sq)
			}

			//Pawn move forward
			if pos.Pieces[sq+10] == empty {
				pos.addWhitePawnMove(sq, sq+10, list)
				//Pawn move 2 forward
				if ranksBoard[sq] == rank2 && pos.Pieces[sq+20] == empty {
					pos.addQuietMove(toMove(sq, sq+20, empty, empty, moveFlagPS), list)
				}
			}

			//Pawn Captures
			if squareOnBoard(sq+9) && getPieceColor(pos.Pieces[sq+9]) == black {
				pos.addWhitePawnCaptureMove(sq, sq+9, pos.Pieces[sq+9], list)
			}

			if squareOnBoard(sq+11) && getPieceColor(pos.Pieces[sq+11]) == black {
				pos.addWhitePawnCaptureMove(sq, sq+11, pos.Pieces[sq+11], list)
			}

			//Pawn EnPassant Capture
			if sq+9 == pos.EnPassant {
				pos.addCaptureMove(toMove(sq, sq+9, empty, empty, moveFlagEP), list)
			}

			if sq+11 == pos.EnPassant {
				pos.addCaptureMove(toMove(sq, sq+11, empty, empty, moveFlagEP), list)
			}
		}
	}

	return nil
}
