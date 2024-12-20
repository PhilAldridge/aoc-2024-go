package main

import (
	"fmt"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/coords"
	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1: ", part1("input.txt"), " in:", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", part2("input.txt"), " in:", time.Since(start))
}

func part1(name string) int {
	grid:= files.ReadLines(name)
	distanceMap := make(map[coords.Coord]int)
	start:= findStart(grid)
	floodFill(grid,distanceMap,start)
	savingsCount:= make(map[int]int)
	glitch(distanceMap,savingsCount,2)	
	return addSavings(savingsCount,100)
}

func part2(name string) int {
	grid:= files.ReadLines(name)
	distanceMap := make(map[coords.Coord]int)
	start:= findStart(grid)
	floodFill(grid,distanceMap,start)
	savingsCount:= make(map[int]int)
	glitch(distanceMap,savingsCount,20)	
	return addSavings(savingsCount,100)
}

func findStart(grid []string) coords.Coord {
	for i,row:= range grid {
		for j,char:= range row {
			if char == 'S' {
				return coords.NewCoord(i,j)
			}
		}
	}
	panic("start not found")
}

func floodFill(grid []string, distanceMap map[coords.Coord]int, start coords.Coord) {
	distance:= 0
	positionsToCheck:= []coords.Coord{start}
	distanceMap[start] = 0
	for {
		distance ++
		nextChecks:= []coords.Coord{}
		for _,p:= range positionsToCheck {
			for _, adj:= range p.GetAdjacent() {
				if adj.I <0 || adj.J<0 || adj.I>=len(grid) || adj.J>=len(grid[0]) {
					continue
				}
				if _,ok:= distanceMap[adj]; !ok && grid[adj.I][adj.J]!='#' {
					distanceMap[adj] = distance
					nextChecks = append(nextChecks, adj)
				}
				if grid[adj.I][adj.J]=='E' {
					return
				}
			}
		}
		positionsToCheck = nextChecks
	}
}

func glitch(distanceMap map[coords.Coord]int, savingsCount map[int]int, maxGlitchDistance int) {
	for k1,v1:= range distanceMap {
		for k2,v2:= range distanceMap {
			manhattan := coords.ManhattanDistance(k1,k2)
			if manhattan > maxGlitchDistance {
				continue
			}
			saves:= v2-v1-manhattan
			if saves<=0 {
				continue
			}
			savingsCount[saves]++
		}
	}
}

func addSavings(savingsCount map[int]int, minSaving int) int {
	total:=0
	for k,v:= range savingsCount {
		if k <minSaving {
			continue
		}
		total +=v
	}
	return total
}