package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	totalPoints := 0
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		loopPoints := 0
		winningNumbers := make([]int, 0)
		trimmedString := fileScanner.Text()[strings.IndexByte(fileScanner.Text(), ':'):]
		trimmedString = strings.TrimPrefix(trimmedString, ": ")
		splitToSlice := strings.Split(trimmedString, "|")
		splitToSlice[0] = strings.Trim(splitToSlice[0], " ")
		splitToSlice[1] = strings.Trim(splitToSlice[1], " ")
		firstSlice := strings.Split(splitToSlice[0], " ")
		secondSlice := strings.Split(splitToSlice[1], " ")

		for i := 0; i < len(firstSlice); i++ {
			{
				if firstSlice[i] == "" {
					continue
				}

				for j := 0; j < len(secondSlice); j++ {
					{
						if secondSlice[j] == "" {
							continue
						}

						if firstSlice[i] == secondSlice[j] {
							match, _ := strconv.Atoi(firstSlice[i])

							winningNumbers = append(winningNumbers, match)

						}
					}
				}
			}

		}
		for i := 0; i < len(winningNumbers); i++ {
			if i == 0 {
				loopPoints = loopPoints + 1
				continue
			}
			loopPoints = loopPoints * 2
		}
		totalPoints = totalPoints + loopPoints
		loopPoints = 0

	}
	readFile.Close()
	fmt.Println(totalPoints)
}
