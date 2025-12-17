package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/ints"
	z3 "github.com/mitchellh/go-z3"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1 answer: ", part1("input.txt"))
	split := time.Now()
	fmt.Println("Part 2 answer: ", part2("input.txt"))

	fmt.Println()
	fmt.Println("Part 1: ", split.Sub(start))
	fmt.Println("Part 2: ", time.Since(split))
}

func part1(name string) int {
	machines := parseInput(name)
	total := 0
	for _, machine := range machines {
		stateMap := make(map[string]int)
		baseState := strings.Repeat(".", len(machine.lights))
		stateMap[baseState] = 0
		for _, nextWiring := range machine.wiring {
			statesToAdd := []lightState{}
			for state, presses := range stateMap {
				newState := getNextState(state, presses, nextWiring)
				if val, ok := stateMap[string(newState.lights)]; !ok || val > newState.presses {
					statesToAdd = append(statesToAdd, newState)
				}
			}

			for _, state := range statesToAdd {
				stateMap[string(state.lights)] = state.presses
			}
		}
		total += stateMap[machine.lights]
	}

	return total
}

func part2(name string) int {
	machines := parseInput(name)

	total:= 0
	for _,machine:= range machines {
		total+= solveMachine(machine)
	}


	return total
}

func solveMachine(machine machine) int {
	config:= z3.NewConfig()
	ctx:= z3.NewContext(config)
	config.Close()
	defer ctx.Close()

	s := ctx.NewSolver()
	defer s.Close()

	xs:= make([]*z3.AST, len(machine.wiring))

	zero:= ctx.Int(0, ctx.IntSort())
	sum:= ctx.Int(0,ctx.IntSort())

	for i := range machine.wiring {
		//create a variable for each wire
		xs[i] = ctx.Const(ctx.Symbol(fmt.Sprintf("x_%d", i)), ctx.IntSort())

		//add variable to sum of all variables
		sum = sum.Add(xs[i])

		//variable non-negative
		s.Assert(xs[i].Ge(zero))
	}

	for i, jolt:= range machine.joltage {
		// create equation
		wires:= []*z3.AST{}
		for j,wire:= range machine.wiring {
			if slices.Contains(wire, i) {
				wires = append(wires, xs[j])
			}
		}

		//sum of all wires affecting joltage index = joltage
		s.Assert(zero.Add(wires...).Eq(ctx.Int(jolt,ctx.IntSort())))
	}

	
	v:=s.Check()
	total:= 0

	//to find minimum solution, create a loop reducing the allowed sum each time, until the problem is unsolvable
	for v== z3.True {
		m:= s.Model()
		total= 0
		for _,assignment:= range m.Assignments() {
			total += assignment.Int()
		}

		newTotal:= ctx.Int(total,ctx.IntSort())
		s.Assert(sum.Lt(newTotal))

		v = s.Check()
	}

	//return lowest valid solution
	return total
}

func parseInput(name string) []machine {
	var machines []machine

	lines := files.ReadLines(name)
	for i, line := range lines {
		split := strings.Split(line, " ")
		machines = append(machines, machine{
			index:   i,
			lights:  split[0][1 : len(split[0])-1],
			wiring:  getWiring(split[1 : len(split)-1]),
			joltage: parseBracketedInts(split[len(split)-1]),
		})
	}
	return machines
}

func getWiring(input []string) [][]int {
	var wiring [][]int
	for _, set := range input {
		wiring = append(wiring, parseBracketedInts(set))
	}

	return wiring
}

func parseBracketedInts(input string) []int {
	split := strings.Split(input[1:len(input)-1], ",")
	return ints.FromStringSlice(split)
}

func getNextState(state string, presses int, wiring []int) lightState {
	newState := lightState{
		lights:  []byte(state),
		presses: presses + 1,
	}
	for _, val := range wiring {
		if newState.lights[val] == '.' {
			newState.lights[val] = '#'
		} else {
			newState.lights[val] = '.'
		}
	}
	return newState
}

type machine struct {
	index   int
	lights  string
	wiring  [][]int
	joltage []int
}

type lightState struct {
	lights  []byte
	presses int
}