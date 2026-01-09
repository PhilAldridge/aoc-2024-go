package main

import (
	"fmt"
	"math"
	"slices"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/coords"
	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/ints"
)

type potentialAnswer struct {
	pos   coords.Coord
	score int
	char  rune
}

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
	maze := files.ReadLines(name)

	start, end := getStartEnd(maze)

	scoreMap := traverseMap(maze, start, end)

	return scoreMap[end]
}

func part2(name string) int {
	maze := files.ReadLines(name)

	start, end := getStartEnd(maze)

	scoreMap := traverseMap(maze, start, end)

	return scoreMap[end]
}

func part3(name string) int {
	maze := files.ReadLines(name)

	start, end := getStartEnd(maze)

	score := traverseMapReverse(maze, start, end)

	return score
}

func getHeightDifference(maze []string, a, b coords.Coord) int {
	aChar := maze[a.I][a.J]
	bChar := maze[b.I][b.J]

	var aVal, bVal int
	switch aChar {
	case 'S':
		aVal = 0
	case 'E':
		aVal = 0
	default:
		aVal = int(aChar - '0')
	}
	switch bChar {
	case 'S':
		bVal = 0
	case 'E':
		bVal = 0
	default:
		bVal = int(bChar - '0')
	}

	return ints.ModularDifference(aVal, bVal, 10)
}

func getStartEnd(maze []string) (coords.Coord, coords.Coord) {
	var start, end coords.Coord

	for i, row := range maze {
		for j, char := range row {
			if char == 'S' {
				start = coords.NewCoord(i, j)
			}
			if char == 'E' {
				end = coords.NewCoord(i, j)
			}
		}
	}

	return start, end
}

func traverseMap(maze []string, start, end coords.Coord) map[coords.Coord]int {
	scoreMap := map[coords.Coord]int{
		start: 0,
	}

	for {
		var nextVals []potentialAnswer

		for pos, val := range scoreMap {
			adj := pos.GetAdjacent()

			for _, a := range adj {
				if _, ok := scoreMap[a]; ok || !a.InInput(maze) || maze[a.I][a.J] == '#' || maze[a.I][a.J] == ' ' {
					continue
				}

				nextVals = append(nextVals, potentialAnswer{
					pos:   a,
					score: val + 1 + getHeightDifference(maze, pos, a),
				})
			}
		}

		slices.SortFunc(nextVals, func(a, b potentialAnswer) int {
			return a.score - b.score
		})

		scoreMap[nextVals[0].pos] = nextVals[0].score

		_, ok := scoreMap[end]
		if ok {
			return scoreMap
		}
	}
}

func traverseMapReverse(maze []string, start, end coords.Coord) int {
	scoreMap := map[coords.Coord]int{
		end: 0,
	}

	for {
		prevVal := 0
		nextVal := potentialAnswer{
			score: math.MaxInt,
		}

		for pos, val := range scoreMap {
			if val >= nextVal.score || val + 11 < prevVal {
				continue
			}

			adj := pos.GetAdjacent()

			for _, a := range adj {
				if _, ok := scoreMap[a]; ok || !a.InInput(maze) || maze[a.I][a.J] == '#' || maze[a.I][a.J] == ' ' {
					continue
				}

				score := val + 1 + getHeightDifference(maze, pos, a)

				if score < nextVal.score {
					nextVal = potentialAnswer{
						pos:   a,
						score: val + 1 + getHeightDifference(maze, pos, a),
						char:  rune(maze[a.I][a.J]),
					}
				}
			}
		}

		scoreMap[nextVal.pos] = nextVal.score
		prevVal = nextVal.score
		
		if nextVal.char == 'S' {
			return nextVal.score
		}
	}
}
