package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/coords"
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
	in := parseInput(name)
	total:= 0
	squares:= findSquares(in)

	for _, square := range squares {
		hRepeat := in.width / (len(in.vOffset)*2)
		if in.width % (len(in.vOffset)*2) > square.I {
			hRepeat++
		}
		vRepeat := in.height / (len(in.hOffset)*2)
		if in.height % (len(in.hOffset)*2) > square.J {
			vRepeat++
		}

		total += hRepeat*vRepeat
	}

	return total
}

func part2(name string) int {
	in := parseInput(name)
	colouredGrid := colourGrid(in)
	counts := make(map[bool]int)
	squares := findSquares(in)

	for _, square := range squares {
		hRepeat := in.width / (len(in.vOffset)*2)
		if in.width % (len(in.vOffset)*2) > square.I {
			hRepeat++
		}
		vRepeat := in.height / (len(in.hOffset)*2)
		if in.height % (len(in.hOffset)*2) > square.J {
			vRepeat++
		}

		counts[colouredGrid[coords.NewCoord(
			square.I%(len(in.vOffset)*2),
			square.J%(len(in.hOffset)*2),			
		)]] += hRepeat*vRepeat
	}

	return ints.Max([]int{counts[false], counts[true]})
}

func part3(name string) int {
	return part2(name)
}

type inputType struct {
	width, height    int
	hOffset, vOffset []bool
}

func parseInput(name string) inputType {
	lines := files.ReadLines(name)

	width := ints.FromStringIgnoreText(lines[0])
	height := ints.FromStringIgnoreText(lines[1])

	hOffsetString := lines[2][strings.Index(lines[2], "=")+1:]
	vOffsetString := lines[3][strings.Index(lines[3], "=")+1:]

	var hOffset, vOffset []bool

	for _, hByte := range hOffsetString {
		if hByte == '1' {
			hOffset = append(hOffset, true)
		} else {
			hOffset = append(hOffset, false)
		}
	}

	for _, vByte := range vOffsetString {
		if vByte == '1' {
			vOffset = append(vOffset, true)
		} else {
			vOffset = append(vOffset, false)
		}
	}

	return inputType{
		width:   width,
		height:  height,
		hOffset: hOffset,
		vOffset: vOffset,
	}
}

// Find squares in two layer repeat of pattern (twice needed for two-color coloring)
func findSquares(in inputType) []coords.Coord {
	var potentialSquares, squares []coords.Coord

	// Find potential squares by hOffset (where there is a stitch along top and bottom of square)
	for i := range len(in.hOffset)*2 {
		hOffset := in.hOffset[i%len(in.hOffset)]

		if in.hOffset[(i+1)%len(in.hOffset)] != hOffset {
			continue
		}

		xPos := 0

		if hOffset {
			xPos++
		}

		for xPos < len(in.vOffset)*2 {
			potentialSquares = append(potentialSquares, coords.NewCoord(xPos, i))

			xPos += 2
		}
	}

	// Find potential squares by vOffset (where there is a stitch along left and right of square)
	for i := range len(in.vOffset)*2 {
		vOffset := in.vOffset[i%len(in.vOffset)]

		if in.vOffset[(i+1)%len(in.vOffset)] != vOffset {
			continue
		}

		yPos := 0

		if vOffset {
			yPos++
		}

		for yPos < len(in.hOffset)*2 {
			square := coords.NewCoord(i, yPos)

			// match potential squares from hOffset
			if slices.ContainsFunc(potentialSquares, func(a coords.Coord) bool {
				return a.Equals(square)
			}) {
				// square from both offsets -> add to output
				squares = append(squares, square)
			}

			yPos += 2
		}
	}

	return squares
}

// get colour for each square in two layer repeat of pattern (twice needed for two-color coloring)
// method - count lines crossed moving from first square to each square
func colourGrid(in inputType) map[coords.Coord]bool {
	colStart := false
	current := false
	out := make(map[coords.Coord]bool)

	for i := range len(in.vOffset)*2 {
		for j := range len(in.hOffset)*2 {
			if in.hOffset[j%len(in.hOffset)] == (i%2 == 1) {
				current = !current
			}

			out[coords.NewCoord(i, j)] = current
		}

		if in.vOffset[(i+1)%len(in.vOffset)] {
			colStart = !colStart
		}
		current = colStart
	}

	return out
}
