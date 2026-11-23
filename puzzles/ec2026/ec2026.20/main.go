package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/bools"
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

	for _, line := range files.ReadLines(name) {
		total += parseLine(line)
	}

	return total
}

func part2(name string) string {
	total := 0

	for _, line := range files.ReadLines(name) {
		total += parseLine(line)
	}

	return ItoaB11(total)
}

func part3(name string) string {
	total := 0

	for _, line := range files.ReadLines(name) {
		total += parseLinePart3(line)
	}

	return ItoaB11(total)
}

// parseSide is for parts 1 and 2
// it adds the scores in base 11 for one side of the equation
func parseSide(str string) (out int) {
	split := strings.Split(str, " ")
	addingMultiplier := 1

	for _, s := range split {
		switch s {
		case "+":
			addingMultiplier = 1
		case "-":
			addingMultiplier = -1
		default:
			out += addingMultiplier * AtoiB11(s)
		}
	}

	return out
}

// parseLine is for parts 1 and 2. If the left == right, add all segments used
func parseLine(line string) int {
	split := strings.Split(line, " = ")
	left := parseSide(split[0])
	right := parseSide(split[1])

	if left == right {
		return parseSegments(line)
	}

	return 0
}

// parseSegments counts the segments used for every character in the line
func parseSegments(line string) (out int) {
	for _, s := range line {
		switch s {
		case '1', '+', '=':
			out += 2
		case '2', '5', '3':
			out += 5
		case '4':
			out += 4
		case '6', '9', '0', 'A':
			out += 6
		case '7':
			out += 3
		case '8':
			out += 7
		case '-':
			out += 1
		case ' ':
			continue
		default:
			panic("unexpected char")
		}
	}

	return out
}

// AtoiB11 converts a base 11 number string into a decimal int
func AtoiB11(s string) (out int) {
	for _, c := range s {
		out *= 11
		if c == 'A' {
			out += 10
		} else {
			out += ints.FromString(string(c))
		}
	}
	return out
}

// ItoaB11 converts a decimal int into a base 11 number string
func ItoaB11(i int) (out string) {
	for i > 0 {
		r := i % 11
		if r == 10 {
			out = "A" + out
		} else {
			out = strconv.Itoa(r) + out
		}

		i /= 11
	}

	return out
}

type TValue struct {
	value                        int
	addedSegment, removedSegment bool
}

// paseSidePart3 gets all potential values for a side that can be achieved
// by either adding, removing or moving a single segment
func parseSidePart3(terms []string) (out []TValue) {
	// first get all potential values for each term
	allTermsPotentialValues := [][]TValue{}

	lastSign := ""

	for _, term := range terms {
		if term == "+" || term == "-" {
			lastSign = term
			continue
		}

		switch lastSign {
		case "":
			allTermsPotentialValues = append(allTermsPotentialValues, parseTerm(term, false, false))
		case "+":
			termPotentialValues := parseTerm(term, false, false)
			potentialNegativeValues := parseTerm(term, false, true)
			for _, v := range potentialNegativeValues {
				v.value = -v.value
				termPotentialValues = append(termPotentialValues, v)
			}

			allTermsPotentialValues = append(allTermsPotentialValues, termPotentialValues)
		case "-":
			termPotentialValues := parseTerm(term, true, false)
			potentialNegativeValues := parseTerm(term, false, false)
			for _, v := range potentialNegativeValues {
				v.value = -v.value
				termPotentialValues = append(termPotentialValues, v)
			}

			allTermsPotentialValues = append(allTermsPotentialValues, termPotentialValues)
		}
	}

	// get all possible combinations of terms where only one segment is moved overall
	return collectValidCombinations(allTermsPotentialValues, false, false)
}

// collectValidCombinations is a recursive function. It takes and array of possible values for each term
// and collects them into valid answers where either a single segment is added, removed or moved (added and removed)
// or no segment is changed
// all other combinations are discarded
func collectValidCombinations(termTValues [][]TValue, added, removed bool) (out []TValue) {
	firstTermValues := termTValues[0]

	for _, v := range firstTermValues {
		if (added && v.addedSegment) || (removed && v.removedSegment) {
			continue
		}

		// Base case - one term - return potential values from single term
		if len(termTValues) == 1 {
			v.addedSegment = added || v.addedSegment
			v.removedSegment = removed || v.removedSegment
			out = append(out, v)
			continue
		}

		// Recursive case, for each potential value of first term, add all
		// combinations of potential values from remaining terms
		next := collectValidCombinations(termTValues[1:], v.addedSegment || added, v.removedSegment || removed)

		for _, n := range next {
			out = append(out, TValue{
				value:          v.value + n.value,
				addedSegment:   added || v.addedSegment || n.addedSegment,
				removedSegment: removed || v.removedSegment || n.removedSegment,
			})
		}
	}

	return out
}

