package main

import (
	"fmt"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/coords"
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

func part1(name string) int {
	barrels := files.ReadLines(name)

	start := coords.NewCoord(0, 0)
	burnedMap := map[coords.Coord]bool{	start: true}
	queue := []coords.Coord{start}

	return len(runRound(barrels,queue,burnedMap))
}

func part2(name string) int {
	barrels := files.ReadLines(name)

	start := coords.NewCoord(0, 0)
	start2 := coords.NewCoord(len(barrels)-1, len(barrels[0])-1)
	burnedMap := make(map[coords.Coord]bool)
	queue := []coords.Coord{start,start2}

	return len(runRound(barrels,queue,burnedMap))
}

func part3(name string) int {
	input := files.ReadLines(name)

	doneMap := make(map[coords.Coord]bool)
	var testMap map[coords.Coord]bool
	roundMax := make(map[coords.Coord]bool)
	var maxPos coords.Coord
	positions:=[]coords.Coord{}

	for range 3 {
		for i, row := range input {
			for j := range row {
				testPos := []coords.Coord{coords.NewCoord(i, j)}

				if _,ok:= doneMap[testPos[0]]; ok {
					continue
				}

				if _,ok:= testMap[testPos[0]]; ok {
					continue
				}

				testMap = runRound(input, testPos, doneMap)
				if len(testMap) > len(roundMax) {
					roundMax = copyMap(testMap)
					maxPos = testPos[0]
				}
			}
		}

		positions = append(positions, maxPos)

		for k,v:= range roundMax {
			doneMap[k]=v
		}

		roundMax= make(map[coords.Coord]bool)
	}

	return len(runRound(input,positions,make(map[coords.Coord]bool)))
}

func runRound(input []string, queue []coords.Coord, prevMap map[coords.Coord]bool) map[coords.Coord]bool {
	newMap := make(map[coords.Coord]bool)

	for _,pos:= range queue {
		newMap[pos] = true
	}

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		directions := coords.DirectionsInOrder
		current := ints.FromString(input[next.I][next.J : next.J+1])

		for _, direction := range directions {
			newPos := next.Add(direction)

			if !newPos.InInput(input) {
				continue
			}

			if _, ok := newMap[newPos]; ok {
				continue
			}

			if _, ok := prevMap[newPos]; ok {
				continue
			}

			newVal := ints.FromString(input[newPos.I][newPos.J : newPos.J+1])

			if newVal > current {
				continue
			}

			newMap[newPos] = true

			queue = append(queue, newPos)
		}
	}

	return newMap
}

func copyMap(src map[coords.Coord]bool) map[coords.Coord]bool {
	newMap := make(map[coords.Coord]bool)
	for k, v := range src {
		newMap[k] = v
	}

	return newMap
}
