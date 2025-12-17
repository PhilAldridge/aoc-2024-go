package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/ints"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1 answer: ", part1("input"))
	split := time.Now()
	fmt.Println("Part 2 answer: ", part2("input"))
	
	fmt.Println()
	fmt.Println("Part 1: ", split.Sub(start))
	fmt.Println("Part 2: ", time.Since(split))
}

func part1(name string) int {
	sections:= files.ReadParagraphs(name)

	total:=0

	for _,input:= range sections[len(sections)-1] {
		grid := parseGrid(input)

		// Cannot fit - not enough area
		if grid.x * grid.y < ints.Sum(grid.presentIndices)*9 {
			continue
		}

		// Must fit, even if you don't try rearranging squares
		if (grid.x/3) * (grid.y/3) >= ints.Sum(grid.presentIndices) {
			total++
			continue
		}

		panic("TODO: figure out how fit presents")
	}

	return total
}

func part2(name string) int {
	return 0
}

func parseInput(name string) ([]presentType, []gridType) {
	sections:= files.ReadParagraphs(name)

	presents:= []presentType{}

	for _,input:= range sections[0:len(sections)-1] {
		presents = append(presents, parsePresent(input))
	}

	grids:= []gridType{}

	for _,input:= range sections[len(sections)-1] {
		grids = append(grids, parseGrid(input))
	}

	return presents, grids
}

func parsePresent(input []string) presentType {
	present:= presentType{
		index: ints.FromString(input[0][0:len(input[0])-1]),
		coords: [][2]int{},
	}

	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			if input[i+1][j] == '#' {
				present.coords = append(present.coords, [2]int{i,j})
			}
		}
	}

	return present
}

func parseGrid(input string) gridType {
	splitOne:= strings.Split(input, ":")

	splitTwo:= ints.FromStringSlice(strings.Split(splitOne[0],"x"))

	splitThree:= ints.FromStringSlice(strings.Split(splitOne[1]," ")[1:])

	return gridType{
		x:splitTwo[0],
		y:splitTwo[1],
		presentIndices: splitThree,
	}
}

type presentType struct {
	index int
	coords [][2]int
}

type gridType struct {
	x int
	y int
	presentIndices []int
}
