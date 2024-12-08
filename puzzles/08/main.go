package main

import (
	"fmt"
	"time"

	"github.com/kylehoehns/aoc-2023-go/pkg/files"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1: ", part1("input.txt"), " in:", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", part2("input.txt"), " in:", time.Since(start))
}

func part1(name string) int {
	antennaMap,iLen,jLen := getAntennaMap(name)
	newAntennaMap := getNewAntennaMap(antennaMap,iLen,jLen)
	return len(newAntennaMap)
}

func part2(name string) int {
	antennaMap, iLen, jLen := getAntennaMap(name)
	return 0
}

func getAntennaMap(name string) (map[rune][][2]int,int,int) {
	mapping:= make(map[rune][][2]int)
	lines:= files.ReadLines(name)
	for i,line:= range lines {
		for j,char:= range line {
			if char != '.' {
				mapping[char] = append(mapping[char], [2]int{i,j})
			}
		}
	}
	return mapping, len(lines),len(lines[0])
}

func getNewAntennaMap(oldMap map[rune][][2]int, iLen int, jLen int) map[[2]int]int {
	newMap := make(map[[2]int]int)
	for _,antennaList:= range oldMap {
		for i,antenna1:= range antennaList {
			for j,antenna2:= range antennaList {
				if i==j {
					continue
				}
				firstAntiPos, secondAntiPos := getAntiPos(antenna1,antenna2)
				if inBounds(firstAntiPos,iLen,jLen) {
					newMap[firstAntiPos] = newMap[firstAntiPos]+1
				}
				if inBounds(secondAntiPos,iLen,jLen) {
					newMap[secondAntiPos] = newMap[secondAntiPos]+1
				}
			}
		}
	} 
	return newMap
}

func getAntiPos(antenna1 [2]int, antenna2 [2]int) ([2]int,[2]int) {
	firstAntiPos := [2]int{2*antenna1[0]-antenna2[0],2*antenna1[1]-antenna2[1]}
	
	secondAntiPos := [2]int{2*antenna2[0]-antenna1[0],2*antenna2[1]-antenna1[1]}

	return firstAntiPos,secondAntiPos
}

func inBounds (pos [2]int, iLen int, jLen int) bool {
	if pos[0]< 0 || pos[1] <0 || pos[0]>= iLen || pos[1] >= jLen {
		return false
	}
	return true
}