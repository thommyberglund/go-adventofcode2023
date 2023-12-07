package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	var lettersToNumbers = strings.NewReplacer(
		"oneight", "18",
		"twone", "21",
		"threeight", "38",
		"fiveight", "58",
		"sevenine", "79",
		"eightwo", "82",
		"eighthree", "83",
		"nineight", "98",
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9").Replace

	/*	lettersAndNUmbersToSlice := map[string][]int{
		"oneight": "18",
		"twone", "21",
		"threeight", "38",
		"fiveight", "58",
		"sevenine", "79",
		"eightwo", "82",
		"eighthree", "83",
		"nineight", "98",
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9"
	}*/

	calibrationValue := 0

	//re := regexp.MustCompile(`\d`)

	reg_2 := regexp.MustCompile(`([0-9]|oneight|twone|threeight|fiveight|sevenine|eightwo|eighthree|nineight|one|two|three|four|five|six|seven|eight|nine)`)

	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	/*	fileScanner := bufio.NewScanner(readFile)
			fileScanner.Split(bufio.ScanLines)

			// Del 1
			for fileScanner.Scan() {
				extrctedNumbers := re.FindAllString(fileScanner.Text(), -1)
				convertedNumber, _ := strconv.Atoi(extrctedNumbers[0] + extrctedNumbers[len(extrctedNumbers)-1])
				calibrationValue = calibrationValue + convertedNumber
			}

			readFile.Close()



		readFile, err = os.Open(filePath)
		if err != nil {
			fmt.Println(err)
		}

	*/
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// Del 2
	for fileScanner.Scan() {
		extrctedNumbers2 := reg_2.FindAllString(fileScanner.Text(), -1)
		for i, e := range extrctedNumbers2 {
			extrctedNumbers2[i] = lettersToNumbers(e)
		}
		convNum, _ := strconv.Atoi(extrctedNumbers2[0])
		convNumEnd, _ := strconv.Atoi(extrctedNumbers2[len(extrctedNumbers2)-1])
		if convNum > 9 {
			convNum = convNum / 10
		}
		if convNumEnd > 9 {
			convNumEnd = convNumEnd % 10
		}
		convertedNumber2, _ := strconv.Atoi(strconv.Itoa(convNum) + strconv.Itoa(convNumEnd))
		calibrationValue = calibrationValue + convertedNumber2
	}
	readFile.Close()
	fmt.Println(calibrationValue)

}
