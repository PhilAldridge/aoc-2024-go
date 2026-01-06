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
	maze:= files.ReadLines(name)
	starts:= getStarts(maze)
	return ints.Max(floodFill(maze,starts))
}

func part2(name string) int {
	return part1(name)
}


func part3(name string) int {
	maze:= files.ReadLines(name)
	var scores []int

	for i,row:= range maze {
		for j,char:= range row {
			if char == '.' {
				scores = append(scores, ints.Sum(floodFill(maze,[]coords.Coord{coords.NewCoord(i,j)})))
			}
		}
	}

	return ints.Min(scores)
}

func getStarts(maze []string) []coords.Coord {
	var starts []coords.Coord
	for i:= range maze[0] {
		if maze[0][i] != '#' {
			starts = append(starts,  coords.NewCoord(0,i))
		}
		if maze[len(maze)-1][i] != '#' {
			starts = append(starts,  coords.NewCoord(len(maze)-1,i))
		}
	}

	for i:= range maze {
		if maze[i][0] != '#' {
			starts = append(starts,  coords.NewCoord(i,0))
		}
		if maze[i][len(maze[i])-1] != '#' {
			starts = append(starts,  coords.NewCoord(i,len(maze[i])-1))
		}
	}

	return starts
}

type state struct {
	position coords.Coord
	time int
}

func floodFill(maze []string, starts []coords.Coord) []int {
	queue:= []state{}
	visitedMap:= map[coords.Coord]bool{}

	for _, start:= range starts {
		queue = append(queue, state{
			position: start,
			time: 0,
		})

		visitedMap[start] = true
	}

	palmTimes:=[]int{}

	for len(queue)>0 {
		next:= queue[0]
		queue = queue[1:]
		
		if maze[next.position.I][next.position.J] == 'P' {
			palmTimes = append(palmTimes, next.time)
		}

		adjacents:= next.position.GetAdjacent()

		for _, adjacent:= range adjacents {
			if visitedMap[adjacent] || !adjacent.InInput(maze) {
				continue
			}

			if maze[adjacent.I][adjacent.J] == '#' {
				continue
			}

			visitedMap[adjacent] = true
			queue = append(queue, state{
				position: adjacent,
				time: next.time+1,
			})
		}
	}

	return palmTimes
}
