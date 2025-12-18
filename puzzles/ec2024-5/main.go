package main

import (
	"fmt"
	"slices"
	"strconv"
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
	total:= 0
	dancers:= parseInput(name)

	for i:=0; i<10; i++ {
		dancers, total = dance(dancers,i)
	}

	return total
}

func part2(name string) int {
	dancers:= parseInput(name)
	shoutCount:= make(map[int]int)
	totalVal:=0
	
	i:=0
	for {
		dancers, totalVal = dance(dancers,i)

		shoutCount[totalVal] ++
		i++

		if shoutCount[totalVal] == 2024 {
			break
		}
	}

	return totalVal * i
}


func part3(name string) int {
	dancers:= parseInput(name)
	shoutCount:= make(map[int]int)
	totalVal:=0
	
	i:=0
	for {
		dancers, totalVal = dance(dancers,i)

		shoutCount[totalVal] ++
		i++

		if shoutCount[totalVal] == 2024 {
			break
		}
	}

	max:=0

	for shout:= range shoutCount {
		if shout > max {
			max = shout
		}
	}

	return max
}

func parseInput(name string) [][]int {
	lines:= files.ReadLines(name)
	horizontal:= make([][]int, len(lines))

	for i, line:= range lines {
		horizontal[i] = ints.FromStringSlice(strings.Split(line, " "))
	}

	result:= make([][]int, len(horizontal[0]))

	for i:=0; i<len(horizontal[0]); i++ {
		for _, h:= range horizontal {
			result[i] = append(result[i], h[i])
		}
	}

	return result
}

func dance(dancers [][]int, index int) ([][]int, int) {
	movingDancer := dancers[index%len(dancers)][0]
	dancers[index%len(dancers)] = dancers[index%len(dancers)][1:]

	lineIndex:= (index+1)%len(dancers)
	quotient:= (movingDancer-1)/len(dancers[lineIndex])
	remainder:= movingDancer % len(dancers[lineIndex])
	if remainder == 0 {
		remainder = len(dancers[lineIndex])
	}

	if quotient%2 == 0 {
		dancers[lineIndex] = slices.Insert(dancers[lineIndex], remainder - 1, movingDancer)
	} else {
		dancers[lineIndex] = slices.Insert(dancers[lineIndex], len(dancers[lineIndex]) - remainder + 1, movingDancer)
	}

	total:= ""
	for _,line:= range dancers {
		total += strconv.Itoa(line[0])
	}

	return dancers, ints.FromString(total)
}