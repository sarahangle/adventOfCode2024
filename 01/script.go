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

func makeMap(input []int) (map[int]int) {
	output := make(map[int]int)
	for _, x := range input {
		x_count, x_present := output[x]
		if !x_present {
			output[x] = 1
		} else {
			output[x] = x_count + 1
		}
	}
	return output
}

func part01_process(input []string) int {
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

func part02_process(input []string) int {
	// Turn string input into two unsorted lists of ints
	input_ints := readFile.SplitByWhiteSpaceToInts(input)
	list1, list2 := splitIntoTwoSlices(input_ints)


	list2_map := makeMap(list2)

	// Compare each set
	var sum = 0
	for _, y := range list1 {
		score := y * list2_map[y]
		sum += score
	}

	// Print and return
	return sum
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
