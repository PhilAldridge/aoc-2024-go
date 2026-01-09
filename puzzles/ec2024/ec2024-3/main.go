package main

import (
	"fmt"
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
	mapping:= firstLevel(name)

	i:=2
	for dig(mapping, i, canDig) {
		i++
	}

	return ints.SumMap(mapping)
}

func part2(name string) int {
	mapping:= firstLevel(name)

	i:=2
	for dig(mapping, i, canDig) {
		i++
	}

	return ints.SumMap(mapping)
}


func part3(name string) int {
	mapping:= firstLevel(name)

	i:=2
	for dig(mapping, i, canDigIncludingDiagonals) {
		i++
	}

	return ints.SumMap(mapping)
}

func firstLevel(name string) map[coords.Coord]int {
	lines:= files.ReadLines(name)
	result:= make(map[coords.Coord]int)

	for i, line:= range lines {
		for j,char:= range line {
			if char == '#' {
				result[coords.NewCoord(i,j)] = 1
			}
		}
	}

	return result
}

func dig(level map[coords.Coord]int, depth int, canDig func(map[coords.Coord]int,int,coords.Coord) bool) (bool) {
	dug:= false
	for coord, val:= range level {
		if val >= depth {
			continue
		}

		if canDig(level, depth, coord) {
			level[coord] = depth
			dug = true
		}
	}

	return dug
}

func canDig(level map[coords.Coord]int, depth int, coord coords.Coord) bool {
	adjacents:= coord.GetAdjacent()

	for _, adjacent:= range adjacents {
		if level[adjacent] < depth -1 {
			return false
		}
	}

	return true
}

func canDigIncludingDiagonals(level map[coords.Coord]int, depth int, coord coords.Coord) bool {
	adjacents:= coord.GetAdjacentIncludingDiagonals()

	for _, adjacent:= range adjacents {
		if level[adjacent] < depth -1 {
			return false
		}
	}

	return true
}