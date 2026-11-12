package main

import (
	"encoding/json"
	"fmt"
	"math"
	"slices"
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
	total := 0

	stacks := parseInput(name)

	for i, stack := range stacks {
		if i == 25 {
			fmt.Print("f")
		}
		scoreMap := map[string]int{"[]": 0}
		score, _ := reflect(stack.positions, math.MaxInt, scoreMap)

		fmt.Println(i, score, stack.positions[0])
		total += score
	}

	return total
}

func part2(name string) int {
	return part1(name)
}

func part3(name string) int {
	return part1(name)
}

type stack struct {
	positions []coords.Coord
}

func lowerBound(ps []coords.Coord) int {
	is, js := make(map[int]bool), make(map[int]bool)

	for _, p := range ps {
		is[p.I] = true
		js[p.J] = true
	}

	if len(is) < len(js) {
		return len(is)
	}

	return len(js)
}

func reflect(positions []coords.Coord, upperBound int, scoreMap map[string]int) (int, bool) {
	positionsBytes, _ := json.Marshal(positions)
	if v, ok := scoreMap[string(positionsBytes)]; ok {
		return v, true
	}

	if lowerBound(positions) >= upperBound {
		return 0, false
	}

	maxRemovals := getMaxRemovals(positions)

	score := upperBound
	scoreFound := false

	for _, removals := range maxRemovals {
		newPos := []coords.Coord{}
		for _, p := range positions {
			if !slices.ContainsFunc(removals, func(pos coords.Coord) bool { return pos.Equals(p) }) {
				newPos = append(newPos, p)
			}
		}

		next, ok := reflect(newPos, score-1, scoreMap)
		if !ok {
			continue
		}
		newPosBytes, _ := json.Marshal(newPos)
		scoreMap[string(newPosBytes)] = next
		scoreFound = true
		newScore := 1 + next
		if newScore == lowerBound(positions) {
			scoreMap[string(positionsBytes)] = newScore
			return newScore, true
		}
		if newScore < score {
			score = newScore
		}
	}

	scoreMap[string(positionsBytes)] = score
	return score, scoreFound
}

func getMaxRemovals(positions []coords.Coord) [][]coords.Coord {
	maxRemovals := [][]coords.Coord{}
	length := 0

	triedH := []coords.Coord{}
	for _, pos := range positions {
		if slices.ContainsFunc(triedH, func(p coords.Coord) bool { return p.Equals(pos) }) {
			continue
		}

		removals := []coords.Coord{}
		newPos := pos
		for slices.ContainsFunc(positions, func(pos coords.Coord) bool { return pos.Equals(newPos) }) {
			removals = append(removals, newPos)
			newPos = newPos.Left(1)
		}

		newPos = pos.Right(1)
		for slices.ContainsFunc(positions, func(pos coords.Coord) bool { return pos.Equals(newPos) }) {
			removals = append(removals, newPos)
			newPos = newPos.Right(1)
		}

		triedH = append(triedH, removals...)

		if len(removals) == length {
			maxRemovals = append(maxRemovals, removals)
			continue
		}

		if len(removals) > length {
			length = len(removals)
			maxRemovals = [][]coords.Coord{removals}
		}
	}

	triedV := []coords.Coord{}
	for _, pos := range positions {
		if slices.ContainsFunc(triedV, func(p coords.Coord) bool { return p.Equals(pos) }) {
			continue
		}

		removals := []coords.Coord{}
		newPos := pos
		for slices.ContainsFunc(positions, func(pos coords.Coord) bool { return pos.Equals(newPos) }) {
			removals = append(removals, newPos)
			newPos = newPos.Up(1)
		}

		newPos = pos.Down(1)
		for slices.ContainsFunc(positions, func(pos coords.Coord) bool { return pos.Equals(newPos) }) {
			removals = append(removals, newPos)
			newPos = newPos.Down(1)
		}

		triedV = append(triedV, removals...)

		if len(removals) == length {
			maxRemovals = append(maxRemovals, removals)
			continue
		}

		if len(removals) > length {
			length = len(removals)
			maxRemovals = [][]coords.Coord{removals}
		}
	}

	return maxRemovals
}

func length(r [][]coords.Coord) int {
	if len(r) == 0 {
		return 0
	}

	return len(r[0])
}

func getRemovals(p coords.Coord, positions []coords.Coord) [][]coords.Coord {
	out := [][]coords.Coord{{}, {}}

	newPos := p
	for slices.ContainsFunc(positions, func(pos coords.Coord) bool { return pos.Equals(newPos) }) {
		out[0] = append(out[0], newPos)
		newPos = newPos.Up(1)
	}

	newPos = p.Down(1)
	for slices.ContainsFunc(positions, func(pos coords.Coord) bool { return pos.Equals(newPos) }) {
		out[0] = append(out[0], newPos)
		newPos = newPos.Down(1)
	}

	newPos = p
	for slices.ContainsFunc(positions, func(pos coords.Coord) bool { return pos.Equals(newPos) }) {
		out[1] = append(out[1], newPos)
		newPos = newPos.Left(1)
	}

	newPos = p.Right(1)
	for slices.ContainsFunc(positions, func(pos coords.Coord) bool { return pos.Equals(newPos) }) {
		out[1] = append(out[1], newPos)
		newPos = newPos.Right(1)
	}

	return out
}

func parseInput(name string) []stack {
	grid := files.ReadLinesAsRunes(name)
	out := []stack{}

	for i, row := range grid {
		for j, char := range row {
			if char != '#' {
				continue
			}
			newPos := coords.NewCoord(i, j)

			if slices.ContainsFunc(out, func(s stack) bool {
				return slices.ContainsFunc(s.positions, func(p coords.Coord) bool {
					return p.Equals(newPos)
				})
			}) {
				continue
			}

			newPositions := []coords.Coord{newPos}
			queue := []coords.Coord{}
			for _, adj := range newPos.GetAdjacent() {
				queue = append(queue, adj)
			}
			for len(queue) > 0 {
				nextQueue := []coords.Coord{}

				for _, p := range queue {
					if p.I >= 0 &&
						p.I < len(grid) &&
						p.J >= 0 &&
						p.J < len(grid[0]) &&
						grid[p.I][p.J] == '#' &&
						!slices.ContainsFunc(newPositions, func(pos coords.Coord) bool {
							return pos.Equals(p)
						}) {
						newPositions = append(newPositions, p)

						for _, adj := range p.GetAdjacent() {
							nextQueue = append(nextQueue, adj)
						}
					}
				}

				queue = nextQueue
			}

			out = append(out, stack{positions: newPositions})
		}
	}

	return out
}
