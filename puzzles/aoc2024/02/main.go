package main

import (
	"fmt"
	"log"
	"strconv"
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
	levels := getLevels(name)
	total := 0
	for _, level := range levels {
		if testSafe(level) {
			total++
		}
	}
	return total
}

func part2(name string) int {
	levels := getLevels(name)
	total := 0
	for _, level := range levels {
		if testSafe2(level) {
			total++
		}
	}
	return total
}

func getLevels(name string) [][]int {
	levels := [][]int{}
	lines := files.ReadLines(name)
	for _, line := range lines {
		vals := strings.Split(line, " ")
		level := []int{}
		for _, v := range vals {
			if i, err := strconv.Atoi(v); err == nil {
				level = append(level, i)
			} else {
				log.Fatal(err)
			}
		}
		levels = append(levels, level)
	}
	return levels
}

func testSafe(level []int) bool {
	if len(level) < 2 {
		return true
	}
	if level[0] == level[1] {
		return false
	}
	ascendingFlag := level[0] < level[1]
	for i := 1; i < len(level); i++ {
		diff := level[i] - level[i-1]
		if !ascendingFlag {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func testSafe2(level []int) bool {
	if len(level) < 3 {
		return true
	}
	if testSafe(level) || testSafe(level[1:]) || testSafe((level[:len(level)-1])) {
		return true
	}
	for i := 1; i < len(level)-1; i++ {
		levelCopy := []int{}
		for j, v := range level {
			if j != i {
				levelCopy = append(levelCopy, v)
			}
		}
		if testSafe(levelCopy) {
			return true
		}
	}
	return false
}
