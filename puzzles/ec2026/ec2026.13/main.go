package main

import (
	"fmt"
	"math"
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

func part1(name string) int {
	total := 0

	systems := parse(name)
	for _, system := range systems {
		subtotal := 0
		for !system.Done() {
			subtotal += system.Move()
		}
		total += subtotal
	}

	return total
}

func part2(name string) int {
	return part1(name)
}

func part3(name string) int {
	return part1(name)
}

type clockType struct {
	position, speed, moves int
	clockwise              bool
}

func (c clockType) NextPosition(base int) int {
	if c.moves <= 0 {
		return c.position
	}

	multiplier := 1
	if !c.clockwise {
		multiplier = -1
	}

	return ints.Mod(c.position+c.speed*multiplier, base)
}

func (c clockType) NextPositions(base int) []int {
	if c.moves <= 0 {
		return []int{c.position}
	}

	out := make([]int, c.speed+1)
	multiplier := 1
	if !c.clockwise {
		multiplier = -1
	}

	for i := range c.speed + 1 {
		out[i] = ints.Mod(c.position+i*multiplier, base)
	}

	return out
}

type clocksType []clockType

type systemType struct {
	base   int
	clocks clocksType
	states [][]int
}

func parse(name string) []systemType {
	paragraphs := files.ReadParagraphs(name)
	out := make([]systemType, len(paragraphs))

	for sysI, lines := range paragraphs {
		hands := ints.FromDelimitedString(lines[0][6:], ":")
		directions := strings.Split(lines[1][10:], ":")
		currents := ints.FromDelimitedString(lines[2][8:], ":")
		moves := ints.FromDelimitedString(lines[3][5:], ":")

		base := 1

		for _, hand := range hands {
			base = ints.LCM(base, hand)
		}

		clocks := []clockType{}

		for i := range hands {
			multiplier := base / hands[i]

			clocks = append(clocks, clockType{
				position:  currents[i] * multiplier,
				speed:     multiplier,
				moves:     moves[i],
				clockwise: directions[i] == "CW",
			})
		}

		out[sysI] = systemType{
			base:   base,
			clocks: clocks,
			states: [][]int{},
		}
	}

	return out
}

func (s *systemType) Move() int {
	collisions := 0
	state := make([]int, len(s.clocks)+1)
	for i, c1 := range s.clocks {
		for j := i + 1; j < len(s.clocks); j++ {
			if collision(c1, s.clocks[j], s.base) {
				collisions++
			}
		}
		s.clocks[i] = clockType{
			position:  c1.NextPosition(s.base),
			speed:     c1.speed,
			moves:     c1.moves - 1,
			clockwise: c1.clockwise,
		}

		state[i] = s.clocks[i].position
	}

	state[len(state)-1] = collisions

	if lengthOfLoop, loopCollisions := stateRecurrence(state, s.states); lengthOfLoop > 0 {
		maxLoops := math.MaxInt
		for _, c := range s.clocks {
			if c.moves <= 0 {
				continue
			}

			loops := (c.moves - 1) / lengthOfLoop
			if loops < maxLoops {
				maxLoops = loops
			}
		}

		if maxLoops <= 0 {
			s.states = [][]int{}
			return collisions
		}

		for i := range s.clocks {
			s.clocks[i].moves = s.clocks[i].moves - lengthOfLoop*maxLoops
		}

		s.states = [][]int{}
		return loopCollisions*maxLoops + collisions
	} else {
		if slices.ContainsFunc(s.clocks, func(c clockType) bool {
			return c.moves == 0
		}) {
			s.states = [][]int{}
		} else {
			s.states = append(s.states, state)
		}
	}

	return collisions
}

func stateRecurrence(state []int, prevStates [][]int) (lengthOfLoop int, collisions int) {
	index := slices.IndexFunc(prevStates, func(ps []int) bool {
		for i := range len(ps) - 1 {
			if state[i] != ps[i] {
				return false
			}
		}

		return true
	})

	if index == -1 {
		return lengthOfLoop, collisions
	}

	for i := index; i < len(prevStates); i++ {
		collisions += prevStates[i][len(state)-1]
	}

	return len(prevStates) - index, collisions
}

func collision(a, b clockType, base int) bool {
	aPos := a.NextPositions(base)
	bPos := b.NextPositions(base)

	if a.clockwise == b.clockwise {
		return sameDirCollision(aPos, bPos)
	}

	for i := range aPos {
		for j := range bPos {
			if i == 0 && j == 0 {
				continue
			}

			if aPos[i] == bPos[j] {
				return true
			}
		}
	}

	return false
}

func sameDirCollision(a, b []int) bool {
	if len(a) > len(b) {
		a, b = b, a
	}

	index := slices.Index(b, a[0])
	if index <= 0 {
		return false
	}
	for i := range a {
		if i+index >= len(b) || a[i] != b[i+index] {
			return false
		}
	}

	return true
}

func (s systemType) Done() bool {
	for _, c := range s.clocks {
		if c.moves > 0 {
			return false
		}
	}

	return true
}
