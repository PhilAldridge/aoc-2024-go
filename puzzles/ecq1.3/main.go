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

type snailType struct {
	ring, startY int
}

func part1(name string) int {
	input:= files.ReadLines(name)

	snails:= []snailType{}

	for _,row:= range input {
		split:= strings.Split(row," ")
		x:= ints.FromString(split[0][2:])
		y:= ints.FromString(split[1][2:])

		snails = append(snails, snailType{
			ring: x+y-1,
			startY: y-1,
		})
	}

	total:= 0

	for _, snail:= range snails {
		newY:= ints.Mod(snail.startY - 100,snail.ring)
		newX:= snail.ring - newY
		total += 100*(newY+1) + newX
	}


	return total
}

func part2(name string) int {
	input:= files.ReadLines(name)

	snails:= []snailType{}

	for _,row:= range input {
		split:= strings.Split(row," ")
		x:= ints.FromString(split[0][2:])
		y:= ints.FromString(split[1][2:])

		snails = append(snails, snailType{
			ring: x+y-1,
			startY: y-1,
		})
	}

	slices.SortFunc(snails, func(a,b snailType) int {
		return b.ring - a.ring
	})

	n:=0

	for {
		time:= snails[0].startY + n*snails[0].ring

		if checkTime(snails[1:],time) {
			return time
		}

		n++
	}
}


func part3(name string) int {
	return part2(name)
}


func checkTime(snails []snailType, time int) bool {
	for _, snail:= range snails {
		if ints.Mod(time-snail.startY, snail.ring) != 0 {
			return false
		}
	}

	return true
}