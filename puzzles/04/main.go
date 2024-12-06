package main

import (
	"fmt"
	"time"

	"github.com/kylehoehns/aoc-2023-go/pkg/files"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1: ", part1("input.txt"), " in:", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", part2("input.txt"), " in:", time.Since(start))
}

func part1(name string) int {
	lines := files.ReadLines(name)
	total := 0
	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			if line[j] == 'X' {
				total += checkXmas(lines, i, j)
			}
		}
	}
	return total
}

func part2(name string) int {
	lines := files.ReadLines(name)
	total := 0
	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			if line[j] == 'A' {
				total += checkMAS(lines, i, j)
			}
		}
	}
	return total
}

func checkXmas(lines []string, i int, j int) int {
	total := 0
	for k := -1; k <= 1; k++ {
		for l := -1; l <= 1; l++ {
			if i+3*k < 0 || i+3*k >= len(lines) {
				continue
			}
			if j+3*l < 0 || j+3*l >= len(lines[0]) {
				continue
			}
			if lines[i+k][j+l] == 'M' && lines[i+2*k][j+2*l] == 'A' && lines[i+3*k][j+3*l] == 'S' {
				total++
			}
		}
	}
	return total
}

func checkMAS(lines []string, i int, j int) int {
	if i == 0 || i == len(lines)-1 || j == 0 || j == len(lines)-1 {
		return 0
	}
	total := 0
	for k := -1; k <= 0; k++ {
		for l := -1; l <= 0; l++ {
			if k == 0 || l == 0 {
				continue
			}
			if (lines[i+k][j+l] == 'M' && lines[i-k][j-l] == 'S') ||
				(lines[i-k][j-l] == 'M' && lines[i+k][j+l] == 'S') {
				if (lines[i+k][j-l] == 'M' && lines[i-k][j+l] == 'S') ||
					(lines[i-k][j+l] == 'M' && lines[i+k][j-l] == 'S') {
					total++
				}
			}
		}
	}
	return total
}
