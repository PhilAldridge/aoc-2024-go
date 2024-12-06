package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/kylehoehns/aoc-2023-go/pkg/files"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1: ", part1("input.txt"), " in:", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", part2("input.txt"), " in:", time.Since(start))
}

func part1(name string) int {
	return multiplyParts(files.Read(name))
}

func part2(name string) int {
	dontdo := regexp.MustCompile(`don't\(\)(.|\n|\r)*?do\(\)`)
	cleanString := dontdo.ReplaceAllString(files.Read(name), "")

	return multiplyParts(cleanString)
}

func multiplyParts(file string) int {
	validString := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	matches := validString.FindAllString(file, -1)
	total := 0
	for _, match := range matches {
		numStrings := strings.Split(strings.Replace(strings.Replace(match, "mul(", "", -1), ")", "", -1), ",")
		if i, err := strconv.Atoi(numStrings[0]); err == nil {
			if j, err := strconv.Atoi(numStrings[1]); err == nil {
				total += i * j
			}
		}

	}
	return total
}
