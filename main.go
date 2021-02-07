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
	fen := "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1"
	var pos engine.BoardStruct
	var list engine.MoveListStruct
	pos.LoadFEN(fen)
	pos.Print()
	err := pos.GenerateAllMoves(&list)
	if err != nil {
		panic(err)
	}
	pos.MakeMove(list.Moves[2].Move)
	list.Print()
	pos.Print()
}
