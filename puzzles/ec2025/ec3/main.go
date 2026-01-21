package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/ints"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1 answer: ", part1("input.txt"))
	split := time.Now()
	fmt.Println("Part 2 answer: ", part2("input2.txt"))
	split2 := time.Now()
	fmt.Println("Part 3 answer: ", part3("input3.txt"))

	
	fmt.Println()
	fmt.Println("Part 1: ", split.Sub(start))
	fmt.Println("Part 2: ", split2.Sub(split))
	fmt.Println("Part 2: ", time.Since(split2))
}

func part1(name string) int {
	crates:= ints.FromStringSlice(strings.Split(files.Read(name),","))
	crateMap:= make(map[int]bool)
	total:= 0

	for _, crate:= range crates {
		if _,ok:= crateMap[crate]; ok {
			continue
		}

		crateMap[crate] = true
		total += crate
	}

	return total
}

func part2(name string) int {
	crates:= ints.FromStringSlice(strings.Split(files.Read(name),","))
	crateMap:= make(map[int]bool)
	total:= 0

	for _, crate:= range crates {
		if _,ok:= crateMap[crate]; ok {
			continue
		}

		crateMap[crate] = true
	}

	crateSlice := []int{}
	for crate,_ := range crateMap {
		crateSlice = append(crateSlice, crate)
	}

	slices.Sort(crateSlice)

	for i:= range 20 {
		total += crateSlice[i]
	}

	return total
}


func part3(name string) int {
	crates:= ints.FromStringSlice(strings.Split(files.Read(name),","))
	crateMap:= make(map[int]int)
	max:= 0

	for _, crate:= range crates {
		crateMap[crate] ++
	}

	for _,count := range crateMap {
		if count > max {
			max = count
		}
	}

	return max
}

