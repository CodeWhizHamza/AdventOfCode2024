package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	a, b int64
}

func main() {
	rules := readRules()
	printSequences := readPrintOrders()

	var total int64 = 0
	for _, sequence := range printSequences {
		if isValidSequence(sequence, rules) {
			mid := len(sequence) / 2
			total += sequence[mid]
			break
		}
	}

	fmt.Println("Task a:", total)
}

func isValidSequence(s []int64, rules []Pair) bool {

	return true
}

func readRules() []Pair {
	rules, err := os.Open("rules.txt")
	if err != nil {
		log.Fatal("Unable to open file: rules.txt")
	}
	defer func() {
		if err := rules.Close(); err != nil {
			log.Fatal("Unable to close file rules.txt")
		}
	}()

	scanner := bufio.NewScanner(rules)
	result := []Pair{}
	for scanner.Scan() {
		ruleString := strings.Split(scanner.Text(), "|")
		first, err := strconv.ParseInt(ruleString[0], 10, 64)
		if err != nil {
			log.Fatalf("Unable to convert string to number: %s\n", ruleString[0])
		}

		second, err := strconv.ParseInt(ruleString[0], 10, 64)
		if err != nil {
			log.Fatalf("Unable to convert string to number: %s\n", ruleString[0])
		}

		result = append(result, Pair{a: first, b: second})
	}

	return result
}

func readPrintOrders() [][]int64 {
	order, err := os.Open("orders.txt")
	if err != nil {
		log.Fatal("Unable to open file: rules.txt")
	}
	defer func() {
		if err := order.Close(); err != nil {
			log.Fatal("Unable to close file rules.txt")
		}
	}()

	scanner := bufio.NewScanner(order)
	result := [][]int64{}
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, ",")
		parsedValues := []int64{}
		for _, v := range values {
			parsed, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				log.Fatalf("Unable to parse to int: %s", v)
			}
			parsedValues = append(parsedValues, parsed)
		}
		result = append(result, [][]int64{parsedValues}...)
	}

	return result
}
