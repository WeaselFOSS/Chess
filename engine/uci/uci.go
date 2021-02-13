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

//EngineInfo holds the info for our engine
type EngineInfo struct {
	Name    string
	Version string
	Author  string
}

//Current board position
var pos board.PositionStruct

//Search Info
var info search.InfoStruct

//Set to true if a position is set
var positionSet bool

//UCI is our main loop for
func UCI(engineInfo EngineInfo) {
	var command []string
	var ready bool
	scanner := bufio.NewScanner(os.Stdin)

	space := regexp.MustCompile(`\s+`) //Used to delete multiple spaces
	uciHander(engineInfo)
	for scanner.Scan() {
		index := 0

		command = strings.Split(space.ReplaceAllString(scanner.Text(), " "), " ")

	top:

		switch command[index] {
		case "uci":
			uciHander(engineInfo)
		case "debug":
			if command[index+1] == "on" {
				board.DEBUG = true
			} else {
				board.DEBUG = false
			}
		case "isready":
			go func() {

				board.Initialize()
				//Init hash tables size with 2 MB's
				pos.HashTable.Init(32) //TODO: Add option to set hash size in mb

				ready = true

				fmt.Println("readyok")
			}()
		case "setoption":
		case "ucinewgame":
			pos.HashTable.Clear()
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
