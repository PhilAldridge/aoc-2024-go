package graphs

import (
	"fmt"
	"math"
	"math/bits"
	"slices"
)

type GraphMatrix[T comparable] struct {
	vertices map[T]int
	indexMap map[int]T
	matrix   [][]int
}

func NewGraphMatrix[T comparable]() *GraphMatrix[T] {
	return &GraphMatrix[T]{
		vertices: make(map[T]int),
		indexMap: make(map[int]T),
		matrix:   [][]int{},
	}
}

func (g *GraphMatrix[T]) AddVertex(v T) {
	if _, exists := g.vertices[v]; !exists {
		g.vertices[v] = len(g.matrix)
		g.indexMap[len(g.matrix)] = v
		for i := range g.matrix {
			g.matrix[i] = append(g.matrix[i], 0)
		}
		g.matrix = append(g.matrix, make([]int, len(g.matrix)+1))
	}
}

func (g *GraphMatrix[T]) AddEdge(from, to T, weight int) {
	g.AddVertex(from)
	g.AddVertex(to)
	i, j := g.vertices[from], g.vertices[to]
	g.matrix[i][j] += weight
	g.matrix[j][i] += weight
}

func (g *GraphMatrix[T]) ReplaceEdge(from, to T, weight int) {
	g.AddVertex(from)
	g.AddVertex(to)
	i, j := g.vertices[from], g.vertices[to]
	g.matrix[i][j] = weight
	g.matrix[j][i] = weight
}

func (g *GraphMatrix[T]) ShortCut(from, to T, weight int) {
	g.AddVertex(from)
	g.AddVertex(to)
	i, j := g.vertices[from], g.vertices[to]

	existingWeight := g.matrix[i][j]

	if existingWeight == 0 || existingWeight > weight {
		g.matrix[i][j] = weight
		g.matrix[j][i] = weight
	}
}

func (g *GraphMatrix[T]) RemoveEdge(from, to T) {
	g.ReplaceEdge(from, to, 0)
}

func (g *GraphMatrix[T]) RemoveVertex(v T) {
	if i, exists := g.vertices[v]; exists {
		g.matrix = append(g.matrix[:i], g.matrix[i+1:]...)
		for j := range g.matrix {
			g.matrix[j] = append(g.matrix[j][:i], g.matrix[j][i+1:]...)
		}

		delete(g.vertices, v)
		delete(g.indexMap, i)

		newMap := make(map[int]T)

		for index, vertex := range g.indexMap {
			if index > i {
				g.vertices[vertex] = index - 1
				newMap[index-1] = vertex
			} else {
				newMap[index] = vertex
			}
		}

		g.indexMap = newMap
	}
}

func (g *GraphMatrix[T]) ShortestPath() int {
	minFound := math.MaxInt
	// Greedy approach for upper bound
	for i := range g.matrix {
		distanceStartingAtI := g.GreedyShortestPath(i)
		if distanceStartingAtI < minFound {
			minFound = distanceStartingAtI
		}
	}

	// Exhaustive approach not going over upper bound
	memo := map[memoType]int{}
	return g.RecursiveShortestPath(minFound, memo)
}

func (g *GraphMatrix[T]) GreedyShortestPath(startIndex int) int {
	visitedIndices := []int{startIndex}
	currentIndex := startIndex
	distance := 0

	for len(visitedIndices) < len(g.matrix) {
		min := math.MaxInt
		nextIndex := 0
		for i, val := range g.matrix[currentIndex] {
			if slices.Contains(visitedIndices, i) || val == 0 {
				continue
			}
			if val < min {
				min = val
				nextIndex = i
			}
		}

		if min == math.MaxInt {
			return min
		}

		visitedIndices = append(visitedIndices, nextIndex)
		currentIndex = nextIndex
		distance += min
	}

	return distance
}

type memoType struct {
	bitMask uint
	lastI   int
}

func (g *GraphMatrix[T]) RecursiveShortestPath(minFound int, memo map[memoType]int) int {
	for i := range g.matrix {
		result, ok := g.shortestPathR(memoType{
			bitMask: 1 << i,
			lastI:   i,
		}, minFound, memo)
		if ok && result < minFound {
			minFound = result
		}
	}

	return minFound
}

func (g *GraphMatrix[T]) shortestPathR(visited memoType, left int, memo map[memoType]int) (int, bool) {
	if bits.OnesCount(visited.bitMask) == len(g.matrix) {
		return 0, true
	}

	if v, exists := memo[visited]; exists {
		return v, v != math.MaxInt
	}

	minFound := left
	found := false

	for i, adjacencies := range g.matrix {
		bit := uint(1 << i)
		edgeWeight := adjacencies[visited.lastI]
		if edgeWeight != 0 && visited.bitMask&bit != bit && edgeWeight <= left {
			if bits.OnesCount(visited.bitMask) == len(g.matrix)-1 {
				return edgeWeight, true
			}

			result, ok := g.shortestPathR(memoType{
				bitMask: visited.bitMask + bit,
				lastI:   i,
			}, left-edgeWeight, memo)
			if ok && result+edgeWeight < minFound {
				minFound = result + edgeWeight
				found = true
			}
		}
	}

	if found {
		memo[visited] = minFound
	} else {
		memo[visited] = math.MaxInt
	}

	return minFound, found
}

func (g *GraphMatrix[T]) Print() {
	for i := range g.matrix {
		for j := i + 1; j < len(g.matrix); j++ {
			fmt.Printf("%c %c: %d\n", g.indexMap[i], g.indexMap[j], g.matrix[i][j])
		}
	}
}

func (g *GraphMatrix[T]) MinimumSpanningTree() int {
	tree := []int{0}
	size := 0

	for len(tree) < len(g.matrix) {
		minEdgeWeight := math.MaxInt
		minEdgeToIndex := 0

		for _, i := range tree {
			for j := range g.matrix {
				edgeWeight := g.matrix[i][j]
				if edgeWeight == 0 || slices.Contains(tree, j) {
					continue
				}

				if edgeWeight < minEdgeWeight {
					minEdgeWeight = edgeWeight
					minEdgeToIndex = j
				}
			}
		}

		if minEdgeToIndex == 0 {
			panic("spanning tree not found!")
		}

		tree = append(tree, minEdgeToIndex)
		size += minEdgeWeight
	}

	return size
}

func (g *GraphMatrix[T]) Bind(a T) (*GraphMatrix[T], []T) {
	newGraph := NewGraphMatrix[T]()

	i := g.vertices[a]

	toBindIndices := []int{i}
	toBind := []T{a}

	for j, edge := range g.matrix[i] {
		if edge > 0 {
			toBindIndices = append(toBindIndices, j)
			toBind = append(toBind, g.indexMap[j])
		}
	}

	for x := range g.matrix {
		for y, edge := range g.matrix[x] {
			if edge > 0 {
				if slices.Contains(toBindIndices, x) {
					if !slices.Contains(toBindIndices, y) {
						newGraph.ShortCut(a, g.indexMap[y], edge)
					}
				} else if slices.Contains(toBindIndices, y) {
					newGraph.ShortCut(a, g.indexMap[x], edge)
				} else {
					newGraph.ShortCut(g.indexMap[y], g.indexMap[x], edge)
				}
			}
		}
	}

	return newGraph, toBind
}
