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

func part1(name string) string {
	maze, key := parseInput(name)

	width := len(maze[0])
	height := len(maze)

	index := 0

	for i, row := range maze {
		for j := range row {
			pos := coords.NewCoord(i, j)
			if len(getNeighbours(pos, width, height)) != 8 {
				continue
			}

			if key[index] == 'L' {
				maze = rotateLeft(maze, pos)
			} else {
				maze = rotateRight(maze, pos)
			}
			index = (index + 1) % len(key)
		}
	}

	return getDecodedString(maze)
}

func part2(name string) string {
	maze, key := parseInput(name)

	width := len(maze[0])
	height := len(maze)
	for range 100 {
		index := 0

		for i, row := range maze {
			for j := range row {
				pos := coords.NewCoord(i, j)
				if len(getNeighbours(pos, width, height)) != 8 {
					continue
				}

				if key[index] == 'L' {
					maze = rotateLeft(maze, pos)
				} else {
					maze = rotateRight(maze, pos)
				}
				index = (index + 1) % len(key)
			}
		}
	}

	return getDecodedString(maze)
}

func part3(name string) string {
	maze, key := parseInput(name)

	width := len(maze[0])
	height := len(maze)
	for range 1048576000 {
		index := 0

		for i, row := range maze {
			for j := range row {
				pos := coords.NewCoord(i, j)
				if len(getNeighbours(pos, width, height)) != 8 {
					continue
				}

				if key[index] == 'L' {
					maze = rotateLeft(maze, pos)
				} else {
					maze = rotateRight(maze, pos)
				}
				index = (index + 1) % len(key)
			}
		}

		answer, ok:= checkDecodedString(maze)
		if ok {
			return answer
		}
	}

	return getDecodedString(maze)
}

func rotateRight(maze [][]rune, locus coords.Coord) [][]rune {
	neighbours := getNeighbours(locus, len(maze[0]), len(maze))
	prevPos := neighbours[len(neighbours)-1]
	prevChar := maze[prevPos.I][prevPos.J]

	for _, pos := range neighbours {
		nextChar := maze[pos.I][pos.J]
		maze[pos.I][pos.J] = prevChar
		prevChar = nextChar
	}

	return maze
}

func rotateLeft(maze [][]rune, locus coords.Coord) [][]rune {
	neighbours := getNeighbours(locus, len(maze[0]), len(maze))
	prevPos := neighbours[0]
	prevChar := maze[prevPos.I][prevPos.J]

	for i := len(neighbours) - 1; i >= 0; i-- {
		pos := neighbours[i]
		nextChar := maze[pos.I][pos.J]
		maze[pos.I][pos.J] = prevChar
		prevChar = nextChar
	}

	return maze
}

func getNeighbours(locus coords.Coord, width, height int) []coords.Coord {
	adjacents := locus.GetAdjacentIncludingDiagonals()

	var toMove []coords.Coord

	for _, adj := range adjacents {
		if adj.I < 0 || adj.J < 0 || adj.I >= height || adj.J >= width {
			continue
		}
		toMove = append(toMove, adj)
	}

	return toMove
}

func parseInput(name string) ([][]rune, []rune) {
	input := files.ReadLinesAsRunes(name)

	key := input[0]

	maze := input[2:]

	return maze, key
}

func printMaze(maze [][]rune) {
	for _, row := range maze {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func getDecodedString(maze [][]rune) string {
	answer := ""

	for _, row := range maze {
		startFound := false
		for _, char := range row {
			if !startFound {
				if char == '>' {
					startFound = true
				}
				continue
			}

			if char == '<' {
				return answer
			}

			answer += string(char)
		}
	}
	printMaze(maze)
	panic("output incorrect")
}

func checkDecodedString(maze [][]rune) (string,bool) {
	answer := ""

	for _, row := range maze {
		startFound := false
		for _, char := range row {
			if !startFound {
				if char == '>' {
					startFound = true
				}
				continue
			}

			if char == '<' {
				return answer,true
			}

			if char < '0' || char > '9' {
				return "", false
			}

			answer += string(char)
		}
		if startFound {
			break
		}
	}
	return "",false
}
