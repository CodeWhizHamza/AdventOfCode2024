package main

import "math"

func isReportIncreasing(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		cur := report[i]
		next := report[i+1]

		if cur >= next {
			return false
		}
	}
	return true
}

func isReportDecreasing(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		cur := report[i]
		next := report[i+1]

		if cur <= next {
			return false
		}
	}
	return true
}

func hasSafeGap(report []int, gap int) bool {
	for i := 0; i < len(report)-1; i++ {
		if math.Abs(float64(report[i])-float64(report[i+1])) > float64(gap) {
			return false
		}
	}
	return true
}

func solution_2a(reports [][]int) int {
	safe_count := 0

	for _, report := range reports {
		if hasSafeGap(report, 3) && (isReportIncreasing(report) || isReportDecreasing(report)) {
			safe_count++
		}
	}

	return safe_count
}
