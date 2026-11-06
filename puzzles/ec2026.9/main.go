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
		seconds, _ := tree.growToLeafLimit()
		total += seconds
	}

	return total
}

func part3(name string) int {
	return part2(name)
}

type tree struct {
	growth, limit int
	branchMap     map[int]int
}

// return seconds, total elements (including branches)
func (t tree) growToSizeLimit() (int, int) {
	seconds := 0
	size := 0
	for {
		seconds++

		nextBranchMap := make(map[int]int)
		for branchLength, copies := range t.branchMap {
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

		t.branchMap = nextBranchMap

		size = t.size()
		if size >= t.limit {
			return seconds, size
		}
	}
}

// return seconds, total leaves
func (t tree) growToLeafLimit() (int, int) {
	seconds := 0
	leaves := 0
	for {
		seconds += t.grow()

		leaves = t.leaves()
		if leaves >= t.limit {
			return seconds, leaves
		}
	}
}

func (t *tree) grow() int {
	nextBranchMap := make(map[int]int)

	max := 0
	for branchLength := range t.branchMap {
		if branchLength > max {
			max = branchLength
		}
	}

	seconds := t.growth - max

	for branchLength, copies := range t.branchMap {
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

	t.branchMap = nextBranchMap

	return seconds
}

func (t tree) size() int {
	total := 0
	for k, v := range t.branchMap {
		if k == -1 {
			total += v
			continue
		}

		total += k * v
	}

	return total
}

func (t tree) leaves() int {
	total := 0
	for k, v := range t.branchMap {
		if k == -1 {
			total -= v
			continue
		}
		total += v
	}

	return total
}

func parseInput(name string) []tree {
	out := []tree{}

	for _, row := range files.ReadLines(name) {
		split := strings.Split(row, ", ")
		out = append(out, tree{
			growth:    ints.FromString(split[1][7:]),
			limit:     ints.FromString(split[2][6:]),
			branchMap: map[int]int{0: 1},
		})
	}

	return out
}
