package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	gameId int
	red    int
	green  int
	blue   int
}

func main() {

	var Games = []Game{}
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		game := strings.TrimPrefix(fileScanner.Text(), "Game ")
		round, _ := strconv.Atoi(game[:strings.IndexByte(game, ':')])
		fmt.Printf("Game %d\n", round)

		gameStruct := Game{round, 0, 0, 0}
		mainString := strings.Split(game, ":")[1]
		getSets := strings.Split(mainString, ";")
		for _, v := range getSets {
			m := make(map[string]int)
			extractedSet := strings.Split(v, ",")
			for _, k := range extractedSet {
				mapping := strings.Split(strings.TrimSpace(k), " ")
				m[mapping[1]], _ = strconv.Atoi(mapping[0])
				if gameStruct.blue < m["blue"] {
					gameStruct.blue = m["blue"]
				}
				if gameStruct.red < m["red"] {
					gameStruct.red = m["red"]
				}
				if gameStruct.green < m["green"] {
					gameStruct.green = m["green"]
				}
			}
		}
		Games = append(Games, gameStruct)
	}

	readFile.Close()
	fmt.Println(Games)
}
