package main

import (
	"fmt"
	"maps"
	"math"
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
	fmt.Println("Part 3 answer: ", part3("input3.txt", 1000000))

	fmt.Println()
	fmt.Println("Part 1: ", split.Sub(start))
	fmt.Println("Part 2: ", split2.Sub(split))
	fmt.Println("Part 2: ", time.Since(split2))
}

func part1(name string) int {
	return solve1and2(name, 5)
}

func part2(name string) int {
	return solve1and2(name, 10)
}

func part3(name string, energyLimit int) int {
	clouds, _ := parseInput(name)

	max := 0
	for id := range clouds {
		cloudsClone := maps.Clone(clouds)

		impulseMap := map[int]map[int]int{id: map[int]int{10: 1}}
		sendImpulse(impulseMap, cloudsClone, energyLimit)

		total := sumEnergy(cloudsClone)
		if total > max {
			max = total
		}
	}

	return max
}

type cloud struct {
	energy      int
	connections []int
}

// parts 1 and 2 have the same logic, with a different initial impulse
// both send that impulse to the cloud with the most connections
// there are no explosions, so energyLimit is set to a number that will not be reached
func solve1and2(name string, impulse int) int {
	clouds, maxID := parseInput(name)
	sendImpulse(map[int]map[int]int{maxID: map[int]int{impulse: 1}}, clouds, math.MaxInt)
	return sumEnergy(clouds)
}

// parseInput makes a map of each cloud's energy and connections against their id
// it also returns the id of the cloud with the most connections
func parseInput(name string) (map[int]cloud, int) {
	out := make(map[int]cloud)
	max := 0
	maxID := 0

	for _, row := range files.ReadLines(name) {
		split := strings.Split(row, ", ")

		id := ints.FromString(split[0][3:])
		energy := ints.FromString(split[1][7:])
		connections := ints.FromStringSlice(strings.Split(split[2][12:], ","))

		if len(connections) > max {
			max = len(connections)
			maxID = id
		}

		out[id] = cloud{
			energy:      energy,
			connections: connections,
		}
	}

	return out, maxID
}

// impulse goes id, power, times
func sendImpulse(impulseMap map[int]map[int]int, clouds map[int]cloud, energyLimit int) {
	// loop until all energy has been dissipated
	for len(impulseMap) > 0 {
		nextImpulseMap := make(map[int]map[int]int)
		toExplode := []int{}

		for id, val := range impulseMap {
			cloud, ok := clouds[id]
			if !ok {
				continue
			}

			// send impulse to connected clouds, lowering power by 1
			for power, amount := range val {
				cloud.energy += power * amount
				if power > 1 {
					for _, connection := range cloud.connections {
						if _, ok := nextImpulseMap[connection]; !ok {
							nextImpulseMap[connection] = make(map[int]int)
						}
						nextImpulseMap[connection][power-1] += amount
					}
				}

			}

			// add impulses from cloud if it will explode this round
			if cloud.energy >= energyLimit {
				toExplode = append(toExplode, id)
				for _, connection := range cloud.connections {
					if _, ok := nextImpulseMap[connection]; !ok {
						nextImpulseMap[connection] = make(map[int]int)
					}
					nextImpulseMap[connection][10] += 1
				}
			}
			clouds[id] = cloud

		}

		// explode clouds - doing this at the end to avoid issues deleting from map whilst looping through it
		for _, id := range toExplode {
			delete(clouds, id)
		}

		impulseMap = nextImpulseMap
	}
}

func sumEnergy(clouds map[int]cloud) int {
	total := 0
	for _, cloud := range clouds {
		total += cloud.energy
	}

	return total
}
