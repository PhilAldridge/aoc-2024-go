package main

import (
	"fmt"
	"slices"
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
	total:= 0

	nails:= parseInput(name)

	min:= ints.Min(nails)

	for _,nail := range nails {
		if nail>min {
			total+= nail-min
		}
	}

	return total
}

func part2(name string) int {
	total:= 0

	nails:= parseInput(name)

	min:= ints.Min(nails)

	for _,nail := range nails {
		if nail>min {
			total+= nail-min
		}
	}

	return total
}


func part3(name string) int {
	total:= 0

	nails:= parseInput(name)
	
	slices.Sort(nails)

	median:= nails[len(nails)/2]

	for _,nail := range nails {
		total += ints.Abs(median-nail)
	}

	return total
}

func parseInput(name string) []int {
	lines:= files.ReadLines(name)
	return ints.FromStringSlice(lines)
}

