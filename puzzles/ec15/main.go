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
	input := strings.Split(files.Read(name), ",")
	wallMap, up, down, left, right, endPosition := produceMap(input)
	startPosition := coords.NewCoord(0, 0)
	visitedMap := map[coords.Coord]int{startPosition: 0}

	queue := []coords.Coord{startPosition}

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		time := visitedMap[next]

		for _, direction := range coords.DirectionsInOrder {
			newPosition:= next.Add(direction)
			if newPosition.Equals(endPosition) {
				return time+1
			}

			if newPosition.I < up || newPosition.I > down || newPosition.J < left || newPosition.J > right {
				continue
			}

			if wallMap[newPosition] {
				continue
			}

			if _, ok := visitedMap[newPosition]; ok {
				continue
			}

			visitedMap[newPosition] = time+1
			queue = append(queue, newPosition)
		}

	}

	panic("exit not found!")
}

func part2(name string) int {
	return part1(name)
}

func part3(name string) int {
	return part1(name)
}

func produceMap(input []string) (map[coords.Coord]bool, int, int, int, int, coords.Coord) {
	left, right, up, down := 0, 0, 0, 0
	startPosition := coords.NewCoord(0, 0)
	currentPosition := startPosition
	currentDirection := coords.NewCoord(-1, 0)
	wallMap := make(map[coords.Coord]bool)

	for _, instruction := range input {
		switch instruction[0] {
		case 'L':
			currentDirection = coords.TurnLeft(currentDirection)
		case 'R':
			currentDirection = coords.TurnRight(currentDirection)
		}

		for range ints.FromString(instruction[1:]) {
			currentPosition = currentPosition.Add(currentDirection)
			wallMap[currentPosition] = true
		}

		if currentPosition.I < up {
			up = currentPosition.I
		}

		if currentPosition.I > down {
			down = currentPosition.I
		}

		if currentPosition.J < left {
			left = currentPosition.J
		}

		if currentPosition.J > right {
			right = currentPosition.J
		}
	}

	delete(wallMap, currentPosition)

	return wallMap, up, down, left, right, currentPosition
}
