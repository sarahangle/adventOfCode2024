package main

import (
	"fmt"
	"github.com/sarahangle/goLib/readFile"
)

func abs(a int) int{
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func part01_process(input []string) int {
	// Turn string input into []][]int
	input_ints := readFile.SplitByWhiteSpaceToInts(input)

	// Iterate over each row to check conditions
	var valid_count = 0
	for _, row := range input_ints {
		var allIncreasing = true
		var allDecreasing = true
		var valid = true
		var current = row[0]

		for _, element := range row[1:] {
			if element > current {
				allDecreasing = false
			} else if element < current {
				allIncreasing = false
			}
			if abs(current-element) > 3 || abs(current-element) < 1 {
				valid = false
				break
			}

			current = element
		}

		if !valid {
			continue
		}
		if allDecreasing || allIncreasing {
			valid_count++
		}
	}

	// Print and return
	return valid_count
}

func evalRow(row []int) bool {
	var allIncreasing = true
	var allDecreasing = true
	var valid = true
	var current = row[0]

	for _, element := range row[1:] {
		if element > current {
			allDecreasing = false
		} else if element < current {
			allIncreasing = false
		}
		if abs(current-element) > 3 || abs(current-element) < 1 {
			valid = false
			break
		}

		current = element
	}

	if !valid {
		return false
	}
	if allDecreasing || allIncreasing {
		return true
	}
	return false
}

func part02_process(input []string) int {
	// Turn string input into []][]int
	input_ints := readFile.SplitByWhiteSpaceToInts(input)

	// Iterate over each row to check conditions
	var valid_count = 0
	for _, row := range input_ints {
		if evalRow(row) {
			valid_count++
		} else {
			for i := range row {
				row_copy := make([]int, len(row))
				copy(row_copy, row)
				if evalRow(append(row_copy[:i], row_copy[(i+1):]...)) {
					valid_count++
					//fmt.Println("Safe by deleting i: ", i)
					break
				}
			}
		}
	}

	// Print and return
	return valid_count
}

func testRun() {
	input, err := readFile.ReadLines("./test-input.txt")
	fmt.Println("---Test Data---")
	fmt.Println("Errors? ", err)
	fmt.Println("Result Part01: ", part01_process(input))
	fmt.Println("Result Part02: ", part02_process(input))
	fmt.Println()
}

func realRun() {
	input, err := readFile.ReadLines("./input.txt")
	fmt.Println("---Input Data---")
	fmt.Println("Errors? ", err)
	fmt.Println("Result Part01: ", part01_process(input))
	fmt.Println("Result Part02: ", part02_process(input))
	fmt.Println()
}

func main() {
	fmt.Println()
	testRun()
	realRun()
}
