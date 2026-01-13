package main

import (
	"fmt"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/coords"
	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1 answer: ", part1("input.txt"))
	split := time.Now()
	fmt.Println("Part 2 answer: ", part2("input2.txt"))
	split2 := time.Now()
	fmt.Println("Part 3 answer: ", part3("input3.txt"))

	fmt.Println()
	fmt.Println("Part 1: ", split.Sub(start))
	fmt.Println("Part 2: ", split2.Sub(split))
	fmt.Println("Part 2: ", time.Since(split2))
}

func part1(name string) int {
	grid := parseGrid(name)
	total := 0
	roundScore := 0

	for range 10 {
		grid, roundScore = runRound(grid)
		total += roundScore
	}

	return total
}

func part2(name string) int {
	grid := parseGrid(name)
	total := 0
	roundScore := 0

	for range 2025 {
		grid, roundScore = runRound(grid)
		total += roundScore
	}

	return total
}

func part3(name string) int {
	grid := [34][34]bool{}

	centreGoal := parseGrid(name)
	total := 0
	gridMap := map[[34][34]bool]int{grid: 0}
	start, end := 0, 0

	for {
		grid = runRound34(grid)
		total++
		if prevRound, ok := gridMap[grid]; ok {
			start = prevRound
			end = total
			break
		} else {
			gridMap[grid] = total
		}
	}

	roundsLeft := 1000000000 % (end - start)
	repeats:= (1000000000 / (end - start))
	timesReached := 0
	repeatScore:=0

	for k, v := range gridMap {
		if centreReached(k, centreGoal) {
			score:= 0
			for _,row:= range k {
				for _,char:= range row {
					if char {
						score++
					}
				}
			}

			repeatScore +=score
			if v-start < roundsLeft {
				timesReached+=score
			}
		}
	}

	timesReached += repeats * repeatScore

	return timesReached
}

func centreReached(grid [34][34]bool, centre [][]bool) bool {
	iStart := 17 - len(centre)/2
	jStart := 17 - len(centre[0])/2

	for i, row := range centre {
		for j, char := range row {
			if char != grid[iStart+i][jStart+j] {
				return false
			}
		}
	}

	return true
}

func parseGrid(name string) [][]bool {
	input := files.ReadLines(name)
	output := make([][]bool, len(input))
	for i, row := range input {
		outputRow := make([]bool, len(row))
		for j, char := range row {
			if char == '#' {
				outputRow[j] = true
			}
		}
		output[i] = outputRow
	}

	return output
}

func runRound(input [][]bool) ([][]bool, int) {
	output := make([][]bool, len(input))
	total := 0
	for i, row := range input {
		outputRow := make([]bool, len(row))
		for j, char := range row {
			countEven := diagonalCountEven(coords.NewCoord(i, j), input)
			if (char && !countEven) || (!char && countEven) {
				outputRow[j] = true
				total++
			}
		}
		output[i] = outputRow
	}

	return output, total
}

func runRound34(input [34][34]bool) [34][34]bool {
	output := [34][34]bool{}
	corners := [2]int{-1, 1}
	for i, row := range input {
		for j, char := range row {
			count := 0
			for _, iC := range corners {
				iVal := i + iC
				if iVal < 0 || iVal >= 34 {
					continue
				}

				for _, jC := range corners {
					jVal := j + jC
					if jVal < 0 || jVal >= 34 {
						continue
					}

					if input[iVal][jVal] {
						count++
					}
				}
			}

			countEven := count%2 == 0
			if (char && !countEven) || (!char && countEven) {
				output[i][j] = true
			}
		}
	}

	return output
}

func diagonalCountEven(position coords.Coord, input [][]bool) bool {
	diagonals := position.GetAdjacentDiagonals()
	count := 0

	for _, diagonal := range diagonals {
		if coords.GenericInInput(diagonal, input) && input[diagonal.I][diagonal.J] {
			count++
		}
	}

	return count%2 == 0
}
