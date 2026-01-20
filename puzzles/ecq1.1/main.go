package main

import (
	"fmt"
	"slices"
	"strconv"
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
	input := files.ReadLines(name)
	rows := parseInput(input)
	return getMax(rows,eni1)
}

func part2(name string) int {
	input := files.ReadLines(name)
	rows := parseInput(input)
	return getMax(rows,eni2)
}

func part3(name string) int {
	input := files.ReadLines(name)
	rows := parseInput(input)
	return getMax(rows,eni3)
}

func getMax(rows []inputType, eni func(a, e, m int) int) int {
	max := 0

	for _, row := range rows {
		score := eni(row.A, row.X, row.M) + eni(row.B, row.Y, row.M) + eni(row.C, row.Z, row.M)

		if score > max {
			max = score
		}
	}

	return max
}

func eni1(n, exp, mod int) int {
	var remainders []string
	score := 1

	for range exp {
		score = (score * n) % mod
		remainders = append(remainders, strconv.Itoa(score))
	}

	slices.Reverse(remainders)

	return ints.FromString(strings.Join(remainders, ""))
}

func eni2(n, exp, mod int) int {
	var remainders []int
	score := 1

	for range exp {
		score = (score * n) % mod
		if loopStartIndex := slices.Index(remainders, score); loopStartIndex != -1 {
			remainders = remainders[loopStartIndex:]
			finalRemainder := (exp - loopStartIndex - 1) % len(remainders)
			resultString := ""
			for i := range ints.Min([]int{5, exp}) {
				resultString += strconv.Itoa(remainders[ints.Mod(finalRemainder-i, len(remainders))])
			}
			return ints.FromString(resultString)
		}
		remainders = append(remainders, score)
	}

	slices.Reverse(remainders)

	resultString := ""

	for i := range ints.Min([]int{5, len(remainders)}) {
		resultString += strconv.Itoa(remainders[i])
	}

	return ints.FromString(resultString)
}

func eni3(n, exp, mod int) int {
	var remainders []int
	score := 1

	for range exp {
		score = (score * n) % mod
		if loopStartIndex := slices.Index(remainders, score); loopStartIndex != -1 {
			result := ints.Sum(remainders[:loopStartIndex])
			leftToAdd := exp - loopStartIndex
			remainders = remainders[loopStartIndex:]
			fullLoops := leftToAdd / len(remainders)
			result += fullLoops * ints.Sum(remainders)
			leftToAdd %= len(remainders)

			for i := range leftToAdd {
				result += remainders[i]
			}

			return result
		}
		remainders = append(remainders, score)
	}

	return ints.Sum(remainders)
}

type inputType struct {
	A, B, C, X, Y, Z, M int
}

func parseInput(input []string) []inputType {
	var out []inputType

	for _, row := range input {
		split := strings.Split(row, " ")

		out = append(out, inputType{
			A: parseSplit(split[0]),
			B: parseSplit(split[1]),
			C: parseSplit(split[2]),
			X: parseSplit(split[3]),
			Y: parseSplit(split[4]),
			Z: parseSplit(split[5]),
			M: parseSplit(split[6]),
		})
	}

	return out
}

func parseSplit(str string) int {
	return ints.FromString(str[2:])
}
