package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}

}

func main() {
	inputFile, readFileErr := os.Open("./input.txt")
	if readFileErr != nil {
		check(readFileErr)
	}

	fileScanner := bufio.NewScanner(inputFile)
	/* PART 1: 12 red cubes, 13 green cubes, 14 blue cubes */

	// store valid game ids for later summation
	// var gameIDs []string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		newLine := strings.Split(line, ":")
		newLine = strings.Split(newLine[1], ";")
		re := regexp.MustCompile(`(\d)*\s(red)?`)
		fmt.Printf("%s\n", newLine)
		fmt.Printf("%s\n", re.FindAll([]byte(newLine[0]), 100))
	}
}
