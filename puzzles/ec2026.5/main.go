package main

import (
	"fmt"
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
	grid := files.ReadLinesAsRunes(name)

	grid = getLand(grid)
	return countLandType(grid, '#')
}

func part2(name string) int {
	grid := files.ReadLinesAsRunes(name)

	grid = getLand(grid)
	grid = getSand(grid)
	return countLandType(grid, '=')
}

func part3(name string) int {
	grid := files.ReadLinesAsRunes(name)

	grid = getLand(grid)
	grid = getSand(grid)
	grid = getTrees(grid)
	return countLandType(grid, '@')
}

func getLand(grid [][]rune) [][]rune {
	grid[len(grid)/2][len(grid[0])/2] = '#'
	changed := []coords.Coord{coords.NewCoord(len(grid)/2, len(grid[0])/2)}

	for len(changed) > 0 {
		nextChanged := []coords.Coord{}

		for _, pos := range changed {
			for _, adj := range pos.GetAdjacentIncludingDiagonals() {
				if !coords.GenericInInput(adj, grid) || grid[adj.I][adj.J] == '#' {
					continue
				}

				count := 0

				for _, adj2 := range adj.GetAdjacentIncludingDiagonals() {
					if coords.GenericInInput(adj2, grid) && grid[adj2.I][adj2.J] != '0' {
						count++
					}
				}

				if count >= 5 {
					nextChanged = append(nextChanged, adj)
					grid[adj.I][adj.J] = '#'
				}
			}
		}

		changed = nextChanged
	}

	return grid
}

func getSand(grid [][]rune) [][]rune {
	changed := []coords.Coord{}

	for i, row := range grid {
		for j, pos := range row {
			if pos != '#' {
				continue
			}

			posCoord := coords.NewCoord(i, j)

			for _, adj := range posCoord.GetAdjacentIncludingDiagonals() {
				if !coords.GenericInInput(adj, grid) || grid[adj.I][adj.J] == '1' || grid[adj.I][adj.J] == '0' {
					changed = append(changed, posCoord)
					grid[i][j] = '='
					break
				}
			}
		}
	}

	for len(changed) > 0 {
		nextChanged := []coords.Coord{}

		for _, pos := range changed {
			for _, adj := range pos.GetAdjacentIncludingDiagonals() {
				if !coords.GenericInInput(adj, grid) || grid[adj.I][adj.J] != '#' {
					continue
				}

				count := 0

				for _, adj2 := range adj.GetAdjacentIncludingDiagonals() {
					if coords.GenericInInput(adj2, grid) && grid[adj2.I][adj2.J] == '=' {
						count++
					}
				}

				if count >= 5 {
					nextChanged = append(nextChanged, adj)
					grid[adj.I][adj.J] = '='
				}
			}
		}

		changed = nextChanged
	}

	return grid
}

func getTrees(grid [][]rune) [][]rune {
	changed := []coords.Coord{}

	for i, row := range grid {
		for j, pos := range row {
			if pos != '#' {
				continue
			}

			if stampSuccess(i,j, grid) {
				changed = append(changed, coords.NewCoord(i,j))
				grid[i][j]= '@'
			}
		}
	}

	for range 2 {
		nextChanged := []coords.Coord{}

		for _, pos := range changed {
			for _, adj := range pos.GetAdjacentIncludingDiagonals() {
				if !coords.GenericInInput(adj, grid) || grid[adj.I][adj.J] != '#' {
					continue
				}

				count := 0
				sandy:= false

				for _, adj2 := range adj.GetAdjacentIncludingDiagonals() {
					if coords.GenericInInput(adj2, grid) {
						if grid[adj2.I][adj2.J] == '=' {
							sandy = true
							break
						}

						if  grid[adj2.I][adj2.J] == '@' {
							count++
						}
					}
				}

				if count >= 4 && !sandy {
					nextChanged = append(nextChanged, adj)
				}
			}
		}

		for _, coord:= range nextChanged {
			grid[coord.I][coord.J] = '@'
		}

		changed = nextChanged
	}

	return grid
}

func stampSuccess(i, j int, grid [][]rune) bool {
	for iOffset := -2; iOffset <= 2; iOffset++ {
		for jOffset := -2; jOffset <= 2; jOffset++ {
			if ints.Abs(iOffset) == 2 && ints.Abs(jOffset) == 2 {
				continue
			}

			if iOffset == 0 && jOffset == 0 {
				continue
			}

			if !coords.GenericInInput(coords.NewCoord(i+iOffset, j+jOffset), grid) {
				continue
			}

			if grid[i+iOffset][j+jOffset] != '#' && grid[i+iOffset][j+jOffset] != '@' {
				return false
			}
		}
	}
	return true
}

func countLandType(grid [][]rune, landType rune) int {
	total := 0
	for _, row := range grid {
		for _, pos := range row {
			if pos == landType {
				total++
			}
		}
	}

	return total
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}

	fmt.Println()
}
