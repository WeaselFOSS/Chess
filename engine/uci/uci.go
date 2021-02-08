package uci

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/WeaselChess/Weasel/engine/board"
)

//EngineInfo holds the info for our engine
type EngineInfo struct {
	Name    string
	Version string
	Author  string
}

//Current board position
var pos board.PositionStruct

//UCI is our main loop for
func UCI(engineInfo EngineInfo) {
	var command []string
	var ready bool

	scanner := bufio.NewScanner(os.Stdin)

	space := regexp.MustCompile(`\s+`) //Used to delete multiple spaces

	for scanner.Scan() {
		index := 0

		command = strings.Split(space.ReplaceAllString(scanner.Text(), " "), " ")

	top:

		switch command[index] {
		case "uci":
			go uciHander(engineInfo)
		case "debug":
			go func() {
				if command[index+1] == "on" {
					board.DEBUG = true
				} else {
					board.DEBUG = false
				}
			}()
		case "isready":
			go func() {
				board.Initialize()

				ready = true

				fmt.Println("readyok")
			}()
		case "setoption":
		case "register":
		case "ucinewgame":
		case "position":
			if ready {
				go positionHandler(command[index:])
			}
		case "go":
		case "stop":
		case "ponderhit":
		case "quit":
			os.Exit(0)
		case "print":
			go pos.Print()
		case "divide":
			if ready {
				go divideHander(command[index:])
			}
		default:
			if len(command) == index+1 {
				break
			}

			index++
			goto top
		}
	}
}

func uciHander(engineInfo EngineInfo) {
	fmt.Println("id name " + engineInfo.Name + engineInfo.Version)
	fmt.Println("id author " + engineInfo.Author)

	//TODO: Add supported options

	fmt.Println("uciok")
}

func positionHandler(command []string) {
	boardSet := false

	if command[1] == "startpos" {
		err := pos.LoadFEN(board.StartPosFEN)
		if err != nil {
			panic(err)
		}
		boardSet = true
	} else if command[1] == "fen" {
		err := pos.LoadFEN(strings.Join(command[2:], " "))
		if err != nil {
			panic(err)
		}
		boardSet = true
	}

	str := strings.Join(command[2:], " ")

	if strings.Contains(str, "moves") && boardSet {
		index := strings.Index(str, "moves")

		str = str[index+6:]

		moves := strings.Split(str, " ")

		for i := 0; i < len(moves); i++ {
			move, err := pos.ParseMove(moves[i])
			if err != nil {
				panic(err)
			}

			if move != board.NoMove {
				var moveMade bool
				moveMade, err := pos.MakeMove(move)
				if err != nil {
					panic(err)
				}
				if moveMade {
					continue
				}
			}

			fmt.Printf("Non-legal move: %s", moves[i])

			break
		}
	}
}

func divideHander(command []string) {
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