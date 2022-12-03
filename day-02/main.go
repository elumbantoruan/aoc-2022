package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var players []Player

	vals := []string{"A Y", "B X", "C Z"}
	for _, v := range vals {
		players = append(players, parsePlayer(v))
	}
	score := gamingRPS(players, compete)
	fmt.Println("score", score)

	players = nil

	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	filescanner := bufio.NewScanner(readFile)
	filescanner.Split(bufio.ScanLines)

	for filescanner.Scan() {
		val := filescanner.Text()
		players = append(players, parsePlayer(val))
	}

	score = gamingRPS(players, compete)
	fmt.Println("score", score)

	score = gamingRPS(players, compete2)
	fmt.Println("score", score)

}

func gamingRPS(players []Player, calculator func(opponent RPS, me RPS) int) int {
	var totalScore int
	for _, player := range players {
		op := rpsMap[player.Opponent]
		self := rpsMap[player.Me]
		totalScore += calculator(op, self)
	}
	return totalScore
}

func parsePlayer(val string) Player {
	return Player{
		Opponent: val[0],
		Me:       val[2],
	}
}

func compete(opRps RPS, myRps RPS) int {
	var total int

	if opRps == Rock && myRps == Rock {
		total = 1 + 3
	} else if opRps == Rock && myRps == Paper {
		total = 2 + 6
	} else if opRps == Rock && myRps == Scissors {
		total = 3 + 0
	} else if opRps == Paper && myRps == Rock {
		total = 1 + 0
	} else if opRps == Paper && myRps == Paper {
		total = 2 + 3
	} else if opRps == Paper && myRps == Scissors {
		total = 3 + 6
	} else if opRps == Scissors && myRps == Rock {
		total = 1 + 6
	} else if opRps == Scissors && myRps == Paper {
		total = 2 + 0
	} else if opRps == Scissors && myRps == Scissors {
		total = 3 + 3
	}
	return total
}

func compete2(opRps RPS, myRps RPS) int {
	var total int

	if opRps == Rock && myRps == Rock {
		total = 3 + 0
	} else if opRps == Rock && myRps == Paper {
		total = 1 + 3
	} else if opRps == Rock && myRps == Scissors {
		total = 2 + 6
	} else if opRps == Paper && myRps == Rock {
		total = 1 + 0
	} else if opRps == Paper && myRps == Paper {
		total = 2 + 3
	} else if opRps == Paper && myRps == Scissors {
		total = 3 + 6
	} else if opRps == Scissors && myRps == Rock {
		total = 2 + 0
	} else if opRps == Scissors && myRps == Paper {
		total = 3 + 3
	} else if opRps == Scissors && myRps == Scissors {
		total = 1 + 6
	}
	return total
}

var rpsMap = map[byte]RPS{
	'A': Rock,
	'X': Rock,
	'B': Paper,
	'Y': Paper,
	'C': Scissors,
	'Z': Scissors,
}

type RPS int

const (
	Rock     RPS = 1
	Paper    RPS = 2
	Scissors RPS = 3
)

type Player struct {
	Opponent byte
	Me       byte
}
