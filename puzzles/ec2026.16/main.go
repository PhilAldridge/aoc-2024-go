package main

import (
	"fmt"
	"slices"
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
	tablets := []byte(files.Read(name))
	flips, _ := sort(tablets, 0)
	return flips
}

func part2(name string) int {
	tablets := []byte(strings.Repeat(files.Read(name), 10))
	_, elements := sort(tablets, 0)
	return elements
}

func part3(name string) int {
	tablets := []byte(strings.Repeat(files.Read(name), 100))
	_, elements := sort(tablets, 0)
	return elements
}

// returns slice, lcursor, moves
func sort(slice []byte, lCursor int) (flips int, elements int) {
	current := slice[lCursor]
	rCursor := len(slice) - 1

	for lCursor < len(slice)-1 {
		for slice[lCursor] == current {
			lCursor++
			if lCursor >= len(slice) {
				return flips, elements
			}
		}

		for slice[rCursor] != current {
			if rCursor == lCursor {
				for slice[lCursor] == current {
					lCursor++
					if lCursor >= len(slice) {
						return flips, elements
					}
				}

				subFlips, subElements := sort(slice, lCursor)

				return flips + subFlips, elements + subElements
			}
			rCursor--
		}

		slices.Reverse(slice[lCursor : rCursor+1])

		fmt.Println(string(slice))

		flips++
		elements += rCursor - lCursor + 1
	}

	return flips, elements
}
