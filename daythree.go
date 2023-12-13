package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	resultSlice := make([]int, 0)
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {

		trimmedString := fileScanner.Text()[strings.IndexByte(fileScanner.Text(), ':'):]
		trimmedString = strings.TrimPrefix(trimmedString, ": ")
		splitToSlice := strings.Split(trimmedString, "|")
		splitToSlice[0] = strings.Trim(splitToSlice[0], " ")
		splitToSlice[1] = strings.Trim(splitToSlice[1], " ")
		firstSlice := strings.Split(splitToSlice[0], " ")
		secondSlice := strings.Split(splitToSlice[1], " ")

		//fmt.Println(firstSlice)
		//fmt.Println(secondSlice)
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
							fmt.Printf("i: %s j: %s\n", firstSlice[i], secondSlice[j])
							match, _ := strconv.Atoi(firstSlice[i])

							resultSlice = append(resultSlice, match)

						}
					}

				}

			}

		}

	}
	readFile.Close()
	fmt.Println(resultSlice)
}
