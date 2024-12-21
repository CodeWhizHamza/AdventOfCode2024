package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := getLines()

	fmt.Println("Solution A:", SolutionA(lines))
	fmt.Println("Solution B:", SolutionB(lines))

}

func SolutionB(lines []string) int64 {
	mulRegex, err := regexp.Compile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	if err != nil {
		log.Fatal("Unable to compile the given regex. Please confirm that the regex is fine.")
	}

	var (
		total    int64 = 0
		isAdding bool  = true
	)
	var totals []int64 = []int64{}
	for _, line := range lines {

		total, isAdding = getTotalForLineB(line, mulRegex, isAdding)
		totals = append(totals, total)
	}

	total = 0
	for _, v := range totals {
		total += v
	}
	return total
}

func getTotalForLineB(line string, regex *regexp.Regexp, isAdding bool) (int64, bool) {
	matches := regex.FindAllString(line, -1)

	var total int64 = 0
	for _, match := range matches {
		if match == "do()" {
			isAdding = true
			continue
		}
		if match == "don't()" {
			isAdding = false
			continue
		}

		if !isAdding {
			continue
		}

		numStrs := strings.Split(match[4:len(match)-1], ",")

		num1 := strToNumber(numStrs[0])
		num2 := strToNumber(numStrs[1])

		total += num1 * num2
	}
	return total, isAdding
}

func SolutionA(lines []string) int64 {
	mulRegex, err := regexp.Compile(`mul\((?<a>\d+),(?<b>\d+)\)`)
	if err != nil {
		log.Fatal("Unable to compile the given regex. Please confirm that the regex is fine.")
	}

	var totals []int64 = []int64{}
	for _, line := range lines {

		totals = append(totals, getTotalForLineA(line, mulRegex))
	}

	var total int64 = 0
	for _, v := range totals {
		total += v
	}
	return total
}

func getTotalForLineA(line string, regex *regexp.Regexp) int64 {
	matches := regex.FindAllString(line, -1)

	var total int64 = 0
	for _, match := range matches {
		numStrs := strings.Split(match[4:len(match)-1], ",")

		num1 := strToNumber(numStrs[0])
		num2 := strToNumber(numStrs[1])

		total += num1 * num2
	}
	return total
}

func strToNumber(s string) int64 {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("Unable to convert %s to int64.", s)
	}
	return num
}

func getLines() []string {
	f, err := os.Open("3.txt")
	if err != nil {
		log.Fatal("Cannot Open the file")
	}

	scanner := bufio.NewScanner(f)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines

}
