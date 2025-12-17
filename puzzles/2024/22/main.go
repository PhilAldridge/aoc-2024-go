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
	calcCache:= make(map[int]int)
	bananaCache:= make(map[[4]int]map[int]int)
	for i,seed:= range seeds {
		seedVals:= []int{seed%10}
		for i:=0;i<2000;i++ {
			if s,ok:= calcCache[seed];ok {
				seed = s
				seedVals = append(seedVals, seed%10)
				continue
			}
			seedNew := calcNext(seed)
			calcCache[seed] = seedNew
			seed = seedNew
			seedVals = append(seedVals, seed%10)
		}
		for j:= 4; j<len(seedVals); j++ {
			diffs:= [4]int{
				seedVals[j-3]-seedVals[j-4],
				seedVals[j-2]-seedVals[j-3],
				seedVals[j-1]-seedVals[j-2],
				seedVals[j]-seedVals[j-1],
			}
			if b,ok:= bananaCache[diffs]; ok {
				if _,ok:= b[i]; !ok {
					b[i] = seedVals[j]
				}
			} else {
				bananaCache[diffs] = make(map[int]int)
				bananaCache[diffs][i] = seedVals[j]%10
			}
		}
	}
	max:=0
	for _,v:=range bananaCache {
		vTotal:=0
		for i:=0;i<len(seeds);i++ {
			vTotal+=v[i]
		}
		if vTotal>max {
			max = vTotal
		}
	}
	return max
}

func calcNext(last int) int {
	mod:=16777216
	a:= ((last<<6)^last)%mod
	a = ((a>>5)^a)%mod
	a = ((a<<11)^a)%mod
	return a
}