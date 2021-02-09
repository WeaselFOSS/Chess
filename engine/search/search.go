package search

import "github.com/WeaselChess/Weasel/engine/board"

type InfoStruct struct {
	StartTime int
	StopTiem  int
	depth     int
	depthSet  bool
	TimeSet   bool
	MovesToGo bool
	Infinite  bool
	Nodes     int64
	Quit      bool
	Stopped   bool
}

//checkUp Check if time up, or interrupt from GUI
func checkUp() {

}

//clearForSearch Reset serach info and PVTables to get ready for a enw search
func (info *InfoStruct) clearForSearch(pos *board.PositionStruct) {

}

//quiescence Search capture moves until a "quite" position is found or the trade has been resolved
//
//Used to counter the horizon effect
func (info *InfoStruct) quiescence(alpha, beta, pos *board.PositionStruct) int {
	return 0
}

//alphaBeta Normal alphabeta searching
func (info *InfoStruct) alphaBeta(alpha, beta, depth int, doNull bool, pos *board.PositionStruct) int {
	return 0
}

//SearchPosition for best move
func (info *InfoStruct) SearchPosition(pos *board.PositionStruct) {

}
