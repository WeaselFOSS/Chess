package main

import (
	"fmt"

	"github.com/WeaselChess/Weasel/board"
)

func main() {
	for i := 0; i < len(board.Sq64ToSq120); i++ {
		fmt.Println(board.Sq64ToSq120[i])
	}

}
