package main

import (
	"slices"
)

func Solution2b(reports [][]int) int {
	safe_count := 0

	for _, report := range reports {
		if isSafe(report) {
			safe_count++
		}
	}

	return safe_count
}

func isSafe(report []int) bool {
	permutations := generateAllPermutations(report)

	for _, r := range permutations {
		if hasSafeGap(r, 3) && (isReportDecreasing(r) || isReportIncreasing(r)) {
			return true
		}
	}

	return false
}

func generateAllPermutations(r []int) [][]int {
	reports := [][]int{}
	reports = append(reports, [][]int{r}...)

	for i := 0; i < len(r); i++ {
		reportCopy := make([]int, len(r))
		copy(reportCopy, r)

		reportCopy = slices.Delete(reportCopy, i, i+1)
		reports = append(reports, [][]int{reportCopy}...)
	}

	return reports
}
