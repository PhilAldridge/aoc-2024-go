package main

import (
	"fmt"
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
	lines:= files.ReadLines(name)

	total:= 0

	beamMap:= make(map[int]bool)

	for i, char:= range lines[0]{
		if char == 'S' {
			beamMap[i]=true
			break
		}
	}

	for i:=1; i<len(lines); i++ {
		for k:= range beamMap {
			if lines[i][k] == '^' {
				total ++
				delete(beamMap,k)
				if k != 0 {
					beamMap[k-1] = true
				}
				if k+1 != len(lines[0]) {
					beamMap[k+1] = true
				}
			}
		}
	}
	return total
}

func part2(name string) int {
	lines:= files.ReadLines(name)

	startPos:=0

	for i, char:= range lines[0]{
		if char == 'S' {
			startPos = i
			break
		}
	}

	memoMap:= make(map[[2]int]int)

	return countTimelines(1,startPos,lines, memoMap)
}

func countTimelines(startLine int, startPos int, lines []string, memoMap map[[2]int]int) int {
	mapTuple:= [2]int{startLine,startPos}

	val, ok := memoMap[mapTuple]
	if ok {
		return val
	}

	if startLine == len(lines)-1 {
		if lines[startLine][startPos] == '^' {
			memoMap[mapTuple] = 2
			return 2
		}

		memoMap[mapTuple] = 1
		return 1
	}

	total:=0

	if lines[startLine][startPos] == '^' {
		if startPos != 0 {
			total+= countTimelines(startLine+1, startPos-1, lines,memoMap)
		}

		if startPos+1 != len(lines[0]) {
			total+= countTimelines(startLine+1,startPos+1, lines,memoMap)
		}


		return total
	} else {
		total += countTimelines(startLine+1,startPos,lines,memoMap)
	}

	memoMap[mapTuple] = total
	return total
}
