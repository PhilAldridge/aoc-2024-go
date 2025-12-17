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
	keys,locks,height:= parseInput(name)
	total:=0
	for _,k:= range keys {
		for _,l:= range locks {
			if matchCheck(k,l,height) {
				total ++
			}
		}
	}
	return total
}

func part2(name string) int {
	return 0
}

func parseInput(name string) ([][]int, [][]int, int) {
	keys:= [][]int{}
	locks:=[][]int{}
	entries:= files.ReadParagraphs(name)
	for _,e:= range entries {
		colCount:= make([]int,len(e[0]))
		for _,row:= range e {
			for i,val:= range row {
				if val == '#' {
					colCount[i]++
				}
			}
		}
		
		if e[0][0] == '#' {
			locks = append(locks, colCount)
		} else {
			keys = append(keys, colCount)
		}
	}
	return keys,locks,len(entries[0])
}

func matchCheck(key []int, lock []int, height int) bool {
	for i:= 0 ; i< len(key); i++ {
		if key[i]+lock[i]> height {
			return false
		}
	}
	return true
}