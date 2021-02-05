package main

import (
	"github.com/WeaselChess/Weasel/engine"
	"github.com/WeaselChess/Weasel/uci"
)

//EngineInfo is the info for our engine
var EngineInfo = uci.EngineInfo{
	Name:    "Weasel",
	Version: "v0.0.1-beta",
	Author:  "WeaselChess Club",
}

func init() {
	println("                                                  ")
	println("██╗    ██╗███████╗ █████╗ ███████╗███████╗██╗     ")
	println("██║    ██║██╔════╝██╔══██╗██╔════╝██╔════╝██║     ")
	println("██║ █╗ ██║█████╗  ███████║███████╗█████╗  ██║     ")
	println("██║███╗██║██╔══╝  ██╔══██║╚════██║██╔══╝  ██║     ")
	println("╚███╔███╔╝███████╗██║  ██║███████║███████╗███████╗")
	println(" ╚══╝╚══╝ ╚══════╝╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝")
	println("                                                  ")
}

func main() {
	//uci.UCI(EngineInfo)
	//TODO: Make engine
	//TEMP for debugging
	fen := "rnbqkb1r/pp1p1pPp/8/2p1pP2/1P1P4/3P3P/P1P1P3/RNBQKBNR w KQkq e6 0 1"
	var pos engine.BoardStruct
	var list engine.MoveListStruct
	pos.LoadFEN(fen)
	pos.GenerateAllMoves(&list)
	list.Print()
}
