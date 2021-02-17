package uci

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/WeaselChess/Weasel/engine/board"
	"github.com/WeaselChess/Weasel/engine/search"
)

// EngineInfo holds the info for our engine
type EngineInfo struct {
	Name    string
	Version string
	Author  string
	Options EngineOptions
}

// timeMargine is the maring of time to subtract from the total time to ensure we do not go over our total time
const timeMargin = 10

// Current board position
var pos board.PositionStruct

// Search Info
var info search.InfoStruct

// Set to true if a position is set
var positionSet bool

// UCI is our main loop for
func UCI(engineInfo EngineInfo) {
	var command []string
	var ready bool
	scanner := bufio.NewScanner(os.Stdin)

	// Used to delete multiple spaces
	space := regexp.MustCompile(`\s+`)
	uciHander(engineInfo)

	board.Initialize()

	// Init hash tables size with the size configured in options, defaults to 32 MBs
	pos.HashTable.Init(uint64(engineInfo.Options.HashSize))

	for scanner.Scan() {
		index := 0

		command = strings.Split(space.ReplaceAllString(scanner.Text(), " "), " ")

	top:

		switch command[index] {
		case "uci":
			uciHander(engineInfo)
		case "debug":
			if len(command) > 1 {
				if command[index+1] == "on" {
					board.DEBUG = true
				} else {
					board.DEBUG = false
				}
			}
		case "isready":
			go func() {
				ready = true
				fmt.Println("readyok")
			}()
		case "setoption":
			engineInfo.optionsHnadler(command[index:])
		case "ucinewgame":
			pos.HashTable.Init(uint64(engineInfo.Options.HashSize))
			err := pos.LoadFEN(board.StartPosFEN)
			if err != nil {
				panic(err)
			}
		case "position":
			if ready {
				positionHandler(command[index:])
			}
		case "go":
			if ready && positionSet {
				go goHandler(command[index:])
			}
		case "stop":
			info.Stopped = true
		case "quit":
			os.Exit(0)
		case "print":
			go pos.Print()
		default:
			if len(command) == index+1 {
				break
			}

			index++
			goto top
		}
	}
}
