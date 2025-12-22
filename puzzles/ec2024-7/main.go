package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/coords"
	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/slices"
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

type finalPower struct {
	letter string
	power  int
}

func part1(name string) string {
	lines := files.ReadLines(name)
	opMap := slices.StringSliceToStringMapStringSlice(lines, ":", ",")
	answerMap := []finalPower{}

	for letter, ops := range opMap {
		total := 0
		power := 10
		for i := 0; i < 10; i++ {
			switch ops[i%len(ops)] {
			case "+":
				power++
			case "-":
				power--
			}
			total += power
		}
		answerMap = append(answerMap, finalPower{
			letter: letter,
			power:  total,
		})
	}

	sort.Slice(answerMap, func(i, j int) bool {
		a := answerMap[i]
		b := answerMap[j]

		return a.power > b.power
	})

	answer := ""

	for _, final := range answerMap {
		answer += final.letter
	}

	return answer
}

func part2(name string) string {
	paragraphs := files.ReadParagraphs(name)
	opMap := slices.StringSliceToStringMapStringSlice(paragraphs[0], ":", ",")
	answerMap := []finalPower{}

	track := parseTrack(paragraphs[1])

	for letter, ops := range opMap {
		total := 0
		power := 10

		for j := 0; j < 10; j++ {
			for i, trackOp := range track {
				switch trackOp {
				case "+":
					power++
				case "-":
					power--
				default:
					switch ops[(j*len(track)+i)%len(ops)] {
					case "+":
						power++
					case "-":
						power--
					}
				}
				if power < 0 {
					power = 0
				}
				total += power
			}
		}

		answerMap = append(answerMap, finalPower{
			letter: letter,
			power:  total,
		})
	}

	sort.Slice(answerMap, func(i, j int) bool {
		a := answerMap[i]
		b := answerMap[j]

		return a.power > b.power
	})

	answer := ""

	for _, final := range answerMap {
		answer += final.letter
	}

	return answer
}

func part3(name string) int {
	paragraphs := files.ReadParagraphs(name)
	opMap := slices.StringSliceToStringMapStringSlice(paragraphs[0], ":", ",")

	track := parseTrack(paragraphs[1])
	plans := generatePlans()

	oponentPower := 0
	count := 0

	for _, ops := range opMap {
		total := 0
		power := 10

		for j := 0; j < 2024; j++ {
			for i, trackOp := range track {
				switch trackOp {
				case "+":
					power++
				case "-":
					power--
				default:
					switch ops[(j*len(track)+i)%len(ops)] {
					case "+":
						power++
					case "-":
						power--
					}
				}
				if power < 0 {
					power = 0
				}
				total += power
			}
		}
		oponentPower = total
	}

	for _, plan := range plans {
		total := 0
		power := 10

		for j := 0; j < 2024; j++ {
			for i, trackOp := range track {
				switch trackOp {
				case "+":
					power++
				case "-":
					power--
				default:
					switch plan[(j*len(track)+i)%len(plan)] {
					case '+':
						power++
					case '-':
						power--
					}
				}
				if power < 0 {
					power = 0
				}
				total += power
			}
		}

		if total > oponentPower {
			count++
		}
	}

	return count
}

func parseTrack(input []string) []string {
	currentDirection := coords.NewCoord(0, 1)
	currentPosition := coords.NewCoord(0, 1)
	lastTrack := string(input[currentPosition.I][currentPosition.J])
	result := []string{lastTrack}

	for lastTrack != "S" {
		adjacents := coords.DirectionsInOrder
		for _, adjacent := range adjacents {
			if adjacent.I == -currentDirection.I && adjacent.J == -currentDirection.J {
				continue
			}

			nextPosition := currentPosition.Add(adjacent)
			if nextPosition.I < 0 || nextPosition.J < 0 ||
				nextPosition.I >= len(input) || nextPosition.J >= len(input[nextPosition.I]) ||
				input[nextPosition.I][nextPosition.J] == ' ' {
				continue
			}

			currentPosition = nextPosition
			currentDirection = adjacent
			break
		}

		lastTrack = string(input[currentPosition.I][currentPosition.J])
		result = append(result, lastTrack)
	}

	return result
}

func generatePlans() []string {
	plusCount := 5
	minusCount := 3
	equalsCount := 3
	results := []string{"+", "-", "="}

	for i := 1; i < 11; i++ {
		nextResults := []string{}

		for _, soFar := range results {
			if strings.Count(soFar, "+") < plusCount {
				nextResults = append(nextResults, soFar+"+")
			}
			if strings.Count(soFar, "-") < minusCount {
				nextResults = append(nextResults, soFar+"-")
			}
			if strings.Count(soFar, "=") < equalsCount {
				nextResults = append(nextResults, soFar+"=")
			}
		}

		results = nextResults
	}

	return results
}
