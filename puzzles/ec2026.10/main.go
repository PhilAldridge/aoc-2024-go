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
	walls, ghosts := parseWallsAndGhosts(name)
	total := 0

	for range 10 {
		for i := range ghosts {
			if collision := ghosts[i].Move(walls); collision {
				total++
			}
		}

		for i := range ghosts {
			if collision := ghosts[i].Overlap(ghosts); collision {
				total++
			}
		}
	}

	return total
}

func part2(name string) int {
	return solveWithLoopCheck(name, 10000)
}

func part3(name string) int {
	return solveWithLoopCheck(name, 1000000000000)
}

func solveWithLoopCheck(name string, limit int) int {
	walls, ghosts := parseWallsAndGhosts(name)
	rounds := 0
	state := stateType{ghosts.copy()}

	for len(ghosts) > 0 {
		rounds++
		for i := len(ghosts) - 1; i >= 0; i-- {
			ghosts[i].Move(walls)
		}

		if next, died := ghosts.Kill(limit); died {
			state = stateType{}
			ghosts = next
		}

		for i := len(ghosts) - 1; i >= 0; i-- {
			ghosts[i].Overlap(ghosts)
		}

		if next, died := ghosts.Kill(limit); died {
			state = stateType{}
			ghosts = next
		}

		if roundsInLoop, collisionsInLoop, loopFound := state.checkLoop(ghosts); loopFound {
			minRepeats := math.MaxInt
			for i, g := range ghosts {
				if collisionsInLoop[i] == 0 {
					continue
				}
				canLoop := (limit - g.collisions) / collisionsInLoop[i]
				if canLoop < minRepeats {
					minRepeats = canLoop
				}
			}
			if len(ghosts) == 1 {
				minRepeats--
			}
			rounds += roundsInLoop * minRepeats
			for i := range ghosts {
				ghosts[i].collisions = ghosts[i].collisions + minRepeats*collisionsInLoop[i]
			}
			state = stateType{}
		} else {
			state = append(state, ghosts.copy())
		}
	}

	return rounds
}

type ghostType struct {
	id                  int
	position, direction coords.Coord
	collisions          int
}

type ghostsType []ghostType

func (gs ghostsType) copy() ghostsType {
	out := make(ghostsType, len(gs))
	for i, g := range gs {
		out[i] = ghostType{
			id:         g.id,
			direction:  g.direction,
			position:   g.position,
			collisions: g.collisions,
		}
	}

	return out
}

func (gs ghostsType) equal(gs2 ghostsType) bool {
	for i, g := range gs {
		g2 := gs2[i]

		if !g.position.Equals(g2.position) || !g.direction.Equals(g2.direction) {
			// fmt.Printf("failed at %d with %v,%v", i, g, g2)
			return false
		}
	}
	return true
}

type stateType []ghostsType

func (s stateType) checkLoop(gs ghostsType) (int, []int, bool) {
	for i, prevState := range s {
		if prevState.equal(gs) {
			collisions := []int{}
			for j, g := range gs {
				collisions = append(collisions, g.collisions-prevState[j].collisions)
			}
			return len(s) - i, collisions, true
		}
	}

	return 0, nil, false
}

func parseWallsAndGhosts(name string) (map[coords.Coord]bool, ghostsType) {
	walls := make(map[coords.Coord]bool)
	ghosts := []ghostType{}

	for i, row := range files.ReadLines(name) {
		for j, char := range row {
			switch char {
			case '#':
				walls[coords.NewCoord(i, j)] = true
			case '>':
				ghosts = append(ghosts, ghostType{
					position:  coords.NewCoord(i, j),
					direction: coords.Right,
					id:        len(ghosts),
				})
			case '<':
				ghosts = append(ghosts, ghostType{
					position:  coords.NewCoord(i, j),
					direction: coords.Left,
					id:        len(ghosts),
				})
			case 'v':
				ghosts = append(ghosts, ghostType{
					position:  coords.NewCoord(i, j),
					direction: coords.Down,
					id:        len(ghosts),
				})
			case '^':
				ghosts = append(ghosts, ghostType{
					position:  coords.NewCoord(i, j),
					direction: coords.Up,
					id:        len(ghosts),
				})
			}
		}
	}

	return walls, ghosts
}

func (g *ghostType) Move(walls map[coords.Coord]bool) bool {
	newPos := g.position.MoveBy(g.direction, 1)

	if walls[newPos] {
		g.direction = coords.TurnRight(g.direction)
		g.collisions++
		return true
	}

	g.position = newPos
	return false
}

func (g *ghostType) Overlap(ghosts []ghostType) bool {
	for _, ghost := range ghosts {
		if g.position.Equals(ghost.position) && g.id != ghost.id {
			g.direction = coords.TurnBack(g.direction)
			g.collisions++
			return true
		}
	}

	return false
}

func (gs ghostsType) Kill(limit int) (ghostsType, bool) {
	out := ghostsType{}

	for _, g := range gs {
		if g.collisions < limit {
			out = append(out, g)
		}
	}

	return out, len(gs) != len(out)
}

func print(walls map[coords.Coord]bool, ghosts []ghostType, total int) {
	for i := range 24 {
		for j := range 24 {
			pos := coords.NewCoord(i, j)
			if walls[pos] {
				fmt.Print("#")
				continue
			}

			if slices.ContainsFunc(ghosts, func(g ghostType) bool {
				return g.position.Equals(pos)
			}) {
				fmt.Print("G")
				continue
			}

			fmt.Print(".")
		}
		fmt.Println()
	}
	fmt.Printf("%d\n\n", total)
}
