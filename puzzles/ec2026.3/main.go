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

	slices.Reverse(iso)

	return countVote(iso, scoreStr)
}

func getVotesWithSmallestCheck(line string) int {
	split := strings.Split(line, " ")
	isoStr := split[0][5:]
	scoreStr := split[1]

	iso := ints.FromStringSlice(strings.Split(isoStr, ""))

	slices.Reverse(iso)

	votes := countVote(iso, scoreStr)

	lowerBound := votes
	for i := ints.FromString(scoreStr[0:1]) - 1; i >= 0; i-- {
		trial := fmt.Sprintf("%d%s", i, strings.Repeat("9", len(scoreStr)-1))
		trialVal := ints.FromString(trial)
		if trialVal < votes {
			break
		}

		if countVote(iso, trial) <= votes {
			lowerBound = trialVal
		}
	}

	for i := lowerBound; i < ints.FromString(scoreStr); i++ {
		if votes == countVote(iso, strconv.Itoa(i)) {
			return 0
		}
	}

	return votes
}

func countVote(iso []int, scoreStr string) int {
	score := ints.FromStringSlice(strings.Split(scoreStr, ""))
	slices.Reverse(score)

	total := 0

	for i, scoreDigit := range score {
		total += scoreDigit * ints.Pow(iso[i%len(iso)], i)
	}

	return total
}
