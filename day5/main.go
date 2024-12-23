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

	// rules := []Pair{
	// 	{a: 47, b: 53},
	// 	{a: 97, b: 13},
	// 	{a: 97, b: 61},
	// 	{a: 97, b: 47},
	// 	{a: 75, b: 29},
	// 	{a: 61, b: 13},
	// 	{a: 75, b: 53},
	// 	{a: 29, b: 13},
	// 	{a: 97, b: 29},
	// 	{a: 53, b: 29},
	// 	{a: 61, b: 53},
	// 	{a: 97, b: 53},
	// 	{a: 61, b: 29},
	// 	{a: 47, b: 13},
	// 	{a: 75, b: 47},
	// 	{a: 97, b: 75},
	// 	{a: 47, b: 61},
	// 	{a: 75, b: 61},
	// 	{a: 47, b: 29},
	// 	{a: 75, b: 13},
	// 	{a: 53, b: 13},
	// }
	// printSequences := [][]int64{
	// 	{75, 47, 61, 53, 29},
	// 	{97, 61, 53, 29, 13},
	// 	{75, 29, 13},
	// 	{75, 97, 47, 61, 53},
	// 	{61, 13, 29},
	// 	{97, 13, 75, 29, 47},
	// }

	var total int64 = 0
	for _, sequence := range printSequences {
		if isValidSequence(sequence, rules) {
			mid := len(sequence) / 2
			total += sequence[mid]

			fmt.Println(sequence)
		}
	}

	fmt.Println("Task a:", total)
}

func isValidSequence(s []int64, rules []Pair) bool {
	for i, v := range s {
		for _, next := range s[i+1:] {
			filtered := []Pair{}
			for _, r := range rules {
				if r.a == next {
					filtered = append(filtered, r)
				}
			}

			for _, r := range filtered {
				if r.b == v {
					return false
				}
			}
		}
	}
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
