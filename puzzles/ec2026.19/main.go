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

	allColours:= []colour{black, cyan,magenta,yellow,blue,green,red}
	primary:= []colour{black, cyan,magenta,yellow}

	// f returns how many cells can become at least one colour in set.
	f := func(set ...colour) int {
		total := 0
		for _, s := range primary {
			for _, t := range set {
				if slices.Contains(potentialMapping[s], t) {
					total += counts[s]
					break
				}
			}
		}
		return total
	}

	for a := 0; a < len(allColours); a++ {
		for b := a + 1; b < len(allColours); b++ {
			for c := b + 1; c < len(allColours); c++ {
				colour1, colour2, colour3:= allColours[a], allColours[b], allColours[c]
				maxCells:= f(colour1,colour2,colour3)
				max1, max2, max3:= f(colour1), f(colour2), f(colour3)
				max12, max13, max23:= f(colour1,colour2), f(colour1,colour3), f(colour2,colour3)

				for colour1Count:= ints.Max([]int{0, maxCells - max23}); colour1Count <= max1; colour1Count++ {
					colour2Min := ints.Max([]int{0, maxCells-colour1Count-max3, maxCells - max13})
					colour2Max:= ints.Min([]int{max2, max12-colour1Count})

					for colour2Count := colour2Min; colour2Count<=colour2Max; colour2Count ++ {
						colour3Count:= maxCells - colour1Count - colour2Count
						res:= colour1Count * colour2Count * colour3Count
						if res > max {
							max = res
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