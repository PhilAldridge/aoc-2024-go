package main

import (
	"fmt"
	"slices"
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
	tablets := []byte(files.Read(name))
	flips, _ := sort(tablets)
	return flips
}

func part2(name string) int {
	tablets := []byte(strings.Repeat(files.Read(name), 10))
	_, elements := sort(tablets)
	return elements
}

func part3(name string) int {
	tablets := []byte(strings.Repeat(files.Read(name), 100))
	_, elements := sort(tablets)
	return elements
}

// Flips is number of times we perform a flip
// Elements is a count of times we move a tablet
func sort(slice []byte) (flips int, elements int) {
	for len(slice) == 0 {
		// Split into alternating groups of current shade/not current shade
		groups := collectGroups(slice)

		// Get flips/elements needed for this shade and part of slice that still needs to be sorted
		groupFlips, groupElements, newSliceSection := sortGroups(groups)
		flips += groupFlips
		elements += groupElements
		slice = newSliceSection
	}

	return flips, elements
}

func collectGroups(slice []byte) [][]byte {
	current := slice[0]
	groups := [][]byte{}
	isCurrent := true
	lCursor:=0

	for {
		group := []byte{}
		for (slice[lCursor] == current) == isCurrent {
			group = append(group, slice[lCursor])
			lCursor++
			if lCursor >= len(slice) {
				if len(group) > 0 {
					groups = append(groups, group)
				}

				return groups
			}
		}
		if len(group) > 0 {
			groups = append(groups, group)
		}

		isCurrent = !isCurrent
	}
}

func sortGroups(groups [][]byte) (flips, elements int, slice []byte) {
	endFlip := len(groups)%2 == 1

	lCursor, rCursor := 1, len(groups)-2

	if endFlip {
		rCursor++
	}

	flips = rCursor / 2

	count := 1
	for lCursor < rCursor {
		if count%2 == 1 {
			groups[lCursor], groups[rCursor] = groups[rCursor], groups[lCursor]
			slices.Reverse(groups[rCursor])
		}
		

		elements += count * (len(groups[lCursor]) + len(groups[rCursor]))
		count++
		lCursor += 1
		rCursor -= 1
	}

	for _, group := range groups {
		if group[0] != groups[0][0] {
			slice = append(slice, group...)
		}
	}

	return flips, elements, slice
}
