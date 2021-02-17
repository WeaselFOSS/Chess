package search

import (
	"github.com/WeaselChess/Weasel/engine/board"
)

// alphaBeta Normal alphabeta searching
func (info *InfoStruct) pvSearch(alpha, beta, depth int, doNull bool, pos *board.PositionStruct) (int, error) {

	if board.DEBUG {
		err := pos.CheckBoard()
		if err != nil {
			return 0, err
		}
	}

	if depth == 0 {
		return info.quiescence(alpha, beta, pos)
	}

	// Checkup ever 2048 nodes
	if info.Nodes&2047 == 0 {
		info.checkUp()
	}

	info.Nodes++

	// Check for repition and fifty move role
	if pos.IsRepition() || pos.FiftyMove >= 100 {
		return 0, nil
	}

	// if we are at our max depth return the eval
	if pos.Ply > board.MaxDepth-1 {
		return pos.Evaluate(), nil
	}

	// Test if the side to move is in check, if so extend the search by 1 depth
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

		score, err = info.pvSearch(-beta, -beta+1, depth-4, false, pos)
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

	foundPV := false

	// Move loop
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

		if foundPV {
			score, err = info.pvSearch(-alpha-1, -alpha, depth-1, true, pos)
			if err != nil {
				return 0, err
			}
			score *= -1

			if score > alpha && score < beta {
				score, err = info.pvSearch(-beta, -alpha, depth-1, true, pos)
				if err != nil {
					return 0, err
				}
				score *= -1
			}
		} else {
			score, err = info.pvSearch(-beta, -alpha, depth-1, true, pos)
			if err != nil {
				return 0, err
			}
			score *= -1
		}

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
			// If score beats alpha set alpha to score
			if score > alpha {
				// If score is better than our beta cutoff return beta
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
				foundPV = true
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
			// Return -mate plus the ply or moves to the mate so later we can take score and subtrace mate to get mate in X val
			return -board.Infinite + pos.Ply, nil
		} else {
			return 0, nil
		}
	}

	// If we found a better move store it in the PV table
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

	// if we did not improve on alpha return alpha
	return alpha, nil
}
