package engine

//loopSlidePiece array used to loop through all sliding pieces of one color
var loopSlidePiece = [8]int{wB, wR, wQ, 0, bB, bR, bQ, 0}

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
	err := pos.generateAllPawnMoves(list)
	if err != nil {
		return err
	}

	return nil
}
