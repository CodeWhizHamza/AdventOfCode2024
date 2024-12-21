package main

import "testing"

func TestSolution2a(t *testing.T) {
	expected := 2
	input := [][]int{{7, 6, 4, 2, 1}, {1, 2, 7, 8, 9}, {9, 7, 6, 2, 1}, {1, 3, 2, 4, 5}, {8, 6, 4, 4, 1}, {1, 3, 6, 7, 9}}
	got := solution_2a(input)

	if got != expected {
		t.Errorf("Got %d, wanted %d", got, expected)
	}
}

func TestIsIncreasing(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	if !isReportIncreasing(input) {
		t.Error("Expected true, got false:", input)
	}

	input = []int{1, 2, 3, 5, 4}
	if isReportIncreasing(input) {
		t.Error("Expected false, got true:", input)
	}

	input = []int{1, 2, 3, 4, 4}
	if isReportIncreasing(input) {
		t.Error("Expected false, got true:", input)
	}
}

func TestIsDecreasing(t *testing.T) {
	input := []int{5, 4, 3, 2, 1}
	if !isReportDecreasing(input) {
		t.Error("Expected true, got false:", input)
	}

	input = []int{1, 2, 3, 5, 4}
	if isReportIncreasing(input) {
		t.Error("Expected false, got true:", input)
	}

	input = []int{5, 4, 3, 2, 2}
	if isReportIncreasing(input) {
		t.Error("Expected false, got true:", input)
	}
}

func TestHasSafeGap(t *testing.T) {
	input := []int{5, 4, 3, 2, 1}
	if !hasSafeGap(input, 3) {
		t.Error("Expected true, got false:", input)
	}

	input = []int{1, 2, 7, 8, 9}
	if hasSafeGap(input, 3) {
		t.Error("Expected false, got true:", input)
	}
}

func TestSolution2b(t *testing.T) {
	expected := 4
	input := [][]int{{7, 6, 4, 2, 1}, {1, 2, 7, 8, 9}, {9, 7, 6, 2, 1}, {1, 3, 2, 4, 5}, {8, 6, 4, 4, 1}, {1, 3, 6, 7, 9}}
	got := Solution2b(input)

	if got != expected {
		t.Errorf("Got %d, wanted %d", got, expected)
	}
}
