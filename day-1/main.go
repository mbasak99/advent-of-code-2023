package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	for i := *index; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			value, ok := stringToNumMap[s[i:j]]

			// fmt.Println(i, j, value, s[i:j], s)
			if ok {
				fmt.Println(i, j, value, s[i:j], s)
				if j < len(s)-1 && !isNumeric(string(s[j+1])) {
					*index = j - 2
				} else {
					*index = j - 1
				}
				// fmt.Printf("index: %d\n", *index)
				return value, true
			}
		}
	}

	return "", false
}

func isNumeric(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	// if err != nil {
	// 	fmt.Printf("Number: %d\n", num)
	// }
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
				// fmt.Printf("index: %d\n", i)
				numStr, isNum := stringToNum(fileScanner.Text(), &i)
				// fmt.Println(numStr)
				if isNum && firstNum == "" {
					firstNum = numStr
				} else if isNum && firstNum != "" {
					lastNum = numStr
				}
				// fmt.Printf("index after: %d\n", i)
			}
		}

		if firstNum != "" && lastNum != "" {
			numVal, _ := strconv.Atoi(firstNum + lastNum)
			// fmt.Printf("numVal: %d\n", numVal)
			totalNum += numVal
		} else if firstNum != "" {
			numVal, _ := strconv.Atoi(firstNum + firstNum)
			// fmt.Printf("numVal: %d\n", numVal)
			totalNum += numVal
		}
		// fmt.Println()
		// fmt.Printf("%s%s \n", firstNum, lastNum)
		// break
	}
	fmt.Printf("Total number: %d\n", totalNum)
}
