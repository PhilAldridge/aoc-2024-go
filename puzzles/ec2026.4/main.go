package main

import (
	"fmt"
	"slices"
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
	return part2(name)
}

func part2(name string) int {
	scores := make(map[int]result)
	lines := files.ReadLines(name)

	for _, line := range lines {
		result := parseRow(line)

		found := scores[result.id]
		found.time += result.time
		found.part++
		scores[result.id] = found
	}

	finalResults := []result{}
	for _, v := range scores {
		finalResults = append(finalResults, v)
	}

	slices.SortFunc(finalResults, func(a, b result) int {
		if b.part == a.part {
			return a.time - b.time
		}

		return b.part - a.part
	})

	total := 0

	for i, result := range finalResults {
		total += (i + 1) * result.time
	}

	return total
}

func part3(name string) int {
	results := []result{}
	lines := files.ReadLines(name)

	for _, line := range lines {
		result := parseRow(line)
		results = append(results, result)
	}

	resultsMap:= mapResults(results)

	scores:= make(map[int]int)

	for k,v:= range resultsMap {
		slices.SortFunc(v, func(a,b idTime) int {
			return a.time - b.time
		})

		currentTime:=0
		maxScore := 50*k.part
		lastIndex:=0

		for i, idRes:= range v {
			if currentTime != idRes.time {
				lastIndex = i
				currentTime = idRes.time

				if lastIndex >= maxScore {
					break
				}
			}

			scores[idRes.id] += maxScore - lastIndex
		}
	}

	finalScores := []int{}
	for _, v := range scores {
		finalScores = append(finalScores, v)
	}

	slices.SortFunc(finalScores, func(a,b int) int { return b-a})

	total := 0

	lastIndex:=0
	currentScore:=0
	for i, score := range finalScores {
		if score != currentScore {
			currentScore = score
			lastIndex = i
		}

		total += (lastIndex + 1) * score
	}

	return total
}

type result struct {
	id, quest, part, time int
}

func parseRow(line string) result {
	split := strings.Split(line, " ")

	return result{
		id:    parseInt(split[0]),
		quest: parseInt(split[1]),
		part:  parseInt(split[2]),
		time:  parseInt(split[3]),
	}
}

func parseInt(split string) int {
	split2 := strings.Split(split, "=")
	if len(split2) == 1 {
		return ints.FromString(split)
	}

	return ints.FromString(split2[1])
}

func mapResults(results []result) map[questPart][]idTime {
	out := make(map[questPart][]idTime)

	for _, res := range results {
		if res.part == 3 {
			for _, res2 := range results {
				if res2.part == 2 && res2.quest == res.quest && res2.id == res.id {
					res.time += res2.time
				}
			}
		}

		if res.part > 1 {
			for _, res2 := range results {
				if res2.part == 1 && res2.quest == res.quest && res2.id == res.id {
					res.time += res2.time
				}
			}
		}

		out[questPart{
			quest: res.quest,
			part:  res.part,
		}] = append(out[questPart{
			quest: res.quest,
			part:  res.part,
		}], idTime{
			id:   res.id,
			time: res.time,
		})
	}

	return out
}

type questPart struct {
	quest, part int
}

type idTime struct {
	id, time int
}
