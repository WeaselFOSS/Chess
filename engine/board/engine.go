package board

//Initialize conversion arrays for engine operation
func Initialize() {
	initSq120To64()
	initFileRanks()
	initBitMasks()
	initHashKeys()
	InitMvvLva()
}

//initSq120To64 Initalize sq120tosq64 and sq64tosq120 arrays
func initSq120To64() {
	for i := 0; i < SquareNumber; i++ {
		sq120ToSq64[i] = 65
	}

	for i := 0; i < 64; i++ {
		sq64ToSq120[i] = 120
	}

	sq64 := 0
	for rank := rank1; rank <= rank8; rank++ {
		for file := fileA; file <= fileH; file++ {
			sq := fileRankToSquare(file, rank)
			sq64ToSq120[sq64] = sq
			sq120ToSq64[sq] = sq64
			sq64++
		}
	}
}

//initFileRanks Initialize File and Rank arrays
func initFileRanks() {
	for i := 0; i < SquareNumber; i++ {
		filesBoard[i] = offBoard
		ranksBoard[i] = offBoard
	}

	for rank := rank1; rank <= rank8; rank++ {
		for file := fileA; file <= fileH; file++ {
			sq := fileRankToSquare(file, rank)
			filesBoard[sq] = file
			ranksBoard[sq] = rank
		}
	}
}

//IsRepition tests for 3 fold repitition
func (pos *PositionStruct) IsRepition() bool {
	if pos.HisPly >= 8 {
		return (pos.PosKey == pos.History[pos.HisPly-4].PosKey && pos.PosKey == pos.History[pos.HisPly-8].PosKey)
	}

	return false
}
