package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/coords"
	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/ints"
	"github.com/PhilAldridge/aoc-2024-go/pkg/sets"
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
	dice := parseDice(input)
	total := 0

	for total < 10000 {
		for _, die := range dice {
			score := die.roll()
			total += score
		}
	}

	return dice[0].rollCount - 1
}

func part2(name string) string {
	input := files.ReadParagraphs(name)
	dice := parseDice(input[0])
	track := ints.FromStringSlice(strings.Split(input[1][0], ""))

	for _, die := range dice {
		die.race(track)
	}

	slices.SortFunc(dice, func(a, b *diceType) int {
		return a.rollCount - b.rollCount
	})

	ids := []string{}

	for _, die := range dice {
		ids = append(ids, strconv.Itoa(die.id))
	}

	return strings.Join(ids, ",")
}

func part3(name string) int {
	input := files.ReadParagraphs(name)
	dice := parseDice(input[0])

	maze := parseMaze(input[1])

	completeMap := sets.NewSet[coords.Coord]()

	for _, die := range dice {
		singleMap := die.navigateMaze(maze)

		completeMap.AddSlice(singleMap.List())
	}

	return completeMap.Size()
}

type diceType struct {
	id                                  int
	faces                               []int
	seed, pulse, rollCount, currentFace int
}

func (dice *diceType) roll() int {
	spin := dice.rollCount * dice.pulse

	dice.currentFace = (dice.currentFace + spin) % len(dice.faces)
	dice.pulse = (dice.pulse+spin)%dice.seed + 1 + dice.rollCount + dice.seed
	dice.rollCount++

	return dice.faces[dice.currentFace]
}

func (dice *diceType) race(track []int) {
	i := 0

	for i < len(track) {
		score := dice.roll()
		if score == track[i] {
			i++
		}
	}
}

func (dice *diceType) navigateMaze(maze map[int][]coords.Coord) *sets.Set[coords.Coord] {
	score := dice.roll()

	resultMap := make(map[coords.Coord]bool)
	visitedMap := sets.NewSet[coords.Coord]()

	for _, coord := range maze[score] {
		resultMap[coord] = true
		visitedMap.Add(coord)
	}

	for len(resultMap) > 0 {
		newMap := make(map[coords.Coord]bool)
		score = dice.roll()

		for _, coord := range maze[score] {
			if toAdd(coord, resultMap) {
				newMap[coord] = true
				visitedMap.Add(coord)
			}
		}

		resultMap = newMap
	}

	return visitedMap
}

func toAdd(pos coords.Coord, mapping map[coords.Coord]bool) bool {
	if mapping[pos] {
		return true
	}

	for _, direction := range coords.DirectionsInOrder {
		if mapping[pos.Add(direction)] {
			return true
		}
	}

	return false
}

func parseDie(input string) diceType {
	split := strings.Split(input, " ")
	id := ints.FromString(strings.ReplaceAll(split[0], ":", ""))
	seed := ints.FromString(split[2][5:])
	split2 := strings.Split(split[1][7:len(split[1])-1], ",")
	faces := ints.FromStringSlice(split2)

	return diceType{
		id:        id,
		seed:      seed,
		faces:     faces,
		pulse:     seed,
		rollCount: 1,
	}
}

func parseDice(input []string) []*diceType {
	dice := []*diceType{}
	for _, row := range input {
		die := parseDie(row)
		dice = append(dice, &die)
	}

	return dice
}

func parseMaze(input []string) map[int][]coords.Coord {
	out := make(map[int][]coords.Coord)

	for i, row := range input {
		for j, char := range row {
			out[int(char-'0')] = append(out[int(char-'0')], coords.NewCoord(i, j))
		}
	}

	return out
}
