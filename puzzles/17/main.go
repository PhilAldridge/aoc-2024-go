package main

import (
	"fmt"
	"strconv"
	"strings"
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

func part1(name string) string {
	pointer := 0
	lines := files.ReadLines(name)
	A := ints.FromString(strings.Replace(lines[0], "Register A: ", "", -1))
	B := ints.FromString(strings.Replace(lines[1], "Register B: ", "", -1))
	C := ints.FromString(strings.Replace(lines[2], "Register C: ", "", -1))
	program := ints.FromStringSlice(strings.Split(
		strings.Replace(lines[4], "Program: ", "", -1),
		","))
	stopper := 0
	output := []string{}
	for pointer < len(program)-1 && stopper < 1000 {
		calculate(&A, &B, &C, program[pointer], program[pointer+1], &pointer, &output)
		stopper++
	}
	return strings.Join(output, ",")
}

func part2(name string) int {
	pointer := 0
	lines := files.ReadLines(name)
	Bstatic := ints.FromString(strings.Replace(lines[1], "Register B: ", "", -1))
	Cstatic := ints.FromString(strings.Replace(lines[2], "Register C: ", "", -1))
	program := ints.FromStringSlice(strings.Split(
		strings.Replace(lines[4], "Program: ", "", -1),
		","))
	
	prev := []int{5}
	output := []string{}
	for len(output) < len(program) {
		newPrev:= []int{}
		for _,p:= range prev {
			for a := p * 8; a < p*8+8; a++ {
				output = []string{}
				A := a
				B := Bstatic
				C := Cstatic
				pointer = 0
				stopper := 0
				
				for pointer < len(program)-1 && stopper < 10000 {
					calculate(&A, &B, &C, program[pointer], program[pointer+1], &pointer, &output)
					stopper++
				}
				if ints.FromString(output[0])==program[len(program)-len(output)] {
					fmt.Println(output)
					newPrev = append(newPrev, a)
				}
		
				fmt.Printf("a: %d, output: %v\n", a, output)
			}
		}
		prev = newPrev
		
	}
	

	panic("not found")
}

func calculate(A *int, B *int, C *int, instruction int, operand int, pointer *int, output *[]string) {
	combo := operand
	switch combo {
	case 4:
		combo = *A
	case 5:
		combo = *B
	case 6:
		combo = *C
	}
	switch instruction {
	case 0:
		*A = *A >> combo
		*pointer += 2
	case 1:
		*B = *B ^ operand
		*pointer += 2
	case 2:
		*B = combo % 8
		*pointer += 2
	case 3:
		if *A == 0 {
			*pointer += 2
		} else {
			*pointer = operand
		}
	case 4:
		*B = *B ^ *C
		*pointer += 2
	case 5:
		*output = append(*output, strconv.Itoa(combo%8))
		*pointer += 2
	case 6:
		*B = *A >> combo
		*pointer += 2
	case 7:
		*C = *A >> combo
		*pointer += 2
	}
}

func calculate2(A *int, B *int, C *int, instruction int, operand int, pointer *int, output *[]string, expected int) int {
	combo := operand
	switch combo {
	case 4:
		combo = *A
	case 5:
		combo = *B
	case 6:
		combo = *C
	}
	switch instruction {
	case 0:
		*A = *A >> combo
		*pointer += 2
	case 1:
		*B = *B ^ operand
		*pointer += 2
	case 2:
		*B = combo % 8
		*pointer += 2
	case 3:
		if *A == 0 {
			*pointer += 2
		} else {
			*pointer = operand
		}
	case 4:
		*B = *B ^ *C
		*pointer += 2
	case 5:
		out := combo % 8
	
		*output = append(*output, strconv.Itoa(out))
		*pointer += 2
		if out == expected {
			return 1
		}
		return 0
	case 6:
		*B = *A >> combo
		*pointer += 2
	case 7:
		*C = *A >> combo
		*pointer += 2
	}
	return -1
}
