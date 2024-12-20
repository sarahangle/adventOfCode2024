package main

import (
	"fmt"
	"github.com/sarahangle/goLib/readFile"
)

func create2DArrayChecker(width, height int) func(int, int) bool {
	return func(x, y int) bool {
		if x < 0 || x >= width {
			return false
		} else if y < 0 || y >= height {
			return false
		}
		return true
	}
}

func getPossibleXMASLocs(x, y int) [8][4][2]int {
	var outputLocs = [8][4][2]int {
		{{x,y},{x+1,y},{x+2,y},{x+3,y}},
		{{x,y},{x-1,y},{x-2,y},{x-3,y}},
		{{x,y},{x,y+1},{x,y+2},{x,y+3}},
		{{x,y},{x,y-1},{x,y-2},{x,y-3}},
		{{x,y},{x+1,y+1},{x+2,y+2},{x+3,y+3}},
		{{x,y},{x+1,y-1},{x+2,y-2},{x+3,y-3}},
		{{x,y},{x-1,y+1},{x-2,y+2},{x-3,y+3}},
		{{x,y},{x-1,y-1},{x-2,y-2},{x-3,y-3}},
	}
	return outputLocs
}

func getPossibleAMMSSLocs(x, y int) [4][5][2]int {
	var outputLocs = [4][5][2]int {
		{{x,y},{x+1,y+1},{x+1,y-1},{x-1,y-1},{x-1,y+1}},
		{{x,y},{x+1,y-1},{x-1,y-1},{x-1,y+1},{x+1,y+1}},
		{{x,y},{x-1,y-1},{x-1,y+1},{x+1,y+1},{x+1,y-1}},
		{{x,y},{x-1,y+1},{x+1,y+1},{x+1,y-1},{x-1,y-1}},
	}
	return outputLocs
}

func part01_process(input []string) int {
	// Initialize function to test if coords are in bounds
	inBounds := create2DArrayChecker(len(input[0]), len(input))

	// Look around all Xs for XMAS
	count := 0
	for j,row := range input {
		for i, letter := range row {
			if letter == 'X' {
				XMAS_coords := getPossibleXMASLocs(i,j)
				for _, coord := range XMAS_coords {
					lastx := coord[3][0]
					lasty := coord[3][1]
					if !inBounds(lastx, lasty) {
						continue
					} else if input[coord[1][1]][coord[1][0]] == 'M' &&
							  input[coord[2][1]][coord[2][0]] == 'A' &&
						   	  input[coord[3][1]][coord[3][0]] == 'S' {
						count++
					}
				}
			}
		}
	}

	return count
}

func part02_process(input []string) int {
	// Initialize function to test if coords are in bounds
	inBounds := create2DArrayChecker(len(input[0]), len(input))

	// Look around all As for AMMSS
	count := 0
	for j,row := range input {
		for i, letter := range row {
			if letter == 'A' {
				AMMSS_coords := getPossibleAMMSSLocs(i,j)
				for _, coord := range AMMSS_coords {
					coord_valid := true
					for _, loc := range coord {
						lastx := loc[0]
						lasty := loc[1]
						if !inBounds(lastx, lasty) {
							coord_valid = false
							break
						}
					}
					if !coord_valid {
						continue
					} else if input[coord[1][1]][coord[1][0]] == 'M' &&
							  input[coord[2][1]][coord[2][0]] == 'M' &&
						   	  input[coord[3][1]][coord[3][0]] == 'S' &&
							  input[coord[4][1]][coord[4][0]] == 'S'{
						count++
					}
				}
			}
		}
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
