package main

import (
	"fmt"
	"slices"
	"strings"
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
	lines:= files.ReadLines(name)

	total := 0

	for _, line:= range lines {
		total += parseAndSort(line)
	}

	return total
}

func part2(name string) int {
	return part1(name)
}

func part3(name string) int {
	return part1(name)
}

func parseAndSort(line string) int {
	comparisons:= strings.Split(line, " ")
	out:= []byte{'1','2','3','4','5'}
	
	swaps:= true

	for swaps {
		swaps = false

		for _, comparison:= range comparisons {
			aIndex:= slices.Index(out, comparison[0])
			if aIndex == -1 {
				aIndex = len(out)
				out = append(out, comparison[0])
			}

			bIndex:= slices.Index(out, comparison[2])
			if bIndex == -1 {
				bIndex = len(out)
				out = append(out, comparison[2])
			}

			if comparison[1] == '>' {
				if aIndex > bIndex {
					swaps = true
					out[aIndex] = comparison[2]
					out[bIndex] = comparison[0]
				}
			} else {
				if bIndex > aIndex {
					swaps = true
					out[aIndex] = comparison[2]
					out[bIndex] = comparison[0]
				}
			}
		}
	}

	return ints.FromString(string(out))
}