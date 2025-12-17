package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
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
	input:= files.Read(name)

	for _, char:= range input {
		total += getScore(char)
	}

	return total
}

func part2(name string) int {
	total:=0
	input:= files.Read(name)

	for i:= 0; i<len(input)/2; i++ {
		fmt.Println(total)
		c1:= input[2*i]
		c2:= input[2*i+1]

		if c1 != 'x' && c2 !='x' {
			total+=2
		}

		s1:= getScore(rune(c1))
		s2:= getScore(rune(c2))

		total += s1 + s2
	}

	return total
}


func part3(name string) int {
	total:=0
	input:= files.Read(name)

	for i:= 0; i<len(input)/3; i++ {
		fmt.Println(total)
		c1:= input[3*i]
		c2:= input[3*i+1]
		c3:= input[3*i+2]

		xCount := strings.Count(input[3*i:3*(i+1)],"x")
		switch xCount {
		case 0:
			total += 6
		case 1:
			total += 2
		}

		s1:= getScore(rune(c1))
		s2:= getScore(rune(c2))
		s3:= getScore(rune(c3))

		total += s1 + s2 + s3
	}

	return total
}

func getScore(char rune) int {
	switch char {
		case 'A':
			return 0
		case 'B':
			return 1
		case 'C':
			return 3
		case 'D':
			return 5
		}
	return 0
}
