package board

import (
	"errors"
	"strconv"
	"strings"
)

// LoadFEN loads the engine with a new board position from a FEN string
func (pos *PositionStruct) LoadFEN(fen string) error {
	if fen == "" {
		return errors.New("FEN String is empty")
	}

	rank := rank8
	file := fileA
	piece := 0
	count := 0

	pos.resetBoard()

	for (rank >= rank1) && len(fen) > 0 {
		count = 1

		switch fen[0] {
		case 'p':
			piece = bP
		case 'n':
			piece = bN
		case 'b':
			piece = bB
		case 'r':
			piece = bR
		case 'q':
			piece = bQ
		case 'k':
			piece = bK

		case 'P':
			piece = wP
		case 'N':
			piece = wN
		case 'B':
			piece = wB
		case 'R':
			piece = wR
		case 'Q':
			piece = wQ
		case 'K':
			piece = wK

		case '1', '2', '3', '4', '5', '6', '7', '8':
			piece = empty
			count = int(fen[0] - '0')

		case '/', ' ':
			rank--
			file = fileA
			fen = fen[1:]
			continue

		default:
			return errors.New("Bad FEN string")
		}

		for i := 0; i < count; i++ {
			sq64 := rank*8 + file
			sq120 := sq64ToSq120[sq64]
			if piece != empty {
				pos.Pieces[sq120] = piece
			}
			file++
		}
		fen = fen[1:]
	}

	if fen[0] != 'w' && fen[0] != 'b' {
		return errors.New("Bad FEN Side To move")
	}

	if fen[0] == 'w' {
		pos.Side = white
	} else {
		pos.Side = black
	}

	if len(fen) < 3 {
		return errors.New("Bad FEN Length")
	}

	fen = fen[2:]

	for i := 0; i < 4; i++ {
		if fen[0] == ' ' {
			break
		}
		switch fen[0] {
		case 'K':
			pos.CastelPerm |= wkcastel
		case 'Q':
			pos.CastelPerm |= wqcastel
		case 'k':
			pos.CastelPerm |= bkcastel
		case 'q':
			pos.CastelPerm |= bqcastel
		}
		fen = fen[1:]
	}

	if len(fen) < 2 {
		return errors.New("Bad FEN Length")
	}

	fen = fen[1:]

	if fen[0] != '-' {
		file = int(fen[0] - 'a')
		rank = int(fen[1] - '1')

		if len(fen) < 4 {
			return errors.New("Bad FEN Length")
		}

		fen = fen[3:]

		if file < fileA || file > fileH {
			return errors.New("Bad FEN EnPas File")
		}

		if rank < rank1 || rank > rank8 {
			return errors.New("Bad FEN EnPas Rank")
		}

		pos.EnPassant = fileRankToSquare(file, rank)
	} else {
		if len(fen) < 3 {
			return errors.New("Bad FEN Length")
		}
		fen = fen[2:]
	}

	nums := strings.Split(fen, " ")

	if len(nums) < 2 {
		return errors.New("Bad FEN Length")
	}

	var err error
	pos.FiftyMove, err = strconv.Atoi(nums[0])
	if err != nil {
		return err
	}

	pos.HisPly, err = strconv.Atoi(nums[1])
	if err != nil {
		return err
	}

	pos.updateMaterialLists()
	pos.PosKey, err = pos.generatePosKey()
	return err
}
