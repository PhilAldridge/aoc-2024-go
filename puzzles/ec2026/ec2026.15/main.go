package main

import (
	"fmt"
	"math"
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
	cavern := parseInput(name)

	allCollectedBitMask := 0
	for i := range cavern.sulphurDeposits {
		allCollectedBitMask |= 1 << (i)
	}

	initialState := stateType{
		lastEdge:    nil,
		currentNode: cavern.start,
	}

	states := map[stateType]int{initialState: 0}
	statesToCheck := []stateType{initialState}

	for {
		nextState := stateType{}
		min := math.MaxInt
		nextStatesToCheck := []stateType{}

		for _, state := range statesToCheck {
			dist := states[state]
			new := false
			for _, edge := range state.currentNode.edges {
				depositCollectedBitmap := state.count
				if edge == state.lastEdge {
					continue
				}
				if edge.sulphurCount > 0 {
					i := slices.Index(cavern.sulphurDeposits, edge)
					if i == -1 {
						panic("what?")
					}
					depositCollectedBitmap |= 1 << (i)
				}

				newNode := edge.to
				if newNode == state.currentNode {
					newNode = edge.from
				}

				testState := stateType{
					count:       depositCollectedBitmap,
					lastEdge:    edge,
					currentNode: newNode,
				}

				_, exists := states[testState]
				if exists {
					continue
				}

				new = true
				newDist := dist + edge.weight
				if newDist < min {
					nextState = testState
					min = newDist
				}
			}

			if new {
				nextStatesToCheck = append(nextStatesToCheck, state)
			}
		}

		if min == math.MaxInt {
			panic("no states added")
		}

		if nextState.count == allCollectedBitMask {
			if nextState.currentNode == cavern.end {
				return min
			}

		}

		states[nextState] = min

		statesToCheck = append(nextStatesToCheck, nextState)
	}
}

func part2(name string) int {
	return solve(name, 2026)
}

func part3(name string) int {
	return solve(name, 202620262026)
}

func solve(name string, limit int) int {
	cavern := parseInput(name)

	initialState := stateType{
		lastEdge:    nil,
		currentNode: cavern.start,
	}

	states := map[stateType]mapOut{initialState: {
		distance: 0,
		bestPath: []*edgeType{},
	}}
	statesToCheck := []stateType{initialState}

	loopFound := []loopInfo{}

	for {
		nextState := stateType{}
		min := math.MaxInt
		nextStatesToCheck := []stateType{}
		path := []*edgeType{}

		for _, state := range statesToCheck {
			stateOut := states[state]
			new := false
			for _, edge := range state.currentNode.edges {
				newCount := state.count + edge.sulphurCount
				if edge == state.lastEdge {
					continue
				}

				newNode := edge.to
				if newNode == state.currentNode {
					newNode = edge.from
				}

				testState := stateType{
					count:       newCount,
					lastEdge:    edge,
					currentNode: newNode,
				}

				_, exists := states[testState]
				if exists {
					continue
				}

				new = true
				newDist := stateOut.distance + edge.weight
				if newDist < min {
					nextState = testState
					min = newDist
					path = append([]*edgeType{}, stateOut.bestPath...)
					path = append(path, edge)
				}
			}

			if new {
				nextStatesToCheck = append(nextStatesToCheck, state)
			}
		}

		if min == math.MaxInt {
			panic("no states added")
		}

		if nextState.count >= limit && nextState.currentNode == cavern.end {
			return min
		}

		states[nextState] = mapOut{
			distance: min,
			bestPath: path,
		}

		statesToCheck = append(nextStatesToCheck, nextState)

		if len(loopFound) < 10 && nextState.count > 600 {
			distance, sulphur, found := getLoop(path)
			if !found {
				continue
			}

			canLoopTimes := (limit - nextState.count - 600) / sulphur

			nextState.count += sulphur * canLoopTimes

			loopFound = append(loopFound, loopInfo{
				s: nextState,
				o: mapOut{
					distance: min + distance*canLoopTimes,
				},
			})

			if len(loopFound) == 9 {
				slices.SortFunc(loopFound, func(a, b loopInfo) int {
					ratioA := float32(a.s.count) / float32(a.o.distance)
					ratioB := float32(b.s.count) / float32(b.o.distance)

					if ratioA > ratioB {
						return -1
					}

					if ratioA < ratioB {
						return 1
					}
					return 0
				})

				states = map[stateType]mapOut{
					loopFound[0].s: loopFound[0].o,
				}

				statesToCheck = []stateType{loopFound[0].s}
			}
		}
	}
}

