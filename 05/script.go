package main

import (
	"fmt"
	"github.com/sarahangle/goLib/readFile"
)

func part01_process(input []string) int {
	input1, input2 := readFile.SplitByBlankLine(input)
	rules := readFile.SplitByDelimiterToInts(input1, "|")
	updates := readFile.SplitByDelimiterToInts(input2, ",")

	rules_map := make(map[[2]int]bool)
	for _,rule := range rules {
		new_rule := [2]int{rule[0], rule[1]}
		rules_map[new_rule] = true
	}

	sum := 0
	for _,update := range updates {
		update_valid := true
		out:
		for i,earlyPage := range update {
			for _,latePage := range update[i:] {
				check_rule := [2]int{latePage, earlyPage}
				_, check_rule_real := rules_map[check_rule]
				if check_rule_real {
					update_valid = false
					break out
				}
			}
		}
		if update_valid {
			midpoint := len(update)/2
			sum += update[midpoint]
		}
	}

	return sum
}

func part02_process(input []string) int {
	input1, input2 := readFile.SplitByBlankLine(input)
	rules := readFile.SplitByDelimiterToInts(input1, "|")
	updates := readFile.SplitByDelimiterToInts(input2, ",")

	rules_map := make(map[[2]int]bool)
	for _,rule := range rules {
		new_rule := [2]int{rule[0], rule[1]}
		rules_map[new_rule] = true
	}

	var failed_updates [][]int
	for _,update := range updates {
		out:
		for i,earlyPage := range update {
			for _,latePage := range update[i:] {
				check_rule := [2]int{latePage, earlyPage}
				_, check_rule_real := rules_map[check_rule]
				if check_rule_real {
					failed_updates = append(failed_updates, update)
					break out
				}
			}
		}
	}

	sum := 0
	for _,update := range failed_updates {
		iEarlyPage := 0
		for iEarlyPage < len(update)-1 {
			earlyPage := update[iEarlyPage]
			jLatePage := iEarlyPage + 1
			for jLatePage < len(update) {
				latePage := update[jLatePage]
				check_rule := [2]int{latePage, earlyPage}
				_, check_rule_real := rules_map[check_rule]
				if check_rule_real {
					update[iEarlyPage] = latePage
					update[jLatePage] = earlyPage
					earlyPage = update[iEarlyPage]
					jLatePage = iEarlyPage + 1
				} else{
					jLatePage++
				}
			}
			iEarlyPage++
		}
		midpoint := len(update)/2
		sum += update[midpoint]
	}

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
