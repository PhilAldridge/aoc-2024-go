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
	turns := ints.FromStringSlice(strings.Split(input[0][0], ","))

	machine := parseMachine(input[1], len(turns))

	positions := make([]int, len(turns))

	for range 100 {
		for i := range turns {
			positions[i] = (positions[i] + turns[i]) % len(machine[i])
		}
	}

	result := []string{}

	for i := range turns {
		result = append(result, machine[i][positions[i]])
	}

	return strings.Join(result, " ") + " "
}

func part2(name string) int {
	totalPulls := 202420242024

	input := files.ReadParagraphs(name)
	turns := ints.FromStringSlice(strings.Split(input[0][0], ","))

	machine := parseMachine(input[1], len(turns))

	positions := [10]int{}

	lcm := 1

	for i := range len(turns) {
		lcm = ints.LCM(lcm, len(machine[i])/ints.GCD(len(machine[i]), turns[i]))
	}

	total := 0

	for range lcm {
		newPositions, score := pull(positions, turns, machine)
		positions = newPositions

		total += score
	}

	total *= totalPulls / lcm

	for range totalPulls % lcm {
		newPositions, score := pull(positions, turns, machine)
		positions = newPositions

		total += score
	}

	return total
}

type state struct {
	positions [10]int
	pullsLeft int
}

func part3(name string) string {
	input := files.ReadParagraphs(name)
	turns := ints.FromStringSlice(strings.Split(input[0][0], ","))

	machine := parseMachine(input[1], len(turns))

	for i := range turns {
		turns[i] = turns[i] % len(machine[i])
	}

	start := state{
		pullsLeft: 256,
	}

	resultMap := make(map[state][2]int)

	totals := runPulls(machine, start, turns, resultMap)

	return fmt.Sprintf("%d %d", totals[1], totals[0])
}

func parseMachine(input []string, wheelCount int) [][]string {
	result := make([][]string, wheelCount)

	for _, row := range input {
		for i := range wheelCount {
			if i*4 >= len(row) {
				break
			}
			if row[i*4] == ' ' {
				continue
			}

			result[i] = append(result[i], row[i*4:(i+1)*4-1])
		}
	}

	return result
}

func pull(positions [10]int, turns []int, machine [][]string) ([10]int, int) {
	total := 0
	symbolCount := make(map[rune]int)
	var positionsCopy [10]int
	positionsCopy = positions

	for i := range turns {
		positionsCopy[i] = (positionsCopy[i] + turns[i]) % len(machine[i])
		for j, char := range machine[i][positionsCopy[i]] {
			if j == 1 {
				continue
			}
			symbolCount[char]++
		}
	}

	for _, count := range symbolCount {
		if count < 3 {
			continue
		}

		total += count - 2
	}
	return positionsCopy, total
}

func runPulls(machine [][]string, lastState state, turns []int, resultMap map[state][2]int) [2]int {
	prevResult, ok := resultMap[lastState]
	if ok {
		return prevResult
	}

	neutralPos, neutralScore := pull(lastState.positions, turns, machine)

	var upPositions, downPositions [10]int

	for i := range turns {
		upPositions[i] = (lastState.positions[i] + len(machine[i]) - 1) % len(machine[i])
	}

	upPos, upScore := pull(upPositions, turns, machine)

	for i := range turns {
		downPositions[i] = (lastState.positions[i] + 1) % len(machine[i])
	}

	downPos, downScore := pull(downPositions, turns, machine)

	scores := []int{}

	if lastState.pullsLeft == 1 {
		scores := []int{neutralScore, upScore, downScore}
		result := [2]int{ints.Min(scores), ints.Max(scores)}

		resultMap[lastState] = result
		return result
	}

	neutralScores := runPulls(machine, state{
		positions: neutralPos,
		pullsLeft: lastState.pullsLeft - 1,
	}, turns, resultMap)

	upScores := runPulls(machine, state{
		positions: upPos,
		pullsLeft: lastState.pullsLeft - 1,
	}, turns, resultMap)

	downScores := runPulls(machine, state{
		positions: downPos,
		pullsLeft: lastState.pullsLeft - 1,
	}, turns, resultMap)

	scores = append(scores,
		neutralScores[0]+neutralScore,
		neutralScores[1]+neutralScore,
		upScores[0]+upScore,
		upScores[1]+upScore,
		downScores[0]+downScore,
		downScores[1]+downScore,
	)

	result := [2]int{ints.Min(scores), ints.Max(scores)}

	resultMap[lastState] = result
	return result
}
