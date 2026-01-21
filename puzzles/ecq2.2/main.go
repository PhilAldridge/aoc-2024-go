package main

import (
	"fmt"
	"strings"
	"time"

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
	boltOrder := []byte{'R', 'G', 'B'}
	balloons := files.Read(name)
	total := 0
	i := 0
	for i < len(balloons) {
		boltColour := boltOrder[total%len(boltOrder)]
		for i < len(balloons) {
			if balloons[i] != boltColour {
				break
			}

			i++
		}

		total++
		i++
	}

	return total
}

func part2(name string) int {

	balloons := strings.Repeat(files.Read(name), 100)
	return countPops(balloons, 0)
}

func part3(name string) int {

	balloons := strings.Repeat(files.Read(name), 100000)
	return countPops(balloons, 0)
}

func countPops(balloons string, startIndex int) int {
	if len(balloons) <=1 {
		return len(balloons)
	}

	balloonsLeft := []byte{}

	halfIndex := (len(balloons) + 1) / 2
	boltIndex := startIndex
	bolts := []byte{'R', 'G', 'B'}

	secondHalfDoubleIndex := len(balloons)

	for i := 0; i < halfIndex; i++ {
		if secondHalfDoubleIndex%2 != 0 {
			secondHalfDoubleIndex++
			boltIndex = (boltIndex + 1) % 3
			continue
		}

		if balloons[i] == bolts[boltIndex] {
			secondHalfDoubleIndex += 2
			boltIndex = (boltIndex + 1) % 3
			continue
		}

		balloonsLeft = append(balloonsLeft, balloons[secondHalfDoubleIndex/2])
		secondHalfDoubleIndex++
		boltIndex = (boltIndex + 1) % 3
	}

	remaining:= string(balloonsLeft)
	finalIndex:= (secondHalfDoubleIndex+1)/2
	if finalIndex < len(balloons) {
		remaining += balloons[finalIndex:]
	}

	return halfIndex + countPops(remaining, boltIndex)
}
