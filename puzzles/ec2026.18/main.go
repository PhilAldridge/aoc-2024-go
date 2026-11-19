package main

import (
	"fmt"
	"time"

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
	total := 0

	stackPositions := parse(name)

	lCursor, rCursor := 0, len(stackPositions)-1
	lastHeight := 0
	currentHeight, rect := getRect(stackPositions[lCursor], stackPositions[rCursor], lastHeight)
	total += rect

	for lCursor < rCursor {
		for stackPositions[lCursor].y <= currentHeight && lCursor < rCursor {
			lCursor++
		}

		for stackPositions[rCursor].y <= currentHeight && lCursor < rCursor {
			rCursor--
		}

		lastHeight = currentHeight
		currentHeight, rect = getRect(stackPositions[lCursor], stackPositions[rCursor], lastHeight)
		total += rect
	}

	return total
}

func part2(name string) int {
	return part1(name)
}

func part3(name string) int {
	return part1(name)
}

type keyType struct {
	x, y  int
	width int
}

func newKey(vals []int) keyType {
	if len(vals) < 3 {
		vals = append(vals, 1)
	}

	return keyType{
		x: vals[0], y: vals[1], width: vals[2],
	}
}

func parse(name string) []keyType {
	out := []keyType{}

	for _, row := range files.ReadLines(name) {
		vals := ints.FromDelimitedString(row, " ")

		out = append(out, newKey(vals))
	}

	return out
}

func getRect(l, r keyType, prevHeight int) (newHeight int, rect int) {
	if l.y < r.y {
		newHeight = l.y
	} else {
		newHeight = r.y
	}

	if newHeight < prevHeight {
		panic("height went down")
	}

	rect = (r.x - l.x + r.width) * (newHeight - prevHeight)

	return newHeight, rect
}
