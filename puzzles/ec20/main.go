package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/coords"
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
	input := files.ReadLines(name)

	directions := []coords.Coord{
		coords.NewCoord(1, 0),
		coords.NewCoord(0, 1),
	}

	total := 0

	for i, row := range input {
		for j, char := range row {
			if char != 'T' {
				continue
			}

			pos := coords.NewCoord(i, j)

			for _, direction := range directions {
				newPos := pos.Add(direction)

				if !newPos.InInput(input) || !validCrossing(pos, newPos) || input[newPos.I][newPos.J] != 'T' {
					continue
				}

				total++
			}
		}
	}

	return total
}

func part2(name string) int {
	input := files.ReadLines(name)
	start := findChar(input, 'S')
	end := findChar(input, 'E')

	queue := []coords.Coord{start}
	visitedMap := map[coords.Coord]int{
		start: 0,
	}

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		score := visitedMap[next]

		for _, direction := range coords.DirectionsInOrder {
			newPos := next.Add(direction)

			if !newPos.InInput(input) || !validCrossing(next, newPos) ||
				(input[newPos.I][newPos.J] != 'T' && input[newPos.I][newPos.J] != 'E') {
				continue
			}

			if _, ok := visitedMap[newPos]; ok {
				continue
			}

			if newPos.Equals(end) {
				return score + 1
			}

			visitedMap[newPos] = score + 1
			queue = append(queue, newPos)
		}
	}

	panic("no route found")
}

func part3(name string) int {
	input := files.ReadLines(name)

	grids := produceRotations(input)

	start := findChar(input, 'S')

	queue := []state{
		{
			position: start,
			phase:    0,
		},
	}
	visitedMap := map[state]int{
		queue[0]: 0,
	}

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		score := visitedMap[next]

		newState := state{
			position: next.position,
			phase:    (next.phase + 1) % 3,
		}

		if _, ok := visitedMap[newState]; !ok &&
		(grids[newState.phase][next.position.I][next.position.J] == 'E' || grids[newState.phase][next.position.I][next.position.J] == 'T') {
			visitedMap[newState] = score + 1
			queue = append(queue, newState)
		}

		if grids[newState.phase][next.position.I][next.position.J] == 'E' {
			return score + 1
		}

		for _, direction := range coords.DirectionsInOrder {
			newPos := next.position.Add(direction)
			newState := state{
				position: newPos,
				phase:    (next.phase + 1) % 3,
			}

			if !newPos.InInput(input) || !validCrossing(next.position, newPos) ||
				(grids[newState.phase][newPos.I][newPos.J] != 'T' && grids[newState.phase][newPos.I][newPos.J] != 'E') {
				continue
			}

			if _, ok := visitedMap[newState]; ok {
				continue
			}

			if grids[newState.phase][newPos.I][newPos.J] == 'E' {
				return score + 1
			}

			visitedMap[newState] = score + 1
			queue = append(queue, newState)
		}
	}

	panic("no route found")
}

func validCrossing(a, b coords.Coord) bool {
	if a.I == b.I {
		return true
	}

	if a.I < b.I {
		return a.I%2+a.J%2 == 1
	}

	return b.I%2+b.J%2 == 1
}

func findChar(grid []string, character rune) coords.Coord {
	for i, row := range grid {
		for j, char := range row {
			if char == character {
				return coords.NewCoord(i, j)
			}
		}
	}

	panic("no char found")
}

type state struct {
	position coords.Coord
	phase    int
}

func produceRotations(grid []string) [3][]string {
	out := [3][]string{grid}

	out[1] = rotateOnce(grid)

	out[2] = rotateOnce(out[1])

	return out
}

func rotateOnce(grid []string) []string {
	out := make([]string, len(grid))

	for i := range out {
		out[i] = strings.Repeat(".", i)
		for j := range len(grid[0])/2 + 1 - i {
			newI := len(grid) - 1 - j - i
			newJ := ints.Mod(len(grid[0])/2-j+i, len(grid[0]))

			out[i] += string(grid[newI][newJ])
			if newI > 0 {
				out[i] += string(grid[newI-1][newJ])
			}

		}
		out[i] += strings.Repeat(".", i)
	}

	return out
}
