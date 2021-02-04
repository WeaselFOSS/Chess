package engine

import (
	"errors"
	"math/rand"
	"time"
)

var pieceKeys [13][120]uint64
var sideKey uint64
var castelKeys [16]uint64

func initHashKeys() {
	rand.Seed(time.Now().UTC().UnixNano())

	for x := 0; x < 13; x++ {
		for y := 0; y < 120; y++ {
			pieceKeys[x][y] = rand.Uint64()
		}
	}
	sideKey = rand.Uint64()
	for i := 0; i < 16; i++ {
		castelKeys[i] = rand.Uint64()
	}
}

//GeneratePosKey generates a unique key for the current position
func (pos *BoardStruct) GeneratePosKey() error {
	var finalKey uint64 = 0
	var piece int = Empty

	//Pieces
	for sq := 0; sq < SquareNumber; sq++ {
		piece = pos.Pieces[sq]
		if piece != NoSquare && piece != Empty {
			if !(piece >= WP && piece <= BK) {
				return errors.New("Piece value out of bounds")
			}
			finalKey ^= pieceKeys[piece][sq]
		}
	}

	//Side
	if pos.Side == White {
		finalKey ^= sideKey
	}

	//EnPassant
	if pos.EnPassant != NoSquare {
		if !(pos.EnPassant >= 0 && pos.EnPassant < SquareNumber) {
			return errors.New("EnPassant value out of bounds")
		}
		finalKey ^= pieceKeys[Empty][pos.EnPassant]
	}

	//CastelPerm
	if !(pos.CastelPerm >= 0 && pos.CastelPerm <= 15) {
		return errors.New("CastelPerm value out of bounds")
	}
	finalKey ^= castelKeys[pos.CastelPerm]

	pos.PosKey = finalKey
	return nil
}
