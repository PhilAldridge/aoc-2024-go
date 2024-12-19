package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1: ", part1("input.txt"), " in:", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", part2("input.txt"), " in:", time.Since(start))
}

func part1(name string) int {
	foundMap:= make(map[string]int)
	towels, patterns:= parseInput(name)
	total:=0
	for _,pattern:= range patterns {
		if checkPossible(pattern, towels, foundMap)>0 {
			total++
		}
	}
	return total
}

func part2(name string) int {
	foundMap:= make(map[string]int)
	towels, patterns:= parseInput(name)
	total:=0
	for _,pattern:= range patterns {
		total += checkPossible(pattern, towels, foundMap)
	}
	return total
}

func parseInput(name string) ([]string, []string) {
	files:= files.ReadLines(name)
	towels:= strings.Split(files[0],", ")
	patterns:= files[2:]
	return towels,patterns
} 

func checkPossible(pattern string, towels []string, doneMap map[string]int) int {
	if len(pattern)==0 {
		return 1
	}
	if done,ok:= doneMap[pattern]; ok {
		return done
	}
	for _,towel:= range towels {
		match:= true
		if len(towel)>len(pattern) {
			continue
		}
		for i,char:= range towel {
			if char != rune(pattern[i]) {
				match = false
				break
			}
		}
		if match  {
			checkScore:= checkPossible(pattern[len(towel):],towels,doneMap)
			doneMap[pattern] +=checkScore
		}
	}
	return doneMap[pattern]
}
