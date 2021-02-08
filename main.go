package main

import (
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

	println("                                                  ")
	println("██╗    ██╗███████╗ █████╗ ███████╗███████╗██╗     ")
	println("██║    ██║██╔════╝██╔══██╗██╔════╝██╔════╝██║     ")
	println("██║ █╗ ██║█████╗  ███████║███████╗█████╗  ██║     ")
	println("██║███╗██║██╔══╝  ██╔══██║╚════██║██╔══╝  ██║     ")
	println("╚███╔███╔╝███████╗██║  ██║███████║███████╗███████╗")
	println(" ╚══╝╚══╝ ╚══════╝╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝")
	println("                                                  ")

	<-make(chan struct{})
}
