package search

import (
	"fmt"
	"time"

	"github.com/WeaselChess/Weasel/engine/board"
)

// InfoStruct The search info struct
type InfoStruct struct {
	StartTime int64
	StopTime  int64
	Depth     int
	DepthSet  bool
	TimeSet   bool
	MovesToGo bool
	Infinite  bool
	Nodes     int64
	LeafNodes int
	Quit      bool
	Stopped   bool

	NullCut int

	// Used to calculate the efficency of our move ordering
	FailHigh      float32
	FailHighFirst float32
}

// avargeBranchingFactor calculated with the average(over multiple positions) of Sqrt(Nodes(currentDepth) / nodes(depth - 2))
const avargeBranchingFactor = 3.73

// pickNextMove Pick the next move base on inital score
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

// checkUp Check if time up, or interrupt from GUI
func (info *InfoStruct) checkUp() {
	currentTime := time.Now().UnixNano() / int64(time.Millisecond)
	if info.TimeSet && currentTime > info.StopTime {
		info.Stopped = true
	}
}

// clearForSearch Reset serach info and PVTables to get ready for a enw search
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

// SearchPosition for best move
func (info *InfoStruct) SearchPosition(pos *board.PositionStruct) error {
	var bestMove int = board.NoMove
	var bestScore int = -board.Infinite //nolint -Infinite by default
	var pvMoves int
	var pvNum int

	var err error

	info.clearForSearch(pos)

	// Iterative deepening loop
	for currentDepth := 1; currentDepth <= info.Depth; currentDepth++ {
		bestScore, err = info.pvSearch(-board.Infinite, board.Infinite, currentDepth, true, pos)
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

		currentTime := time.Now().UnixNano() / int64(time.Millisecond)

		//timeEstimate := int(avargeBranchingFactor * float64(currentTime-startTime))

		// Sending infor to GUI
		if bestScore >= board.IsMate {
			fmt.Printf("info score mate %d depth %d nodes %d time %d ",
				board.Infinite-bestScore, currentDepth, info.Nodes, currentTime-info.StartTime)
		} else if bestScore <= -board.IsMate {
			fmt.Printf("info score mate -%d depth %d nodes %d time %d ",
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
	if err != nil {
		return err
	}
	err = pos.ClearMoveFromHash()
	if err != nil {
		return err
	}
	fmt.Printf("bestmove %s\n", board.MoveToString(bestMove))
	err = pos.TakeMove()
	return err
}
