package board

import "fmt"

// GenerateAllMoves Generate all moves
func (pos *PositionStruct) GenerateAllCaptureMoves(list *MoveListStruct) error {
	if DEBUG {
		err := pos.CheckBoard()
		if err != nil {
			return err
		}
	}
	var err error

	if pos.Side == white {
		// White pawn moves
		for pieceNum := 0; pieceNum < pos.PieceNum[wP]; pieceNum++ {
			sq := pos.PieceList[wP][pieceNum]

			if DEBUG && !squareOnBoard(sq) {
				return fmt.Errorf("Square: %d not on board", sq)
			}

			// Pawn Captures
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
				// Pawn EnPassant Capture
				if sq+9 == pos.EnPassant {
					list.addEnPasMove(ToMove(sq, sq+9, empty, empty, MoveFlagEP))
				}

				if sq+11 == pos.EnPassant {
					list.addEnPasMove(ToMove(sq, sq+11, empty, empty, MoveFlagEP))
				}
			}
		}
	} else {
		// Black pawn moves
		for pieceNum := 0; pieceNum < pos.PieceNum[bP]; pieceNum++ {
			sq := pos.PieceList[bP][pieceNum]

			if DEBUG && !squareOnBoard(sq) {
				return fmt.Errorf("Square: %d not on board", sq)
			}

			// Pawn Captures
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
				// Pawn EnPassant Capture
				if sq-9 == pos.EnPassant {
					list.addEnPasMove(ToMove(sq, sq-9, empty, empty, MoveFlagEP))
				}

				if sq-11 == pos.EnPassant {
					list.addEnPasMove(ToMove(sq, sq-11, empty, empty, MoveFlagEP))
				}
			}
		}
	}
	// Slider moves
	pieceIndex := loopSlideIndex[pos.Side]
	piece := loopSlidePiece[pieceIndex]
	pieceIndex++
	for piece != 0 {
		if DEBUG && !pieceValid(piece) {
			return fmt.Errorf("Invalid piece %d", piece)
		}

		for pieceNum := 0; pieceNum < pos.PieceNum[piece]; pieceNum++ {
			sq := pos.PieceList[piece][pieceNum]
			if DEBUG && !squareOnBoard(sq) {
				return fmt.Errorf("Square not on board %d", sq)
			}
			for i := 0; i < pieceNumDir[piece]; i++ {
				dir := pieceDirection[piece][i]
				tSq := sq + dir

				for squareOnBoard(tSq) {
					if pos.Pieces[tSq] != empty {
						// BLACK ^ 1 == WHITe || WHITE ^ 1 == BLACK
						if getPieceColor(pos.Pieces[tSq]) == pos.Side^1 {
							list.addCaptureMove(ToMove(sq, tSq, pos.Pieces[tSq], empty, 0), pos)
						}
						break
					}
					tSq += dir
				}
			}
		}
		piece = loopSlidePiece[pieceIndex]
		pieceIndex++
	}

	// NonSlider moves
	pieceIndex = loopNonSlideIndex[pos.Side]
	piece = loopNonSlidePiece[pieceIndex]
	pieceIndex++
	for piece != 0 {
		if DEBUG && !pieceValid(piece) {
			return fmt.Errorf("Invalid piece %d", piece)
		}
		for pieceNum := 0; pieceNum < pos.PieceNum[piece]; pieceNum++ {
			sq := pos.PieceList[piece][pieceNum]
			if DEBUG && !squareOnBoard(sq) {
				return fmt.Errorf("Square not on board %d", sq)
			}
			for i := 0; i < pieceNumDir[piece]; i++ {
				dir := pieceDirection[piece][i]
				tSq := sq + dir

				if !squareOnBoard(tSq) {
					continue
				}

				if pos.Pieces[tSq] != empty {
					// BLACK ^ 1 == WHITe || WHITE ^ 1 == BLACK
					if getPieceColor(pos.Pieces[tSq]) == pos.Side^1 {
						list.addCaptureMove(ToMove(sq, tSq, pos.Pieces[tSq], empty, 0), pos)
					}
					continue
				}
			}
		}
		piece = loopNonSlidePiece[pieceIndex]
		pieceIndex++
	}

	return nil
}
