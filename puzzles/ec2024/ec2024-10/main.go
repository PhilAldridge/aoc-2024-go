package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
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
	lines := files.ReadLinesAsRunes(name)
	offsets := []int{0, 1, 6, 7}

	var total int

	changesMade := true

	//loop until no changes to lines made
	for changesMade {
		changesMade = false
		total = 0
		for i, row := range lines {
			for j, char := range row {

				//Start of block to check
				if char == '*' && i+7 < len(lines) && j+7 < len(lines[i+1]) && lines[i+1][j+1] == '*' {
					word := 0
					complete := true

					//Loop through middle of block
					for i2 := 2; i2 < 6; i2++ {
						for j2 := 2; j2 < 6; j2++ {
							//Map of runes on edge of block for that row/column
							rowSet := sets.NewSet[rune]()
							columnSet := sets.NewSet[rune]()

							for _, offset := range offsets {
								rowSet.Add(lines[i+i2][j+offset])
								columnSet.Add(lines[i+offset][j+j2])
							}

							//Already filled in
							if lines[i+i2][j+j2] != '.' {
								word += (j2 - 1 + 4*(i2-2)) * (int(lines[i+i2][j+j2]-'A') + 1)

								//Use filled in value to complete row
								if rowSet.Contains('?') && !columnSet.Contains('?') && !rowSet.Contains(lines[i+i2][j+j2]) {
									for _, offset:= range offsets {
										if lines[i+i2][j+offset] == '?' {
											lines[i+i2][j+offset] = lines[i+i2][j+j2]
											break
										}
									}
									changesMade = true
									continue
								}

								//Use filled in value to complete column
								if columnSet.Contains('?') && !rowSet.Contains('?') && !columnSet.Contains(lines[i+i2][j+j2]) {									
									for _, offset:= range offsets {
										if lines[i+offset][j+j2] == '?' {
											lines[i+offset][j+j2] = lines[i+i2][j+j2]
											break
										}
									}
									changesMade = true
									continue
								}

								continue
							}

							complete = false

							intersection := sets.Intersection(rowSet, columnSet)

							//No ? marks and one answer == fill it in
							if intersection.Size() == 1 && intersection.List()[0] != '?' {
								lines[i+i2][j+j2] = intersection.List()[0]
								changesMade = true
								continue
							}

							//All but one value in column already used == fill in that value
							for i3 := 2; i3 < 6; i3++ {
								columnSet.Remove(lines[i+i3][j+j2])
							}

							if !columnSet.Contains('?') && columnSet.Size() == 1 {
								lines[i+i2][j+j2] = columnSet.List()[0]
								changesMade = true
								continue
							}

							//All but one value in row already used == fill in that value
							for j3 := 2; j3 < 6; j3++ {
								rowSet.Remove(lines[i+i2][j+j3])
							}

							if !rowSet.Contains('?') && rowSet.Size() == 1 {
								lines[i+i2][j+j2] = rowSet.List()[0]
								changesMade = true
								continue
							}
						}
					}

					if complete {
						total += word
					}
				}
			}
		}
	}

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
