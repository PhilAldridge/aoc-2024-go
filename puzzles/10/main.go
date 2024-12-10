package main

import (
	"fmt"
	"slices"
	"strings"
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
	mapping := getMapping(name)
	scoreMap := createScoreMap(len(mapping), len(mapping[0]))
	total := 0
	for check := 9; check >= 0; check-- {
		for i, line := range mapping {
			for j, val := range line {
				if val != check {
					continue
				}
				if check == 9 {
					scoreMap[i][j] = [][2]int{{i, j}}
					continue
				}
				newScore := [][2]int{}
				if i > 0 && mapping[i-1][j] == check+1 {
					newScore = addUnique(newScore, scoreMap[i-1][j])
				}
				if i < len(mapping)-1 && mapping[i+1][j] == check+1 {
					newScore = addUnique(newScore, scoreMap[i+1][j])
				}
				if j > 0 && mapping[i][j-1] == check+1 {
					newScore = addUnique(newScore, scoreMap[i][j-1])
				}
				if j < len(mapping[0])-1 && mapping[i][j+1] == check+1 {
					newScore = addUnique(newScore, scoreMap[i][j+1])
				}
				scoreMap[i][j] = newScore
				if check == 0 {
					total += len(newScore)
				}
			}
		}
	}
	return total
}

func part2(name string) int {
	mapping := getMapping(name)
	scoreMap := createScoreMap2(len(mapping), len(mapping[0]))
	total := 0
	for check := 9; check >= 0; check-- {
		for i, line := range mapping {
			for j, val := range line {
				if val != check {
					continue
				}
				if check == 9 {
					scoreMap[i][j] = 1
					continue
				}
				newScore := 0
				if i > 0 && mapping[i-1][j] == check+1 {
					newScore += scoreMap[i-1][j]
				}
				if i < len(mapping)-1 && mapping[i+1][j] == check+1 {
					newScore += scoreMap[i+1][j]
				}
				if j > 0 && mapping[i][j-1] == check+1 {
					newScore += scoreMap[i][j-1]
				}
				if j < len(mapping[0])-1 && mapping[i][j+1] == check+1 {
					newScore += scoreMap[i][j+1]
				}
				scoreMap[i][j] = newScore
				if check == 0 {
					total += newScore
				}
			}
		}
	}
	return total
}

func getMapping(name string) [][]int {
	mapping := [][]int{}
	lines := files.ReadLines(name)
	for _, line := range lines {
		mapping = append(mapping, ints.FromStringSlice(strings.Split(line, "")))
	}
	return mapping
}

func createScoreMap(y int, x int) [][][][2]int {
	result := [][][][2]int{}
	for i := 0; i < y; i++ {
		newLine := [][][2]int{}
		for j := 0; j < x; j++ {
			newLine = append(newLine, [][2]int{})
		}
		result = append(result, newLine)
	}
	return result
}

func createScoreMap2(y int, x int) [][]int {
	result := [][]int{}
	for i := 0; i < y; i++ {
		newLine := []int{}
		for j := 0; j < x; j++ {
			newLine = append(newLine, 0)
		}
		result = append(result, newLine)
	}
	return result
}

func addUnique(newScore [][2]int, toAppend [][2]int) [][2]int {
	for _, coord := range toAppend {
		if slices.Contains(newScore, coord) {
			continue
		}
		newScore = append(newScore, coord)
	}
	return newScore
}
