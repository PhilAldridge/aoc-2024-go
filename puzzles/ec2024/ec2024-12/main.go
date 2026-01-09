package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/coords"
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

	problemSpace := files.ReadLinesAsRunes(name)

	rankMap := createRankMap(problemSpace)

	for i, row := range problemSpace {
		for j, char := range row {
			if char == 'T' {
				score, ok := rankMap[coords.NewCoord(i, j)]
				if !ok {
					panic("no score found¬")
				}

				total += score
			}
		}
	}

	return total
}

func part2(name string) int {
	total := 0

	problemSpace := files.ReadLinesAsRunes(name)

	rankMap := createRankMap(problemSpace)

	for i, row := range problemSpace {
		for j, char := range row {
			if char == 'T' {
				score, ok := rankMap[coords.NewCoord(i, j)]
				if !ok {
					panic("no score found¬")
				}

				total += score
			}

			if char == 'H' {
				score, ok := rankMap[coords.NewCoord(i, j)]
				if !ok {
					panic("no score found¬")
				}

				total += score * 2
			}
		}
	}

	return total
}

func part3(name string) int {
	problemSpace := files.ReadLines(name)

	meteorCoords := parseMeteorCoords(problemSpace)

	maxX := 0

	for _, coord := range meteorCoords {
		if coord.I > maxX {
			maxX = coord.I
		}
	}

	rankMap := createInvertedRankMap(maxX)

	total := 0

	for _, coord := range meteorCoords {
		time:= 0
		currentPos := [2]int{coord.I, coord.J}

		for {
			currentPos[0]--
			currentPos[1]--
			time++

			if currentPos[1] < 0 || currentPos[0] < 0 {
				panic(fmt.Sprintf("you died: %v", currentPos))
			}

			// Projectile not had time to reach meteor yet
			if time < currentPos[1] {
				continue
			}

			score, ok := rankMap[currentPos]

			if ok {
				total += score
				break
			}
		}
	}

	return total
}

func createRankMap(input [][]rune) map[coords.Coord]int {
	rankMap := make(map[coords.Coord]int)

	for i := range len(input) + 1 {
		yIndex := len(input) - i - 1
		if i < 1 || yIndex < 0 || input[yIndex][1] == '.' {
			continue
		}

		for power := 1; power < len(input[0])/2; power++ {
			currentPos := coords.NewCoord(yIndex, 1)

			for upDiag := 1; upDiag <= power; upDiag++ {
				currentPos = currentPos.MoveBy(coords.NewCoord(-1, 1), 1)
				score, ok := rankMap[currentPos]
				if !ok || power*i < score {
					rankMap[currentPos] = power * i
				}
			}

			for right := 1; right <= power; right++ {
				currentPos = currentPos.MoveBy(coords.NewCoord(0, 1), 1)
				score, ok := rankMap[currentPos]
				if !ok || power*i < score {
					rankMap[currentPos] = power * i
				}
			}

			for currentPos.I < len(input) && currentPos.J < len(input[len(input)-1]) {
				currentPos = currentPos.MoveBy(coords.NewCoord(1, 1), 1)
				score, ok := rankMap[currentPos]
				if !ok || power*i < score {
					rankMap[currentPos] = power * i
				}
			}
		}
	}

	return rankMap
}

func createInvertedRankMap(xSize int) map[[2]int]int {
	rankMap := make(map[[2]int]int)

	for i := range 3 {
		for power := 1; power <= xSize/2; power++ {
			newScore := power * (i + 1)
			currentPos := [2]int{i, 0}

			for upDiag := 1; upDiag <= power; upDiag++ {
				currentPos[0]++
				currentPos[1]++
				updateScore(rankMap,currentPos,newScore)
			}

			for right := 1; right <= power; right++ {
				currentPos[1]++
				updateScore(rankMap,currentPos,newScore)
			}

			for currentPos[0] > 0 && currentPos[1] < xSize {
				currentPos[0]--
				currentPos[1]++
				updateScore(rankMap,currentPos,newScore)
			}
		}
	}

	return rankMap
}

func updateScore(scoreMap map[[2]int]int, currentPos [2]int, newScore int) {
	score, ok := scoreMap[currentPos]
	if !ok || newScore < score {
		scoreMap[currentPos] = newScore
	}
}

func parseMeteorCoords(input []string) []coords.Coord {
	result := []coords.Coord{}

	for _, row := range input {
		coordSlice := ints.FromStringSlice(strings.Split(row, " "))
		result = append(result, coords.NewCoord(coordSlice[1], coordSlice[0]))
	}

	return result
}
