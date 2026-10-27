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
	cones:= ints.FromStringSlice(strings.Split(files.Read(name),","))
	n:=1
	nSquared:=n*n
	total := 0

	for _, cone:= range cones {
		for cone > nSquared {
			n++
			nSquared = n*n
		}

		total += nSquared - cone
	}

	return total
}

func part2(name string) int {
	cones:= ints.FromStringSlice(strings.Split(files.Read(name),","))
	n:=8
	total := 0

	for _, cone:= range cones {
		for cone > n {
			n+=4
		}

		total += n - cone
	}

	return total
}

func part3(name string) int {
	cones:= ints.FromStringSlice(strings.Split(files.Read(name),","))
	n:=3
	nSquared:=n*n-1
	total := 0

	for _, cone:= range cones {
		for cone > nSquared {
			n+=2
			nSquared = n*n-1
		}

		total += nSquared - cone
	}

	return total
}
