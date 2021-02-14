package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/WeaselChess/Weasel/engine/board"
	"github.com/WeaselChess/Weasel/engine/search"
	"github.com/WeaselChess/Weasel/engine/uci"
)

func start(engineInfo uci.EngineInfo) {
	var command []string
	scanner := bufio.NewScanner(os.Stdin)
	// Used to delete multiple spaces
	space := regexp.MustCompile(`\s+`)

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

// weaselConsol Weasel console mode
func weaselConsol() {

	fmt.Print("Welcome to weasel in console mode\n")
	fmt.Print("Type help for commands\n\n")

	var command []string
	scanner := bufio.NewScanner(os.Stdin)

	var pos board.PositionStruct
	board.Initialize()
	// Init hash tables size with 2 MB's
	pos.HashTable.Init(16)

	err := pos.LoadFEN(board.StartPosFEN)
	if err != nil {
		fmt.Println(err.Error())
	}

	pos.Print()

	var searchInfo search.InfoStruct
	searchInfo.Depth = board.MaxDepth
	var moveTime = 3000

	var forceMode = false

	fmt.Print("\nEnter Move > ")

	for scanner.Scan() {
		var moveMade = false
		index := 0
		command = strings.Split(scanner.Text(), " ")

		switch command[index] {
		case "print":
			pos.Print()
		case "eval":
			fmt.Println("Evaluation in the side to move's POV in 100ths of a pawn")
			fmt.Printf("The current Eveluation is %d\n", pos.Evaluate())
		case "startpos":
			pos.HashTable.Clear()
			err := pos.LoadFEN(board.StartPosFEN)
			if err != nil {
				fmt.Println(err.Error())
			}

			pos.Print()
		case "setboard":
			pos.HashTable.Clear()
			err := pos.LoadFEN(strings.Join(command[index+1:], " "))
			if err != nil {
				fmt.Println(err.Error())
			}

			if pos.PosKey != 0 {
				pos.Print()
			} else {
				err := pos.LoadFEN(board.StartPosFEN)
				if err != nil {
					fmt.Println(err.Error())
				}

				fmt.Printf("Unkown FEN string '%s'\n", strings.Join(command[index+1:], " "))
			}
		case "force":
			forceMode = true
			fmt.Println("Weasel will no longr make moves until you type 'go'")
		case "go":
			forceMode = false
			moveMade = true
		case "divide":
			divideHander(command[index:], &pos)
		case "mirror":
			err := pos.MirrorBoard()
			if err != nil {
				panic(err)
			}
		case "quit":
			os.Exit(0)
		case "help":
			fmt.Println("print - print the current board state")
			fmt.Println("eval - give the current evaluation score")
			fmt.Println("startpos - set the boards position to the starting position")
			fmt.Println("setboard x - set the board position to the FEN x")
			fmt.Println("force - weasel will not make moves")
			fmt.Println("go - let weasel make moves for current side to move")
			fmt.Println("divide x - runs a perft divide to the depth of x")
			fmt.Println("quit - exit weasel")
		default:
			move, _ := pos.ParseMove(command[index])

			var err error
			if move != board.NoMove {
				moveMade, err = pos.MakeMove(move)
				if err != nil {
					fmt.Println(err.Error())
				}
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
				fmt.Println(err.Error())
			}
			_, err = pos.MakeMove(pos.PvArray[0])
			if err != nil {
				fmt.Println(err.Error())
			}

			pos.Print()
			fmt.Printf("Weasel plays move: %s\n", board.MoveToString(pos.PvArray[0]))
		}

		fmt.Print("\nEnter Move > ")

	}

}

func divideHander(command []string, pos *board.PositionStruct) {
	if len(command) > 1 {
		if unicode.IsDigit(rune(command[1][0])) {
			var depth int = int(rune(command[1][0]) - '0')
			err := pos.PerftDivide(depth)
			if err != nil {
				panic(err)
			}
		}
	}
}
