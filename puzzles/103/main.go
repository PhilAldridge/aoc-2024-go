package main

import (
	"fmt"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/ints"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1: ", part1("input.txt"), " in:", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", part2("input.txt"), " in:", time.Since(start))
}

func part1(name string) int {
	lines:= files.ReadLines(name)
	total:=0
	for _,line:= range lines {
		total += findMaxVoltage(line, 2)
	}
	return total
}

func part2(name string) int {
	lines:= files.ReadLines(name)
	total:=0

	for _, line:= range lines {
		total += findMaxVoltage(line,12)
	}

	return total
}

func findMaxVoltage(input string, digitsRemaining int) (int) {
	max:=-1
	maxIndex:=0
	for i:=0; i<=len(input)-digitsRemaining; i++ {
		next:= ints.FromString(string(input[i]))
		if next>max {
			max = next
			maxIndex = i
		}
	}

	if digitsRemaining == 1 {
		return max
	}

	return max*ints.Pow(10,digitsRemaining-1) + findMaxVoltage(input[maxIndex+1:],digitsRemaining-1)
} 
