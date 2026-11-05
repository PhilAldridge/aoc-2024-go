package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/ints"
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
	ents := parseInput(name)

	return countEnergyDownRight(ents)
}

func part2(name string) int {
	ents := parseInput(name)

	return countEnergyDownRight(ents) +
		countEnergyDownLeft(ents) +
		countEnergyLeft(ents) + 
		countEnergyUpLeft(ents) + 
		countEnergyUpRight(ents) + 
		countEnergyRight(ents)
}

func part3(name string) int {
	total := 0

	return total
}

func parseInput(name string) [][]int {
	out := [][]int{}

	for _, row := range files.ReadLines(name) {
		row = strings.ReplaceAll(row, ".", "")
		out = append(out, ints.FromStringSlice(strings.Split(row, "")))
	}

	return out
}

func countEnergyDownRight(ents [][]int) int {
	total := 0
	for i := range ents[0] {
		total += countEnergyGivenStartAndDirection(ents, 0, i, downRight)
	}

	for i := 1; i <= len(ents)/2; i++ {
		total += countEnergyGivenStartAndDirection(ents, i, 0, downRight)
	}
	return total
}

func countEnergyDownLeft(ents [][]int) int {
	total := 0
	for i := range ents[0] {
		total += countEnergyGivenStartAndDirection(ents, 0, i, downLeft)
	}

	for i := 1; i <= len(ents)/2; i++ {
		total += countEnergyGivenStartAndDirection(ents, i, len(ents[i])-1, downLeft)
	}
	return total
}

func countEnergyLeft(ents [][]int) int {
	total := 0
	for i := range ents {
		total += countEnergyGivenStartAndDirection(ents, i, len(ents[i])-1, left)
	}
	return total
}

func countEnergyUpLeft(ents [][]int) int {
	total := 0
	for i := range ents[0] {
		total += countEnergyGivenStartAndDirection(ents, len(ents)-1, i, upLeft)
	}

	for i := len(ents)/2; i < len(ents)-1; i++ {
		total += countEnergyGivenStartAndDirection(ents, i, len(ents[i])-1, upLeft)
	}
	return total
}

func countEnergyUpRight(ents [][]int) int {
	total := 0
	for i := range ents[0] {
		total += countEnergyGivenStartAndDirection(ents, len(ents)-1, i, upRight)
	}

	for i := len(ents)/2; i < len(ents)-1; i++ {
		total += countEnergyGivenStartAndDirection(ents, i, 0, upRight)
	}
	return total
}

func countEnergyRight(ents [][]int) int {
	total := 0
	for i := range ents {
		total += countEnergyGivenStartAndDirection(ents, i, 0, right)
	}
	return total
}

func downRight(i, j, sideLength int) (int, int, bool) {
	newI := i + 1
	if newI+1 >= sideLength*2 {
		return 0, 0, false
	}

	if newI < sideLength {
		return newI, j + 1, true
	}

	length := sideLength*3 - newI - 2
	if j >= length {
		return 0, 0, false
	}

	return newI, j, true
}

func downLeft(i, j, sideLength int) (int, int, bool) {
	newI := i + 1
	if newI+1 >= sideLength*2 {
		return 0, 0, false
	}

	if newI < sideLength {
		return newI, j, true
	}
	if j < 1 {
		return 0, 0, false
	}

	return newI, j - 1, true
}

func left(i, j, sideLength int) (int, int, bool) {
	if j < 1 {
		return 0, 0, false
	}

	return i, j - 1, true
}

func upLeft(i, j, sideLength int) (int, int, bool) {
	newI := i - 1
	if newI < 0 {
		return 0, 0, false
	}

	if newI < sideLength-1 {
		if j < 1 {
			return 0, 0, false
		}

		return newI, j - 1, true
	}

	return newI, j, true
}

func upRight(i, j, sideLength int) (int, int, bool) {
	newI := i - 1
	if newI < 0 {
		return 0, 0, false
	}

	if newI < sideLength-1 {
		length := sideLength+newI
		if j >= length {
			return 0, 0, false
		}

		return newI, j, true
	}

	return newI, j + 1, true
}

func right(i, j, sideLength int) (int, int, bool) {
	length := sideLength+i
	if i >= sideLength {
		length = sideLength*3 - i - 2
	}
	if j+1 >= length {
		return 0, 0, false
	}

	return i, j + 1, true
}

func countEnergyGivenStartAndDirection(ents [][]int, j, k int, moveFun func(j, k, sideLength int) (int, int, bool)) int {
	total := 1
	ok := true
	shadowHeight := ents[j][k]
	for {
		j, k, ok = moveFun(j, k, len(ents[0]))
		if !ok {
			break
		}

		shadowHeight--
		if ents[j][k] > shadowHeight {
			shadowHeight = ents[j][k]
			total++
		}
	}

	return total
}
