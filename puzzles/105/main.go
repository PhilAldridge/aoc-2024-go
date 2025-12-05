package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/ints"
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
	split:= files.ReadParagraphs(name)
	freshRanges := [][2]int{}
	for _,rangeStr:= range split[0] {
		next:= strings.Split(rangeStr,"-")
		nextVals:= ints.FromStringSlice(next)
		freshRanges = append(freshRanges, [2]int{nextVals[0],nextVals[1]})
	}
	ingredients:= ints.FromStringSlice(split[1])

	slices.Sort(ingredients)
	sort.Slice(freshRanges, func(i, j int) bool {
		a:= freshRanges[i][0]
		b:= freshRanges[j][0]
		return a < b
	})

	total:=0

	i:=0
	j:=0
	for i<len(ingredients) && j<len(freshRanges) {
		if ingredients[i] < freshRanges[j][0] {
			//spoiled
			i++
			continue
		}

		if ingredients[i] > freshRanges[j][1] {
			//out of range
			j++
			continue
		}

		total++
		i++
	}

	return total
}

func part2(name string) int {
	split:= files.ReadParagraphs(name)
	freshRanges := [][2]int{}
	for _,rangeStr:= range split[0] {
		next:= strings.Split(rangeStr,"-")
		nextVals:= ints.FromStringSlice(next)
		freshRanges = append(freshRanges, [2]int{nextVals[0],nextVals[1]})
	}

	sort.Slice(freshRanges, func(i, j int) bool {
		a:= freshRanges[i][0]
		b:= freshRanges[j][0]
		return a < b
	})

	total:=0
	maxReached:=0

	for _,freshRange := range freshRanges {
		if maxReached>=freshRange[1] {
			continue
		}

		if maxReached < freshRange[0] {
			maxReached = freshRange[0] - 1
		}
		
		total += freshRange[1] - maxReached
		maxReached = freshRange[1]
	}

	return total
}
