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
	fen := "rnbqkbnr/p1p1p3/3p3p/1p1p4/2P1Pp2/8/PP1P1PpP/RNBQKB1R b KQkq e3 0 1"
	var pos engine.BoardStruct
	var list engine.MoveListStruct
	pos.LoadFEN(fen)
	pos.Print()
	err := pos.GenerateAllMoves(&list)
	if err != nil {
		panic(err)
	}
	list.Print()
}
