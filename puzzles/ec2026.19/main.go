package main

import (
	"fmt"
	"slices"
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
	counts := map[colour]int{}

	input := files.ReadParagraphs(name)

	for i, row := range input[0] {
		for j := range row {
			a, b := aToColour(input[0][i][j]), aToColour(input[1][i][j])
			merged := a | b
			if merged > black {
				merged = black
			}

			counts[merged]++
		}
	}

	return countWhiteCrystals(counts)
}

func part2(name string) int {
	counts := map[p2]int{}

	input := files.ReadParagraphs(name)

	for i, row := range input[0] {
		for j := range row {
			a, b := aToColour(input[0][i][j]), ints.FromString(string(input[1][i][j]))

			counts[p2{c: a, i: b}]++
		}
	}

	return tryCombinations(counts)
}

func part3(name string) int {
	counts := map[colour]int{}

	input := files.ReadLines(name)

	for i, row := range input {
		for j := range row {
			a := aToColour(input[i][j])

			counts[a]++
		}
	}

	max:=0

	for i := range counts[cyan] + 1 {
		for i2 := 0; i2+i <= counts[cyan]; i2++ {
			i3:= counts[cyan]-i-i2
			for j := range counts[magenta] + 1 {
				for j2 := j; j2+j <= counts[magenta]; j2++ {
					j3:= counts[magenta] - j-j2
					for k := range counts[yellow] + 1 {
						for k2:=0; k2+k<=counts[yellow]; k2++ {
							k3:= counts[yellow] - k - k2
							finalMapping:= map[colour]int{
								black: counts[black],
								cyan: i,
								magenta: j,
								yellow: k,
								blue: i2 + j2,
								green: i3 + k2,
								red: j3 + k3,
							}

							res:= countWhiteCrystals(finalMapping)

							if res> max {
								max = res
							}
						}
					}
				}

			}
		}

	}

	return max
}

type colour int

const (
	cyan colour = 1 << iota
	magenta
	yellow
	black
)

var (
	blue             colour = cyan | magenta
	green            colour = cyan | yellow
	red              colour = magenta | yellow
	potentialMapping        = map[colour][]colour{
		black:   []colour{black},
		cyan:    []colour{cyan, blue, green},
		magenta: []colour{magenta, blue, red},
		yellow:  []colour{yellow, green, red},
	}
)

func aToColour(a byte) colour {
	switch a {
	case 'C':
		return cyan
	case 'M':
		return magenta
	case 'Y':
		return yellow
	case 'K':
		return black
	}

	panic("unexpected colour")
}

type p2 struct {
	c colour
	i int
}

func tryCombinations(counts map[p2]int) int {
	allColours := []colour{cyan, magenta, yellow, black}
	max := 0

	for i := range allColours {
		for j := range allColours {
			if i == j {
				continue
			}
			for k := range allColours {
				if i == k || j == k {
					continue
				}
				l := 0
				for l == i || l == j || l == k {
					l++
				}

				combo := map[int]colour{
					1: allColours[i],
					2: allColours[j],
					3: allColours[k],
					4: allColours[l],
				}

				comboResult := tryCombination(counts, combo)
				if comboResult > max {
					max = comboResult
				}
			}

		}
	}

	return max
}

func tryCombination(counts map[p2]int, combo map[int]colour) int {
	countsMapped := map[colour]int{}

	for k, count := range counts {
		merged := k.c | combo[k.i]
		if merged > black {
			merged = black
		}

		countsMapped[merged] += count
	}

	return countWhiteCrystals(countsMapped)
}

func countWhiteCrystals(counts map[colour]int) int {
	vals := []int{}

	for _, v := range counts {
		vals = append(vals, v)
	}

	slices.Sort(vals)

	return vals[len(vals)-1] * vals[len(vals)-2] * vals[len(vals)-3]
}