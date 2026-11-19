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
	grid := files.ReadLines(name)
	fill := floodFill(grid, coords.NewCoord(0, 0))

	return fill[coords.NewCoord(len(grid)-1, len(grid[0])-1)]
}

func part2(name string) int {
	grid := files.ReadLines(name)

	return getPerimeter(grid)
}

func part3(name string) int {
	grid := files.ReadLines(name)

	perimeter := getPerimeter(grid)

	left := followLeftEdge(grid)

	right := perimeter - left

	if left < right {
		return left
	}

	return right
}

func floodFill(grid []string, start coords.Coord) map[coords.Coord]int {
	dist := 0
	queue := map[coords.Coord]bool{start: true}
	fill := map[coords.Coord]int{start: dist}

	for len(queue) > 0 {
		dist++
		next := make(map[coords.Coord]bool)

		for pos := range queue {
			for _, adj := range pos.GetAdjacent() {
				if _, ok := fill[adj]; ok {
					continue
				}

				if adj.InInput(grid) && grid[adj.I][adj.J] == '#' {
					fill[adj] = dist
					next[adj] = true
				}
			}
		}

		queue = next
	}

	return fill
}

func getPerimeter(grid []string) int {
	total := 0
	for i, row := range grid {
		for j, char := range row {
			if char == '#' {
				for _, adj := range coords.NewCoord(i, j).GetAdjacent() {
					if !adj.InInput(grid) || grid[adj.I][adj.J] != '#' {
						total++
					}
				}
			}
		}
	}
	return total
}

func followLeftEdge(grid []string) int {
	total := 0
	pos := coords.NewCoord(0, 0)
	direction := coords.Down

	for pos.I != len(grid) || pos.J != len(grid[0]) {
		if validPath(pos, direction, grid) {
			total++
			pos = pos.Add(direction)
			continue
		}

		direction = coords.TurnLeft(direction)
		if validPath(pos, direction, grid) {
			total++
			pos = pos.Add(direction)
			continue
		}

		direction = coords.TurnBack(direction)
		if validPath(pos, direction, grid) {
			total++
			pos = pos.Add(direction)
			continue
		}

		panic("lost")
	}

	return total
}

func validPath(pos, direction coords.Coord, grid []string) bool {
	switch direction {
	case coords.Down:
		return isMagma(pos, grid) && !isMagma(pos.Left(1), grid)
	case coords.Left:
		return isMagma(pos.Left(1), grid) && !isMagma(pos.Add(coords.NewCoord(-1, -1)), grid)
	case coords.Right:
		return !isMagma(pos, grid) && isMagma(pos.Up(1), grid)
	case coords.Up:
		return !isMagma(pos.Up(1), grid) && isMagma(pos.Add(coords.NewCoord(-1, -1)), grid)
	default:
		panic("direction not found")
	}
}

func isMagma(p coords.Coord, grid []string) bool {
	return p.InInput(grid) && grid[p.I][p.J] == '#'
}
