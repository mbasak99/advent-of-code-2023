package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}

}

func stringToNum(s string, index *int) (string, bool) {
	stringToNumMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for wordNum, charNum := range stringToNumMap {
		numIndex := strings.Index(s[*index:], wordNum)
		if numIndex > -1 {
			fmt.Println(*index, numIndex, s[*index:*index+len(wordNum)], s)
			// fmt.Println(*index, numIndex, s[*index:*index+len(wordNum)])
			*index = (*index + numIndex + len(wordNum) - 1)
			return charNum, true
		}
	}

	// substring := []rune(s[index : len(s)-1])

	// var substring string
	// for i := *index; i < len(s); i++ {
	// 	substring = s[*index:i]

	// 	num, ok := stringToNumMap[substring]
	// 	fmt.Println(substring)
	// 	if ok {
	// 		*index = i
	// 		return num, true
	// 	}

	// 	if isNumeric(s[i:i]) {
	// 		*index = i
	// 		return "", false
	// 	}
	// }

	// no string num
	return "", false
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func main() {
	inputFile, readFileErr := os.Open("./input.txt")
	check(readFileErr)

	fileScanner := bufio.NewScanner(inputFile)

	var firstNum string
	var lastNum string
	var totalNum int
	for fileScanner.Scan() {
		firstNum = ""
		lastNum = ""

		// for i := 0; i < len(fileScanner.Text()); i++ {
		// 	fmt.Println(strconv.Itoa(int(fileScanner.Text()[i])))
		// }
		for i := 0; i < len(fileScanner.Text()); i++ {
			char := fileScanner.Text()[i]
			// fmt.Printf("%c %d\n", char, i)
			// Numeric path
			if isNumeric(string(char)) {
				if firstNum == "" && i == 0 {
					firstNum = string(char)
				} else if firstNum != "" && i == len(fileScanner.Text())-1 {
					lastNum = string(char)
				}
			} else {
				numStr, isNum := stringToNum(fileScanner.Text(), &i)
				// fmt.Println(numStr)
				if isNum && firstNum == "" {
					firstNum = numStr
				} else if isNum && firstNum != "" {
					lastNum = numStr
				}
			}
		}

		if firstNum != "" && lastNum != "" {
			numVal, _ := strconv.Atoi(firstNum + lastNum)
			totalNum += numVal
		} else if firstNum != "" {
			numVal, _ := strconv.Atoi(firstNum + firstNum)
			totalNum += numVal
		}
		// fmt.Println()
		fmt.Printf("%s%s \n", firstNum, lastNum)
		// break
	}
	fmt.Printf("Total number: %d\n", totalNum)
}
