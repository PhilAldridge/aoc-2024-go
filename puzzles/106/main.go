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
	fmt.Println("Part 2 answer: ", part2("input.txt"))
	
	fmt.Println()
	fmt.Println("Part 1: ", split.Sub(start))
	fmt.Println("Part 2: ", time.Since(split))
}

func part1(name string) int {
	lines:= files.ReadLines(name)	
	operations:= strings.Fields(lines[len(lines)-1])

	values:= make([][]int, len(lines)-1)

	for i:=0; i<len(lines)-1; i++ {
		values[i] = ints.FromStringSlice(strings.Fields(lines[i]))
	}


	total:=0
	for i, operation:= range operations {
		if operation == "*" {
			total += multiplyValues(values, i)
			continue
		}
		total += addValues(values,i)
	}

	return total
}

func part2(name string) int {
	lines:= files.ReadLines(name)
	total:= 0
	var operation rune
	var subTotal int

	for i:=0; i<len(lines[0]); i++ {
		switch lines[len(lines)-1][i] {
		case '*':
			operation = '*'
			subTotal = 1
		case '+':
			operation = '+'
			subTotal = 0
		}

		nextValString:= ""
		for j:=0; j<len(lines)-1; j++ {
			if lines[j][i]==' ' {
				continue
			}
			nextValString += string(lines[j][i])
		}

		if nextValString == "" {
			total += subTotal
			continue
		}

		if operation == '*' {
			subTotal *= ints.FromString(nextValString)
		} else {
			subTotal += ints.FromString(nextValString)
		}

		if i==len(lines[0])-1 {
			total += subTotal
		}
	}

	return total
}

func multiplyValues(values [][]int, i int) int {
	total:=1
	for _,value:=range values {
		total *= value[i]
	}
	return total
}

func addValues(values [][]int, i int) int {
	total:=0
	for _,value:= range values {
		total += value[i]
	}
	return total
}