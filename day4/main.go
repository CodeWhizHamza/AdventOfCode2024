package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	lines := getLines()

	// lines := []string{
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

	// SolutionA(lines)
	SolutionB(lines)
}

func SolutionB(lines []string) {
	charactersMatrix := getCharactersMatrix(lines)
	count := 0
	for i := 1; i < len(charactersMatrix)-2; i++ {
		for j := 1; j < len(charactersMatrix[0])-2; j++ {
			if charactersMatrix[i][j] != "A" {
				continue
			}

			diagonal := charactersMatrix[i-1][j-1] + charactersMatrix[i+1][j+1]
			secondary := charactersMatrix[i-1][j+1] + charactersMatrix[i+1][j-1]

			if (diagonal == "MS" || diagonal == "SM") && (secondary == "MS" || secondary == "SM") {
				count += 1
			}
		}
	}

	fmt.Println("Count:", count)
}

func SolutionA(lines []string) {
	charactersMatrix := getCharactersMatrix(lines)

	var (
		horizontalCount        int = countHorizontal(lines)
		verticalCount          int = countVertical(charactersMatrix)
		mainDiagonalCount      int = countMainDiagonal(charactersMatrix)
		secondaryDiagonalCount int = countSecondaryDiagonal(charactersMatrix)
	)

	fmt.Println("Horizontal Count:", horizontalCount)
	fmt.Println("Vertical Count:", verticalCount)
	fmt.Println("Main Diagonal Count:", mainDiagonalCount)
	fmt.Println("Secondary Diagonal Count:", secondaryDiagonalCount)
	fmt.Println("==================================================")
	fmt.Println("Total:", horizontalCount+verticalCount+mainDiagonalCount+secondaryDiagonalCount)
}

func countHorizontal(lines []string) int {
	count := 0
	for _, l := range lines {
		count += countXMAS(l)
	}
	return count
}

func countVertical(charactersMatrix [][]string) int {
	count := 0
	for i := 0; i < len(charactersMatrix); i++ {
		l := ""
		for j := 0; j < len(charactersMatrix[0]); j++ {
			l += charactersMatrix[j][i]
		}
		count += countXMAS(l)
	}

	return count
}

func countMainDiagonal(charactersMatrix [][]string) int {
	/**
	My whole theory behind writing this function is that
	I am trying to move through all the diagonals of the
	character matrix
	Add all of it in a single string, and then count xmas
	forward and backward with overlapping in it.
	*/
	diagonalStrings := []string{}
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

		diagonalString := ""
		for row < len(charactersMatrix) && col < len(charactersMatrix[0]) {
			diagonalString += charactersMatrix[row][col]
			row += 1
			col += 1
		}
		if len(diagonalString) < 4 {
			continue
		}
		diagonalStrings = append(diagonalStrings, diagonalString)
	}

	count := 0
	for _, d := range diagonalStrings {
		count += countXMAS(d)
	}
	return count
}

func countSecondaryDiagonal(charactersMatrix [][]string) int {
	diagonalStrings := []string{}
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

		diagonalString := ""
		for row < len(charactersMatrix) && col >= 0 {
			diagonalString += charactersMatrix[row][col]
			row += 1
			col -= 1
		}
		if len(diagonalString) < 4 {
			continue
		}
		diagonalStrings = append(diagonalStrings, diagonalString)
	}

	count := 0
	for _, d := range diagonalStrings {
		count += countXMAS(d)
	}
	return count
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
