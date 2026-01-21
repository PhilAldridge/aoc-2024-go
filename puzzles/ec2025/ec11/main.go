package main

import (
	"fmt"
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
	duckCounts := ints.FromStringSlice(files.ReadLines(name))
	rounds := 0

	for rounds < 10 {
		rounds++
		newCounts, moved := roundOne(duckCounts)
		duckCounts = newCounts
		if !moved {
			rounds--
			break
		}
	}

	for rounds < 10 {
		rounds++
		duckCounts, _ = roundTwo(duckCounts)
	}

	return checkSum(duckCounts)
}

func part2(name string) int {
	duckCounts := ints.FromStringSlice(files.ReadLines(name))
	rounds := 0

	for {
		rounds++
		newCounts, moved := roundOne(duckCounts)
		duckCounts = newCounts
		if !moved {
			rounds--
			break
		}
	}

	avg:= ints.Mean(duckCounts)

	for _,count:= range duckCounts {
		if count < avg {
			rounds += avg-count
		}
	}

	return rounds
}

func part3(name string) int {
	return part2(name)
}

func roundOne(input []int) ([]int, bool) {
	output := make([]int, len(input))
	movement := false

	for i := 1; i < len(input); i++ {
		if input[i] >= input[i-1] || input[i-1] == 0 {
			output[i-1] = input[i-1]
			continue
		}

		output[i-1] = input[i-1] - 1
		input[i]++
		movement = true
	}

	output[len(input)-1] = input[len(input)-1]

	return output, movement
}

func roundTwo(input []int) ([]int, bool) {
	output := make([]int, len(input))
	movement := false

	for i := 1; i < len(input); i++ {
		if input[i] <= input[i-1] || input[i] == 0 {
			output[i-1] = input[i-1]
			continue
		}

		output[i-1] = input[i-1] + 1
		input[i]--
		movement = true
	}

	output[len(input)-1] = input[len(input)-1]

	return output, movement
}

func checkSum(input []int) int {
	total := 0

	for i, val := range input {
		total += (i + 1) * val
	}

	return total
}
