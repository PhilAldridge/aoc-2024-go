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
	reachableSummitMap := createReachableSummitMap(len(mapping), len(mapping[0]))
	total := 0
	//Start at peaks and work down
	for check := 9; check >= 0; check-- {
		//At each height add the reachable peaks from the adjacent squares that are one above the current height
		for i, line := range mapping {
			for j, val := range line {
				if val != check {
					continue
				}
				if check == 9 {
					//summits can only reach themselves
					reachableSummitMap[i][j] = [][2]int{{i, j}}
					continue
				}
				//list of unique summits reachable at this location
				reachableSummitMap[i][j] = getReachableSummits(mapping,reachableSummitMap,check,i,j)
				if check == 0 {
					//at the base, add number of reachable summits to total
					total += len(reachableSummitMap[i][j])
				}
			}
		}
	}
	return total
}

func part2(name string) int {
	mapping := getMapping(name)
	scoreMap := createScoreMap(len(mapping), len(mapping[0]))
	total := 0
	//Start at peaks and work down
	for check := 9; check >= 0; check-- {
		//At each height add the number of paths to peaks from the adjacent squares that are one above the current height
		for i, line := range mapping {
			for j, val := range line {
				if val != check {
					continue
				}
				if check == 9 {
					//start at summit, with one way to get to it
					scoreMap[i][j] = 1
					continue
				}
				//add up the four adjacent number of paths where they are 1 above current height
				scoreMap[i][j] = getNewScore(mapping,scoreMap,i,j,check)
				if check == 0 {
					//at base, add total paths to summits to the running total
					total += scoreMap[i][j]
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

func createReachableSummitMap(y int, x int) [][][][2]int {
	//create array of empty coordinate arrays in the same size as the main map
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

func getReachableSummits(mapping [][]int, reachableSummitMap [][][][2]int, check int, i int, j int) [][2]int {
	reachableSummits := [][2]int{}
	//add unique reachable summits in all four directions
	//given that direction stays in bounds
	//and direction has a height of 1 more than the location being checked
	if i > 0 && mapping[i-1][j] == check+1 {
		reachableSummits = addUnique(reachableSummits, reachableSummitMap[i-1][j])
	}
	if i < len(mapping)-1 && mapping[i+1][j] == check+1 {
		reachableSummits = addUnique(reachableSummits, reachableSummitMap[i+1][j])
	}
	if j > 0 && mapping[i][j-1] == check+1 {
		reachableSummits = addUnique(reachableSummits, reachableSummitMap[i][j-1])
	}
	if j < len(mapping[0])-1 && mapping[i][j+1] == check+1 {
		reachableSummits = addUnique(reachableSummits, reachableSummitMap[i][j+1])
	}
	return reachableSummits
}

func createScoreMap(y int, x int) [][]int {
	//create array of 0s, in the same size of the main map
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

func getNewScore(mapping [][]int, scoreMap [][]int, i int, j int, check int) int {
	newScore := 0
	//add number of paths to summits in all four directions
	//given that direction stays in bounds
	//and direction has a height of 1 more than the location being checked
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
	return newScore
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
