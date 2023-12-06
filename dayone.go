package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	calibrationValue := 0

	re := regexp.MustCompile(`\d`)

	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		extrctedNumbers := re.FindAllString(fileScanner.Text(), -1)
		convertedNumber, _ := strconv.Atoi(extrctedNumbers[0] + extrctedNumbers[len(extrctedNumbers)-1])
		calibrationValue = calibrationValue + convertedNumber
	}

	readFile.Close()
	fmt.Println(calibrationValue)
}
