package main

import (
	"fmt"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/coords"
	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1 answer: ", part1("input.txt", 4))
	split := time.Now()
	fmt.Println("Part 2 answer: ", part2("input2.txt", 20))
	split2 := time.Now()
	fmt.Println("Part 3 answer: ", part3("input3.txt"))

	fmt.Println()
	fmt.Println("Part 1: ", split.Sub(start))
	fmt.Println("Part 2: ", split2.Sub(split))
	fmt.Println("Part 2: ", time.Since(split2))
}

func part1(name string, moves int) int {
	input := files.ReadLines(name)
	dragon, sheep, _ := parseInput(input)
	total := 0

	queue := []coords.Coord{dragon}
	positionsReached := map[coords.Coord]int{
		dragon: 0,
	}

	for i := range moves {
		newQueue := []coords.Coord{}

		for _, position := range queue {
			moves := position.GetKnightMoves()

			for _, move := range moves {
				if !move.InInput(input) {
					continue
				}

				if _, ok := positionsReached[move]; ok {
					continue
				}

				positionsReached[move] = i + 1
				newQueue = append(newQueue, move)
			}
		}
		queue = newQueue
	}

	for _, s := range sheep {
		if _, ok := positionsReached[s]; ok {
			total++
		}
	}

	return total
}

type stateType struct {
	position coords.Coord
	move     int
}

func part2(name string, moves int) int {
	input := files.ReadLines(name)
	dragon, sheep, hides := parseInput(input)
	total := 0

	queue := []coords.Coord{dragon}
	positionsReached := map[stateType]bool{{
		position: dragon,
		move:     0,
	}: true,
	}

	for i := range moves {
		newQueue := []coords.Coord{}

		for _, position := range queue {
			moves := position.GetKnightMoves()

			for _, move := range moves {
				if !move.InInput(input) {
					continue
				}

				dragonState := stateType{
					position: move,
					move:     i + 1,
				}

				if _, ok := positionsReached[dragonState]; ok {
					continue
				}

				positionsReached[dragonState] = true
				newQueue = append(newQueue, move)
			}
		}
		queue = newQueue
	}

	for _, s := range sheep {
		sheepState := stateType{
			position: s,
			move:     0,
		}
		for range moves {
			sheepState.move++

			if !hides[sheepState.position] && positionsReached[sheepState] {
				total++
				break
			}

			sheepState.position = sheepState.position.Down(1)

			if !hides[sheepState.position] && positionsReached[sheepState] {
				total++
				break
			}
		}
	}

	return total
}

type gameVariant struct {
	dragon coords.Coord
	sheep  sheepsType
}

type sheepType struct {
	position coords.Coord
	eaten    bool
	void     bool
}

type sheepsType [5]sheepType

func (s sheepsType) eaten() int {
	total := 0

	for _, sheep := range s {
		if sheep.eaten {
			total++
		}
	}

	return total
}

func (s sheepsType) getEaten(dragon coords.Coord, hides map[coords.Coord]bool) sheepsType {
	if hides[dragon] {
		return s
	}

	new := sheepsType{}
	for i := range s {
		if s[i].eaten || s[i].void || !s[i].position.Equals(dragon) {
			new[i] = s[i]
		} else {
			new[i] = sheepType{
				position: s[i].position,
				eaten:    true,
				void:     s[i].void,
			}
		}
	}

	return new
}

func (s sheepsType) move(index int) sheepsType {
	new := sheepsType{}
	for i := range s {
		if i == index {
			new[i] = sheepType{
				position: s[i].position.Down(1),
				eaten:    s[i].eaten,
				void:     s[i].void,
			}
		} else {
			new[i] = s[i]
		}
	}

	return new
}

func (s sheepsType) noneLeft(height int) bool {
	for _,sheep:= range s {
		if !sheep.void && sheep.position.I < height && !sheep.eaten {
			return false
		}
	}

	return true
}

func NewSheepsType(s []coords.Coord) sheepsType {
	output := sheepsType{}

	for i := range output {
		if i >= len(s) {
			output[i] = sheepType{
				void: true,
			}
		} else {
			output[i] = sheepType{
				position: s[i],
			}
		}
	}

	return output
}

func part3(name string) int {
	input := files.ReadLines(name)
	dragon, sheep, hides := parseInput(input)

	initialState:= gameVariant{
		dragon: dragon,
		sheep: NewSheepsType(sheep),
	}

	memo:= make(map[gameVariant]int)

	return moveRecursive(initialState, hides, memo, len(sheep), input)
}

func parseInput(input []string) (coords.Coord, []coords.Coord, map[coords.Coord]bool) {
	var dragon coords.Coord
	sheep := []coords.Coord{}
	hides := make(map[coords.Coord]bool)

	for i, row := range input {
		for j, char := range row {
			pos := coords.NewCoord(i, j)
			switch char {
			case 'D':
				dragon = pos
			case 'S':
				sheep = append(sheep, pos)
			case '#':
				hides[pos] = true
			}
		}
	}

	return dragon, sheep, hides
}

func moveRecursive(current gameVariant, hides map[coords.Coord]bool, memo map[gameVariant]int, toEat int, input []string) int {
	if score, ok := memo[current]; ok {
		return score
	}

	if current.sheep.eaten() == toEat {
		return 1
	}

	if current.sheep.noneLeft(len(input)) {
		return 0
	}

	total := 0
	sheepMoves := []gameVariant{}

	for i, s := range current.sheep {
		if s.void || s.eaten || s.position.I >= len(input) {
			continue
		}

		newPos := s.position.Down(1)
		if (newPos.Equals(current.dragon) && !hides[newPos]) {
			continue
		}

		sheepMoves = append(sheepMoves, gameVariant{
			dragon: current.dragon,
			sheep:  current.sheep.move(i),
		})
	}

	if len(sheepMoves) == 0 {
		sheepMoves = append(sheepMoves, current)
	}

	for _, sheepMove := range sheepMoves {
		dragonMoves := current.dragon.GetKnightMoves()

		for _, dragonMove := range dragonMoves {
			if !dragonMove.InInput(input) {
				continue
			}

			newState := gameVariant{
				dragon: dragonMove,
				sheep:  sheepMove.sheep.getEaten(dragonMove,hides),
			}

			if newState.sheep.eaten() == toEat {
				total++

				memo[newState] = 1

				continue
			}

			total += moveRecursive(newState, hides, memo, toEat, input)
		}
	}

	memo[current] = total

	return total

}
