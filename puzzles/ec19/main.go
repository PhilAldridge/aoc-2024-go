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

type state struct {
	y, flaps int
}

func part1(name string) int {
	input := files.ReadLines(name)
	gaps := parseInput(input)
	queue := []state{{y:0, flaps:0}}
	prevX := 0

	for _, gap:= range gaps {
		nextQueue:= []state{}

		for _, gapY:= range gap.gap {
			min:= math.MaxInt

			for _, prev:= range queue {
				dx:= gap.x - prevX
				dy:= gapY - prev.y

				if ints.Abs(dy) > dx {
					continue
				}

				upAndAcross:= dx + dy
				if upAndAcross%2 ==1 {
					continue
				}

				flaps:= prev.flaps + upAndAcross/2

				if flaps < min {
					min = flaps
				}
			}

			if min != math.MaxInt {
				nextQueue = append(nextQueue, state{y:gapY, flaps: min})
			}
		}

		prevX = gap.x
		queue = nextQueue
	}

	min:= math.MaxInt 

	for _, q:= range queue {
		if q.flaps < min {
			min = q.flaps
		}
	}

	return min
}

func part2(name string) int {
	return part1(name)
}

func part3(name string) int {
	return part1(name)
}

type gapType struct {
	x int
	gap []int
}

func parseInput(input []string) ([]gapType) {
	out := []gapType{}
	gapMap := make(map[int][]int)

	for _, row := range input {
		vals := ints.FromStringSlice(strings.Split(row, ","))
		gapMap[vals[0]] = append(gapMap[vals[0]], ints.GetIntsBetweenInclusive(vals[1], vals[1] + vals[2]-1)...)
	}

	for k,v:= range gapMap {
		out = append(out, gapType{
			x: k,
			gap: v,
		})
	}

	slices.SortFunc(out, func(a,b gapType) int {
		return a.x - b.x
	})

	return out
}
