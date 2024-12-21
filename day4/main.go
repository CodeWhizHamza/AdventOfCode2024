package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	// lines := getLines()
	// charactersMatrix := getCharactersMatrix(lines)

	testLines := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}

	forwardAndBackwardCount := countHorizontal(testLines)
	fmt.Println(forwardAndBackwardCount)
}

func countHorizontal(lines []string) int {
	str := ""
	for _, l := range lines {
		str += l
	}
	return countXMAS(str)
}

func countXMAS(str string) int {
	forwardRegex, err := regexp.Compile(`XMAS`)
	if err != nil {
		log.Fatal("Cannot compile regex")
	}
	backwardRegex, err := regexp.Compile("SAMX")
	if err != nil {
		log.Fatal("Cannot compile regex")
	}

	count := 0
	count += len(forwardRegex.FindAllString(str, -1))
	count += len(backwardRegex.FindAllString(str, -1))
	return count
}
