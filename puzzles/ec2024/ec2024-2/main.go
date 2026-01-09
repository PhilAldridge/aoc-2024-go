package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/mystrings"
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

	words, input := parseInput(name)

	for _, line := range input {
		for i := 0; i < len(line); i++ {
			for _, word := range words {
				if word[0] == line[i] &&
					i+len(word) <= len(line) &&
					word == line[i:i+len(word)] {
					total++
				}
			}
		}
	}

	return total
}

func part2(name string) int {
	words, input := parseInput(name)
	words = addReversals(words)

	total:=0

	for _, line := range input {
		runicChars:= make(map[int]bool)
		for i := 0; i < len(line); i++ {
			for _, word := range words {
				if word[0] == line[i] &&
					i+len(word) <= len(line) &&
					word == line[i:i+len(word)] {
					for j:=0; j<len(word); j++ {
						runicChars[i+j] = true
					}
				}
			}
		}
		total += len(runicChars)
	}

	return total
}

func part3(name string) int {
	words, input := parseInput(name)
	words = addReversals(words)

	
	runicChars:= make(map[[2]int]bool)

	for lineIndex, line := range input {
		for i := 0; i < len(line); i++ {
			for _, word := range words {
				if word[0] == line[i]  {
					runesToAdd:= [][2]int{}
					for j:=0; j<len(word); j++ {
						if word[j]==line[(i+j)%len(line)] {
							runesToAdd = append(runesToAdd, [2]int{lineIndex,(i+j)%len(line)})
							continue
						}
						break
					}
					if len(runesToAdd) == len(word) {
						for _,runeToAdd:= range runesToAdd {
							runicChars[runeToAdd] = true
						}
					}

					runesToAdd= [][2]int{}
					for j:=0; j<len(word); j++ {
						if lineIndex+j >= len(input) {
							break
						}

						if word[j]==input[(lineIndex+j)][i] {
							runesToAdd = append(runesToAdd, [2]int{(lineIndex+j),i})
							continue
						}
						break
					}
					if len(runesToAdd) == len(word) {
						for _,runeToAdd:= range runesToAdd {
							runicChars[runeToAdd] = true
						}
					}
				}
			}
		}
	}

	return len(runicChars)
}

func parseInput(name string) ([]string, []string) {
	lines := files.ReadLines(name)

	words := strings.Replace(lines[0], "WORDS:", "", 1)

	return strings.Split(words, ","), lines[2:]
}

func addReversals(words []string) []string {
	newWords := []string{}

	for _, word := range words {
		if mystrings.IsPalindrome(word) {
			continue
		}
		newWords = append(newWords, mystrings.Reverse(word))
	}

	return append(words, newWords...)
}
