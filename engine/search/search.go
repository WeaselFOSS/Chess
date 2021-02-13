package search

import (
	"fmt"
	"time"

	"github.com/WeaselChess/Weasel/engine/board"
)

//InfoStruct The search info struct
type InfoStruct struct {
	StartTime int64
	StopTime  int64
	Depth     int
	DepthSet  bool
	TimeSet   bool
	MovesToGo bool
	Infinite  bool
	Nodes     int64
	Quit      bool
	Stopped   bool

	NullCut int

	//Used to calculate the efficency of our move ordering
	FailHigh      float32
	FailHighFirst float32
}

//pickNextMove Pick the next move base on inital score
func pickNextMove(moveNum int, list *board.MoveListStruct) {
	var temp board.MoveStruct
	bestScore := 0
	bestNum := moveNum

	for i := moveNum; i < list.Count; i++ {
		if list.Moves[i].Score > bestScore {
			bestScore = list.Moves[i].Score
			bestNum = i
		}
	}
	temp = list.Moves[moveNum]
	list.Moves[moveNum] = list.Moves[bestNum]
	list.Moves[bestNum] = temp
}

//checkUp Check if time up, or interrupt from GUI
func (info *InfoStruct) checkUp() {
	currentTime := time.Now().UnixNano() / int64(time.Millisecond)
	if info.TimeSet && currentTime > info.StopTime {
		info.Stopped = true
	}
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

	pos.HashTable.OverWrite = 0
	pos.HashTable.Hit = 0
	pos.HashTable.Cut = 0

	pos.Ply = 0

	info.Stopped = false
	info.Nodes = 0
	info.FailHigh = 0
	info.FailHighFirst = 0
}

