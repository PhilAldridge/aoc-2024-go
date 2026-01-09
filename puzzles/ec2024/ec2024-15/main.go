package main

import (
	"fmt"
	"math"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/coords"
	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/sets"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1 answer: ", part1("input.txt"))
	split := time.Now()
	fmt.Println("Part 2 answer: ", part2("input2.txt"))
	split2 := time.Now()
	fmt.Println("Part 3 answer: ", part2("input3.txt"))

	
	fmt.Println()
	fmt.Println("Part 1: ", split.Sub(start))
	fmt.Println("Part 2: ", split2.Sub(split))
	fmt.Println("Part 2: ", time.Since(split2))
}

type nextValType struct {
	pos coords.Coord
	val int
	plantsCollected *sets.Set[rune]
}

func part1(name string) int {
	maze:= files.ReadLines(name)

	costMap:= make(map[coords.Coord]nextValType)
	for i, char := range maze[0] {
		if char == '.' {
			costMap[coords.NewCoord(0,i)] = nextValType{val:0,plantsCollected: sets.NewSet[rune]()}
			break
		}
	}

	for {
		nextVal := nextValType{val:math.MaxInt,plantsCollected: sets.NewSet[rune]()}
		for pos, prevVal:= range costMap {
			if prevVal.val+1 >= nextVal.val {
				continue
			}

			adjacents := pos.GetAdjacent()
			for _, adjacent:= range adjacents {
				if _,ok:= costMap[adjacent]; ok {
					continue
				}

				if !adjacent.InInput(maze) {
					continue
				}

				if maze[adjacent.I][adjacent.J] == '#' {
					continue
				}

				nextVal = nextValType{
					pos: adjacent,
					val: prevVal.val+1,
					plantsCollected: prevVal.plantsCollected,
				}
				
				if maze[adjacent.I][adjacent.J] == 'H' {
					nextVal.plantsCollected.Add(rune(maze[adjacent.I][adjacent.J]))
				}
			}
		}

		costMap[nextVal.pos] = nextVal
		if nextVal.plantsCollected.Contains('H') {
			return nextVal.val*2
		}
	}
}

type state struct {
	x,y,collectedMask int
}

type queueType struct {
	state
	distance int
}

func part2(name string) int {
	maze := files.ReadLines(name)

	visited:= make(map[state]bool)

	var queue []queueType

	for i, char := range maze[0] {
		if char == '.' {
			queue = []queueType{{
				state: state{x:i},
			}}
			break
		}
	}

	fullMask:= getFullMask(maze)

	visited[queue[0].state] = true
	startX:= queue[0].x

	for len(queue) >0 {
		stateToVisit:= queue[0]
		queue = queue[1:]
		
		if stateToVisit.y == 0 && stateToVisit.collectedMask == fullMask && stateToVisit.x == startX {
			return stateToVisit.distance
		}

		for _, adj:= range coords.NewCoord(stateToVisit.y, stateToVisit.x).GetAdjacent() {
			if !adj.InInput(maze) {
				continue
			}

			char := rune(maze[adj.I][adj.J])

			if char == '#' || char == '~' {
				continue
			}

			newState:= state{
				x:adj.J,
				y:adj.I,
				collectedMask: stateToVisit.collectedMask,
			
			}

			if char != '.' {
				bit:= runeToMask(char)
				newState.collectedMask = newState.collectedMask | bit
			}

			if visited[newState] {
				continue
			}

			visited[newState] = true
			queue = append(queue, queueType{
				state: newState,
				distance: stateToVisit.distance+1,
			})
		}
	}

	return -1
}

func runeToMask(char rune) int {
	if char >= 'A' && char <= 'Z' {
		return 1 << (char-'A')
	}

	return 0
}

func getFullMask(input []string) int {
	runesFound:= make(map[rune]bool)
	for _,row:=range input {
		for _,char:= range row {
			runesFound[char] = true
		}
	}

	result:=0
	for char,_:=range runesFound {
		result += runeToMask(char)
	}

	return result
}

