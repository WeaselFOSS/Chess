package engine

import "fmt"

var knightDir = [8]int{-8, -19, -21, -12, 8, 19, 21, 12}
var rookDir = [4]int{-1, -10, 1, 10}
var bishopDir = [4]int{-9, -11, 11, 9}
var kingDir = [8]int{-1, -10, 1, 10, -9, -11, 11, 9}

func (pos *BoardStruct) isAttacked(sq, side int) bool {
	if pos.PosKey == 0 {
		return false
	}
	//pawns
	if side == white {
		if pos.Pieces[sq-11] == wP || pos.Pieces[sq-9] == wP {
			return true
		}
	} else {
		if pos.Pieces[sq+11] == bP || pos.Pieces[sq+9] == bP {
			return true
		}
	}

	//knights
	for i := 0; i < 8; i++ {
		piece := pos.Pieces[sq+knightDir[i]]
		if (piece == wN && side == white) || (piece == bN && side == black) {
			return true
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
					return true
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
					return true
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
			return true
		}
	}

	return false
}

func (pos *BoardStruct) printSqAttacked(side int) {
	fmt.Printf("\nSquares attacked by: %c\n", sideChar[side])
	for rank := rank8; rank >= rank1; rank-- {
		for file := fileA; file <= fileH; file++ {
			sq := fileRankToSquare(file, rank)
			if pos.isAttacked(sq, side) {
				fmt.Print(" X ")
			} else {
				fmt.Print(" - ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
