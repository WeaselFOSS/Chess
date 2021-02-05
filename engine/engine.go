package engine

func init() {
	initSq120To64()
	initFileRanks()
	initBitMasks()
	initHashKeys()
}

func initSq120To64() {
	for i := 0; i < squareNumber; i++ {
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

func initFileRanks() {
	for i := 0; i < squareNumber; i++ {
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
