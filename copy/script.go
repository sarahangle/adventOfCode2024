package main

import (
	"fmt"
	"github.com/sarahangle/goLib/readFile"
)

func part01_process(input []string) int {
	return 0
}

func part02_process(input []string) int {
	return 0
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
