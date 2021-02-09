package board

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

var victimScore = [13]int{0, 100, 200, 300, 400, 500, 600, 100, 200, 300, 400, 500, 600}
var mvvLvaScores [13][13]int

//InitMvvLva MvvLva (Most valuable victim least valuable attacker)
func InitMvvLva() {
	for attacker := wP; attacker <= bK; attacker++ {
		for victim := wP; victim <= bK; victim++ {
			mvvLvaScores[victim][attacker] = victimScore[victim] + 6 - (victimScore[attacker] / 100)
		}
	}
}

//addQuietMove Add a normal non capture mofe
func (list *MoveListStruct) addQuietMove(move int, pos *PositionStruct) {
	list.Moves[list.Count].Move = move

	if pos.Ply <= 63 && pos.SearchKillers[0][pos.Ply] == move {
		list.Moves[list.Count].Score = 900000
	} else if pos.Ply <= 63 && pos.SearchKillers[1][pos.Ply] == move {
		list.Moves[list.Count].Score = 800000
	} else {
		list.Moves[list.Count].Score = pos.SearchHistory[pos.Pieces[GetFrom(move)]][GetTo(move)]
	}
	list.Count++
}

//addCaptureMove add a capture move
func (list *MoveListStruct) addCaptureMove(move int, pos *PositionStruct) {
	list.Moves[list.Count].Move = move
	list.Moves[list.Count].Score = mvvLvaScores[GetCapture(move)][pos.Pieces[GetFrom(move)]] + 1000000
	list.Count++
}

//GenerateAllMoves Generate all moves
func (pos *PositionStruct) GenerateAllMoves(list *MoveListStruct) error {
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
				e1A, err := pos.IsAttacked(e1, black)
				if err != nil {
					return err
				}
				var f1A bool
				f1A, err = pos.IsAttacked(f1, black)
				if err != nil {
					return err
				}

				if !e1A && !f1A {
					list.addQuietMove(ToMove(e1, g1, empty, empty, MoveFlagCA), pos)
				}
			}
		}

		//White queen side castel
		if pos.CastelPerm&wqcastel != 0 {
			if pos.Pieces[d1] == empty && pos.Pieces[c1] == empty && pos.Pieces[b1] == empty {
				e1A, err := pos.IsAttacked(e1, black)
				if err != nil {
					return err
				}
				var d1A bool
				d1A, err = pos.IsAttacked(d1, black)
				if err != nil {
					return err
				}

				if !e1A && !d1A {
					list.addQuietMove(ToMove(e1, c1, empty, empty, MoveFlagCA), pos)
				}
			}
		}
	} else {
		//Black king side castel
		if pos.CastelPerm&bkcastel != 0 {
			if pos.Pieces[f8] == empty && pos.Pieces[g8] == empty {
				e8A, err := pos.IsAttacked(e8, white)
				if err != nil {
					return err
				}
				var f8A bool
				f8A, err = pos.IsAttacked(f8, white)
				if err != nil {
					return err
				}

				if !e8A && !f8A {
					list.addQuietMove(ToMove(e8, g8, empty, empty, MoveFlagCA), pos)
				}
			}
		}

		//Black queen side castel
		if pos.CastelPerm&bqcastel != 0 {
			if pos.Pieces[d8] == empty && pos.Pieces[c8] == empty && pos.Pieces[b8] == empty {
				e8A, err := pos.IsAttacked(e8, white)
				if err != nil {
					return err
				}
				var d8A bool
				d8A, err = pos.IsAttacked(d8, white)
				if err != nil {
					return err
				}

				if !e8A && !d8A {
					list.addQuietMove(ToMove(e8, c8, empty, empty, MoveFlagCA), pos)
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
							list.addCaptureMove(ToMove(sq, tSq, pos.Pieces[tSq], empty, 0), pos)
						}
						break
					}
					list.addQuietMove(ToMove(sq, tSq, empty, empty, 0), pos)
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
						list.addCaptureMove(ToMove(sq, tSq, pos.Pieces[tSq], empty, 0), pos)
					}
					continue
				}
				list.addQuietMove(ToMove(sq, tSq, empty, empty, 0), pos)
			}
		}
		piece = loopNonSlidePiece[pieceIndex]
		pieceIndex++
	}

	return nil
}
