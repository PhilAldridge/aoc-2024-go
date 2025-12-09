package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/ints"
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
	coords,_,_ := parseCoords(name)

	max:=0

	for i:=0;i<len(coords); i++ {
		for j:=i+1; j<len(coords); j++ {
			area:= (ints.Abs(coords[i][0] - coords[j][0])+1) *
				 	(ints.Abs(coords[i][1] - coords[j][1])+1) 
			
			if area > max {
				max = area
			}
		}
	}

	return max
}

func part2(name string) int {
	coords, _, _:= parseCoords(name)

	grid:= make(map[[2]int]rune)

	lastCoord:=coords[len(coords)-1]
	grid[lastCoord] = '#'
	for i:=0; i<len(coords); i++ {
		grid[coords[i]] = '#'

		if coords[i][0] != lastCoord[0] {
			yFill:= ints.GetIntsBetween(coords[i][0], lastCoord[0])
			for _,y:= range yFill {
				grid[[2]int{y,lastCoord[1]}] = 'X'
			}
		} else {
			xFill:= ints.GetIntsBetween(coords[i][1], lastCoord[1])
			for _,x:= range xFill {
				grid[[2]int{lastCoord[0],x}] = 'X'
			}
		}

		lastCoord = coords[i]
	}

	max:=0

	for i:=0;i<len(coords); i++ {
		for j:=i+1; j<len(coords); j++ {
			ok, area:= isValid(coords[i],coords[j],grid)

			if !ok {
				continue
			}
			
			if area > max {
				max = area
			}
		}
		fmt.Println(i)
	}

	return max
}

func parseCoords(name string) ([][2]int, int, int) {
	lines:=files.ReadLines(name)
	coords:= [][2]int{}
	maxX,maxY:=0,0

	for _,line:= range lines {
		sliced:= strings.Split(line, ",")
		vals:= ints.FromStringSlice(sliced)
		if vals[0] > maxY {
			maxY = vals[0]
		}
		if vals[1] > maxX {
			maxX = vals[1]
		}
		coords = append(coords, [2]int{vals[0],vals[1]})
	}

	return coords, maxY, maxX
}

func isValid(a [2]int, b [2]int, grid map[[2]int]rune) (bool,int) {
	xInts:= ints.GetIntsBetween(a[1], b[1])
	yInts:= ints.GetIntsBetween(a[0],b[0])
	for _,x:=range xInts {
		for _,y:= range yInts {
			if _,ok:= grid[[2]int{y,x}]; ok {
				return false,0
			}
		}
	}
	return true, (len(xInts)+2)*(len(yInts)+2)
}