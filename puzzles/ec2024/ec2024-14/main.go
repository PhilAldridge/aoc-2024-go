package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/ints"
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

type nextValType struct {
	pos [3]int
	score int
	leaf bool
}

func part1(name string) int {
	input := files.Read(name)
	growths := strings.Split(input, ",")
	height := 0

	max := 0

	for _, growth := range growths {
		amount := ints.FromString(growth[1:])
		switch growth[0] {
		case 'U':
			height += amount
			if height > max {
				max = height
			}
		case 'D':
			height -= amount
		}
	}

	return max
}

func part2(name string) int {
	input := files.ReadLines(name)

	segments := make(map[[3]int]bool)

	for _, branch := range input {
		grow(branch, segments)
	}

	return len(segments)
}

func part3(name string) int {
	input := files.ReadLines(name)

	segments := make(map[[3]int]bool)

	for _, branch := range input {
		grow(branch, segments)
	}

	maxHeight:= 0
	leafCount:= 0

	for segment, leaf := range segments {
		if segment[0] > maxHeight {
			maxHeight = segment[0]
		}
		if leaf {
			leafCount++
		}
	}

	minMurkiness:= math.MaxInt

	for height:= range maxHeight {
		if _,ok:= segments[[3]int{height+1,0,0}]; !ok {
			continue
		}

		murkiness,ok:= calculateMurkiness(segments, height+1, leafCount, minMurkiness)
		if !ok {
			fmt.Println(height, " skipped")
			continue
		}

		if murkiness < minMurkiness {
			minMurkiness = murkiness
			fmt.Println(height, ": ",murkiness)
		}
	}

	return minMurkiness
}

func grow(branch string, segments map[[3]int]bool) {
	currentPos := [3]int{}
	growths := strings.Split(branch, ",")
	for _, growth := range growths {
		amount := ints.FromString(growth[1:])

		switch growth[0] {
		case 'U':
			for range amount {
				currentPos[0]++
				if _,ok:= segments[currentPos];!ok {
					segments[currentPos] = false
				}
			}
		case 'D':
			for range amount {
				currentPos[0]--
				if _,ok:= segments[currentPos];!ok {
					segments[currentPos] = false
				}
			}
		case 'L':
			for range amount {
				currentPos[1]--
				if _,ok:= segments[currentPos];!ok {
					segments[currentPos] = false
				}
			}
		case 'R':
			for range amount {
				currentPos[1]++
				if _,ok:= segments[currentPos];!ok {
					segments[currentPos] = false
				}
			}
		case 'F':
			for range amount {
				currentPos[2]++
				if _,ok:= segments[currentPos];!ok {
					segments[currentPos] = false
				}
			}
		case 'B':
			for range amount {
				currentPos[2]--
				if _,ok:= segments[currentPos];!ok {
					segments[currentPos] = false
				}
			}
		}
	}
	segments[currentPos] = true
}

func calculateMurkiness(tree map[[3]int]bool, height int, leafCount int, minSoFar int) (int, bool) {
	currentPos:= [3]int{height,0,0}
	shortestDistances:= map[[3]int]int{
		currentPos:0,
	}

	directions:= [6][3]int{
		{1,0,0},
		{-1,0,0},
		{0,1,0},
		{0,-1,0},
		{0,0,1},
		{0,0,-1},
	}

	total:= 0
	leavesReached:=0
	lastShortest:=0

	for leavesReached < leafCount {
		nextVal:= nextValType{
			score: math.MaxInt,
		}

		if total + (lastShortest*(leafCount-leavesReached)) > minSoFar {
			return 0,false
		}

		for startPos, score := range shortestDistances {
			if score+1 >= nextVal.score {
				continue
			}

			for _, direction:= range directions {
				newPos:= [3]int{
					startPos[0]+direction[0],
					startPos[1]+direction[1],
					startPos[2]+direction[2],
				}

				if _,ok:= shortestDistances[newPos]; ok {
					continue
				} 

				leaf,ok:= tree[newPos]

				if !ok {
					continue
				}

				nextVal = nextValType{
					pos: newPos,
					score: score+1,
					leaf: leaf,
				}
			}
		}
		if nextVal.pos == [3]int{} {
			panic("no new paths found!")
		}

		shortestDistances[nextVal.pos] = nextVal.score
		if nextVal.leaf {
			leavesReached++
			total += nextVal.score
		}
		lastShortest = nextVal.score
	}

	return total,true
}