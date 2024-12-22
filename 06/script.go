package main

import (
	"fmt"
	"github.com/sarahangle/goLib/readFile"
)

// Make closure to generate function that checks if index is out of bounds for array
func makeOBChecker(input []string) func(int, int) bool {
	return func(x, y int) bool {
		if x < 0 || x >= len(input[0]) {
			return true
		} else if y < 0 || y >= len(input) {
			return true
		}
		return false
	}
}

func part01_process(input []string) int {
	isOB := makeOBChecker(input)

	var field [][]byte
	count := 1
	dir := 0 //0=N, 1=E, 2=S, 3=W
	var loc [2]int
	var next_loc [2]int
	for j,row := range input {
		field = append(field, []byte(row))
		for i,char := range row {
			if char == '^' {
				loc = [2]int{i,j}
				field[j][i] = 'X'
			}
		}
	}

	for !isOB(loc[0], loc[1]) {
		if dir == 0 {
			next_loc = [2]int{loc[0], loc[1]-1}
		} else if dir == 1 {
			next_loc = [2]int{loc[0]+1, loc[1]}
		} else if dir == 2 {
			next_loc = [2]int{loc[0], loc[1]+1}
		} else if dir == 3 {
			next_loc = [2]int{loc[0]-1, loc[1]}
		}

		if isOB(next_loc[0], next_loc[1]) {
			loc = next_loc
			continue
		} else if field[next_loc[1]][next_loc[0]] == '#' {
			dir++
			if dir == 4 {dir = 0}
		} else if field[next_loc[1]][next_loc[0]] == 'X' {
			loc = next_loc
		} else if field[next_loc[1]][next_loc[0]] == '.' {
			field[next_loc[1]][next_loc[0]] = 'X'
			count++
			loc = next_loc
		}
	}

	return count
}

func isLoop(field [][]byte, startPos [3]int, isOB func(int, int) bool) bool {
	haveBeen := make(map[[3]int]bool)
	currentPos := [3]int{startPos[0], startPos[1], startPos[2]}
	var nextPos [3]int

	for !isOB(currentPos[0], currentPos[1]) {
		_,pres := haveBeen[currentPos]
		if pres {return true}

		haveBeen[currentPos] = true
		if currentPos[2] == 0 {
			nextPos = [3]int{currentPos[0], currentPos[1]-1, currentPos[2]}
		} else if currentPos[2] == 1 {
			nextPos = [3]int{currentPos[0]+1, currentPos[1], currentPos[2]}
		} else if currentPos[2] == 2 {
			nextPos = [3]int{currentPos[0], currentPos[1]+1, currentPos[2]}
		} else if currentPos[2] == 3 {
			nextPos = [3]int{currentPos[0]-1, currentPos[1], currentPos[2]}
		}

		if isOB(nextPos[0], nextPos[1]) {
			return false
		} else if field[nextPos[1]][nextPos[0]] == '#' {
			currentPos[2] = currentPos[2]+1
			if currentPos[2] == 4 {currentPos[2] = 0}
		} else {
			currentPos = nextPos
		}
	}
	return false
}

func part02_process(input []string) int {
	isOB := makeOBChecker(input)
	count := 0

	var field [][]byte
	path := make(map[[2]int]bool)
	dir := 0 //0=N, 1=E, 2=S, 3=W
	var start [3]int
	var loc [2]int
	var next_loc [2]int
	for j,row := range input {
		field = append(field, []byte(row))
		for i,char := range row {
			if char == '^' {
				loc = [2]int{i,j}
				field[j][i] = 'X'
				start = [3]int{loc[0], loc[1], 0}
			}
		}
	}

	for !isOB(loc[0], loc[1]) {
		path[loc] = true
		if dir == 0 {
			next_loc = [2]int{loc[0], loc[1]-1}
		} else if dir == 1 {
			next_loc = [2]int{loc[0]+1, loc[1]}
		} else if dir == 2 {
			next_loc = [2]int{loc[0], loc[1]+1}
		} else if dir == 3 {
			next_loc = [2]int{loc[0]-1, loc[1]}
		}

		if isOB(next_loc[0], next_loc[1]) {
			loc = next_loc
			continue
		} else if field[next_loc[1]][next_loc[0]] == '#' {
			dir++
			if dir == 4 {dir = 0}
		} else if field[next_loc[1]][next_loc[0]] == 'X' {
			loc = next_loc
		} else if field[next_loc[1]][next_loc[0]] == '.' {
			field[next_loc[1]][next_loc[0]] = 'X'
			loc = next_loc
		}
	}

	for loc := range path{
		if loc[0] == start[0] && loc[1] == start[1] {continue}
		field[loc[1]][loc[0]] = '#'
		if isLoop(field, start, isOB) {count++}
		field[loc[1]][loc[0]] = '.'
	}

	return count
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
