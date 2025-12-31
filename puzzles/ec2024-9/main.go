package main

import (
	"fmt"
	"math"
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
	total := 0
	stamps, targets := parseInput(name)

	minSteps:= mapMinimumSteps(stamps, ints.Max(targets))
	for _, target := range targets {
		total += minSteps[target]
	}

	return total
}

func part2(name string) int {
	return part1(name)
}

func part3(name string) int {
	total := 0
	stamps, targets := parseInput(name)

	minSteps:= mapMinimumSteps(stamps, ints.Max(targets))

	for _, target:= range targets {
		total += getBestSplit(target, minSteps)
	}

	return total
}

func parseInput(name string) ([]int, []int) {
	input := files.ReadParagraphs(name)
	stamps := ints.FromStringSlice(strings.Split(input[0][0], ", "))
	targets := ints.FromStringSlice(input[1])

	slices.SortFunc(stamps, func(a, b int) int {
		return b - a
	})

	return stamps, targets
}

func mapMinimumSteps(stamps []int, maxTarget int) map[int]int {
	minSteps := make(map[int]int)

	for _, stamp := range stamps {
		minSteps[stamp] = 1
	}

	for i := 1; i <= maxTarget; i++ {
		steps, ok := minSteps[i]
		if !ok {
			continue
		}

		for _, stamp := range stamps {
			next := i + stamp
			nextSteps, ok := minSteps[next]
			if !ok || nextSteps > steps+1 {
				minSteps[next] = steps + 1
			}
		}
	}

	return minSteps
}

func getBestSplit(target int, minSteps map[int]int) int {
	best:= math.MaxInt

	for i:=target/2 - 49; i<=target/2; i++ {
		// if target - i - i > 100 {
		// 	continue
		// }

		split1, ok:= minSteps[i]
		if !ok {
			continue
		}
		split2,ok := minSteps[target-i]
		if !ok {
			continue
		}

		total:= split1 + split2

		if total < best {
			best = total
		}
	}

	if best == math.MaxInt {
		panic(fmt.Sprintf("could not find split %d",target))
	}

	return best
}