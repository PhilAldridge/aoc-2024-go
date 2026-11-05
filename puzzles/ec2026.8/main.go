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

	directionStruct := startingDownRight(ents)

	total, _ := genericRound(ents, directionStruct, false)

	return total
}

func part2(name string) int {
	ents := parseInput(name)

	directions := getDirections(ents)

	total := 0
	for _, d := range directions {
		t, _ := genericRound(ents, d, false)
		total += t
	}

	return total
}

func part3(name string) int {
	total := 0
	grownWarriors := 0
	ents := parseInput(name)
	lastTotal := 0
	lastGrownWarriors := 0
	round := 0
	directions := getDirections(ents)

	for {
		lastTotal, lastGrownWarriors = genericRound(ents, directions[round%len(directions)], true)
		round++
		grownWarriors += lastGrownWarriors
		total += lastTotal
		if grownWarriors >= 100000 {
			return total
		}
	}
}

type position struct {
	i, j int
}
type direction struct {
	moveFunc          func(j, k, sideLength int) (int, int, bool)
	startingPositions []position
	name              string
}

func genericRound(ents [][]int, d direction, growth bool) (int, int) {
	toGrow := []position{}
	warriorsGrown := 0
	energy := 0

	for _, pos := range d.startingPositions {
		toGrow = append(toGrow, findGrowths(ents, pos.i, pos.j, d.moveFunc)...)
	}

	for _, grown := range toGrow {
		if x := ents[grown.i][grown.j]; x >= 9 {
			if growth {
				ents[grown.i][grown.j] = -2
			}

			warriorsGrown++
			energy++
		} else {
			if growth {
				ents[grown.i][grown.j] = x + 1
			}

			if x > 0 {
				energy++
			}
		}
	}

	return energy, warriorsGrown
}

func parseInput(name string) [][]int {
	out := [][]int{}

	for _, row := range files.ReadLines(name) {
		row = strings.ReplaceAll(row, ".", "")
		out = append(out, ints.FromStringSlice(strings.Split(row, "")))
	}

	return out
}

func startingDownRight(ents [][]int) direction {
	out := []position{}
	for i := range ents[0] {
		out = append(out, position{i: 0, j: i})
	}

	for i := 1; i <= len(ents)/2; i++ {
		out = append(out, position{i: i, j: 0})
	}
	return direction{
		startingPositions: out,
		moveFunc:          downRight,
		name:              "down-right",
	}
}

func startingDownLeft(ents [][]int) direction {
	out := []position{}
	for i := range ents[0] {
		out = append(out, position{i: 0, j: i})
	}

	for i := 1; i <= len(ents)/2; i++ {
		out = append(out, position{i: i, j: len(ents[i]) - 1})
	}
	return direction{
		startingPositions: out,
		moveFunc:          downLeft,
		name:              "down-left",
	}
}

func startingLeft(ents [][]int) direction {
	out := []position{}
	for i := range ents {
		out = append(out, position{i: i, j: len(ents[i]) - 1})
	}
	return direction{
		startingPositions: out,
		moveFunc:          left,
		name:              "left",
	}
}

func startingUpLeft(ents [][]int) direction {
	out := []position{}
	for i := range ents[0] {
		out = append(out, position{i: len(ents) - 1, j: i})
	}

	for i := len(ents) / 2; i < len(ents)-1; i++ {
		out = append(out, position{i: i, j: len(ents[i]) - 1})
	}
	return direction{
		startingPositions: out,
		moveFunc:          upLeft,
		name:              "up-left",
	}
}

func startingUpRight(ents [][]int) direction {
	out := []position{}
	for i := range ents[0] {
		out = append(out, position{i: len(ents) - 1, j: i})
	}

	for i := len(ents) / 2; i < len(ents)-1; i++ {
		out = append(out, position{i: i, j: 0})
	}
	return direction{
		startingPositions: out,
		moveFunc:          upRight,
		name:              "up-right",
	}
}

func startingRight(ents [][]int) direction {
	out := []position{}
	for i := range ents {
		out = append(out, position{i: i, j: 0})
	}
	return direction{
		startingPositions: out,
		moveFunc:          right,
		name:              "right",
	}
}

func getDirections(ents [][]int) []direction {
	return []direction{
		startingDownRight(ents),
		startingDownLeft(ents),
		startingLeft(ents),
		startingUpLeft(ents),
		startingUpRight(ents),
		startingRight(ents),
	}
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
		length := sideLength + newI
		if j >= length {
			return 0, 0, false
		}

		return newI, j, true
	}

	return newI, j + 1, true
}

func right(i, j, sideLength int) (int, int, bool) {
	length := sideLength + i
	if i >= sideLength {
		length = sideLength*3 - i - 2
	}
	if j+1 >= length {
		return 0, 0, false
	}

	return i, j + 1, true
}

func findGrowths(ents [][]int, j, k int, moveFun func(j, k, sideLength int) (int, int, bool)) []position {
	ok := true
	shadowHeight := ents[j][k]
	toGrow := []position{{i: j, j: k}}

	for {
		j, k, ok = moveFun(j, k, len(ents[0]))
		if !ok {
			break
		}

		shadowHeight--

		if ents[j][k] < 1 {
			toGrow = append(toGrow, position{i: j, j: k})
			continue
		}

		if ents[j][k] > shadowHeight {
			shadowHeight = ents[j][k]
			toGrow = append(toGrow, position{i: j, j: k})
		}
	}

	return toGrow
}
