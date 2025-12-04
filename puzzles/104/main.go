package main

import (
	"fmt"
	"time"

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
	paperMap:= createPaperMap(lines)

	total:=0
	for paperLocation := range paperMap {
		if len(getAdjacents(paperLocation[0],paperLocation[1],paperMap)) <4 {
			total++
		}
	}
	return total
}

func part2(name string) int {
	lines:= files.ReadLines(name)
	paperMap:= createPaperMap(lines)
	mapToTest:= createPaperMap(lines)

	total:=0
	for len(mapToTest)>0 {
		newMap := make(map[[2]int]int)
		for paperLocation := range mapToTest {
			if _,ok := paperMap[paperLocation]; !ok {
				continue
			}
			adjacents:= getAdjacents(paperLocation[0],paperLocation[1],paperMap)
			if len(adjacents) <4 {
				total++
				delete(paperMap,paperLocation)
				for _,coord:= range adjacents {
					newMap[coord] = 1
				}
			}
		}
		mapToTest = newMap
	}
	
	return total
}

func createPaperMap(input []string) map[[2]int]int {
	res:= make(map[[2]int]int)

	for y,row:=range input {
		for x,col:= range row {
			if col == '@' {
				res[[2]int{x,y}] = 1
			}
		}
	}

	return res
}

func getAdjacents(x int, y int, paperMap map[[2]int]int) [][2]int {
	affected:= [][2]int{}
	for i:=-1;i<=1;i++ {
		for j:=-1;j<=1;j++ {
			coord:= [2]int{x+i,y+j}
			if _,ok := paperMap[coord];ok && !(i==0 && j==0) {
				affected = append(affected,coord)
			}
		}
	}
	return affected
}