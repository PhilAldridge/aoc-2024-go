package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/ints"
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
	wg := &sync.WaitGroup{}
	results := make(chan int)
	for _, machine := range machines {
		if machine.index != 103 {
			continue
		}
		wg.Add(1)
		go getPresses(machine, wg, results)
	}

	// goroutine that closes channel after work is done
	go func() {
		wg.Wait()
		close(results)
	}()

	return TotalFromChan(results)
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

func getPresses(machine machine, wg *sync.WaitGroup, channel chan int) {
	defer wg.Done()
	lights := getSortedLights(machine)

	buttonOrder := []int{}
	solutionsArray := [][]int{}

	combinationMap := make(map[combinationParams][][]int)

	result := math.MaxInt

	for i, nextLight := range lights {
		if i == 0 {
			buttonOrder = nextLight.buttons
			solutionsArray = getCombinations(combinationParams{
				total:       nextLight.joltage,
				buttonCount: len(nextLight.buttons),
			}, combinationMap)
			continue
		}

		nextSolutions := [][]int{}
		buttonsLeft := []int{}

		for _, button := range nextLight.buttons {
			if !slices.Contains(buttonOrder, button) {
				buttonsLeft = append(buttonsLeft, button)
			}
		}

		for _, solution := range solutionsArray {
			total := nextLight.joltage

			for _, button := range nextLight.buttons {
				index := slices.Index(buttonOrder, button)

				if index != -1 {
					total -= solution[index]
				}
			}

			presses := ints.Sum(solution)

			if total < 0 || (len(buttonsLeft) == 0 && total != 0) || presses >= result {
				continue
			}

			combinations := getCombinations(combinationParams{
				total: total, buttonCount: len(buttonsLeft),
			}, combinationMap)

			for _, combination := range combinations {
				nextSolution := []int{}
				nextSolution = append(nextSolution, solution...)
				nextSolution = append(nextSolution, combination...)
				nextSolutions = append(nextSolutions, nextSolution)
			}

		}

		buttonOrder = append(buttonOrder, buttonsLeft...)
		solutionsArray = nextSolutions
		// fmt.Println(nextLight.joltage)
		// fmt.Println(buttonOrder)
		// fmt.Println(solutionsArray)

		// fmt.Println()
	}

	result = calculateMinimumSolution(solutionsArray, lights, buttonOrder)
	fmt.Println(machine.index)
	channel <- result
}

func calculateMinimumSolution(solutions [][]int, lights []light, buttonOrder []int) int {
	min := math.MaxInt
	for _, solution := range solutions {
		presses := ints.Sum(solution)
		if presses > min {
			continue
		}
		min = presses
	}

	return min
}

type combinationParams struct {
	total       int
	buttonCount int
}

func getCombinations(p combinationParams, combinationMap map[combinationParams][][]int) [][]int {
	prevResult, ok := combinationMap[p]

	if ok {
		return prevResult
	}

	if p.total == 0 {
		result := make([]int, p.buttonCount)
		combinationMap[p] = [][]int{result}
		return [][]int{result}
	}

	if p.buttonCount == 1 {
		combinationMap[p] = [][]int{{p.total}}
		return [][]int{{p.total}}
	}

	var combinations [][]int

	for i := 0; i <= p.total; i++ {
		nextCombinations := getCombinations(combinationParams{
			total:       p.total - i,
			buttonCount: p.buttonCount - 1,
		}, combinationMap)

		for _, nextCombination := range nextCombinations {
			nextCombination = append([]int{i}, nextCombination...)
			combinations = append(combinations, nextCombination)
		}
	}

	combinationMap[p] = combinations
	return combinations
}

func getSortedLights(machine machine) []light {
	var lights []light

	for lightIndex, joltage := range machine.joltage {
		nextLight := light{
			joltage: joltage,
			buttons: []int{},
		}

		for buttonIndex, button := range machine.wiring {
			if slices.Contains(button, lightIndex) {
				nextLight.buttons = append(nextLight.buttons, buttonIndex)
			}
		}

		lights = append(lights, nextLight)
	}

	sort.Slice(lights, func(i, j int) bool {
		a, b := lights[i], lights[j]

		aPower := ints.Pow(a.joltage, len(a.buttons)-1) / ints.Factorial(len(a.buttons)-1)
		bPower := ints.Pow(b.joltage, len(b.buttons)-1) / ints.Factorial(len(b.buttons)-1)

		return aPower < bPower
	})

	return lights
}

type light struct {
	joltage int
	buttons []int
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

func TotalFromChan(ch <-chan int) int {
	total := 0

	for v := range ch {
		total += v
		fmt.Println(total)
	}

	return total
}
