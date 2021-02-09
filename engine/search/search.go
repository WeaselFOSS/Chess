package search

import (
	"fmt"
	"time"

	"github.com/WeaselChess/Weasel/engine/board"
)

//InfoStruct The search info struct
type InfoStruct struct {
	StartTime int64
	StopTiem  int64
	Depth     int
	DepthSet  bool
	TimeSet   bool
	MovesToGo bool
	Infinite  bool
	Nodes     int64
	Quit      bool
	Stopped   bool

	//Used to calculate the efficency of our move ordering
	FailHigh      float32
	FailHighFirst float32
}

//infinitie Largest score value
const infinitie = 30000
const mate = 29000

//checkUp Check if time up, or interrupt from GUI
func checkUp() {

}

//clearForSearch Reset serach info and PVTables to get ready for a enw search
func (info *InfoStruct) clearForSearch(pos *board.PositionStruct) {
	for x := 0; x < 13; x++ {
		for y := 0; y < board.SquareNumber; y++ {
			pos.SearchHistory[x][y] = 0
		}
	}

	for x := 0; x < 2; x++ {
		for y := 0; y < board.MaxDepth; y++ {
			pos.SearchKillers[x][y] = 0
		}
	}

	pos.PVTable.Clear()
	pos.Ply = 0

	info.StartTime = time.Now().UnixNano() / int64(time.Millisecond)
	info.Stopped = false
	info.Nodes = 0
	info.FailHigh = 0
	info.FailHighFirst = 0
}

//quiescence Search capture moves until a "quite" position is found or the trade has been resolved
//
//Used to counter the horizon effect
func (info *InfoStruct) quiescence(alpha, beta, pos *board.PositionStruct) int {
	return 0
}

//alphaBeta Normal alphabeta searching
func (info *InfoStruct) alphaBeta(alpha, beta, depth int, doNull bool, pos *board.PositionStruct) (int, error) {

	if board.DEBUG {
		err := pos.CheckBoard()
		if err != nil {
			return 0, err
		}
	}

	if depth == 0 {
		info.Nodes++
		return pos.Evaluate(), nil
	}

	info.Nodes++

	//Check for repition and fifty move role
	if pos.IsRepition() || pos.FiftyMove >= 100 {
		return 0, nil
	}

	//if we are at our max depth return the eval
	if pos.Ply > board.MaxDepth-1 {
		return pos.Evaluate(), nil
	}

	var list board.MoveListStruct
	err := pos.GenerateAllMoves(&list)
	if err != nil {
		return 0, err
	}

	legal := 0
	oldAlpha := alpha
	bestMove := board.NoMove
	score := -infinitie

	//Move loop
	for i := 0; i < list.Count; i++ {
		moveMade := false
		moveMade, err = pos.MakeMove(list.Moves[i].Move)
		if err != nil {
			return 0, err
		}

		if !moveMade {
			continue
		}

		legal++
		score, err = info.alphaBeta(-beta, -alpha, depth-1, true, pos)
		//Flipping score for the other sides POV
		score *= -1
		err = pos.TakeMove()
		if err != nil {
			return 0, err
		}

		//If score beats alpha set alpha to score
		if score > alpha {
			//If score is better than our beta cutoff return beta
			if score >= beta {
				if legal == 1 {
					info.FailHighFirst++
				}
				info.FailHigh++
				return beta, nil
			}

			alpha = score
			bestMove = list.Moves[i].Move
		}
	}

	if legal == 0 {
		//pos.Side^1 WHITE^1 == black and BLACK^1 == white
		isAttacked := false
		isAttacked, err = pos.IsAttacked(pos.KingSquare[pos.Side], pos.Side^1)
		if err != nil {
			return 0, err
		}

		if isAttacked {
			//Return -mate plus the ply or moves to the mate so later we can take score and subtrace mate to get mate in X val
			return -mate + pos.Ply, nil
		} else {
			return 0, nil
		}
	}

	//If we found a better move store it in the PV table
	if alpha != oldAlpha {
		err = pos.StorePVMove(bestMove)
		if err != nil {
			return 0, err
		}
	}

	//if we did not improve on alpha return alpha
	return alpha, nil
}

//SearchPosition for best move
func (info *InfoStruct) SearchPosition(pos *board.PositionStruct) error {
	var bestMove int = board.NoMove
	var bestScore int = -infinitie
	var pvMoves = 0
	var pvNum = 0

	var err error

	info.clearForSearch(pos)

	//Iterative deepening loop
	for currentDepth := 1; currentDepth <= info.Depth; currentDepth++ {
		bestScore, err = info.alphaBeta(-infinitie, infinitie, currentDepth, true, pos)
		if err != nil {
			return err
		}

		//TODO: Check if out of time

		pvMoves, err = pos.GetPvLine(currentDepth)
		if err != nil {
			return err
		}
		bestMove = pos.PvArray[0]

		fmt.Printf("Depth: %d, Scroe: %d, Move: %s, Nodes %d, ",
			currentDepth, bestScore, board.MoveToString(bestMove), info.Nodes)

		fmt.Print("PV: ")
		for pvNum = 0; pvNum < pvMoves; pvNum++ {
			fmt.Printf(" %s", board.MoveToString(pos.PvArray[pvNum]))
		}
		fmt.Print("\n")
		fmt.Printf("Ordering: %.2f\n", info.FailHighFirst/info.FailHigh)
	}

	return err
}
