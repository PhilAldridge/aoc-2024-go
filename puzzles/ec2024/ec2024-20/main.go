package main

import (
	"fmt"
	"math"
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
	input := files.ReadLines(name)
	startingState := startState(input)

	changeMap := map[state]int{
		startingState: 1000,
	}

	queue := []state{startingState}

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		if next.time == 100 {
			continue
		}

		queue = append(queue, getNext(next, input, changeMap)...)
	}

	max := 0
	for k, v := range changeMap {
		if k.time == 100 && v > max {
			max = v
		}
	}

	return max
}

func part2(name string) int {
	input := files.ReadLines(name)
	startingState := startState(input)

	changeMap := map[state]int{
		startingState: 10000,
	}

	queue := []state{startingState}

	min := math.MaxInt

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		if !next.aCollected && (next.bCollected || next.cCollected) {
			continue
		}

		if !next.bCollected && (next.cCollected) {
			continue
		}

		if next.time >= min {
			continue
		}

		if next.aCollected &&
			next.bCollected &&
			next.cCollected &&
			next.position.Equals(startingState.position) {
			if changeMap[next] >= 10000 {
				min = next.time
			}
			continue
		}

		queue = append(queue, getNext(next, input, changeMap)...)
	}

	return min
}

func part3(name string) int {
	// Assumptions:
	// - + and # are sparse, turning takes too much altitude.
	// - therefore the best longterm action is to find the column where you lose the least altitude and stay there
	// - since # are sparse, assume we can get to the first + using manhattan distance
	// Therefore time = distanceSouth + distanceToBestColumn

	// Find best column closest to start column (Count = -s subtract +s)
	input := files.ReadLines(name)
	columnScore := make(map[int]int)
	startCol := 0
	altitude := 384400

	for _, row := range input {
		for j, char := range row {
			switch char {
			case 'S':
				startCol = j
				columnScore[j] += 1
			case '.':
				columnScore[j] += 1
			case '#':
				columnScore[j] += 100
			case '-':
				columnScore[j] += 2
			case '+':
				columnScore[j] -= 1
			}
		}
	}

	bestColumn := 0
	bestScore := math.MaxInt

	for column, score := range columnScore {
		if score < bestScore {
			bestColumn = column
			bestScore = score
		} else if score == bestScore && ints.Abs(startCol-column) < ints.Abs(startCol-bestColumn) {
			bestColumn = column
		}
	}

	// Subtract distance from start to column from altitude
	altitude -= ints.Abs(startCol - bestColumn)

	// Distance = (altitude/count)*len(input) + how far you can get with the remaining altitude = altitude%count
	distance := (altitude/bestScore - 1) * len(input)

	altitude = altitude%bestScore + bestScore

	for altitude > 0 {
		for i := range input {
			switch input[(i+1)%len(input)][bestColumn] {
			case 'S', '.':
				altitude--
			case '-':
				altitude -= 2
			case '+':
				altitude++
			}
			distance++
			if altitude <= 0 {
				break
			}
		}
	}

	return distance
}

type state struct {
	position   coords.Coord
	direction  coords.Coord
	time       int
	aCollected bool
	bCollected bool
	cCollected bool
}

func startState(input []string) state {
	for i, row := range input {
		for j, char := range row {
			if char == 'S' {
				return state{
					position:  coords.NewCoord(i, j),
					direction: coords.NewCoord(1, 0),
				}
			}
		}
	}

	panic("no start provided")
}

func getNext(next state, input []string, changeMap map[state]int) []state {
	directions := coords.DirectionsInOrder
	queue := []state{}
	for _, direction := range directions {
		if direction.IsOposite(next.direction) {
			continue
		}

		newState := state{
			position:   next.position.Add(direction),
			direction:  direction,
			time:       next.time + 1,
			aCollected: next.aCollected,
			bCollected: next.bCollected,
			cCollected: next.cCollected,
		}

		height := changeMap[next]

		if !newState.position.InInput(input) {
			continue
		}

		switch input[newState.position.I][newState.position.J] {
		case '#':
			continue
		case '.', 'S':
			height--
		case '+':
			height++
		case '-':
			height -= 2
		case 'A':
			height--
			newState.aCollected = true
		case 'B':
			height--
			newState.bCollected = true
		case 'C':
			height--
			newState.cCollected = true
		default:
			fmt.Println(string(input[newState.position.I][newState.position.J]))
			panic("character not accounted for!")
		}

		existing, ok := changeMap[newState]

		if ok && existing >= height {
			continue
		}

		changeMap[newState] = height
		queue = append(queue, newState)
	}
	return queue
}
