package main

import (
	"fmt"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/memo"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1 answer: ", part1("input.txt"))
	split := time.Now()
	fmt.Println("Part 2 answer: ", part2("input.txt"))

	fmt.Println()
	fmt.Println("Part 1: ", split.Sub(start))
	fmt.Println("Part 2: ", time.Since(split))
}

func part1(name string) int {
	lines := files.ReadLines(name)

	total := 0

	beamMap := make(map[int]bool)

	for i, char := range lines[0] {
		if char == 'S' {
			beamMap[i] = true
			break
		}
	}

	for i := 1; i < len(lines); i++ {
		for k := range beamMap {
			if lines[i][k] == '^' {
				total++
				delete(beamMap, k)
				if k != 0 {
					beamMap[k-1] = true
				}
				if k+1 != len(lines[0]) {
					beamMap[k+1] = true
				}
			}
		}
	}
	return total
}

func part2(name string) int {
	lines := files.ReadLines(name)

	startX := 0

	for i, char := range lines[0] {
		if char == 'S' {
			startX = i
			break
		}
	}

	type params struct {
		xIndex, yIndex int
	}

	memoTimeLines := memo.MemoRec(func(f func(params) int, p params) int {
		if p.yIndex == len(lines)-1 {
			if lines[p.yIndex][p.xIndex] == '^' {
				return 2
			}
			return 1
		}

		total := 0

	if lines[p.yIndex][p.xIndex] == '^' {
		if p.xIndex != 0 {
			total += f(params{yIndex: p.yIndex+1, xIndex: p.xIndex-1})
		}

		if p.xIndex+1 != len(lines[0]) {
			total += f(params{yIndex: p.yIndex+1, xIndex: p.xIndex+1})
		}

		return total
	} else {
		total += f(params{yIndex:  p.yIndex+1, xIndex: p.xIndex})
	}

	return total
	})

	return memoTimeLines(params{yIndex: 1, xIndex: startX})
}