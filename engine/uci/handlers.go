package uci

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/WeaselChess/Weasel/engine/board"
)

func uciHander(engineInfo EngineInfo) {
	fmt.Println("id name " + engineInfo.Name + " " + engineInfo.Version)
	fmt.Println("id author " + engineInfo.Author)

	//TODO: Add supported options

	fmt.Println("uciok")
}

func positionHandler(command []string) {
	var boardSet bool
	positionSet = false
	if command[1] == "startpos" {
		err := pos.LoadFEN(board.StartPosFEN)
		if err != nil {
			panic(err)
		}
		boardSet = true
	} else if command[1] == "fen" {
		err := pos.LoadFEN(strings.Join(command[2:], " "))
		if err != nil {
			panic(err)
		}
		boardSet = true
	}

	str := strings.Join(command[2:], " ")

	if strings.Contains(str, "moves") && boardSet {

		index := strings.Index(str, "moves")

		str = str[index+6:]

		moves := strings.Split(str, " ")
		tmp := 0
		for i := 0; i < len(moves); i++ {
			if len(moves[i]) < 4 {
				continue
			}
			move, err := pos.ParseMove(moves[i])
			if err != nil {
				panic(err)
			}

			if move != board.NoMove {
				var moveMade bool
				moveMade, err := pos.MakeMove(move)
				if err != nil {
					panic(err)
				}
				if moveMade {
					tmp++
					continue
				}
			}

			fmt.Printf("Non-legal move: %s index: %d", moves[i], i)

			break
		}
	}
	positionSet = true

}

func goHandler(command []string) {
	var err error
	depth := -1
	movesToGo := 30
	moveTime := -1
	timeV := -1
	inc := 0
	info.TimeSet = false
	for i := 0; i < len(command); i++ {
		switch command[i] {
		case "binc":
			if pos.Side == 1 {
				inc, err = strconv.Atoi(command[i+1])
				if err != nil {
					panic("Failed to parse binc time " + err.Error())
				}
			}
		case "winc":
			if pos.Side == 0 {
				inc, err = strconv.Atoi(command[i+1])
				if err != nil {
					panic("Failed to parse winc time " + err.Error())
				}
			}
		case "wtime":
			if pos.Side == 0 {
				fmt.Println(command[i+1])
				timeV, err = strconv.Atoi(command[i+1])
				if err != nil {
					panic("Failed to parse wtime time " + err.Error())
				}
			}
		case "btime":
			if pos.Side == 1 {
				fmt.Println(command[i+1])
				timeV, err = strconv.Atoi(command[i+1])
				if err != nil {
					panic("Failed to parse btime time " + err.Error())
				}
			}

		case "movestogo":
			movesToGo, err = strconv.Atoi(command[i+1])
			if err != nil {
				panic("Failed to parse binc time " + err.Error())
			}

		case "movetime":
			moveTime, err = strconv.Atoi(command[i+1])
			if err != nil {
				panic("Failed to parse move time " + err.Error())
			}
		case "depth":
			depth, err = strconv.Atoi(command[i+1])
			if err != nil {
				panic("Failed to parse depth " + err.Error())
			}
		}
	}

	if moveTime != -1 {
		timeV = moveTime
		movesToGo = 1
	}

	info.StartTime = time.Now().UnixNano() / int64(time.Millisecond)
	info.Depth = depth

	if timeV != -1 {
		info.TimeSet = true

		if timeV >= 200 {
			timeV /= movesToGo
		}
		info.StopTime = info.StartTime + int64(timeV+inc)
		if info.StopTime <= 50 {
			info.StopTime = 500
		}
	}

	if depth == -1 {
		info.Depth = board.MaxDepth
	}

	fmt.Printf("time: %d start: %d stop: %d depth %d timeset %v\n",
		timeV, info.StartTime, info.StopTime, info.Depth, info.TimeSet)

	err = info.SearchPosition(&pos)
	if err != nil {
		panic(err)
	}
}
