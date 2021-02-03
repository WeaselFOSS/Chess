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

var DebugModeOn bool = false

//UCI is our main loop for
func UCI(engineInfo EngineInfo) {
	for {
		var command string

		fmt.Scanln(&command)

		switch command {
		case "uci":
			identify(engineInfo)
		case "debug":
			if getCommandArguments(command)[0] == "on" {
				DebugModeOn = true
			} else {
				DebugModeOn = false
			}
		case "isready":
		case "setoption":
		case "register":
		case "ucinewgame":
		case "position":
		case "go":
		case "stop":
		case "ponderhit":
		case "quit":
		case "print":
		case "divide":
		}
	}
}

func getCommandArguments(command string) []string {
	return strings.SplitAfter(command, " ")
}

func identify(engineInfo EngineInfo) {
	println("id name " + engineInfo.Name + engineInfo.Version)
	println("id author " + engineInfo.Author)

	//TODO: Add supported options

	println("uciok")
}
