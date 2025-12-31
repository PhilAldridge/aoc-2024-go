package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
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
	lines := files.ReadLines(name)

	return parseWord(lines)
}

func part2(name string) int {
	lines := files.ReadParagraphs(name)
	total := 0

	input := parseInput(lines)

	for _, tablet := range input {
		word := parseWord(tablet)

		total += calculatePower(word)
	}

	return total
}

func part3(name string) int {
	total := 0

	return total
}

func parseWord(input []string) string {
	columns := make(map[int]string)
	rows := make(map[int]string)

	for j := 2; j < len(input[0])-2; j++ {
		columns[j-2] = string(input[0][j]) + string(input[1][j]) + string(input[len(input)-2][j]) + string(input[len(input)-1][j])
		rows[j-2] = string(input[j][0]) + string(input[j][1]) + string(input[j][len(input[j])-2]) + string(input[j][len(input[j])-1])
	}

	result := ""

	for i := 0; i < len(input)-4; i++ {
		row := rows[i]

		for j := 0; j < len(input[i])-4; j++ {
			column := columns[j]

			for k := 0; k < len(row); k++ {
				if strings.Contains(column, row[k:k+1]) {
					result += row[k : k+1]
					break
				}
			}

		}
	}

	return result
}

func parseInput(input [][]string) [][]string {
	output := [][]string{}

	for _, chunk := range input {
		lot := make([][]string, strings.Count(chunk[0], " ")+1)
		for _, row := range chunk {
			split := strings.Split(row, " ")

			for i, rowChunk := range split {
				lot[i] = append(lot[i], rowChunk)
			}
		}

		output = append(output, lot...)
	}

	return output
}

func calculatePower(word string) int {
	total := 0

	for i, char := range word {
		total += (i + 1) * (int(char-'A') + 1)
	}

	return total
}
