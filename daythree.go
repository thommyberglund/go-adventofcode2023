package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	inputSlice := make([][]string, 0)
	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	sum := 0

	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		splitToSlice := strings.Split(fileScanner.Text(), "")
		inputSlice = append(inputSlice, [][]string{splitToSlice}...)
	}
	readFile.Close()

	for i := 0; i < len(inputSlice); i++ {
		current := ""
		part := false
		for j := 0; j < len(inputSlice[i]); j++ {
			if slices.Contains(numbers, inputSlice[i][j]) {
				current = current + inputSlice[i][j]

				js := len(inputSlice) - 1
				is := len(inputSlice[0]) - 1
				if (i > 0 && j > 0 && inputSlice[i-1][j-1] != ".") ||
					(i > 0 && inputSlice[i-1][j] != ".") ||
					(i > 0 && j < js && inputSlice[i-1][j+1] != ".") ||

					(j > 0 && inputSlice[i][j-1] != ".") && !slices.Contains(numbers, inputSlice[i][j-1]) ||
					(j < js && inputSlice[i][j+1] != ".") && !slices.Contains(numbers, inputSlice[i][j+1]) ||

					(i < is && j > 0 && inputSlice[i+1][j-1] != ".") ||
					(i < is && inputSlice[i+1][j] != ".") ||
					(i < is && j < js && inputSlice[i+1][j+1] != ".") {

					part = true
				}
			} else {
				if current != "" && part {
					convertedStr, _ := strconv.Atoi(current)
					fmt.Println(convertedStr)
					sum = sum + convertedStr

				}
				current = ""
				part = false
			}
		}
		if current != "" && part {
			convertedStr, _ := strconv.Atoi(current)
			fmt.Println(convertedStr)

			sum = sum + convertedStr
			current = ""
			part = false
		}
	}
	fmt.Println(sum)
}
