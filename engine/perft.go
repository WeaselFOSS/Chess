package engine

import (
	"fmt"
	"time"
)

func (pos *BoardStruct) Perft(depth int) (int, error) {
	leafNodes := 0
	if DEBUG {
		err := pos.CheckBoard()
		if err != nil {
			return 0, err
		}
	}

	if depth == 0 {
		return 1, nil
	}

	var list MoveListStruct
	err := pos.GenerateAllMoves(&list)
	if err != nil {
		return 0, err
	}

	for i := 0; i < list.Count; i++ {
		var makeMove bool
		makeMove, err = pos.MakeMove(list.Moves[i].Move)
		if err != nil {
			return 0, err
		}

		if !makeMove {
			continue
		}

		var foundNodes int
		foundNodes, err = pos.Perft(depth - 1)
		leafNodes += foundNodes

		if err != nil {
			return 0, err
		}

		err = pos.TakeMove()
		if err != nil {
			return 0, err
		}
	}

	return leafNodes, nil
}

func (pos *BoardStruct) PerftDivide(depth int) error {
	start := time.Now()
	if DEBUG {
		err := pos.CheckBoard()
		if err != nil {
			return err
		}
	}

	pos.Print()

	fmt.Printf("\nStarting perft Divide To Depth: %d\n", depth)

	var list MoveListStruct
	err := pos.GenerateAllMoves(&list)
	if err != nil {
		return err
	}

	rootNodes := 0
	leafNodes := 0
	for i := 0; i < list.Count; i++ {
		rootNodes++
		var makeMove bool
		makeMove, err = pos.MakeMove(list.Moves[i].Move)
		if err != nil {
			return err
		}

		if !makeMove {
			continue
		}
		var foundNodes int
		foundNodes, err = pos.Perft(depth - 1)
		if err != nil {
			return err
		}
		leafNodes += foundNodes

		err = pos.TakeMove()
		if err != nil {
			return err
		}

		fmt.Printf("%s: %d\n", MoveToString(list.Moves[i].Move), foundNodes)
	}
	if leafNodes == 0 {
		leafNodes = rootNodes
	}
	fmt.Printf("Perft Complete\nNodes visited: %d\nCalculation Time: %v\n", leafNodes, time.Since(start))

	return nil
}
