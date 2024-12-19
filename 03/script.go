package main

import (
	"fmt"
	"github.com/sarahangle/goLib/readFile"
	"strings"
	"regexp"
	"strconv"
)

func SplitByMulSpace(input []string) ([]string) {
	var output []string
	for _, row := range input {
		output = append(output, strings.Split(row, "mul")...)
	}
	return output
}

func GetDoInstructions(input []string) (string) {
	var combinedInput string
	for _, row := range input {
		combinedInput += row
	}

	var doInstructions string
	var next_dont = strings.Index(combinedInput, "don't()")
	var next_do = 0
	for next_dont > -1 {
		doInstructions += combinedInput[:next_dont]
		combinedInput = combinedInput[next_dont+7:]
		next_do = strings.Index(combinedInput, "do()")
		if next_do == -1 {
			break
		} else {
			combinedInput = combinedInput[next_do+4:]
			next_dont = strings.Index(combinedInput, "don't()")
			if next_dont == -1 {
				doInstructions += combinedInput
			}
		}
	}

	return doInstructions
}

func part01_process(input []string) int {
	// Turn string input into []string split by 'mul'
	possible_instructions := SplitByMulSpace(input)

	sum := 0

	// Check each instruction against regex
	for _, instruction := range possible_instructions {
		isMatch, _ := regexp.MatchString("^\\((\\d*,\\d*)\\)[\\s\\S]*", instruction)
		if isMatch {
			instruction = instruction[:strings.Index(instruction, ")")]
			instruction = instruction[1:]
			nums := strings.Split(instruction, ",")
			numa, _ := strconv.Atoi(nums[0])
			numb, _ := strconv.Atoi(nums[1])
			sum += numa * numb
		}
	}

	// Print and return
	return sum
}

func part02_process(input []string) int {
	// Turn string input into []string split by 'mul'
	var possible_instructions []string
	possible_instructions = append(possible_instructions, GetDoInstructions(input))

	// Turn string input into []string split by 'mul'
	possible_instructions = SplitByMulSpace(possible_instructions)

	sum := 0

	// Check each instruction against regex
	for _, instruction := range possible_instructions {
		isMatch, _ := regexp.MatchString("^\\((\\d*,\\d*)\\)[\\s\\S]*", instruction)
		if isMatch {
			instruction = instruction[:strings.Index(instruction, ")")]
			instruction = instruction[1:]
			nums := strings.Split(instruction, ",")
			numa, _ := strconv.Atoi(nums[0])
			numb, _ := strconv.Atoi(nums[1])
			sum += numa * numb
		}
	}

	// Print and return
	return sum

	// Print and return
	fmt.Println(possible_instructions)
	return 0
}

func testRun() {
	input, err := readFile.ReadLines("./test-input.txt")
	fmt.Println("---Test Data---")
	fmt.Println("Errors? ", err)
	fmt.Println("Result Part01: ", part01_process(input))
	input2, _ := readFile.ReadLines("./test-input2.txt")
	fmt.Println("Result Part02: ", part02_process(input2))
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
