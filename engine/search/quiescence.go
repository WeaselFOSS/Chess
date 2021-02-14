package search

import "github.com/WeaselChess/Weasel/engine/board"

// quiescence Search capture moves until a "quite" position is found or the trade has been resolved
//
// Used to counter the horizon effect
func (info *InfoStruct) quiescence(alpha, beta int, pos *board.PositionStruct) (int, error) {

	// Checkup ever 2048 nodes
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
	score = -board.Infinite // nolint

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
		if err != nil {
			return 0, err
		}
		// Flipping score for the other sides POV
		score *= -1
		err = pos.TakeMove()
		if err != nil {
			return 0, err
		}

		if info.Stopped {
			return 0, nil
		}

		// If score beats alpha set alpha to score
		if score > alpha {
			// If score is better than our beta cutoff return beta
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
