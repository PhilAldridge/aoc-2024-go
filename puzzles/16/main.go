package main

import (
	"fmt"
	//"sort"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/bools"
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
	start, end:= getStartAndEnd(&lines)
	paths := getPaths(&lines)
	// sort.Slice(paths, func(i, j int) bool {
	// 	return paths[i].distance < paths[j].distance
	// })
	return calcRoute(paths, start, end)
}

func part2(name string) int {
	return 0
}

type path struct {
	start coords.Coord
	end coords.Coord
	startDir coords.Coord 
	endDir coords.Coord
	distance int
	squares int
}

func getStartAndEnd(lines *[]string) (coords.Coord,coords.Coord) {
	var start,end coords.Coord
	for i,line:= range *lines {
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

func getPaths(lines *[]string) []path {
	paths:= []path{}
	for i,line:= range *lines {
		for j,char:= range line {
			if char == '#' {
				continue
			}
			if bools.CountTrues(
				(*lines)[i][j-1]!='#',
				(*lines)[i][j+1]!='#',
				(*lines)[i-1][j]!='#',
				(*lines)[i+1][j]!='#',
			)>2 || char != '.' {
				paths = append(paths, getPathsFromNode(lines, coords.NewCoord(i,j))...)
			}
		}
	}
	return paths
}

func getPathsFromNode(lines *[]string,start coords.Coord) []path {
	paths:= []path{}
	dirs := start.GetAdjacent()
	for _,dir:= range dirs {
		if (*lines)[dir.I][dir.J] != '#' {
			end, pathCoords, endDir, turns := followPath(lines,dir,[]coords.Coord{start,dir},dir.Subtract(start),0)
			paths = append(paths, path{
				start: start,
				end:end,
				startDir: dir.Subtract(start),
				endDir: endDir,
				distance: len(pathCoords)-1 + turns*1000,
				squares: len(pathCoords)-1,
			})
		}
	}
	return paths
}

func followPath(lines *[]string,pos coords.Coord, pathSoFar []coords.Coord, dir coords.Coord, turns int) (coords.Coord, []coords.Coord, coords.Coord, int) {
	nextPoss := pos.GetAdjacent()
	var nextPosition coords.Coord
	nextPositionFound:= false
	prevPos := pos.Subtract(dir)
	for _,nextPos:= range nextPoss {
		if (prevPos.I != nextPos.I || prevPos.J != nextPos.J) && (*lines)[nextPos.I][nextPos.J] != '#' {
			if nextPositionFound {
				return pos, pathSoFar, dir, turns
			}
			nextPosition = nextPos
			nextPositionFound = true
		}
	}
	if !nextPositionFound || (*lines)[pos.I][pos.J] != '.' {
		return pos, pathSoFar, dir, turns
	}
	straightAhead := pos.Add(dir)
	if straightAhead.I == nextPosition.I && straightAhead.J == nextPosition.J {
		return followPath(lines,nextPosition,append(pathSoFar,nextPosition),nextPosition.Subtract(pos),turns)
	}
	
	return followPath(lines,nextPosition,append(pathSoFar,nextPosition),nextPosition.Subtract(pos),turns+1)
}

type node struct {
	pos coords.Coord
	dir coords.Coord
	Cost int
	Locked bool
	Checked bool
	squares int
}

func calcRoute(paths []path, start coords.Coord, end coords.Coord) int {
	nodeArr:= []node{
		{
			pos: start,
			dir: coords.NewCoord(0,1),
			Cost:0,
			Locked:true,
			Checked: false,
		},
		{
			pos: start,
			dir: coords.NewCoord(-1,0),
			Cost:1000,
			Locked:true,
			Checked: false,
		},
		{
			pos: start,
			dir: coords.NewCoord(0,-1),
			Cost:2000,
			Locked:true,
			Checked: false,
		},
		{
			pos: start,
			dir: coords.NewCoord(1,0),
			Cost:1000,
			Locked:true,
			Checked: false,
		},
	}

	for {
		for i,v:= range nodeArr {
			if v.Checked || !v.Locked {continue}
			nodeArr[i].Checked = true
			for _,path:=range paths {
				if path.start.I != v.pos.I || path.start.J != v.pos.J ||
				!path.startDir.IsSameDirectionAs(v.dir) {continue}
				nodeFound:= false
				for j,v2:=range nodeArr {
					if v2.Locked {continue}
					if v2.pos.I != path.end.I || v2.pos.J !=path.end.J {
						continue
					}
					nodeFound = true
					newCost:= v.Cost + path.distance + turn(path.endDir,v2.dir)
					if newCost < v2.Cost {
						nodeArr[j].Cost = newCost
						nodeArr[j].squares = v.squares + path.squares
					}
					if newCost == v2.Cost {
						nodeArr[j].squares += v.squares + path.squares
					}
					
				}
				if !nodeFound {
					newCost:= v.Cost + path.distance
					nodeArr = append(nodeArr, 
						node{
							pos:path.end,
							dir:path.endDir,
							Cost:newCost,
							Locked: false,
							Checked: false,
							squares : v.squares + path.squares,
						},
						node{
							pos:path.end,
							dir:coords.TurnLeft(path.endDir),
							Cost:newCost + 1000,
							Locked: false,
							Checked: false,
							squares : v.squares + path.squares,
						},
						node{
							pos:path.end,
							dir:coords.TurnRight(path.endDir),
							Cost:newCost + 1000,
							Locked: false,
							Checked: false,
							squares : v.squares + path.squares,
						},
						node{
							pos:path.end,
							dir:coords.TurnBack(path.endDir),
							Cost:2000 + newCost,
							Locked: false,
							Checked: false,
							squares : v.squares + path.squares,
						},
					)
				}
			}
		}

		minThisRound:= 8238947329874329
		keysInMin := []int{}

		for i,v:= range nodeArr {
			if v.Locked {continue}
			if v.Cost < minThisRound {
				minThisRound = v.Cost
				keysInMin = []int{i}
			} else if v.Cost == minThisRound {
				keysInMin = append(keysInMin,i)
			}
		}
		total:=0
		for _,v:=range keysInMin {
			if nodeArr[v].pos.I == end.I && nodeArr[v].pos.J==end.J {
				fmt.Println(total)
				total+= nodeArr[v].squares
			}
			nodeArr[v].Locked =true
		}
		if total != 0 {
			fmt.Println(total)
			return minThisRound
		}
		//fmt.Println(pathCosts)
	}


}

func turn(dir coords.Coord, targetDir coords.Coord) int {
	newDir := coords.NewCoord(dir.I,dir.J)
	if newDir.I == targetDir.I && newDir.J == targetDir.J {return 0}
	newDir = coords.TurnLeft(dir)
	if newDir.I == targetDir.I && newDir.J == targetDir.J {return 1000}
	newDir = coords.TurnRight(dir)
	if newDir.I == targetDir.I && newDir.J == targetDir.J {return 1000}
	if dir.I == -targetDir.I && dir.J == -targetDir.J {return 2000}
	panic("turn not found")

}