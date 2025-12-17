package main

import (
	"fmt"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
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
	lines:= files.ReadLines(name)
	paperMap:= createAdjacencyMap(lines)

	total:=0
	for _, adjacents := range paperMap {
		if len(adjacents) <4 {
			total++
		}
	}
	return total
}

func part2(name string) int {
	lines:= files.ReadLines(name)
	paperMap:= createAdjacencyMap(lines)

	total:=0
	for pos, adjs:= range paperMap {
		if len(adjs)<4 {
			total += removeAdjacents(pos,adjs,paperMap)
		}
	}
	
	return total
}

func createAdjacencyMap(input []string) (map[[2]int][][2]int ) {
	resMap:= make(map[[2]int][][2]int)	

	for y,row:=range input {
		for x,col:= range row {
			if col == '@' {
				pos:= [2]int{x,y}
				for i:=-1; i<=1;i++ {
					for j:=0; j<=1; j++ {
						if (i==0 && j==0) {
							continue
						}

						if _, ok:= resMap[pos]; !ok {
							resMap[pos] = [][2]int{}
						}

						adj:= [2]int{x-i,y-j}

						if _,ok:= resMap[adj]; ok {
							resMap[adj] = append(resMap[adj], pos)
							resMap[pos] = append(resMap[pos], adj)
						}
					}
				}
			}
		}
	}

	return resMap
}

func removeAdjacents(pos [2]int, adjs [][2]int, paperMap map[[2]int][][2]int) int {
	total := 1
	delete(paperMap,pos)

	for _,adj:= range adjs {
		adjTwos, ok := paperMap[adj]
		if !ok {
			continue
		}

		if len(adjTwos) < 5 {
			total += removeAdjacents(adj, adjTwos, paperMap)
			continue
		}

		newAdjTwo:= [][2]int{}
		for _,adjTwo := range adjTwos {
			if adjTwo == pos {
				continue
			}
			newAdjTwo = append(newAdjTwo, adjTwo)
		}
		paperMap[adj] = newAdjTwo
	}

	return total
}