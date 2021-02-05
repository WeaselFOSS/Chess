package uci

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/WeaselChess/Weasel/engine"
)

//EngineInfo holds the info for our engine
type EngineInfo struct {
	Name    string
	Version string
	Author  string
}

//DebugModeOn tells us if we should output debug messages

var pos engine.BoardStruct

//UCI is our main loop for
func UCI(engineInfo EngineInfo) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")

		switch command[0] {
		case "uci":
			identify(engineInfo)
		case "debug":
			if command[1] == "on" {
				engine.DEBUG = true
			} else {
				engine.DEBUG = false
			}
		case "isready":
		case "setoption":
		case "register":
		case "ucinewgame":
		case "position":
			if command[1] == "startpos" {
				if len(command) > 2 {
					//boardInitWithMove(strings.Join(command[2:], " ")) //moves e1e2
				} else {
					err := pos.LoadFEN(engine.StartPosFEN)
					if err != nil {
						panic(err)
					}
				}
			} else if command[1] == "fen" {
				err := pos.LoadFEN(strings.Join(command[2:], " "))
				if err != nil {
					panic(err)
				}
			}
		case "go":
		case "stop":
		case "ponderhit":
		case "quit":
			os.Exit(0)
		case "print":
			pos.Print()
		case "divide":
		}
	}
}

func identify(engineInfo EngineInfo) {
	fmt.Println("id name " + engineInfo.Name + engineInfo.Version)
	fmt.Println("id author " + engineInfo.Author)

	//TODO: Add supported options

	fmt.Println("uciok")
}
