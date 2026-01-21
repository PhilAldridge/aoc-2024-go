package main

import (
	"fmt"
	"slices"
	"time"
	"unicode"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1 answer: ", part1("input.txt"))
	split := time.Now()
	fmt.Println("Part 2 answer: ", part2("input2.txt"))
	split2 := time.Now()
	fmt.Println("Part 3 answer: ", part3("input3.txt",1000,1000))

	
	fmt.Println()
	fmt.Println("Part 1: ", split.Sub(start))
	fmt.Println("Part 2: ", split2.Sub(split))
	fmt.Println("Part 2: ", time.Since(split2))
}

func part1(name string) int {
	input:= files.Read(name)
	mentorCount:=0
	total:= 0

	for _, char:= range input {
		if char == 'A' {
			mentorCount++
		}
		if char == 'a' {
			total+=mentorCount
		}
	}

	return total
}

func part2(name string) int {
	input:= files.Read(name)
	mentorCount:=make(map[rune]int)
	total:=0

	for _, char:= range input {
		if char >= 'A' && char <='Z' {
			mentorCount[char]++
			continue
		}
		total+=mentorCount[unicode.ToUpper(char)]
	}


	return total
}


func part3(name string, repeats, distanceLimit int) int {
	input:= files.Read(name)
	total:=0
	mentorPos:=make(map[rune][]int)
	novicePos:=make(map[rune][]int)

	for i, char:= range input {
		if char >= 'A' && char <='Z' {
			addToMap(char, i,len(input),repeats,mentorPos)
			continue
		}
		addToMap(unicode.ToUpper(char),i,len(input),repeats,novicePos)
	}

	for novice, positions:= range novicePos {
		mentorPositions:= mentorPos[novice]
		slices.Sort(positions)
		slices.Sort(mentorPositions)
		fmt.Println(string(novice),len(positions),len(mentorPositions))
		
		mentorMinI:=0

		for _, position:= range positions {
			min:= position - distanceLimit
			max:= position + distanceLimit

			for i:= mentorMinI; i<len(mentorPositions); i++ {
				mentorPosition := mentorPositions[i]
				if mentorPosition<min {
					mentorMinI = i
					continue
				}
				if mentorPosition>max {
					break
				}
				
				total ++
			}
		}
	}

	return total
}

func addToMap(char rune, pos, length, repeats int, mapping map[rune][]int) {
	max:= length*repeats
	for pos<max {
		mapping[char] = append(mapping[char], pos)
		pos += length
	}
}

