package main

import (
	"fmt"
	"time"

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
	seeds:= ints.FromStringSlice(files.ReadLines(name))
	total:=0
	calcCache:= make(map[int]int)
	for _,seed:= range seeds {
		for i:=0;i<2000;i++ {
			if s,ok:= calcCache[seed];ok {
				seed = s
				continue
			}
			seedNew := calcNext(seed)
			calcCache[seed] = seedNew
			seed = seedNew
		}
		total += seed
	}
	
	return total
}

func part2(name string) int {
	seeds:= ints.FromStringSlice(files.ReadLines(name))
	total:=0
	calcCache:= make(map[int]int)
	for _,seed:= range seeds {
		for i:=0;i<2000;i++ {
			if s,ok:= calcCache[seed];ok {
				seed = s
				continue
			}
			seedNew := calcNext(seed)
			calcCache[seed] = seedNew
			seed = seedNew
		}
		total += seed
	}
	return total
}

func calcNext(last int) int {
	a:= ((last<<6)^last)%16777216
	a = ((a>>5)^a)%16777216
	a = ((a<<11)^a)%16777216
	return a
}