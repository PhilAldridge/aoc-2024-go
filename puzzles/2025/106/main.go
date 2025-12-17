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
	fmt.Println("Part 2 answer: ", part2("input.txt"))
	
	fmt.Println()
	fmt.Println("Part 1: ", split.Sub(start))
	fmt.Println("Part 2: ", time.Since(split))
}

func part1(name string) int {
	lines:= files.ReadLines(name)	

	operations := getOperations(lines[len(lines)-1])

	total:=0

	for _, operation:= range operations {
		values := collectValuesHorizontal(lines[:len(lines)-1], operation.startIndex, operation.endIndex)
		total += calculateSubtotal(values, operation.operator)
	}

	return total
}

func part2(name string) int {
	lines:= files.ReadLines(name)	

	operations := getOperations(lines[len(lines)-1])

	total:=0

	for _, operation:= range operations {
		values := collectValuesVertical(lines[:len(lines)-1], operation.startIndex, operation.endIndex)
		total += calculateSubtotal(values, operation.operator)
	}

	return total
}

type operation struct {
	operator rune
	startIndex int
	endIndex int
}

func getOperations(finalLine string) []operation {
	var operations []operation

	for i,str:= range finalLine {
		if str != ' ' {
			if i != 0 {
				operations[len(operations)-1].endIndex = i
			}

			operations = append(operations, operation{
				operator: str,
				startIndex: i,
				endIndex: len(finalLine),
			})
		}
	}

	return operations
}

func collectValuesHorizontal(lines []string, start int, end int) []int {
	var values []int 
	for _, line:= range lines {
		nextValueString := ""

		for i:=start; i<end; i++ {
			if line[i] == ' ' {
				continue
			}

			nextValueString += string(line[i])
		}

		values = append(values, ints.FromString(nextValueString))
	}

	return values
}

func collectValuesVertical(lines []string, start int, end int) []int {
	var values []int 
	
	for i:=start; i<end; i++ {
		nextValueString := ""

		for _,line:= range lines {
			if line[i] == ' ' {
				continue
			}

			nextValueString += string(line[i])
		}

		if nextValueString == "" {
			continue
		}

		values = append(values, ints.FromString(nextValueString))
	}

	return values
}

func calculateSubtotal(values []int, operator rune) int {
	if operator == '*' {
		subtotal:= 1

		for _,value:= range values {
			subtotal *= value
		}

		return subtotal
	}

	subtotal:= 0

	for _,value:= range values {
		subtotal += value
	}

	return subtotal
}