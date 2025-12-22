package main

import (
	"fmt"
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
	blocksGiven:= ints.FromString(files.Read(name))

	blocksUsed:= 0
	targetWidth:= -1

	for blocksUsed < blocksGiven {
		targetWidth +=2
		blocksUsed += targetWidth
	}

	return targetWidth * (blocksUsed - blocksGiven)
}

func part2(name string) int {
	input:= ints.FromStringSlice( files.ReadLines(name))
	multiplier:= input[0]
	modulo:= input[1]
	blocks:= input[2]

	blocksUsed:= 1
	targetWidth:= 1
	thickness:=1

	for blocksUsed < blocks {
		targetWidth +=2
		thickness = (thickness * multiplier) % modulo
		blocksUsed += targetWidth * thickness
	}

	return targetWidth * (blocksUsed - blocks)
}


func part3(name string) int {
	input:= ints.FromStringSlice( files.ReadLines(name))
	multiplier:= input[0]
	modulo:= input[1]
	blocks:= input[2]

	columns:= []int{1}
	blocksRemoved:=0
	thickness:= 1

	for ints.Sum(columns)- blocksRemoved < blocks {
		nextLayer:= make([]int,len(columns)+2)
		thickness = (thickness * multiplier) % modulo + modulo
		blocksRemoved = 0
		for i:= range nextLayer {
			columnIndex:= i-1
			nextLayer[i] = thickness
			if columnIndex >=0 && columnIndex <len(columns) {
				nextLayer[i] += columns[columnIndex]
				blocksRemoved += (multiplier * len(nextLayer) * nextLayer[i]) % modulo
			}
		}

		columns = nextLayer
	}

	return ints.Sum(columns)- blocksRemoved - blocks
}

