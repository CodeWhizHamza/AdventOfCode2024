package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func getLines() []string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Unable to open file input.txt")
	}

	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func getCharactersMatrix(lines []string) [][]string {
	charsMatrix := [][]string{}
	for _, line := range lines {
		chars := strings.Split(line, "")
		charsMatrix = append(charsMatrix, [][]string{chars}...)
	}
	return charsMatrix
}
