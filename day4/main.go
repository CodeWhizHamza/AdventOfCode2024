package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	// lines := getLines()
	// charactersMatrix := getCharactersMatrix(lines)

	// testLines := []string{
	// 	"MMMSXXMASM",
	// 	"MSAMXMSMSA",
	// 	"AMXSXMAAMM",
	// 	"MSAMASMSMX",
	// 	"XMASAMXAMM",
	// 	"XXAMMXXAMA",
	// 	"SMSMSASXSS",
	// 	"SAXAMASAAA",
	// 	"MAMMMXMMMM",
	// 	"MXMXAXMASX",
	// }
	testLines := []string{
		"ABCD",
		"BCDA",
		"CDAB",
		"DABC",
	}
	testCharMatrix := getCharactersMatrix(testLines)

	var (
		horizontalCount        int = countHorizontal(testLines)
		verticalCount          int = countVertical(testCharMatrix)
		mainDiagonalCount      int = countMainDiagonal(testCharMatrix)
		secondaryDiagonalCount int = countSecondaryDiagonal(testCharMatrix)
	)

	fmt.Println("Horizontal Count:", horizontalCount)
	fmt.Println("Vertical Count:", verticalCount)
	fmt.Println("Main Diagonal Count:", mainDiagonalCount)
	fmt.Println("Secondary Diagonal Count:", secondaryDiagonalCount)

}

func countHorizontal(lines []string) int {
	str := ""
	for _, l := range lines {
		str += l
	}
	return countXMAS(str)
}

func countVertical(charactersMatrix [][]string) int {
	str := ""

	for i := 0; i < len(charactersMatrix); i++ {
		for j := 0; j < len(charactersMatrix[0]); j++ {
			str += charactersMatrix[j][i]
		}
	}

	return countXMAS(str)
}

func countMainDiagonal(charactersMatrix [][]string) int {
	/**
	My whole theory behind writing this function is that
	I am trying to move through all the diagonals of the
	character matrix
	Add all of it in a single string, and then count xmas
	forward and backward with overlapping in it.
	*/
	str := ""

	diagonalsCount := len(charactersMatrix) + len(charactersMatrix[0]) - 1
	for i := 0; i < diagonalsCount; i++ {
		var (
			row = 0
			col = 0
		)
		if i < len(charactersMatrix[0]) {
			col = i
		} else {
			row = i - len(charactersMatrix[0]) + 1
		}

		for row < len(charactersMatrix) && col < len(charactersMatrix[0]) {
			str += charactersMatrix[row][col]
			row += 1
			col += 1
		}
	}

	return countXMAS(str)
}

func countSecondaryDiagonal(charactersMatrix [][]string) int {
	str := ""

	diagonalsCount := len(charactersMatrix) + len(charactersMatrix[0]) - 1
	for i := 0; i < diagonalsCount; i++ {
		var (
			row = 0
			col = 0
		)
		if i < len(charactersMatrix[0]) {
			col = i
		} else {
			row = i - len(charactersMatrix[0]) + 1
			col = len(charactersMatrix[0]) - 1
		}

		for row < len(charactersMatrix) && col >= 0 {
			str += charactersMatrix[row][col]
			row += 1
			col -= 1
		}
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
