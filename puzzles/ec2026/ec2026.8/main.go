package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/coords"
	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/hexgrids"
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
	ents := parseInput(name)

	rows := ents.DownRightRows()

	total, _ := genericRound(ents, rows, false)

	return total
}

func part2(name string) int {
	ents := parseInput(name)

	directionRows := ents.GetAllDirectionRows()

	total := 0
	for _, rows := range directionRows {
		t, _ := genericRound(ents, rows, false)
		total += t
	}

	return total
}

func part3(name string) int {
	total := 0
	grownWarriors := 0
	ents := parseInput(name)
	lastTotal := 0
	lastGrownWarriors := 0
	round := 0
	directionRows := ents.GetAllDirectionRows()
	jsonStrings := []string{}

	for {
		lastTotal, lastGrownWarriors = genericRound(ents, directionRows[round%len(directionRows)], true)
		round++
		grownWarriors += lastGrownWarriors
		total += lastTotal

		jsonBytes, err := json.Marshal(roundJson{
			Ents: ents,
			Sun:  round % len(directionRows),
		})
		if err != nil {
			panic(err)
		}

		jsonStrings = append(jsonStrings, string(jsonBytes))

		if grownWarriors >= 100000 {
			file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println("File does not exists or cannot be created")
				os.Exit(1)
			}
			defer file.Close()

			w := bufio.NewWriter(file)
			fmt.Fprintf(w, "[%s]", strings.Join(jsonStrings, ","))

			w.Flush()
			return total
		}
	}
}

type roundJson struct {
	Sun  int     `json:"sun"`
	Ents [][]int `json:"ents"`
}

func genericRound(ents hexgrids.HexGrid, rows [][]coords.Coord, growth bool) (int, int) {
	toGrow := []coords.Coord{}
	warriorsGrown := 0
	energy := 0

	for _, row := range rows {
		toGrow = append(toGrow, findGrowths(ents, row)...)
	}

	for _, grown := range toGrow {
		if x := ents[grown.I][grown.J]; x >= 9 {
			if growth {
				ents[grown.I][grown.J] = -2
			}

			warriorsGrown++
			energy++
		} else {
			if growth {
				ents[grown.I][grown.J] = x + 1
			}

			if x > 0 {
				energy++
			}
		}
	}

	return energy, warriorsGrown
}

func parseInput(name string) hexgrids.HexGrid {
	out := hexgrids.HexGrid{}

	for _, row := range files.ReadLines(name) {
		row = strings.ReplaceAll(row, ".", "")
		out = append(out, ints.FromStringSlice(strings.Split(row, "")))
	}

	return out
}

func findGrowths(ents hexgrids.HexGrid, row []coords.Coord) []coords.Coord {
	shadowHeight := 0
	toGrow := []coords.Coord{}

	for _, pos := range row {
		shadowHeight--

		if height := ents[pos.I][pos.J]; height < 1 {
			toGrow = append(toGrow, pos)
		} else if height > shadowHeight {
			toGrow = append(toGrow, pos)
			shadowHeight = height
		}
	}

	return toGrow
}
