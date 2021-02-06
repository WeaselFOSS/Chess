package engine

func (list *MoveListStruct) addQuietMove(move int) {
	list.Moves[list.Count].Move = move
	list.Moves[list.Count].Score = 0
	list.Count++
}

func (list *MoveListStruct) addCaptureMove(move int) {
	list.Moves[list.Count].Move = move
	list.Moves[list.Count].Score = 0
	list.Count++
}

func (list *MoveListStruct) addEnPasMove(move int) {
	list.Moves[list.Count].Move = move
	list.Moves[list.Count].Score = 0
	list.Count++
}

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
