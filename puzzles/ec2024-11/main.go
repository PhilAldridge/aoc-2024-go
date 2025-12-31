package main

import (
	"fmt"
	"math"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/slices"
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
	input := files.ReadLines(name)
	termiteMap := slices.StringSliceToStringMapStringSlice(input, ":", ",")
	termiteCount := map[string]int{
		"A": 1,
	}

	for range 4 {
		termiteCount = populate(termiteMap, termiteCount)
	}

	total := 0
	for _, count := range termiteCount {
		total += count
	}

	return total
}

func part2(name string) int {
	input := files.ReadLines(name)
	termiteMap := slices.StringSliceToStringMapStringSlice(input, ":", ",")
	termiteCount := map[string]int{
		"Z": 1,
	}

	for range 10 {
		termiteCount = populate(termiteMap, termiteCount)
	}

	total := 0
	for _, count := range termiteCount {
		total += count
	}

	return total
}

func part3(name string) int {
	input := files.ReadLines(name)
	termiteMap := slices.StringSliceToStringMapStringSlice(input, ":", ",")
	min := math.MaxInt
	max := math.MinInt

	for k, _ := range termiteMap {
		termiteCount := map[string]int{
			k: 1,
		}
		for range 20 {
			termiteCount = populate(termiteMap, termiteCount)
		}

		total := 0
		for _, count := range termiteCount {
			total += count
		}

		if total < min {
			min = total
		}

		if total > max {
			max = total
		}
	}

	return max - min
}

func populate(termiteMap map[string][]string, termiteCount map[string]int) map[string]int {
	nextDay := make(map[string]int)

	for termite, count := range termiteCount {
		nextTermites := termiteMap[termite]

		for _, nextTermite := range nextTermites {
			nextDay[nextTermite] += count
		}
	}

	return nextDay
}
