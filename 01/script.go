package main

import (
	"fmt"
	"github.com/sarahangle/goLib/readFile"
	"sort"
)

func splitIntoTwoSlices(input [][]int) ([]int, []int) {
	var list1, list2 []int
	for _, row := range input {
        list1 = append(list1, row[0])
		list2 = append(list2, row[1])
    }
	return list1, list2
}

func process(input []string) int {
	// Turn string input into two unsorted lists of ints
	input_ints := readFile.SplitByWhiteSpaceToInts(input)
	list1, list2 := splitIntoTwoSlices(input_ints)

	// Sort lists
	sort.Ints(list1)
	sort.Ints(list2)

	// Compare each set
	var sum = 0
	for i := range list1 {
		var difference = max(list2[i] - list1[i], list1[i] - list2[i])
		sum += difference
	}

	// Print and return
	return sum
}

func testRun() {
	input, err := readFile.ReadLines("./test-input.txt")
	fmt.Println("---Test Data---")
	fmt.Println("Errors? ", err)
	fmt.Println("Result: ", process(input))
	fmt.Println()
}

func realRun() {
	input, err := readFile.ReadLines("./input.txt")
	fmt.Println("---Input Data---")
	fmt.Println("Errors? ", err)
	fmt.Println("Result: ", process(input))
	fmt.Println()
}

func main() {
	fmt.Println()
	testRun()
	realRun()
}
