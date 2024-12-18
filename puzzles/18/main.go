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
	grid,gridSize:= createGrid(name)
	leftToDrop:= leftToDrop(name)
	for _,pos:= range leftToDrop {
		grid[pos.I][pos.J] = '#'
		if floodFill(grid,gridSize)==-1 {
			return fmt.Sprintf("%d,%d",pos.J,pos.I)
		}
	}
	for _,row:= range grid {
		
		fmt.Println(row)
	}
	panic("pos not found")
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

func leftToDrop(name string) []coords.Coord {
	linesSplit:= files.ReadParagraphs(name)
	memory:= ints.FromString(linesSplit[1][0])
	positions:= []coords.Coord{}
	for i, pos:= range linesSplit[2] {
		if i< memory {
			continue
		}
		vals:= ints.FromStringSlice(strings.Split(pos,","))
		positions = append(positions, coords.NewCoord(vals[1],vals[0]))
	}
	return positions
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