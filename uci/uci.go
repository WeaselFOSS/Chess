package uci

import (
	"fmt"
	"strings"
)

//EngineInfo holds the info for our engine
type EngineInfo struct {
	Name    string
	Version string
	Author  string
}

//DebugModeOn tells us if we should output debug messages
var DebugModeOn bool = false

//UCI is our main loop for
func UCI(engineInfo EngineInfo) {
	for {
		var cmd string

		fmt.Scanln(&cmd)

		command := strings.Split(cmd, " ")

		switch command[0] {
		case "uci":
			identify(engineInfo)
		case "debug":
			if command[1] == "on" {
				DebugModeOn = true
			} else {
				DebugModeOn = false
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
					//boardInit()
				}
			} else if command[1] == "fen" {
				//loadFEN(strings.Join(command[2:], " "))
			}
		case "go":
		case "stop":
		case "ponderhit":
		case "quit":
		case "print":
		case "divide":
		}
	}
}

func identify(engineInfo EngineInfo) {
	println("id name " + engineInfo.Name + engineInfo.Version)
	println("id author " + engineInfo.Author)

	//TODO: Add supported options

	println("uciok")
}
