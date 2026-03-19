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
	total := 0

	lines := files.ReadLines(name)

	for _, line := range lines {
		duckDna := parseDNA(line)
		if duckDna.g > duckDna.r && duckDna.g > duckDna.b {
			total += duckDna.id
		}
	}

	return total
}

func part2(name string) int {
	lines := files.ReadLines(name)
	var shiniest dna

	for _, line := range lines {
		duckDna := parseDNA(line)
		if duckDna.s > shiniest.s ||
			(duckDna.s == shiniest.s && duckDna.brightness() < shiniest.brightness()) {
			shiniest = duckDna
		}
	}

	return shiniest.id
}

func part3(name string) int {
	groups := make(map[string][]int)

	max := 0
	total := 0

	lines := files.ReadLines(name)

	for _, line := range lines {
		duckDna := parseDNA(line)
		group, ok := duckDna.group()

		if ok {
			groups[group] = append(groups[group], duckDna.id)
		}
	}

	for _, v := range groups {
		if len(v) > max {
			max = len(v)
			total = ints.Sum(v)
		}
	}

	return total
}

type dna struct {
	id, r, g, b, s int
}

func (d dna) brightness() int {
	return d.r + d.g + d.b
}

func (d dna) dominantColour() string {
	if d.b > d.g {
		if d.b > d.r {
			return "blue"
		}

		if d.b == d.r {
			return ""
		}
	} else {
		if d.g > d.r {
			if d.g == d.b {
				return ""
			}

			return "green"
		}

		if d.g == d.r {
			return ""
		}
	}

	return "red"
}

func (d dna) shinyMatte() string {
	if d.s >= 33 {
		return "shiny"
	}

	if d.s <= 30 {
		return "matte"
	}

	return ""
}

func (d dna) group() (string, bool) {
	shinyString := d.shinyMatte()
	if shinyString == "" {
		return "", false
	}

	colourString := d.dominantColour()
	if colourString == "" {
		return "", false
	}

	return shinyString + "-" + colourString, true
}

func parseDNA(input string) dna {
	colonIndex := strings.Index(input, ":")
	colourStrings := strings.Split(input[colonIndex+1:], " ")
	if len(colourStrings) == 4 {
		return dna{
			id: ints.FromString(input[:colonIndex]),
			r:  parseColour(colourStrings[0]),
			g:  parseColour(colourStrings[1]),
			b:  parseColour(colourStrings[2]),
			s:  parseColour(colourStrings[3]),
		}
	}

	return dna{
		id: ints.FromString(input[:colonIndex]),
		r:  parseColour(colourStrings[0]),
		g:  parseColour(colourStrings[1]),
		b:  parseColour(colourStrings[2]),
	}
}

func parseColour(input string) int {
	result := 0
	for _, char := range input {
		result *= 2
		if char >= 'A' && char <= 'Z' {
			result++
		}
	}

	return result
}
