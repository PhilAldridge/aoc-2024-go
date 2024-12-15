package main

import (
	"fmt"
	"sort"
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
	parts:= files.ReadParagraphs(name)
	warehouse := [][]byte{}
	for _,row := range parts[0] {
		warehouse = append(warehouse, []byte(row))
	}
	robotPos := findRobot(warehouse)
	for _,directions:= range parts[1] {
		for _,dir:= range directions {
			robotPos = moveRobot(&warehouse, dir, robotPos)
			// for _,line:= range warehouse {
			// 	fmt.Println(string(line))
			// }
		}
	}
	
	return calcGPS(warehouse)
}

func part2(name string) int {
	parts:= files.ReadParagraphs(name)
	warehouse := [][]byte{}
	for _,row := range parts[0] {
		newRow := []byte{}
		for _,pos := range row {
			switch pos {
			case '#':
				newRow = append(newRow, '#')
				newRow = append(newRow, '#')
			case '.':
				newRow = append(newRow, '.')
				newRow = append(newRow, '.')
			case '@':
				newRow = append(newRow, '@')
				newRow = append(newRow, '.')
			case 'O':
				newRow = append(newRow, '[')
				newRow = append(newRow, ']')
			}
		}
		warehouse = append(warehouse, newRow)
	}
	robotPos := findRobot(warehouse)
	for _,directions:= range parts[1] {
		for _,dir:= range directions {
			robotPos = moveRobot(&warehouse, dir, robotPos)
			
		}
	}
	// for _,line:= range warehouse {
	// 	fmt.Println(string(line))
	// }
	return calcGPS(warehouse)
}

func findRobot(warehouse [][]byte) coords.Coord {
	for iPos,row:= range warehouse {
		for jPos,point := range row {
			if point == '@' {
				return coords.NewCoord(iPos,jPos)
			}
		}
	}
	panic("robot not found!")
}

func moveRobot(warehouse *[][]byte, dir rune, robotLoc coords.Coord) coords.Coord {
	nextPos := getNextPos(robotLoc,dir)
	moveBox(warehouse, dir, nextPos)
	if (*warehouse)[nextPos.I][nextPos.J] == '.' {
		(*warehouse)[nextPos.I][nextPos.J], (*warehouse)[robotLoc.I][robotLoc.J] = 
			(*warehouse)[robotLoc.I][robotLoc.J],(*warehouse)[nextPos.I][nextPos.J]
		return nextPos
	}
	return robotLoc
}

func moveBox(warehouse *[][]byte, dir rune, pos coords.Coord) bool {
	nextPos:= getNextPos(pos,dir)
	switch (*warehouse)[pos.I][pos.J] {
	case 'O':
		if moveBox(warehouse,dir,nextPos) {
			(*warehouse)[nextPos.I][nextPos.J], (*warehouse)[pos.I][pos.J] = 
				(*warehouse)[pos.I][pos.J],(*warehouse)[nextPos.I][nextPos.J]
			return true
		}
		return false
	case '.':
		return true
	case '#':
		return false
	case '[':
		return moveBigBox(warehouse,dir,pos) 
	case ']':
		return moveBigBox(warehouse,dir,pos) 
	} 
	return false
}

func moveBigBox(warehouse *[][]byte, dir rune, pos coords.Coord) bool {
	blocksToMove, ok:= getBlocksToMove(warehouse,dir,pos)
	if ok {
		for _,block:= range blocksToMove {
			nextPos:= getNextPos(block,dir)
			(*warehouse)[nextPos.I][nextPos.J], (*warehouse)[block.I][block.J] = 
				(*warehouse)[block.I][block.J],(*warehouse)[nextPos.I][nextPos.J]
		}
		return true
	}
	return false
}

func getBlocksToMove(warehouse *[][]byte, dir rune, pos coords.Coord) ([]coords.Coord,bool) {
	blocksToMove:= []coords.Coord{pos}
	if (*warehouse)[pos.I][pos.J] == '[' {
		blocksToMove = append(blocksToMove, pos.Right(1))
	} else {
		blocksToMove = append(blocksToMove, pos.Left(1))
	}
	for {
		changesMade:= false
		for _, block:= range blocksToMove {
			blockSpace := getNextPos(block,dir)
			if coords.CoordInSlice(blockSpace,blocksToMove) {
				continue
			}
			switch (*warehouse)[blockSpace.I][blockSpace.J] {
			case '#':
				return []coords.Coord{},false
			case '[':
				blocksToMove = append(blocksToMove, blockSpace)
				blocksToMove = append(blocksToMove, blockSpace.Right(1))
				changesMade = true
			case ']':
				blocksToMove = append(blocksToMove, blockSpace)
				blocksToMove = append(blocksToMove, blockSpace.Left(1))
				changesMade = true
			}
		}
		if !changesMade {
			break
		} 
	}
	sort.Slice(blocksToMove, func(i,j int) bool {
		switch dir {
		case '^':
			return blocksToMove[i].I < blocksToMove[j].I
		case '<':
			return blocksToMove[i].J < blocksToMove[j].J
		case '>':
			return blocksToMove[i].J > blocksToMove[j].J
		case 'v':
			return blocksToMove[i].I > blocksToMove[j].I
	}
	panic("invalid direction: " + string(dir))
	})

	return blocksToMove,true
}

func getNextPos(pos coords.Coord, dir rune) coords.Coord {
	switch dir {
		case '^':
			return pos.Up(1)
		case '<':
			return pos.Left(1)
		case '>':
			return pos.Right(1)
		case 'v':
			return pos.Down(1)
	}
	panic("invalid direction: "+string(dir))
}

func calcGPS(warehouse [][]byte) int { 
	total:= 0
	for i,row:= range warehouse {
		for j,pos := range row {
			if pos == 'O'|| pos == '[' {
				total += 100*i + j
			}
		}
	}
	return total
}