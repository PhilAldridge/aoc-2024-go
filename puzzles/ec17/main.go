package main

import (
	"fmt"
	"slices"
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
	grid := files.ReadLines(name)

	volcano := findChar(grid, '@')
	radius := 10

	return calculateDestruction(grid, volcano, radius)
}

func part2(name string) int {
	grid := files.ReadLines(name)

	volcano := findChar(grid, '@')
	radius := 1
	maxDestruction, maxDestructionRadius, prevDestruction := 0, 0, 0

	for radius <= len(grid)/2 {
		destruction := calculateDestruction(grid, volcano, radius)
		roundDestruction := destruction - prevDestruction
		if roundDestruction > maxDestruction {
			maxDestruction = roundDestruction
			maxDestructionRadius = radius
		}

		prevDestruction = destruction
		radius++
	}

	return maxDestruction * maxDestructionRadius
}

type queueType struct {
	state
	score int
}

type state struct {
	position          coords.Coord
	left, right, down bool
}

func part3(name string) int {
	grid := files.ReadLines(name)
	volcano := findChar(grid, '@')
	start := findChar(grid, 'S')

	r := 1
	for {
		availableSquares := getAvailableSquares(volcano, grid, r)

		if len(availableSquares) == 1 {
			panic("path not found")
		}

		queue := []queueType{{
			state: state{position: start},
			score:    0,
		}}

		visitedMap := map[state]int{{
			position: start,
		}: 0}

		maxTime := (r+1) * 30

		for len(queue) > 0 {
			next := queue[0]
			queue = queue[1:]

			for _, direction := range coords.DirectionsInOrder {
				newPos := next.position.Add(direction)

				if _, ok := availableSquares[newPos]; !ok {
					continue
				}

				newState := state{
					position: newPos,
					left:     next.left || (newPos.I == volcano.I && newPos.J < volcano.J),
					right:    next.right || (newPos.I == volcano.I && newPos.J > volcano.J),
					down:     next.down || (newPos.J == volcano.J && newPos.I > volcano.I),
				}

				if newPos.Equals(start) {
					if next.left && next.right && next.down {

						if score, ok := visitedMap[newState]; !ok || score > next.score {
							visitedMap[newState] = next.score
						}
					}
					continue
				}

				newScore := next.score + ints.FromString(string(grid[newPos.I][newPos.J]))

				if newScore > maxTime {
					continue
				}

				if score, ok := visitedMap[newState]; ok && score <= newScore {
					continue
				}

				visitedMap[newState] = newScore
				queue = append(queue, queueType{
					state: newState,
					score:    newScore,
				})
			}

			slices.SortFunc(queue, func(a, b queueType) int {
				return a.score - b.score
			})
		}

		if score := visitedMap[state{
			position: start,
			left: true,
			down: true,
			right: true,
		}]; score > 0 {
			fmt.Println(score,r)
			return score * r
		}

		r++
	}
}

func calculateDestruction(grid []string, volcano coords.Coord, radius int) int {
	total := 0
	squareDistance := radius * radius

	for i, row := range grid {
		for j, char := range row {
			pos := coords.NewCoord(i, j)
			if pos.Equals(volcano) {
				continue
			}

			if coords.PythagoreanSquareDistance(pos, volcano) <= squareDistance {
				total += ints.FromString(string(char))
			}
		}
	}

	return total
}

func getAvailableSquares(volcano coords.Coord, grid []string, radius int) map[coords.Coord]int {
	availableSquares := make(map[coords.Coord]int)
	squareDistance := radius * radius

	for i, row := range grid {
		for j, char := range row {
			pos := coords.NewCoord(i, j)
			if pos.Equals(volcano) {
				continue
			}

			if char == 'S' {
				availableSquares[pos] = 0
				continue
			}

			if coords.PythagoreanSquareDistance(pos, volcano) <= squareDistance {
				continue
			}

			availableSquares[pos] = ints.FromString(string(char))
		}
	}

	return availableSquares
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
