package engine

func (pos *BoardStruct) addQuietMove(move int, list *MoveListsStruct) {
	list.Moves[list.Count].Move = move
	list.Moves[list.Count].Score = 0
	list.Count++
}

func (pos *BoardStruct) addCaptureMove(move int, list *MoveListsStruct) {
	list.Moves[list.Count].Move = move
	list.Moves[list.Count].Score = 0
	list.Count++
}

func (pos *BoardStruct) addEnPasMove(move int, list *MoveListsStruct) {
	list.Moves[list.Count].Move = move
	list.Moves[list.Count].Score = 0
	list.Count++
}

func (pos *BoardStruct) generateAllMoves(list *MoveListsStruct) {

}
