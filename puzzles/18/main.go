package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/coords"
	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/ints"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1: ", part1("input.txt"), " in:", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", part2("input.txt"), " in:", time.Since(start))
}

func part1(name string) int {
	grid,gridSize:= createGrid(name)
	return floodFill(grid, gridSize)
}

func part2(name string) string {
	file:= files.ReadParagraphs(name)
	gridSize := ints.FromString(file[0][0])
	grouping:= make(map[coords.Coord]string)
	for _, pos:= range file[2] {
		vals:= ints.FromStringSlice(strings.Split(pos,","))
		position:= coords.NewCoord(vals[1],vals[0])
		if position.I ==0 || position.J == gridSize-1 {
			grouping[position] = "topright"
		} else if position.J == 0 || position.I == gridSize-1 {
			grouping[position] = "bottomleft"
		} else {
			grouping[position] = "ungrouped"
		}
		if updateAdjacent(position,grouping) {
			return pos
		}
	}
	panic("no blockage found")
}

func updateAdjacent(pos coords.Coord, grouping map[coords.Coord]string) bool {
	groupFoundName:= grouping[pos]
	adjacentSquares := [8]coords.Coord{
		pos.Add(coords.NewCoord(1,1)),
		pos.Add(coords.NewCoord(1,-1)),
		pos.Add(coords.NewCoord(-1,1)),
		pos.Add(coords.NewCoord(-1,-1)),
		pos.Add(coords.NewCoord(0,1)),
		pos.Add(coords.NewCoord(0,-1)),
		pos.Add(coords.NewCoord(1,0)),
		pos.Add(coords.NewCoord(-1,0)),
	}
	for _, adj:= range adjacentSquares {
		if adjGroup, ok:= grouping[adj];ok && adjGroup != "ungrouped" {
			if groupFoundName != "ungrouped" && groupFoundName != adjGroup {
				return true
			}
			groupFoundName = adjGroup
			grouping[pos] = groupFoundName
		}
	}
	if groupFoundName == "ungrouped" {
		return false
	}
	for _, adj:= range adjacentSquares {
		if adjGroup, ok:= grouping[adj];ok && adjGroup=="ungrouped" {
			grouping[adj] = groupFoundName
			if updateAdjacent(adj,grouping) {
				return true
			}
		}
	}
	return false
}

func createGrid(name string) ([][]byte,int) {
	linesSplit:= files.ReadParagraphs(name)
	gridSize:= ints.FromString(linesSplit[0][0])
	memory:= ints.FromString(linesSplit[1][0])
	grid:= [][]byte{}
	positions:= []coords.Coord{}
	for i, pos:= range linesSplit[2] {
		if i>= memory {
			break
		}
		vals:= ints.FromStringSlice(strings.Split(pos,","))
		positions = append(positions, coords.NewCoord(vals[1],vals[0]))
	}
	for i:=0; i<gridSize; i++ {
		row:= []byte{}
		for j:=0; j<gridSize; j++ {
			if coords.CoordInSlice(coords.NewCoord(i,j),positions) {
				row = append(row, '#')
			} else {
				row = append(row, '.')
			}
		}
		grid = append(grid, row)
	}
	return grid, gridSize
}

func floodFill(grid [][]byte, gridSize int) int {
	distance:= 0
	positionsReached:= []coords.Coord{coords.NewCoord(0,0)}
	positionsToCheck:= []coords.Coord{coords.NewCoord(0,0)}
	for distance < gridSize*gridSize {
		distance ++
		nextChecks:= []coords.Coord{}
		for _,p:= range positionsToCheck {
			for _, adj:= range p.GetAdjacent() {
				if adj.I == gridSize-1 && adj.J == gridSize-1 {
					return distance
				}
				if adj.I <0 || adj.J<0 || adj.I>=gridSize || adj.J>=gridSize {
					continue
				}
				if grid[adj.I][adj.J]=='.' && !coords.CoordInSlice(adj,positionsReached) {
					positionsReached = append(positionsReached, adj)
					nextChecks = append(nextChecks, adj)
				}
			}
		}
		positionsToCheck = nextChecks
	}
	return -1	
}