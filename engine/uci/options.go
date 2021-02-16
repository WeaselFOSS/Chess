package uci

import (
	"fmt"
	"strconv"
)

// EngineOptions The struct for the engine options
type EngineOptions struct {
	HashSize int
}

// optionsHnadler Handler for options command
func (info *EngineInfo) optionsHnadler(command []string) {
	var err error
	if len(command) > 4 {
		if command[1] == "name" {
			switch command[2] {
			case "Hash":
				if command[3] == "value" {
					info.Options.HashSize, err = strconv.Atoi(command[4])
					if err != nil {
						fmt.Println("Failed to parse hash size " + err.Error())
					}

					if info.Options.HashSize > 1024 {
						info.Options.HashSize = 1024
					} else if info.Options.HashSize < 1 {
						info.Options.HashSize = 1
					}
				}
			default:
				fmt.Printf("Unkown option %s\n", command[1])
			}
		}
	}
}

// printOptions Print out the options that are available
func printOptions() {
	fmt.Println("option name Hash type spin default 32 min 1 max 1024")
}
