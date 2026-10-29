package main

import (
	"fmt"
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
	lines := files.ReadLines(name)

	total := 0

	for _, line := range lines {
		total += getVotes(line)
	}

	return total
}

func part2(name string) int {
	return part1(name)
}

func part3(name string) int {
	lines := files.ReadLines(name)

	total := 0

	for _, line := range lines {
		total += getVotesWithSmallestCheck(line)
	}

	return total
}

func getVotes(line string) int {
	split := strings.Split(line, " ")
	isoStr := split[0][4:]
	scoreStr := split[1]

	iso := ints.FromStringSlice(strings.Split(isoStr, ""))

	isoMultipliers := make([]int, len(scoreStr))

	isoIndex := len(iso) - 1

	for i := range isoMultipliers {
		isoMultipliers[len(isoMultipliers)-1-i] = ints.Pow(iso[isoIndex], i)

		isoIndex--
		if isoIndex < 0 {
			isoIndex += len(iso)
		}
	}

	return countVote(isoMultipliers, scoreStr)
}

func getVotesWithSmallestCheck(line string) int {
	split := strings.Split(line, " ")
	isoStr := split[0][5:]
	scoreStr := split[1]

	iso := ints.FromStringSlice(strings.Split(isoStr, ""))
	score := ints.FromStringSlice(strings.Split(scoreStr, ""))

	isoMultipliers := make([]int, len(scoreStr))
	isoIndex := len(iso) - 1

	for i := range isoMultipliers {
		isoMultipliers[len(isoMultipliers)-1-i] = ints.Pow(iso[isoIndex], i)

		isoIndex--
		if isoIndex < 0 {
			isoIndex += len(iso)
		}
	}

	votes := countVote(isoMultipliers, scoreStr)

	lowerBound := votes
	for i := score[0] - 1; i >= 0; i-- {
		trial := fmt.Sprintf("%d%s", i, strings.Repeat("9", len(scoreStr)-1))
		trialVal := ints.FromString(trial)
		if trialVal < votes {
			break
		}

		if countVote(isoMultipliers, trial) <= votes {
			lowerBound = trialVal
		}
	}

	for i := lowerBound; i < ints.FromString(scoreStr); i++ {
		if votes == countVote(isoMultipliers, strconv.Itoa(i)) {
			return 0
		}
	}

	return votes
}

func countVote(isoMultipliers []int, scoreStr string) int {
	score := ints.FromStringSlice(strings.Split(scoreStr, ""))

	total := 0

	for i := range score {
		total += score[len(score)-1-i] * isoMultipliers[len(isoMultipliers)-1-i]
	}

	return total
}
