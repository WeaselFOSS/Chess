package engine

import "fmt"

var knightDir = [8]int{-8, -19, -21, -12, 8, 19, 21, 12}
var rookDir = [4]int{-1, -10, 1, 10}
var bishopDir = [4]int{-9, -11, 11, 9}
var kingDir = [8]int{-1, -10, 1, 10, -9, -11, 11, 9}

func (pos *BoardStruct) isAttacked(sq, side int) (bool, error) {

	if DEBUG && !squareOnBoard(sq) {
		return false, fmt.Errorf("Square not on board %s", SquareToString(sq))
	}

	if DEBUG && !sideValid(side) {
		return false, fmt.Errorf("Invalid side %d", side)
	}

	if DEBUG {
		err := pos.CheckBoard()
		if err != nil {
			return false, err
		}
	}

	//pawns
	if side == white {
		if pos.Pieces[sq-11] == wP || pos.Pieces[sq-9] == wP {
			return true, nil
		}
	} else {
		if pos.Pieces[sq+11] == bP || pos.Pieces[sq+9] == bP {
			return true, nil
		}
	}

	//knights
	for i := 0; i < 8; i++ {
		piece := pos.Pieces[sq+knightDir[i]]
		if (piece == wN && side == white) || (piece == bN && side == black) {
			return true, nil
		}
	}

	//rooks and queens
	for i := 0; i < 4; i++ {
		dir := rookDir[i]
		tSq := sq + dir
		piece := pos.Pieces[tSq]
		for piece != offBoard {
			if piece != empty {
				if ((piece == wR || piece == wQ) && side == white) ||
					((piece == bR || piece == bQ) && side == black) {
					return true, nil
				}
				break
			}
			tSq += dir
			piece = pos.Pieces[tSq]
		}
	}

	//bishops and queens
	for i := 0; i < 4; i++ {
		dir := bishopDir[i]
		tSq := sq + dir
		piece := pos.Pieces[tSq]
		for piece != offBoard {
			if piece != empty {
				if ((piece == wB || piece == wQ) && side == white) ||
					((piece == bB || piece == bQ) && side == black) {
					return true, nil
				}
				break
			}
			tSq += dir
			piece = pos.Pieces[tSq]
		}

	}

	//kings
	for i := 0; i < 8; i++ {
		piece := pos.Pieces[sq+kingDir[i]]
		if (piece == wK && side == white) || (piece == bK && side == black) {
			return true, nil
		}
	}

	return false, nil
}

func (pos *BoardStruct) printSqAttacked(side int) error {
	fmt.Printf("\nSquares attacked by: %c\n", sideChar[side])
	for rank := rank8; rank >= rank1; rank-- {
		for file := fileA; file <= fileH; file++ {
			sq := fileRankToSquare(file, rank)
			isAttacked, err := pos.isAttacked(sq, side)
			if err != nil {
				return err
			}
			if isAttacked {
				fmt.Print(" X ")
			} else {
				fmt.Print(" - ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
	return nil
}
