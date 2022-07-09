package main

//go:generate goversioninfo

import (
	"fmt"

	"github.com/WeaselFOSS/Chess/engine/uci"
)

var EngineOptions = uci.EngineOptions{
	HashSize: 32,
}

// EngineInfo is the info for our engine
var EngineInfo = uci.EngineInfo{
	Name:    "Weasel",
	Version: "v1.0.2-beta",
	Author:  "WeaselChess Club",
	Options: EngineOptions,
}

func main() {

	fmt.Println("                                                  ")
	fmt.Println("██╗    ██╗███████╗ █████╗ ███████╗███████╗██╗     ")
	fmt.Println("██║    ██║██╔════╝██╔══██╗██╔════╝██╔════╝██║     ")
	fmt.Println("██║ █╗ ██║█████╗  ███████║███████╗█████╗  ██║     ")
	fmt.Println("██║███╗██║██╔══╝  ██╔══██║╚════██║██╔══╝  ██║     ")
	fmt.Println("╚███╔███╔╝███████╗██║  ██║███████║███████╗███████╗")
	fmt.Println(" ╚══╝╚══╝ ╚══════╝╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝")
	fmt.Println("                                                  ")

	start(EngineInfo)
}
