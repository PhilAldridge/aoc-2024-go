package main

import (
	"fmt"
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
	total := 0

	trees := parseInput(name)
	for _, tree := range trees {
		seconds, _ := tree.growToSizeLimit()
		total += seconds
	}

	return total
}

func part2(name string) int {
	total := 0

	trees := parseInput(name)
	for _, tree := range trees {
		seconds := tree.growToLeafLimit()
		total += seconds
	}

	return total
}

func part3(name string) int {
	return part2(name)
}

type tree struct {
	growth, limit      int
	branchLengthCopies map[int]int
}

// return seconds, total elements (including branches)
func (t tree) growToSizeLimit() int {
	seconds := 1
	for t.size() < t.limit {
		seconds++

		nextBranchMap := make(map[int]int)
		for branchLength, copies := range t.branchLengthCopies {
			// Node
			if branchLength == -1 {
				nextBranchMap[branchLength] += copies
				continue
			}

			// Split branch in to two halves, node and branch of size 1
			if branchLength+1 == t.growth {
				nextBranchMap[-1] += copies
				nextBranchMap[t.growth/2] += copies * 2
				nextBranchMap[1] += copies
				continue
			}

			nextBranchMap[branchLength+1] += copies
		}

		t.branchLengthCopies = nextBranchMap
	}

	return seconds
}

// return seconds, total leaves
func (t tree) growToLeafLimit() int {
	seconds := 1
	for t.leaves() < t.limit {
		seconds += t.grow()
	}

	return seconds
}

func (t *tree) grow() int {
	nextBranchMap := make(map[int]int)

	max := 0
	for branchLength := range t.branchLengthCopies {
		if branchLength > max {
			max = branchLength
		}
	}

	seconds := t.growth - max

	for branchLength, copies := range t.branchLengthCopies {
		// Node
		if branchLength == -1 {
			nextBranchMap[branchLength] += copies
			continue
		}

		// Split branch in to two halves, node and branch of size 1
		if branchLength == max {
			nextBranchMap[-1] += copies
			nextBranchMap[t.growth/2] += copies * 2
			nextBranchMap[1] += copies
			continue
		}

		nextBranchMap[branchLength+seconds] += copies
	}

	t.branchLengthCopies = nextBranchMap

	return seconds
}

func (t tree) size() int {
	total := 0
	for branchLength, copies := range t.branchLengthCopies {
		if branchLength == -1 {
			total += copies
			continue
		}

		total += branchLength * copies
	}

	return total
}

func (t tree) leaves() int {
	total := 0
	for branchLength, copies := range t.branchLengthCopies {
		if branchLength == -1 {
			total -= copies
			continue
		}
		total += copies
	}

	return total
}

func parseInput(name string) []tree {
	out := []tree{}

	for _, row := range files.ReadLines(name) {
		split := strings.Split(row, ", ")
		out = append(out, tree{
			growth:             ints.FromString(split[1][7:]),
			limit:              ints.FromString(split[2][6:]),
			branchLengthCopies: map[int]int{1: 1},
		})
	}

	return out
}
