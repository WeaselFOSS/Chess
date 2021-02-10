package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/WeaselChess/Weasel/engine/board"
	"github.com/WeaselChess/Weasel/engine/search"
	"github.com/WeaselChess/Weasel/engine/uci"
)

func start(engineInfo uci.EngineInfo) {
	var command []string
	scanner := bufio.NewScanner(os.Stdin)
	space := regexp.MustCompile(`\s+`) //Used to delete multiple spaces

	fmt.Println("Welcome to the Weasel chess engine!")
	fmt.Println("For console mode type 'weasel'")

	for scanner.Scan() {
		index := 0
		command = strings.Split(space.ReplaceAllString(scanner.Text(), " "), " ")
	top:
		switch command[index] {
		case "uci":
			uci.UCI(EngineInfo)
		case "weasel":
			weaselConsol()
		default:
			if len(command) == index+1 {
				break
			}

			index++
			goto top
		}
	}
}

func weaselConsol() {

	fmt.Print("Welcome to weasel in console mode\n")
	fmt.Print("Type help for commands\n\n")

	var command []string
	scanner := bufio.NewScanner(os.Stdin)

	var pos board.PositionStruct
	board.Initialize()
	//Init hash tables size with 2 MB's
	pos.PVTable.Init(2)

	pos.LoadFEN(board.StartPosFEN)
	pos.Print()

	var searchInfo search.InfoStruct
	searchInfo.Depth = board.MaxDepth
	var moveTime = 3000

	var forceMode = false

	for scanner.Scan() {
		var moveMade = false
		index := 0
		command = strings.Split(scanner.Text(), " ")

		switch command[index] {
		case "print":
			pos.Print()
		case "eval":
			fmt.Printf("The current Eveluation is %d\n", pos.Evaluate())
		case "startpos":
			pos.LoadFEN(board.StartPosFEN)
		case "setboard":
			pos.LoadFEN(strings.Join(command[index+1:], " "))
		case "force":
			forceMode = true
		case "go":
			forceMode = false
			moveMade = true
		case "quit":
			os.Exit(0)
		case "help":
			fmt.Println("print - print the current board state")
			fmt.Println("eval - give the current evaluation score")
			fmt.Println("startpos - set the boards position to the starting position")
			fmt.Println("setboard x - set the board position to the FEN x")
			fmt.Println("force - make the engine not make moves")
			fmt.Println("go - let the engine make moves")
			fmt.Println("quit - exit the program")
		default:
			move, err := pos.ParseMove(command[index])
			if err != nil {
				panic(err)
			}

			moveMade, err = pos.MakeMove(move)
			if err != nil {
				panic(err)
			}

			if !moveMade {
				fmt.Printf("Unkown command '%s' type help for help\n", command[index])
			} else {
				pos.Print()
			}
		}

		if moveMade && !forceMode {
			searchInfo.StartTime = time.Now().UnixNano() / int64(time.Millisecond)
			searchInfo.StopTime = searchInfo.StartTime + int64(moveTime)
			searchInfo.TimeSet = true
			err := searchInfo.SearchPosition(&pos)
			if err != nil {
				panic(err)
			}
			_, err = pos.MakeMove(pos.PvArray[0])
			if err != nil {
				panic(err)
			}
			pos.Print()
		}

	}

}
