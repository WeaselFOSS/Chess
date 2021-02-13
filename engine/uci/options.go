package uci

import (
	"fmt"
	"strconv"
)

type EngineOptions struct {
	HashSize int
}

func (info *EngineInfo) optionsHnadler(command []string) {
	var err error
	switch command[1] {
	case "Hash":
		info.Options.HashSize, err = strconv.Atoi(command[2])
		if err != nil {
			fmt.Println("Failed to parse hash size " + err.Error())
		}

		if info.Options.HashSize > 1024 {
			info.Options.HashSize = 1024
		} else if info.Options.HashSize < 1 {
			info.Options.HashSize = 1
		}
	default:
		fmt.Printf("Unkown option %s", command[1])
	}
}

func printOptions() {
	fmt.Println("option name Hash type spin default 32 min 1 max 1024")
}
