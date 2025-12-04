package main

import (
	"fmt"
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
	lines:= files.ReadLines(name)
	paperMap:= createPaperMap(lines)

	total:=0
	for paperLocation := range paperMap {
		if countAdjacents(paperLocation[0],paperLocation[1],paperMap) <4 {
			total++
		}
	}
	return total
}

func part2(name string) int {
	lines:= files.ReadLines(name)
	paperMap:= createPaperMap(lines)

	total:=0
	removedOne := true
	for removedOne {
		removedOne = false
		for paperLocation := range paperMap {
			if countAdjacents(paperLocation[0],paperLocation[1],paperMap) <4 {
				removedOne = true
				total++
				delete(paperMap,paperLocation)
			}
		}
	}
	
	return total
}

func createPaperMap(input []string) map[[2]int]int {
	res:= make(map[[2]int]int)

	for y,row:=range input {
		for x,col:= range row {
			if col == '@' {
				res[[2]int{x,y}] = 1
			}
		}
	}

	return res
}

func countAdjacents(x int, y int, paperMap map[[2]int]int) int {
	total:=0
	for i:=-1;i<=1;i++ {
		for j:=-1;j<=1;j++ {
			total += paperMap[[2]int{x+i,y+j}]
		}
	}
	//one fewer than total because this counts its own position
	return total - 1
}