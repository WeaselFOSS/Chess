package engine

import (
	"fmt"
)

//loopSlidePiece array used to loop through all sliding pieces of one color
var loopSlidePiece = [8]int{wB, wR, wQ, 0, bB, bR, bQ, 0}

//loopSlidePiece array used to loop through all non sliding pieces of one color
var loopNonSlidePiece = [8]int{wN, wK, 0, bN, bK}

//loopSlideIndex Side to loop index for sliding pieces
var loopSlideIndex = [2]int{0, 4}

//loopNonSlideIndex Side to loop index for non sliding pieces
var loopNonSlideIndex = [2]int{0, 3}

//pieceDirection Get directions a piece can move
var pieceDirection = [13][8]int{
	{0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
	{-8, -19, -21, -12, 8, 19, 21, 12},
	{-9, -11, 11, 9, 0, 0, 0, 0},
	{-1, -10, 1, 10, 0, 0, 0, 0},
	{-1, -10, 1, 10, -9, -11, 11, 9},
	{-1, -10, 1, 10, -9, -11, 11, 9},
	{0, 0, 0, 0, 0, 0, 0},
	{-8, -19, -21, -12, 8, 19, 21, 12},
	{-9, -11, 11, 9, 0, 0, 0, 0},
	{-1, -10, 1, 10, 0, 0, 0, 0},
	{-1, -10, 1, 10, -9, -11, 11, 9},
	{-1, -10, 1, 10, -9, -11, 11, 9},
}

//pieceNumDir get the number of directions for a piece
var pieceNumDir = [13]int{
	0, 0, 8, 4, 4, 8, 8, 0, 8, 4, 4, 8, 8,
}

//addQuietMove Add a normal non capture mofe
func (list *MoveListStruct) addQuietMove(move int) {
	list.Moves[list.Count].Move = move
	list.Moves[list.Count].Score = 0
	list.Count++
}

//addCaptureMove add a capture move
func (list *MoveListStruct) addCaptureMove(move int) {
	list.Moves[list.Count].Move = move
	list.Moves[list.Count].Score = 0
	list.Count++
}

//GenerateAllMoves Generate all moves
func (pos *BoardStruct) GenerateAllMoves(list *MoveListStruct) error {
	if DEBUG {
		err := pos.CheckBoard()
		if err != nil {
			return err
		}
	}

	list.Count = 0

	//Pawn moves
	err := pos.generateAllPawnMoves(list)
	if err != nil {
		return err
	}

	//Castling
	if pos.Side == white {
		//White king side castel
		if pos.CastelPerm&wkcastel != 0 {
			if pos.Pieces[f1] == empty && pos.Pieces[g1] == empty {
				e1A, err := pos.isAttacked(e1, black)
				if err != nil {
					return err
				}
				var f1A bool
				f1A, err = pos.isAttacked(f1, black)
				if err != nil {
					return err
				}

				if !e1A && !f1A {
					list.addQuietMove(toMove(e1, g1, empty, empty, moveFlagCA))
				}
			}
		}

		//White queen side castel
		if pos.CastelPerm&wqcastel != 0 {
			if pos.Pieces[d1] == empty && pos.Pieces[c1] == empty && pos.Pieces[b1] == empty {
				e1A, err := pos.isAttacked(e1, black)
				if err != nil {
					return err
				}
				var d1A bool
				d1A, err = pos.isAttacked(d1, black)
				if err != nil {
					return err
				}

				if !e1A && !d1A {
					list.addQuietMove(toMove(e1, c1, empty, empty, moveFlagCA))
				}
			}
		}
	} else {
		//Black king side castel
		if pos.CastelPerm&bkcastel != 0 {
			if pos.Pieces[f8] == empty && pos.Pieces[g8] == empty {
				e8A, err := pos.isAttacked(e8, white)
				if err != nil {
					return err
				}
				var f8A bool
				f8A, err = pos.isAttacked(f8, white)
				if err != nil {
					return err
				}

				if !e8A && !f8A {
					list.addQuietMove(toMove(e8, g8, empty, empty, moveFlagCA))
				}
			}
		}

		//Black queen side castel
		if pos.CastelPerm&bqcastel != 0 {
			if pos.Pieces[d8] == empty && pos.Pieces[c8] == empty && pos.Pieces[b8] == empty {
				e8A, err := pos.isAttacked(e8, white)
				if err != nil {
					return err
				}
				var d8A bool
				d8A, err = pos.isAttacked(d8, white)
				if err != nil {
					return err
				}

				if !e8A && !d8A {
					list.addQuietMove(toMove(e8, c8, empty, empty, moveFlagCA))
				}
			}
		}
	}

	//Slider moves
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
							list.addCaptureMove(toMove(sq, tSq, pos.Pieces[tSq], empty, 0))
						}
						break
					}
					list.addQuietMove(toMove(sq, tSq, empty, empty, 0))
					tSq += dir
				}
			}
		}
		piece = loopSlidePiece[pieceIndex]
		pieceIndex++
	}

	//NonSlider moves
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
						list.addCaptureMove(toMove(sq, tSq, pos.Pieces[tSq], empty, 0))
					}
					continue
				}
				list.addQuietMove(toMove(sq, tSq, empty, empty, 0))
			}
		}
		piece = loopNonSlidePiece[pieceIndex]
		pieceIndex++
	}

	return nil
}
