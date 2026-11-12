package main

import (
	"fmt"
	"slices"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/coords"
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

func part1(name string) int {
	total := 0

	stacks := parseInput(name)
	for i, stack := range stacks {
		score:= stack.runKuhns()
		// score, _ := reflect(stack.positions, math.MaxInt, scoreMap)

		fmt.Println(i, score, stack.positions[0])
		total += score
	}

	return total
}

func part2(name string) int {
	return part1(name)
}

func part3(name string) int {
	return part1(name)
}

type stack struct {
	positions                    []positionType
	horizontalRuns, verticalRuns [][]int
	blockVRunMap map[int]int
}

type positionType struct {
	pos            coords.Coord
	hIndex, vIndex int
}

func (s stack) runKuhns() int {
	owners:=map[int]int{}
	total:=0

	for i:= range s.horizontalRuns {
		visited:=map[int]bool{}

		if tryClaim(i, s.horizontalRuns, s.verticalRuns, visited, owners, s.blockVRunMap) {
			total ++
		}

	}

	return total
}

func tryClaim(i int, hRuns [][]int, vRuns [][]int, visited map[int]bool, owners map[int]int, hvmap map[int]int) bool {
	for _, blockIndex:= range hRuns[i] {
		vRunIndex:= hvmap[blockIndex]
		if visited[vRunIndex] {
			continue
		}

		visited[vRunIndex] = true

		if owner, ok:=owners[vRunIndex]; !ok {
			owners[vRunIndex] = i
			return true
		} else {
			if tryClaim(owner, hRuns, vRuns, visited, owners, hvmap) {
				owners[vRunIndex] = i
				return true
			}
		}
	} 

	return false
}

func (s *stack) populateRuns() {
	s.blockVRunMap = map[int]int{}
	slices.SortFunc(s.positions, func(a,b positionType) int {
		if a.pos.I == b.pos.I {
			return a.pos.J - b.pos.J
		}

		return a.pos.I - b.pos.I
	})
	for i, position := range s.positions {
		if position.hIndex == -1 {
			hRun := []int{}
			newPos := position.pos
			j := i
			for j != -1 {
				hRun = append(hRun, j)
				s.positions[j].hIndex = len(s.horizontalRuns)
				newPos = newPos.Right(1)
				j = slices.IndexFunc(s.positions, func(p positionType) bool { return p.pos.Equals(newPos) })
			}

			s.horizontalRuns = append(s.horizontalRuns, hRun)
		}

		if position.vIndex == -1 {
			vRun := []int{}
			newPos := position.pos
			j := i
			for j != -1 {
				vRun = append(vRun, j)
				s.positions[j].vIndex = len(s.verticalRuns)
				newPos = newPos.Down(1)
				j = slices.IndexFunc(s.positions, func(p positionType) bool { return p.pos.Equals(newPos) })
			}

			s.verticalRuns = append(s.verticalRuns, vRun)
		}

		s.blockVRunMap[i] = s.positions[i].vIndex
	}
}

func parseInput(name string) []stack {
	grid := files.ReadLinesAsRunes(name)
	out := []stack{}

	for i, row := range grid {
		for j, char := range row {
			if char != '#' {
				continue
			}
			newPos := coords.NewCoord(i, j)

			if slices.ContainsFunc(out, func(s stack) bool {
				return slices.ContainsFunc(s.positions, func(p positionType) bool {
					return p.pos.Equals(newPos)
				})
			}) {
				continue
			}

			newPositions := []positionType{{pos: newPos, hIndex: -1, vIndex: -1}}
			queue := []coords.Coord{}
			for _, adj := range newPos.GetAdjacent() {
				queue = append(queue, adj)
			}
			for len(queue) > 0 {
				nextQueue := []coords.Coord{}

				for _, p := range queue {
					if p.I >= 0 &&
						p.I < len(grid) &&
						p.J >= 0 &&
						p.J < len(grid[0]) &&
						grid[p.I][p.J] == '#' &&
						!slices.ContainsFunc(newPositions, func(pos positionType) bool {
							return pos.pos.Equals(p)
						}) {
						newPositions = append(newPositions, positionType{pos: p, hIndex: -1, vIndex: -1})

						for _, adj := range p.GetAdjacent() {
							nextQueue = append(nextQueue, adj)
						}
					}
				}

				queue = nextQueue
			}

			newStack := stack{positions: newPositions}
			newStack.populateRuns()

			out = append(out, newStack)
		}
	}

	return out
}
