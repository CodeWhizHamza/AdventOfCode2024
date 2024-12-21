package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2.txt")
	if err != nil {
		log.Fatal("Unable to open the file")
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	reports := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		report := []int{}

		for _, item := range items {
			number, err := strconv.Atoi(item)
			if err != nil {
				log.Fatal("Unable to parse Integer")
			}
			report = append(report, number)
		}

		reports = append(reports, [][]int{report}...)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Solution 2a:", solution_2a(reports))
	fmt.Println("Solution 2b:", Solution2b(reports))

}
