package main

import (
	"fmt"
	"math"
	"slices"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/graphs"
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

	grid := files.ReadLines(name)
	total += len(grid) + len(grid[0])
	total *= 2

	for i, row := range grid {
		for j, _ := range row {
			if j+1 < len(row) && grid[i][j] != grid[i][j+1] {
				total++
			}

			if i+1 < len(grid) && grid[i][j] != grid[i+1][j] {
				total++
			}
		}
	}

	return total
}

func part2(name string) int {
	graph := parseGraph(name)

	return graph.MinimumSpanningTree()
}

func part3(name string) int {
	graph := parseGraph(name)
	grid := files.ReadLines(name)

	used := usedBytes(grid)

	minFound := math.MaxInt
	for a := range used {
		bindGraph, bound := graph.Bind(a)
		for b := range used {
			if a < b && !slices.Contains(bound, b) {
				bindGraph2, _ := bindGraph.Bind(b)
				tree := bindGraph2.MinimumSpanningTree()
				if tree < minFound {
					minFound = tree
				}
			}
		}
	}

	return minFound
}

func parseGraph(name string) *graphs.GraphMatrix[byte] {
	grid := files.ReadLines(name)
	graph := graphs.NewGraphMatrix[byte]()
	edgeVertex := unusedByte(grid)
	graph.AddVertex(edgeVertex)

	for i, row := range grid {
		for j, _ := range row {
			vertex := grid[i][j]
			if vertex == edgeVertex {
				panic("choose a new edge vertex!")
			}
			graph.AddVertex(vertex)

			if i == 0 {
				graph.AddEdge(vertex, edgeVertex, 1)
			}
			if j == 0 {
				graph.AddEdge(vertex, edgeVertex, 1)
			}

			if i+1 >= len(grid) {
				graph.AddEdge(vertex, edgeVertex, 1)
			} else if grid[i+1][j] != vertex {
				graph.AddEdge(vertex, grid[i+1][j], 1)
			}

			if j+1 >= len(grid[0]) {
				graph.AddEdge(vertex, edgeVertex, 1)
			} else if grid[i][j+1] != vertex {
				graph.AddEdge(vertex, grid[i][j+1], 1)
			}

		}
	}

	return graph
}

func unusedByte(grid []string) byte {
	used := usedBytes(grid)
	for b := range 256 {
		if !used[byte(b)] {
			return byte(b)
		}
	}
	panic("every byte value is used in the grid")
}

func usedBytes(grid []string) map[byte]bool {
	used := make(map[byte]bool)
	for _, row := range grid {
		for j := 0; j < len(row); j++ {
			used[row[j]] = true
		}
	}

	return used
}
