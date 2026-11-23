package main

import (
	"fmt"
	"math"
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
	portals, shortcuts := parseInput(name, 1)

	total += totalStrain(portals, nil)
	total -= minTotalStrain(portals, [][][2]int{shortcuts})

	return total
}

func part2(name string) int {
	total := 0
	portals, shortcuts := parseInput(name, 3)
	total += totalStrain(portals, nil)
	addShortcuts(portals, shortcuts)

	combinations := part2Crystals(len(portals))

	total -= minTotalStrain(portals, combinations)

	return total
}

func part3(name string) int {
	total := 0
	portals, shortcuts := parseInput(name, 3)
	total += totalStrain(portals, nil)
	addShortcuts(portals, shortcuts)

	combinations := part3Crystals(len(portals))

	total -= minTotalStrain(portals, combinations)

	return total
}

func parseInput(name string, neighbourDistance int) (map[int][]int, [][2]int) {
	paragraphs := files.ReadParagraphs(name)

	portals := ints.FromStringIgnoreText(paragraphs[0][0])

	portalMap := make(map[int][]int)
	shortcuts := [][2]int{}

	for i := range portals {
		for d := range neighbourDistance {
			offset := d + 1
			portalMap[i] = append(portalMap[i], (i+offset)%(portals), ints.Mod(i-offset, portals))
		}
	}

	for _, shortcutStr := range paragraphs[1] {
		vals := ints.FromDelimitedString(shortcutStr, " ")

		shortcuts = append(shortcuts,
			[2]int{vals[0] - 1, vals[1] - 1},
		)
	}

	return portalMap, shortcuts
}

func totalStrain(portals map[int][]int, preExistingMap map[int]map[int]int) int {
	total := 0

	for portal, connections := range portals {
		shortestDistances := map[int]int{portal: 0}

		if preExistingMap != nil {
			if m, ok := preExistingMap[portal]; ok {
				shortestDistances = m
			}
		} else {
			preExistingMap = make(map[int]map[int]int)
		}

		findShortestDistances(portals, 1, connections, shortestDistances)

		preExistingMap[portal] = shortestDistances

		total += score(shortestDistances)
	}

	return total
}

func addShortcuts(portals map[int][]int, shortcuts [][2]int) {
	for _, shortcut := range shortcuts {
		portals[shortcut[0]] = append(portals[shortcut[0]], shortcut[1])
		portals[shortcut[1]] = append(portals[shortcut[1]], shortcut[0])
	}
}

func findShortestDistances(portalMap map[int][]int, depth int, connections []int, distanceMap map[int]int) {
	nextConns := []int{}

	for _, c := range connections {
		if d, ok := distanceMap[c]; ok && d < depth {
			continue
		}

		distanceMap[c] = depth

		nextConns = append(nextConns, portalMap[c]...)
	}

	if len(nextConns) == 0 {

		return
	}

	findShortestDistances(portalMap, depth+1, nextConns, distanceMap)
}

func score(distanceMap map[int]int) int {
	total := 0
	for _, v := range distanceMap {
		total += v
	}

	return total
}

func minTotalStrain(portalMap map[int][]int, combinations [][][2]int) int {
	min := math.MaxInt
	var preExistingMap map[int]map[int]int
	totalStrain(portalMap, preExistingMap)
	for _, combination := range combinations {

		portalCopy := make(map[int][]int)
		for k, v := range portalMap {
			portalCopy[k] = append(portalCopy[k], v...)
		}

		existingMapCopy := make(map[int]map[int]int)
		for k, v := range preExistingMap {
			for k2, v2 := range v {
				existingMapCopy[k] = map[int]int{k2: v2}
			}
		}

		addShortcuts(portalCopy, combination)

		strain := totalStrain(portalCopy, existingMapCopy)
		if strain < min {
			min = strain
		}
	}

	return min
}

func part2Crystals(portals int) [][][2]int {
	out := [][][2]int{}

	for i := range portals {
		for j := range portals {
			if i >= j {
				continue
			}

			out = append(out, createCombinations([]int{i, j})...)
		}
	}

	return out
}

func part3Crystals(portals int) [][][2]int {
	out := [][][2]int{}

	for i := range portals / 8 {
		in := []int{}

		for j := range 8 {
			in = append(in, (i+(j*portals/8))%portals)
		}

		out = append(out, createCombinations(in)...)
	}

	return out
}

func createCombinations(in []int) [][][2]int {
	out := [][][2]int{}

	if len(in) == 0 {
		return out
	}

	for i := range in {
		if i == 0 {
			continue
		}

		combination := [][2]int{{in[0], in[i]}}

		newIn := []int{}
		for j := range in {
			if j == 0 || j == i {
				continue
			}

			newIn = append(newIn, in[j])
		}

		if len(newIn) == 0 {
			out = append(out, combination)
			continue
		}

		subCombinations := createCombinations(newIn)

		for i := range subCombinations {
			subCombinations[i] = append(subCombinations[i], combination...)
		}

		out = append(out, subCombinations...)
	}

	return out
}