type loopInfo struct {
	s stateType
	o mapOut
}

type nodeType struct {
	edges    []*edgeType
	position coords.Coord
}

type edgeType struct {
	from, to     *nodeType
	weight       int
	sulphurCount int
	direction    int
}

type cavernType struct {
	start, end      *nodeType
	sulphurDeposits []*edgeType
}

type stateType struct {
	currentNode *nodeType
	lastEdge    *edgeType
	count       int
}

type mapOut struct {
	distance int
	bestPath []*edgeType
}

func getLoop(edges []*edgeType) (distance, sulphur int, found bool) {
	lastEdge := edges[len(edges)-1]

	for i := len(edges) - 2; i >= 0; i-- {
		if edges[i] == lastEdge {
			for j := i + 1; j < len(edges); j++ {
				distance += edges[j].weight
				sulphur += edges[j].sulphurCount
			}

			return distance, sulphur, true
		}
	}

	return 0, 0, false
}

func parseInput(name string) cavernType {
	grid := files.ReadLines(name)

	start := &nodeType{}
	allNodesFound := map[coords.Coord]*nodeType{}
	sulphurDeposits := []*edgeType{}
	end := &nodeType{}

	for j, char := range grid[len(grid)/2] {
		pos := coords.NewCoord(len(grid)/2, j)
		node := &nodeType{}

		if n, exists := allNodesFound[pos]; exists {
			node = n
		}

		switch char {
		case '#', '@':
			panic("not middle")
		case 'S':
			start = node
		case 'E':
			end = node
		case '.':
			paths := 0
			for _, adj := range pos.GetAdjacent() {
				if adj.InInput(grid) && grid[adj.I][adj.J] != '#' {
					paths++
				}
			}

			if paths <= 2 {
				continue
			}
		}

		node.position = pos

		for _, adj := range pos.GetAdjacent() {
			if adj.InInput(grid) && grid[adj.I][adj.J] != '#' {
				direction := 0
				if adj.I < pos.I {
					direction = 1
				}
				if adj.I > pos.I {
					direction = 2
				}
				endPos, weight, sulphurCollected := followPathToNode([]coords.Coord{pos, adj}, grid)
				if n, exists := allNodesFound[endPos]; exists {
					i := slices.IndexFunc(n.edges, func(e *edgeType) bool {
						return e.to.position == pos && e.direction == direction
					})
					if i != -1 {
						node.edges = append(node.edges, n.edges[i])
					} else {
						newEdge := edgeType{
							from:         node,
							to:           n,
							weight:       weight,
							sulphurCount: sulphurCollected,
							direction:    direction,
						}
						node.edges = append(node.edges, &newEdge)

						if sulphurCollected > 0 {
							sulphurDeposits = append(sulphurDeposits, &newEdge)
						}
					}
				} else {
					endNode := &nodeType{position: endPos}
					allNodesFound[endPos] = endNode
					newEdge := edgeType{
						from:         node,
						to:           endNode,
						weight:       weight,
						sulphurCount: sulphurCollected,
						direction:    direction,
					}
					node.edges = append(node.edges, &newEdge)
					if sulphurCollected > 0 {
						sulphurDeposits = append(sulphurDeposits, &newEdge)
					}
				}
			}
		}

		allNodesFound[pos] = node
	}

	return cavernType{
		start:           start,
		end:             end,
		sulphurDeposits: sulphurDeposits,
	}
}

func followPathToNode(path []coords.Coord, grid []string) (endPos coords.Coord, weight int, sulphurCount int) {
	last := path[len(path)-1]

	if grid[last.I][last.J] == 'S' || grid[last.I][last.J] == 'E' {
		return last, 1, 0
	}

	next := coords.Coord{}
	found := false

	for _, adj := range last.GetAdjacent() {
		if !adj.InInput(grid) || slices.Contains(path, adj) || grid[adj.I][adj.J] == '#' {
			continue
		}
		if found {
			return last, 1, 0 // junction
		}
		found = true
		next = adj
	}

	if !found {
		panic("lost!")
	}

	here := 0
	if grid[last.I][last.J] == '@' {
		here = 1
	}

	endPos, weight, s := followPathToNode(append(path, next), grid)
	return endPos, weight + 1, here + s
}
