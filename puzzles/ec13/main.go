package main

import (
	"fmt"
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
	input:= files.ReadLines(name)
	vals:= ints.FromStringSlice(input)
	dial:= make([]int,len(input)+1)
	evenI, oddI:= 1,len(input)
	dial[0] = 1

	for i,val:= range vals {
		if i%2==0 {
			dial[evenI] = val
			evenI++
		} else {
			dial[oddI] = val
			oddI --
		}
	}

	return dial[2025%len(dial)]
}

func part2(name string) int {
	input:= files.ReadLines(name)
	dial:= createDialGivenRanges(input)

	return dial[20252025%len(dial)]
}


func part3(name string) int {
	input:= files.ReadLines(name)
	dial:= createDialGivenRanges(input)

	return dial[202520252025%len(dial)]
}

func parseRange(input string) (int,int) {
	split:= strings.Split(input,"-")
	vals:= ints.FromStringSlice(split)

	return vals[0],vals[1]
}

func createDialGivenRanges(input []string) []int {
	dialLeft, dialRight:= []int{}, []int{1}

	for i,rangeString:= range input {
		start,end:= parseRange(rangeString)

		if i%2==0 {
			dialRight = append(dialRight, ints.GetIntsBetweenInclusive(start,end)...)
		} else {
			dialLeft = append(dialLeft, ints.GetIntsBetweenInclusive(start,end)...)
		}
	}

	for i:= len(dialLeft)-1; i>=0; i-- {
		dialRight = append(dialRight, dialLeft[i])
	}

	return dialRight
}
