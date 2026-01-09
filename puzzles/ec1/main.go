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

func part1(name string) string {
	input := files.ReadParagraphs(name)

	names := strings.Split(input[0][0], ",")
	instructions := strings.Split(input[1][0], ",")

	index := 0

	for _, instruction := range instructions {
		if instruction[0] == 'L' {
			index -= ints.FromString(instruction[1:])
		} else {
			index += ints.FromString(instruction[1:])
		}

		if index < 0 {
			index = 0
		}
		if index >= len(names) {
			index = len(names) - 1
		}
	}

	return names[index]
}

func part2(name string) string {
	input := files.ReadParagraphs(name)

	names := strings.Split(input[0][0], ",")
	instructions := strings.Split(input[1][0], ",")

	index := 0

	for _, instruction := range instructions {
		if instruction[0] == 'L' {
			index -= ints.FromString(instruction[1:])
		} else {
			index += ints.FromString(instruction[1:])
		}
	}

	index = ints.Mod(index, len(names))

	return names[index]
}

func part3(name string) string {
	input := files.ReadParagraphs(name)

	names := strings.Split(input[0][0], ",")
	instructions := strings.Split(input[1][0], ",")

	nameMap := make(map[int]string)

	for i, name := range names {
		nameMap[i] = name
	}

	for _, instruction := range instructions {
		index := ints.Mod(ints.FromString(instruction[1:]), len(names))
		if instruction[0] == 'L' {
			index = ints.Mod(-index, len(names))
		}

		nameMap[0], nameMap[index] = nameMap[index], nameMap[0]
	}

	return nameMap[0]
}