// parseLinePart3 gets valid values for the left and right side
// then sees if there is a matching left/right value where either no segment is moved
// or a single segment is moved
func parseLinePart3(line string) (out int) {
	split := strings.Split(line, " = ")
	left := parseSidePart3(strings.Split(split[0], " "))
	right := parseSidePart3(strings.Split(split[1], " "))

	for _, vL := range left {
		if slices.ContainsFunc(right, func(vR TValue) bool {
			added := bools.CountTrues(vL.addedSegment, vR.addedSegment)
			removed := bools.CountTrues(vL.removedSegment, vR.removedSegment)

			return added < 2 && removed < 2 && added == removed && vL.value == vR.value
		}) {
			return parseSegments(line)
		}
	}

	return out
}

// parseTerm is a recursive function, returning all potential values for a term
// that can be achieved by adding/removing/moving at most one segment
func parseTerm(str string, added, removed bool) (out []TValue) {
	firstChar := str[0]

	// Get potential values of first character
	firstVs := parseChar(firstChar)

	for _, v := range firstVs {
		if (added && v.addedSegment) || (removed && v.removedSegment) {
			continue
		}

		// Base case - single character, return potential values
		if len(str) == 1 {
			out = append(out, TValue{value: v.value, addedSegment: added || v.addedSegment, removedSegment: removed || v.removedSegment})
			continue
		}

		// Recursion potential value = potentialV of first char + any potential value of remaining chars
		next := parseTerm(str[1:], added || v.addedSegment, removed || v.removedSegment)

		for _, n := range next {
			out = append(out,
				TValue{value: ints.Pow(11, len(str)-1)*v.value + n.value,
					addedSegment:   added || v.addedSegment || n.addedSegment,
					removedSegment: removed || v.removedSegment || n.removedSegment,
				})
		}
	}

	return out
}

// parseChar takes a single character and returns any values it can represent after moving
// at most one segment
func parseChar(c byte) (out []TValue) {
	switch c {
	case '1':
		out = append(out, TValue{value: 1}, TValue{value: 7, addedSegment: true})
	case '2':
		out = append(out, TValue{value: 2}, TValue{value: 3, addedSegment: true, removedSegment: true})
	case '5':
		out = append(out, TValue{value: 5},
			TValue{value: 6, addedSegment: true},
			TValue{value: 9, addedSegment: true},
			TValue{value: 3, addedSegment: true, removedSegment: true},
		)
	case '3':
		out = append(out, TValue{value: 3},
			TValue{value: 9, addedSegment: true},
			TValue{value: 5, addedSegment: true, removedSegment: true},
			TValue{value: 2, addedSegment: true, removedSegment: true},
		)
	case '4':
		out = append(out, TValue{value: 4})
	case '6':
		out = append(out, TValue{value: 6},
			TValue{value: 9, addedSegment: true, removedSegment: true},
			TValue{value: 5, removedSegment: true},
			TValue{value: 8, addedSegment: true},
			TValue{value: 0, addedSegment: true, removedSegment: true},
			TValue{value: 10, addedSegment: true, removedSegment: true},
		)
	case '9':
		out = append(out, TValue{value: 9},
			TValue{value: 6, addedSegment: true, removedSegment: true},
			TValue{value: 5, removedSegment: true},
			TValue{value: 3, removedSegment: true},
			TValue{value: 8, addedSegment: true},
			TValue{value: 0, addedSegment: true, removedSegment: true},
			TValue{value: 10, addedSegment: true, removedSegment: true},
		)
	case '0':
		out = append(out, TValue{value: 0},
			TValue{value: 6, addedSegment: true, removedSegment: true},
			TValue{value: 8, addedSegment: true},
			TValue{value: 9, addedSegment: true, removedSegment: true},
			TValue{value: 10, addedSegment: true, removedSegment: true},
		)
	case 'A':
		out = append(out, TValue{value: 10},
			TValue{value: 9, addedSegment: true, removedSegment: true},
			TValue{value: 8, addedSegment: true},
			TValue{value: 6, addedSegment: true, removedSegment: true},
			TValue{value: 0, addedSegment: true, removedSegment: true},
		)
	case '7':
		out = append(out, TValue{value: 7}, TValue{value: 1, removedSegment: true})
	case '8':
		out = append(out, TValue{value: 8},
			TValue{value: 0, removedSegment: true},
			TValue{value: 9, removedSegment: true},
			TValue{value: 6, removedSegment: true},
			TValue{value: 10, removedSegment: true},
		)
	default:
		panic("unexpected char")
	}

	return out
}
