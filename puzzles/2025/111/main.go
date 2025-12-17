package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
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
	input:= parseInput(name)

	return wrappedCountPaths(input, "you", "out")
}

func part2(name string) int {
	input:= parseInput(name)

	fftToDac:= wrappedCountPaths(input,"fft","dac")

	if fftToDac != 0 {
		return wrappedCountPaths(input, "svr","fft") * fftToDac * wrappedCountPaths(input, "dac","out")
	}

	return wrappedCountPaths(input, "svr", "dac") * wrappedCountPaths(input, "dac", "fft") * wrappedCountPaths(input,"fft","out")
}

func parseInput(name string) map[string][]string {
	lines:= files.ReadLines(name)

	result:= make(map[string][]string)
	for _, line:= range lines {
		split:= strings.Split(line," ")
		result[split[0][:len(split[0])-1]] = split[1:]
	}

	return result
}

func countPaths(input map[string][]string, currentLocation string, memo map[string]int) int {
	val,ok:= memo[currentLocation]
	if ok {
		return val
	}

	next:= input[currentLocation]

	total:=0
	for _,nextVal:= range next {
		total += countPaths(input, nextVal, memo)
	}
	memo[currentLocation] = total

	return total
}

func wrappedCountPaths(input map[string][]string, start string, end string) int {
	memo:= make(map[string]int)
	memo[end] = 1

	return countPaths(input,start,memo)
} 