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

	forwardAndBackwardCount := countForwardAndBackwardWords(testLines)
	fmt.Println(forwardAndBackwardCount)
}

func countForwardAndBackwardWords(lines []string) int {
	forwardRegex, err := regexp.Compile(`XMAS`)
	if err != nil {
		log.Fatal("Cannot compile regex")
	}
	backwardRegex, err := regexp.Compile("SAMX")
	if err != nil {
		log.Fatal("Cannot compile regex")
	}

	count := 0

	for _, line := range lines {
		count += len(forwardRegex.FindAllString(line, -1))
		count += len(backwardRegex.FindAllString(line, -1))
	}

	return count
}
