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
	gears:= ints.FromStringSlice(files.ReadLines(name))
	total:= (2025 * gears[0])/gears[len(gears)-1]

	return total
}

func part2(name string) int {
	gears:= ints.FromStringSlice(files.ReadLines(name))
	total:= (10000000000000 * gears[len(gears)-1])/gears[0]

	if (10000000000000 * gears[len(gears)-1])%gears[0] !=0 {
		total++
	}


	return total
}


func part3(name string) int {
	input:=files.ReadLines(name)

	top:=1
	bottom:=1

	for i,gear:= range input {
		switch i {
		case 0:
			top *= ints.FromString(gear)
		case len(input)-1:
			bottom *= ints.FromString(gear)
		default:
			split:= strings.Split(gear,"|")
			bottom *=ints.FromString(split[0])
			top *= ints.FromString(split[1])
		}
		hcf := ints.GCD(top,bottom)
		top /= hcf
		bottom /= hcf

	}
	total:=(100*top)/bottom

	return total
}

