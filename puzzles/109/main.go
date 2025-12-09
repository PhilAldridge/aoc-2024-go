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

	grid:= make(map[[2]int]bool)
	
	lastCoord:=coords[len(coords)-1]
	grid[lastCoord] = true
	for i:=0; i<len(coords); i++ {
		grid[coords[i]] = true

		if coords[i][0] != lastCoord[0] {
			yFill:= ints.GetIntsBetween(coords[i][0], lastCoord[0])
			for _,y:= range yFill {
				grid[[2]int{y,lastCoord[1]}] = true
			}
		} else {
			xFill:= ints.GetIntsBetween(coords[i][1], lastCoord[1])
			for _,x:= range xFill {
				grid[[2]int{lastCoord[0],x}] = true
			}
		}

		lastCoord = coords[i]
	}

	max:=0

	IsValid := func (i,j int) bool {
		for k:=0; k<len(coords); k++ {
			if k==i || k==j {
				continue
			}
			if ints.IsBetween(coords[k][0], coords[i][0], coords[j][0]) &&
				ints.IsBetween(coords[k][1], coords[i][1], coords[j][1]) {
					return false
				}
		}

		count:=0
		midY:= (coords[i][0]+coords[j][0])/2
		midX:= (coords[i][1]+coords[j][1])/2
		if coords[i][0] != coords[j][0] {
			for k:=0; k<=midY; k++ {
				if grid[[2]int{k,midX}] {
					count ++
				}
			}
		} else {
			for k:=0; k<=midX; k++ {
				if grid[[2]int{midY,k}] {
					count ++
				}
			}
		}
		if count %2 == 0 {
			return false
		}
		
		yInts :=ints.GetIntsBetween(coords[i][0], coords[j][0])
		xInts :=ints.GetIntsBetween(coords[i][1],coords[j][1])

		if len(yInts)>0 {
			for _,x:= range xInts {
				if grid[[2]int{yInts[0],x}] {
					return false
				}
				if grid[[2]int{yInts[len(yInts)-1],x}] {
					return false
				}
			}
		}
		
		if len(xInts)>0 {
			for _,y:= range yInts {
			if grid[[2]int{y,xInts[0]}] {
				return false
			}
			if grid[[2]int{y,xInts[len(xInts)-1]}] {
				return false
			}

		}
		}
		

		return true
	}

	for i:=0;i<len(coords); i++ {
		for j:=i+1; j<len(coords); j++ {
			if IsValid(i,j) {
				area:= (ints.Abs(coords[i][0] - coords[j][0])+1) *
				 	(ints.Abs(coords[i][1] - coords[j][1])+1) 
			
				if area > max {
					max = area
				}
			}
		}
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