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
	current, beacons, moves := parseInput(name)

	visited := map[coords.Coord]bool{current: true}

	for _, move := range moves {
		beacon := beacons[move]

		current = coords.NewCoord(
			(current.I+beacon.I)/2,
			(current.J+beacon.J)/2,
		)

		visited[current] = true
	}

	return len(visited)
}

func part2(name string) int {
	current, beacons, moves := parseInput(name)

	visited := map[coords.Coord]bool{current: true}

	for _, move := range moves {
		beacon := beacons[move]

		current = coords.NewCoord(
			(current.I+beacon.I)/2,
			(current.J+beacon.J)/2,
		)

		visited[current] = true
	}

	fireflies := make(map[coords.Coord]bool)
	for k, _ := range visited {
		adj := k.GetAdjacent()

		for _, f := range adj {
			if !visited[f] {
				fireflies[f] = true
			}
		}
	}

	return len(fireflies)
}

func part3(name string) int {
	current, beacons, _ := parseInput(name)

	visited := map[coords.Coord]bool{current: true}

	for {
		visitedCount := len(visited)

		for beetle := range visited {
			for _, beacon := range beacons {
				current:= coords.NewCoord(
					(beetle.I+beacon.I)/2,
					(beetle.J+beacon.J)/2,
				)

				visited[current] = true
			}
		}

		if len(visited) == visitedCount {
			break
		}
	}

	fireflies := make(map[coords.Coord]bool)
	for k, _ := range visited {
		adj := k.GetAdjacent()

		for _, f := range adj {
			if !visited[f] {
				fireflies[f] = true
			}
		}
	}

	return len(fireflies)
}

func parseInput(name string) (coords.Coord, map[byte]coords.Coord, []byte) {
	lines := files.ReadLines(name)
	start := parseCoord(lines[0])
	beacons := make(map[byte]coords.Coord)

	for i := 1; i < len(lines)-1; i++ {
		beacons[lines[i][0]] = parseCoord(lines[i])
	}

	moves := []byte{}

	finalLine := lines[len(lines)-1]
	if !strings.HasPrefix(finalLine, "MOVES") {
		beacons[finalLine[0]] = parseCoord(finalLine)

		return start, beacons, moves
	}

	movesStr := finalLine[strings.Index(finalLine, "=")+1:]
	for i := range movesStr {
		moves = append(moves, movesStr[i])
	}

	return start, beacons, moves
}

func parseCoord(in string) coords.Coord {
	vals := ints.FromStringSlice(
		strings.Split(
			in[strings.Index(in, "[")+1:len(in)-1],
			","),
	)
	return coords.NewCoord(vals[0], vals[1])
}
