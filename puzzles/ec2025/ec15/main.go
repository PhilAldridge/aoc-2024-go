package main

import (
	"fmt"
	"slices"
	"strings"
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

type queueType struct {
	position coords.Coord
	time int
}

func part1(name string) int {
	input := strings.Split(files.Read(name), ",")
	lines, endPosition:= getLines(input)
	yMap, xMap:= getImportantYandXvals(lines)
	startPosition := coords.NewCoord(0, 0)
	visitedMap := map[coords.Coord]int{startPosition: 0}

	queue := []queueType{{position: startPosition, time:0}}

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		time := visitedMap[next.position]

		for yVal:= range yMap {
			if yVal == next.position.I {
				continue
			}

			newPosition:= coords.NewCoord(yVal,next.position.J)
			newTime:= time + ints.Abs(yVal-next.position.I)

			if score,ok:= visitedMap[newPosition]; ok && newTime>=score {
				continue
			}

			if hitsAWall([2]coords.Coord{next.position,newPosition},lines) {
				continue
			}

			visitedMap[newPosition] = newTime

			if newPosition.Equals(endPosition) {
				continue
			}

			queue = append(queue, queueType{position: newPosition, time: newTime})
		}

		for xVal:= range xMap {
			if xVal == next.position.J {
				continue
			}

			newPosition:= coords.NewCoord(next.position.I,xVal)
			newTime:= time + ints.Abs(xVal-next.position.J)

			if score,ok:= visitedMap[newPosition]; ok && newTime>=score {
				continue
			}

			if hitsAWall([2]coords.Coord{next.position,newPosition},lines) {
				continue
			}

			visitedMap[newPosition] = newTime

			if newPosition.Equals(endPosition) {
				continue
			}

			queue = append(queue, queueType{position: newPosition, time: newTime})
		}

		slices.SortFunc(queue, func(a,b queueType) int {
			return a.time - b.time
		})
	}

	return visitedMap[endPosition]
}

func part2(name string) int {
	return part1(name)
}

func part3(name string) int {
	return part1(name)
}

func getLines(input []string) ([][2]coords.Coord, coords.Coord) {
	startPosition := coords.NewCoord(0, 0)
	currentPosition := startPosition
	currentDirection := coords.NewCoord(-1, 0)
	var endPosition coords.Coord
	lines := [][2]coords.Coord{}

	for i, instruction := range input {
		switch instruction[0] {
		case 'L':
			currentDirection = coords.TurnLeft(currentDirection)
		case 'R':
			currentDirection = coords.TurnRight(currentDirection)
		}

		newPosition:= currentPosition.MoveBy(currentDirection, ints.FromString(instruction[1:]))

		if i==0 {
			currentPosition = currentPosition.Add(currentDirection)
		}

		if i==len(input)-1 {
			endPosition = newPosition
			newPosition = newPosition.Add(coords.TurnBack(currentDirection))
		}

		lines = append(lines, [2]coords.Coord{currentPosition,newPosition})

		currentPosition = newPosition
	}

	return lines, endPosition
}

func getImportantYandXvals(lines [][2]coords.Coord) (map[int]bool, map[int]bool) {
	yMap, xMap:= make(map[int]bool), make(map[int]bool)

	for _,line:= range lines {
		for i:= -1; i<=1; i++ {
			yMap[line[0].I+i] = true
			yMap[line[1].I+i] = true
			xMap[line[0].J+i] = true
			xMap[line[1].J+i] = true
		}
	}

	return yMap,xMap
}

func hitsAWall(line [2]coords.Coord, walls [][2]coords.Coord) bool {
	for _, wall:= range walls {
		if segmentsIntersect(line, wall) {
			return true
		}
	}

	return false
}

func segmentsIntersect(a,b [2]coords.Coord) bool {
	aTop, aBottom:= ints.MinMax([]int{a[0].I,a[1].I})
	aLeft, aRight:= ints.MinMax([]int{a[0].J,a[1].J})
	bTop, bBottom:= ints.MinMax([]int{b[0].I,b[1].I})
	bLeft, bRight:= ints.MinMax([]int{b[0].J,b[1].J})

	if aTop<bTop && aBottom<bTop {
		return false
	}

	if aBottom>bBottom && aTop>bBottom {
		return false
	}

	if aLeft<bLeft && aRight<bLeft {
		return false
	}

	if aRight>bRight && aLeft>bRight {
		return false
	}

	return true
}