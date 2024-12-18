package main

import (
	"fmt"
	"github.com/sarahangle/goLib/readFile"
)

func showProcessedInput(data []string) {
	fmt.Println("Processed input: ")
	for _, row := range data {
        fmt.Println("  ", row)
    }
}

func process(input []string) int {
	return 0
}

func testRun() {
	input, err := readFile.ReadLines("./test-input.txt")
	fmt.Println("---Test Data---")
	fmt.Println("Errors? ", err)
	//showProcessedInput(input)
	fmt.Println("Result: ", process(input))
}

func realRun() {
	input, err := readFile.ReadLines("./input.txt")
	fmt.Println("---Input Data---")
	fmt.Println("Errors? ", err)
	//showProcessedInput(input)
	fmt.Println("Result: ", process(input))
}

func main() {
	fmt.Println()
	testRun()
	fmt.Println()
	realRun()
	fmt.Println()
}
