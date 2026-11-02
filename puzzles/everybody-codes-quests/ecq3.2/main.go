package main

import (
	"fmt"
	"math"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/coords"
	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/utils"
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

	position, bones := parseInput(name)
	visited := visitedMap{position: true}

	index := 0

	for {
		newPosition := position.Add(path[index])
		index = (index + 1) % len(path)
		if visited[newPosition] {
			continue
		}

		total++

		if _, ok := bones[newPosition]; ok {
			break
		}

		position = newPosition
		visited[position] = true
	}

	return total
}

func part2(name string) int {
	total := 0

	position, bones := parseInput(name)
	visited := visitedMap{position: true}
	for bone, _ := range bones {
		visited[bone] = true
	}

	index := 0

	for {
		newPosition := position.Add(path[index])
		index = (index + 1) % len(path)
		if visited[newPosition] {
			continue
		}

		if _, ok := bones[newPosition]; ok {
			continue
		}
		total++

		position = newPosition
		visited[position] = true

		if visited.markEnclosed(position, bones) || bones.visit(position) {
			break
		}
	}

	return total
}

func part3(name string) int {
	total := 0

	position, bones := parseInput(name)
	visited := visitedMap{position: true}
	for bone, _ := range bones {
		visited[bone] = true
		visited.markEnclosed(bone, bones)
	}

	index := 0

	for {
		newPosition := position.Add(pathPart3[index])
		index = (index + 1) % len(pathPart3)
		if visited[newPosition] {
			continue
		}

		if _, ok := bones[newPosition]; ok {
			continue
		}
		total++

		position = newPosition
		visited[position] = true

		if visited.markEnclosed(position, bones) || bones.visit(position) {
			break
		}
	}

	utils.CallClear()
	visited.print(bones)

	return total
}

var path = []coords.Coord{
	coords.NewCoord(-1, 0),
	coords.NewCoord(0, 1),
	coords.NewCoord(1, 0),
	coords.NewCoord(0, -1),
}

var pathPart3 = []coords.Coord{
	coords.NewCoord(-1, 0),
	coords.NewCoord(-1, 0),
	coords.NewCoord(-1, 0),
	coords.NewCoord(0, 1),
	coords.NewCoord(0, 1),
	coords.NewCoord(0, 1),
	coords.NewCoord(1, 0),
	coords.NewCoord(1, 0),
	coords.NewCoord(1, 0),
	coords.NewCoord(0, -1),
	coords.NewCoord(0, -1),
	coords.NewCoord(0, -1),
}

type visitedMap map[coords.Coord]bool

func (v visitedMap) markEnclosed(a coords.Coord, bones bonesType) bool {
	minY, maxY, minX, maxX := v.bounds()

	for _, dir := range path {
		pos := a.Add(dir)
		if v[pos] {
			continue
		}

		toFill := make(visitedMap)
		if fill(minX, maxX, minY, maxY, v, toFill, pos) {
			for p, _ := range toFill {
				v[p] = true
				if bones.visit(p) {
					return true
				}
			}
		}
	}

	return false
}

func (v visitedMap) bounds() (int, int, int, int) {
	minY := math.MaxInt
	minX := math.MaxInt
	var maxY, maxX int

	for p, _ := range v {
		if p.I < minY {
			minY = p.I
		}
		if p.I > maxY {
			maxY = p.I
		}

		if p.J < minX {
			minX = p.J
		}
		if p.J > maxX {
			maxX = p.J
		}
	}

	return minY, maxY, minX, maxX
}

func (v visitedMap) print(bones bonesType) {
	minY, maxY, minX, maxX := v.bounds()
	for i := minY; i <= maxY; i++ {
		var line string
		for j := minX; j <= maxX; j++ {
			p := coords.NewCoord(i, j)
			if visited, ok := bones[p]; ok {
				if visited.done() {
					line += "@"
				} else {
					line += "X"
				}
				continue
			}

			if v[p] {
				line += "#"
			} else {
				line += " "
			}
		}

		fmt.Println(line)
	}
}

func fill(minX, maxX, minY, maxY int, currentlyVisited, currentlyFilled visitedMap, positionToCheck coords.Coord) bool {
	if positionToCheck.I < minY ||
		positionToCheck.I > maxY ||
		positionToCheck.J < minX ||
		positionToCheck.J > maxX {
		return false
	}

	currentlyFilled[positionToCheck] = true

	for _, dir := range path {
		pos := positionToCheck.Add(dir)
		if !currentlyVisited[pos] && !currentlyFilled[pos] {
			if !fill(minX, maxX, minY, maxY, currentlyVisited, currentlyFilled, pos) {
				return false
			}
		}
	}

	return true
}

type bonesType map[coords.Coord]visitedMap

func (b bonesType) visit(a coords.Coord) bool {
	done := true

	for pos, visited := range b {
		if coords.ManhattanDistance(a, pos) == 1 {
			visited[a] = true
		}
		if len(visited) != 4 {
			done = false
		}
	}

	return done
}

func (v visitedMap) done() bool {
	return len(v) == 4
}

func parseInput(name string) (coords.Coord, bonesType) {
	lines := files.ReadLines(name)

	var start coords.Coord
	bones := make(bonesType)

	for i, line := range lines {
		for j, char := range line {
			switch char {
			case '@':
				start = coords.NewCoord(i, j)
			case '#':
				bones[coords.NewCoord(i, j)] = make(visitedMap)
			}
		}
	}

	for bone, visited := range bones {
		for bone2, _ := range bones {
			if coords.ManhattanDistance(bone, bone2) == 1 {
				visited[bone2] = true
			}
		}

		if coords.ManhattanDistance(bone, start) == 1 {
			visited[start] = true
		}
	}

	return start, bones
}
