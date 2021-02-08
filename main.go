package main

import (
	"fmt"

	"github.com/WeaselChess/Weasel/engine/uci"
)

//EngineInfo is the info for our engine
var EngineInfo = uci.EngineInfo{
	Name:    "Weasel",
	Version: "v0.0.1-beta",
	Author:  "WeaselChess Club",
}

func main() {
	go uci.UCI(EngineInfo)

	fmt.Println("                                                  ")
	fmt.Println("██╗    ██╗███████╗ █████╗ ███████╗███████╗██╗     ")
	fmt.Println("██║    ██║██╔════╝██╔══██╗██╔════╝██╔════╝██║     ")
	fmt.Println("██║ █╗ ██║█████╗  ███████║███████╗█████╗  ██║     ")
	fmt.Println("██║███╗██║██╔══╝  ██╔══██║╚════██║██╔══╝  ██║     ")
	fmt.Println("╚███╔███╔╝███████╗██║  ██║███████║███████╗███████╗")
	fmt.Println(" ╚══╝╚══╝ ╚══════╝╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝")
	fmt.Println("                                                  ")

	<-make(chan struct{})
}
