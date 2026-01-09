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
	fmt.Println("Part 2 answer: ", part2("input2.txt"))
	split2 := time.Now()
	fmt.Println("Part 3 answer: ", part3("input3.txt"))

	
	fmt.Println()
	fmt.Println("Part 1: ", split.Sub(start))
	fmt.Println("Part 2: ", split2.Sub(split))
	fmt.Println("Part 2: ", time.Since(split2))
}

func part1(name string) string {
	tree:= parseInput(name)

	finalPaths:= getPaths(tree)

	for _, paths := range finalPaths {
		if len(paths) == 1 {
			return combinePath(paths[0])
		}
	} 

	return "not found"
}

func part2(name string) string {
	tree:= parseInput(name)

	finalPaths:= getPaths(tree)

	for _, paths := range finalPaths {
		if len(paths) == 1 {
			return combinePathShort(paths[0])
		}
	} 

	return "not found"
}


func part3(name string) string {
	tree:= parseInput(name)

	finalPaths:= getPaths(tree)

	for _, paths := range finalPaths {
		if len(paths) == 1 {
			return combinePathShort(paths[0])
		}
	} 

	return "not found"
}

func parseInput(name string) map[string][]string {
	lines:= files.ReadLines(name)

	result:= make(map[string][]string)

	for _,line:= range lines {
		split1:= strings.Split(line,":")
		split2:= strings.Split(split1[1],",")

		result[split1[0]] = split2
	}

	return result
}

func combinePath(path []string) string {
	result:= ""

	for _, branch:= range path {
		result += branch
	}

	return result
}

func combinePathShort(path []string) string {
	result:= ""

	for _, branch:= range path {
		result += branch[0:1]
	}

	return result
}

func getPaths(tree map[string][]string) map[int][][]string {
	paths:= [][]string{{"RR"}}

	finalPaths := make(map[int][][]string)

	for len(paths) > 0 {
		nextPaths:= [][]string{}

		for _, path:= range paths {
			branches:= tree[path[len(path)-1]]
			for _,branch := range branches {
				nextPath := make([]string, len(path))
				copy(nextPath,path)
				nextPath = append(nextPath, branch)

				if branch == "ANT" || branch == "BUG" {
					continue
				}

				if branch != "@" {
					nextPaths = append(nextPaths, nextPath)
					continue
				}

				finalPaths[len(nextPath)] = append(finalPaths[len(nextPath)],nextPath)
			}
		}

		paths = nextPaths
	}

	return finalPaths
}