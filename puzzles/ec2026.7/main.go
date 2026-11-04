package main

import (
	"fmt"
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
	return solve(name, 5)
}

func part2(name string) int {
	return solve(name, 10)
}

func part3(name string) int {
	clouds, _ := parseInput(name)

	max := 0
	for id := range clouds {
		cloudsClone := cloneClouds(clouds)

		impulseMap := map[int]int{id: 1}
		for len(impulseMap) > 0 {
			sendImpulse(impulseMap, 10, cloudsClone)
			impulseMap = explodeClouds(cloudsClone)
		}

		total:= sumEnergy(cloudsClone)
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

func solve(name string, impulse int) int {
	clouds, maxID := parseInput(name)
	sendImpulse(map[int]int{maxID: 1}, impulse, clouds)
	return sumEnergy(clouds)
}

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

func sendImpulse(impulseMap map[int]int, power int, clouds map[int]cloud) {
	for power > 0 {
		nextMap := make(map[int]int)
		for id, val := range impulseMap {
			cloud, ok := clouds[id]
			if !ok {
				continue
			}

			cloud.energy += power * val
			clouds[id] = cloud

			for _, connection := range cloud.connections {
				nextMap[connection] += val
			}
		}

		power--
		impulseMap = nextMap
	}
}

func explodeClouds(clouds map[int]cloud) map[int]int {
	impulseMap := make(map[int]int)

	toDelete := []int{}

	for id, cloud := range clouds {
		if cloud.energy >= 100 {
			for _, con := range cloud.connections {
				impulseMap[con]++
			}
			toDelete = append(toDelete, id)
		}
	}

	for _, id := range toDelete {
		delete(clouds, id)
	}

	return impulseMap
}

func cloneClouds(clouds map[int]cloud) map[int]cloud {
	out := make(map[int]cloud)

	for k, v := range clouds {
		out[k] = v
	}

	return out
}

func sumEnergy(clouds map[int]cloud) int {
	total := 0
	for _, cloud := range clouds {
		total += cloud.energy
	}

	return total
}
