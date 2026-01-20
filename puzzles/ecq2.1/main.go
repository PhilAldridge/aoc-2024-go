package main

import (
	"fmt"
	"math"
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
	input := files.ReadParagraphs(name)

	total := 0

	width := len(input[0][0])

	for i, coinToken := range input[1] {
		startSlot := (i + 1) % width
		total += calculateScore(coinToken, input[0], startSlot)
	}

	return total
}

func part2(name string) int {
	input := files.ReadParagraphs(name)

	total := 0

	for _, coinToken := range input[1] {
		total += calculateMaxScore(coinToken, input[0])
	}

	return total
}

type state struct {
	coinTokenIndex int
	startSlot      int
}

func part3(name string) string {
	input := files.ReadParagraphs(name)

	width := len(input[0][0])

	resultsMap := make(map[state]int)

	for i, coinToken := range input[1] {
		for j := range width/2 + 1 {
			startSlot := j + 1
			score := calculateScore(coinToken, input[0], startSlot)
			resultsMap[state{
				coinTokenIndex: i,
				startSlot:      startSlot,
			}] = score
		}
	}

	comboMap := make(map[[6]int][2]int)

	min, max := calculateMinMaxCombination([6]int{}, resultsMap, comboMap)

	return fmt.Sprintf("%d %d", min, max)
}

func calculateMaxScore(coinToken string, machine []string) int {
	max := 0

	for i := range len(machine[0])/2 + 1 {
		score := calculateScore(coinToken, machine, i+1)

		if score > max {
			max = score
		}
	}

	return max
}

func calculateScore(coinToken string, machine []string, startSlot int) int {
	width := len(machine[0])

	xPos := ((startSlot - 1) * 2) % width
	bounceCount := 0
	for _, row := range machine {
		if row[xPos] == '*' {
			if xPos == 0 {
				xPos++
			} else if xPos == len(row)-1 {
				xPos--
			} else {
				if coinToken[bounceCount%len(coinToken)] == 'L' {
					xPos--
				} else {
					xPos++
				}
			}
			bounceCount++
		}
	}

	endSlot := (xPos/2 + 1) % width

	return ints.Max([]int{endSlot*2 - startSlot, 0})
}

func calculateMinMaxCombination(slots [6]int, resultsMap map[state]int, comboMap map[[6]int][2]int) (int, int) {
	if minMax, ok := comboMap[slots]; ok {
		return minMax[0], minMax[1]
	}

	min := math.MaxInt
	max := 0

	slotsUsed := []int{}

	for _, slot := range slots {
		slotsUsed = append(slotsUsed, slot)
	}

	for i, slot := range slots {
		if slot != 0 {
			continue
		}

		if i == 5 {
			for state, result := range resultsMap {
				if state.coinTokenIndex == i && !slices.Contains(slotsUsed, state.startSlot) {
					if result > max {
						max = result
					}
					if result < min {
						min = result
					}
				}
			}

			comboMap[slots] = [2]int{min, max}
			return min, max
		}

		for state, result := range resultsMap {
			if state.coinTokenIndex == i && !slices.Contains(slotsUsed, state.startSlot) {
				newSlots := [6]int{}
				for j := range slots {
					if i == j {
						newSlots[j] = state.startSlot
					} else {
						newSlots[j] = slots[j]
					}
				}

				minScore, maxScore := calculateMinMaxCombination(newSlots, resultsMap, comboMap)
				if maxScore+result > max {
					max = maxScore + result
				}
				if minScore+result < min {
					min = minScore + result
				}
			}
		}

		comboMap[slots] = [2]int{min, max}
		return min, max
	}

	panic("something went wrong")
}