//quiescence Search capture moves until a "quite" position is found or the trade has been resolved
//
//Used to counter the horizon effect
func (info *InfoStruct) quiescence(alpha, beta int, pos *board.PositionStruct) (int, error) {

	//Checkup ever 2048 nodes
	if info.Nodes&2047 == 0 {
		info.checkUp()
	}

	info.Nodes++

	if pos.IsRepition() || pos.FiftyMove >= 100 {
		return 0, nil
	}

	if pos.Ply > board.MaxDepth-1 {
		return pos.Evaluate(), nil
	}

	score := pos.Evaluate()

	if score >= beta {
		return beta, nil
	}

	if score > alpha {
		alpha = score
	}

	var list board.MoveListStruct
	err := pos.GenerateAllCaptureMoves(&list)
	if err != nil {
		return 0, err
	}

	legal := 0
	score = -board.Infinite

	for i := 0; i < list.Count; i++ {

		pickNextMove(i, &list)

		moveMade := false

		moveMade, err := pos.MakeMove(list.Moves[i].Move)
		if err != nil {
			return 0, nil
		}

		if !moveMade {
			continue
		}

		legal++

		score, err = info.quiescence(-beta, -alpha, pos)
		//Flipping score for the other sides POV
		score *= -1
		err = pos.TakeMove()
		if err != nil {
			return 0, err
		}

		if info.Stopped {
			return 0, nil
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
		}
	}
	return alpha, nil
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
		return info.quiescence(alpha, beta, pos)
	}

	//Checkup ever 2048 nodes
	if info.Nodes&2047 == 0 {
		info.checkUp()
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

	//Test if the side to move is in check, if so extend the search by 1 depth
	inCheck, err := pos.IsAttacked(pos.KingSquare[pos.Side], pos.Side^1)
	if err != nil {
		return 0, err
	}

	if inCheck {
		depth++
	}

	var found bool
	pvMove := board.NoMove
	score := -board.Infinite

	found, err = pos.ProbeHashEntry(&pvMove, &score, alpha, beta, depth)
	if err != nil {
		return 0, err
	}

	if found {
		pos.HashTable.Cut++

		return score, nil
	}

	if doNull && !inCheck && pos.Ply != 0 && pos.BigPieces[pos.Side] > 0 && depth >= 4 {
		err = pos.MakeNullMove()
		if err != nil {
			return 0, err
		}

		score, err = info.alphaBeta(-beta, -beta+1, depth-4, false, pos)
		if err != nil {
			return 0, err
		}
		score *= -1

		err = pos.TakeNullMove()
		if err != nil {
			return 0, err
		}

		if info.Stopped {
			return 0, nil
		}

		if score >= beta && score < board.IsMate {
			info.NullCut++
			return beta, nil
		}

	}

	var list board.MoveListStruct
	err = pos.GenerateAllMoves(&list)
	if err != nil {
		return 0, err
	}

	legal := 0
	oldAlpha := alpha
	bestMove := board.NoMove
	bestScore := -board.Infinite
	score = -board.Infinite

	if pvMove != board.NoMove {
		for i := 0; i < list.Count; i++ {
			if list.Moves[i].Move == pvMove {
				list.Moves[i].Score = 2000000
				break
			}
		}
	}

	//Move loop
	for i := 0; i < list.Count; i++ {
		pickNextMove(i, &list)

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

		if info.Stopped {
			return 0, nil
		}

		if score > bestScore {
			bestScore = score
			bestMove = list.Moves[i].Move
			//If score beats alpha set alpha to score
			if score > alpha {
				//If score is better than our beta cutoff return beta
				if score >= beta {
					if legal == 1 {
						info.FailHighFirst++
					}
					info.FailHigh++

					if list.Moves[i].Move&board.MoveFlagCA == 0 {
						pos.SearchKillers[1][pos.Ply] = pos.SearchKillers[0][pos.Ply]
						pos.SearchKillers[0][pos.Ply] = list.Moves[i].Move
					}

					err = pos.StoreHashEntry(bestMove, beta, board.HFBETA, depth)
					return beta, err
				}
				alpha = score
				bestMove = list.Moves[i].Move

				if list.Moves[i].Move&board.MoveFlagCAP == 0 {
					pos.SearchHistory[pos.Pieces[board.GetFrom(bestMove)]][board.GetTo(bestMove)] += depth
				}
			}
		}
	}

	if legal == 0 {
		if inCheck {
			//Return -mate plus the ply or moves to the mate so later we can take score and subtrace mate to get mate in X val
			return -board.Infinite + pos.Ply, nil
		} else {
			return 0, nil
		}
	}

	//If we found a better move store it in the PV table
	if alpha != oldAlpha {
		err = pos.StoreHashEntry(bestMove, bestScore, board.HFEXACT, depth)
		if err != nil {
			return 0, err
		}
	} else {
		err = pos.StoreHashEntry(bestMove, alpha, board.HFALPHA, depth)
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
	var bestScore int = -board.Infinite
	var pvMoves = 0
	var pvNum = 0

	var err error

	info.clearForSearch(pos)

	//Iterative deepening loop
	for currentDepth := 1; currentDepth <= info.Depth; currentDepth++ {
		bestScore, err = info.alphaBeta(-board.Infinite, board.Infinite, currentDepth, true, pos)
		if err != nil {
			return err
		}

		if info.Stopped {
			break
		}

		pvMoves, err = pos.GetPvLine(currentDepth)
		if err != nil {
			return err
		}
		bestMove = pos.PvArray[0]

		//Sending infor to GUI
		currentTime := time.Now().UnixNano() / int64(time.Millisecond)

		if bestScore >= board.IsMate {
			fmt.Printf("info score mate %d depth %d nodes %d time %d ",
				board.Infinite-bestScore, currentDepth, info.Nodes, currentTime-info.StartTime)
		} else if bestScore <= -board.IsMate {
			fmt.Printf("info score mate %d depth %d nodes %d time %d ",
				board.Infinite+bestScore, currentDepth, info.Nodes, currentTime-info.StartTime)
		} else {
			fmt.Printf("info score cp %d depth %d nodes %d time %d ",
				bestScore, currentDepth, info.Nodes, currentTime-info.StartTime)
		}
		fmt.Print("pv")
		for pvNum = 0; pvNum < pvMoves; pvNum++ {
			fmt.Printf(" %s", board.MoveToString(pos.PvArray[pvNum]))
		}
		fmt.Print("\n")
	}
	_, err = pos.MakeMove(bestMove)
	pos.ClearMoveFromHash()
	fmt.Printf("bestmove %s\n", board.MoveToString(bestMove))
	return err
}
