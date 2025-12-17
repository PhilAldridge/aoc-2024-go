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
	lines:= files.ReadLines(name)
	start, end:= getStartAndEnd(lines)
	return calcRoute(lines, start, end, false)
}

func part2(name string) int {
	lines:= files.ReadLines(name)
	start, end:= getStartAndEnd(lines)
	return calcRoute(lines, start, end, true)
}

func getStartAndEnd(lines []string) (coords.Coord,coords.Coord) {
	var start,end coords.Coord
	for i,line:= range lines {
		for j,char:= range line {
			if char == 'S' {
				start = coords.NewCoord(i,j)
			} else if char == 'E' {
				end = coords.NewCoord(i,j)
			}
		}
	}
	return start, end
}

type pathCost struct {
	minCost int
	squaresVisited []coords.Coord
}


func calcRoute(grid []string, start coords.Coord, end coords.Coord,outputPart2 bool) int {
	pathCostMap:= make(map[[3]int]pathCost)
	pathCostMap[[3]int{start.I,start.J,0}] = pathCost{
		minCost: 0,
		squaresVisited: []coords.Coord{},
	}

	lockedCoords:= [][3]int{{start.I,start.J,0}}
	coordsToCheck:= [][3]int{{start.I,start.J,0}}

	for len(coordsToCheck) >0 {
		//Check paths
		for _,coord:= range coordsToCheck {
			for i:=0; i<4; i++ {
				nextCoord:= coords.NewCoord(coord[0],coord[1])
				cost:= 1000
				if coord[2]-i ==2 || i-coord[2]==2 {
					cost= 2000
				}
				if i == coord[2] {
					nextCoord = nextCoord.Add(coords.DirectionsInOrder[i])
					cost =1
				}
				if inMap([3]int{nextCoord.I,nextCoord.J,i},lockedCoords) {
					continue
				}
				cost+= pathCostMap[coord].minCost

				if grid[nextCoord.I][nextCoord.J] == '#' {
					continue
				}
				_, ok := pathCostMap[[3]int{nextCoord.I,nextCoord.J,i}]
				if !ok {
					//add to pathCostMap
					addToMap(pathCostMap, nextCoord, coord,i,cost)
				} else {
					//update pathCostMap if cheaper
					updateMap(pathCostMap, nextCoord, coord,i,cost)
				}
			}
		}
		
		//prune
		for _,c:= range coordsToCheck {
			if c[0]==end.I && c[1]==end.J {
				continue
			}
			delete(pathCostMap,c)
		}
		//Find min coords not locked and add to ToCheck and Locked
		coordsToCheck = getMinUnlockedCoords(pathCostMap,lockedCoords)
		lockedCoords = append(coordsToCheck, lockedCoords...)
	}
	
	minCost:= 297489323298474893
	squares:= []coords.Coord{}
	for i:=0; i<4; i++ {
		nextCost := pathCostMap[[3]int{end.I, end.J, i}].minCost
		if nextCost == 0 {
			continue
		}
		if nextCost < minCost {
			minCost = nextCost
			squares = pathCostMap[[3]int{end.I, end.J, i}].squaresVisited
		}
		if nextCost == minCost {
			for _,path:= range pathCostMap[[3]int{end.I, end.J, i}].squaresVisited {
				if !coords.CoordInSlice(path,squares) {
					squares = append(squares, path)
				}
			}
		}
	}
	
	if outputPart2 {
		return len(squares) + 1
	}
	return minCost
}

func getMinUnlockedCoords(pathCostMap map[[3]int]pathCost, lockedCoords [][3]int) [][3]int {
	coordsToCheck:= [][3]int{}
	currentMin := 94723792435789843
	for k,v:= range pathCostMap {
		if inMap(k,lockedCoords) {
			continue
		}
		if v.minCost < currentMin {
			currentMin = v.minCost
			coordsToCheck = [][3]int{k}
		} else if v.minCost == currentMin {
			if !inMap(k,coordsToCheck) {
				coordsToCheck = append(coordsToCheck, k)
			}
		}
	}
	return coordsToCheck
}

func inMap(a [3]int, b [][3]int) bool {
	for _,v:= range b {
		if a[0]==v[0] && a[1]==v[1] && a[2]==v[2] {
			return true
		}
	}
	return false
}

func updateMap(pathCostMap map[[3]int]pathCost, nextCoord coords.Coord, prevPos [3]int, newDir int,cost int) {
	lastPathCost := pathCostMap[prevPos]
	squares:= append([]coords.Coord{coords.NewCoord(prevPos[0],prevPos[1])},lastPathCost.squaresVisited...)
		currentPathCost:= pathCostMap[[3]int{nextCoord.I,nextCoord.J,newDir}]
		currentPathSquares:= append([]coords.Coord{},currentPathCost.squaresVisited...)
		if currentPathCost.minCost < cost {
			return
		}
		if currentPathCost.minCost == cost {
			for _,path:= range squares {
				if !coords.CoordInSlice(path,currentPathSquares) {
					currentPathSquares = append(currentPathSquares, path)
				}
			}
			pathCostMap[[3]int{nextCoord.I,nextCoord.J, newDir}] = pathCost{
				minCost: cost,
				squaresVisited: currentPathSquares,
			}
		}
		if currentPathCost.minCost > cost {
			pathCostMap[[3]int{nextCoord.I,nextCoord.J, newDir}] = pathCost{
				minCost: cost,
				squaresVisited: squares,
			}
		}
}

func addToMap(pathCostMap map[[3]int]pathCost, nextCoord coords.Coord, prevPos [3]int, startDir int, cost int) {
	lastPathCost := pathCostMap[prevPos]
		prevCoord:= coords.NewCoord(prevPos[0],prevPos[1])
		newSquaresVisited := append([]coords.Coord{},lastPathCost.squaresVisited...)
		if !coords.CoordInSlice(prevCoord,newSquaresVisited) {
			newSquaresVisited = append(newSquaresVisited, prevCoord)
		}
		pathCostMap[[3]int{nextCoord.I,nextCoord.J, startDir}] = pathCost{
			minCost: cost,
			squaresVisited: newSquaresVisited,
		}
}