package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/ints"
	"github.com/PhilAldridge/aoc-2024-go/pkg/sets"
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
	dnaStrands := parseInput(name)
	total := 0
	for id, child := range dnaStrands {
		p1id, p2id, ok := getParents(id, dnaStrands)
		if ok {
			total += similarityScore(child, dnaStrands[p1id]) * similarityScore(child, dnaStrands[p2id])
		}
	}

	return total
}

func part2(name string) int {
	return part1(name)
}

func part3(name string) int {
	dnaStrands := parseInput(name)
	
	families := []*sets.Set[int]{}

	for id := range dnaStrands {
		p1id, p2id, ok := getParents(id, dnaStrands)
		if ok {
			familiesFoundIds := []int{}
			for i, family := range families {
				if family.Contains(id) || family.Contains(p1id) || family.Contains(p2id) {
					family.Add(id)
					family.Add(p1id)
					family.Add(p2id)
					familiesFoundIds = append(familiesFoundIds, i)
				}
			}

			if len(familiesFoundIds) == 0 {
				newFamily := sets.NewSet[int]()
				newFamily.Add(id)
				newFamily.Add(p1id)
				newFamily.Add(p2id)
				families = append(families, newFamily)
			}

			families = combineFamilies(families, familiesFoundIds)
		}
	}

	maxMembers := 0
	score:=0

	for _, family := range families {
		if family.Size() > maxMembers {
			maxMembers = family.Size()
			score = ints.Sum(family.List())
		}
	}

	return score
}

func parseInput(name string) map[int]string {
	input := files.ReadLines(name)

	output := make(map[int]string)

	for _, row := range input {
		split := strings.Split(row, ":")
		id := ints.FromString(split[0])
		output[id] = split[1]

	}

	return output
}

func checkParentage(child, parent1, parent2 string) bool {
	for i := range child {
		if child[i] != parent1[i] && child[i] != parent2[i] {
			return false
		}
	}

	return true
}

func similarityScore(child, parent string) int {
	total := 0

	for i := range child {
		if child[i] == parent[i] {
			total++
		}
	}

	return total
}

func getParents(childid int, dnaStrands map[int]string) (int, int, bool) {
	child := dnaStrands[childid]

	for p1id, p1 := range dnaStrands {
		if p1id == childid {
			continue
		}

		for p2id, p2 := range dnaStrands {
			if p2id == childid || p1id >= p2id {
				continue
			}

			if checkParentage(child, p1, p2) {
				return p1id, p2id, true
			}
		}
	}

	return 0, 0, false
}

func combineFamilies(families []*sets.Set[int], toCombineIds []int) []*sets.Set[int] {
	if len(toCombineIds) <= 1 {
		return families
	}

	output := []*sets.Set[int]{}
	combined := sets.NewSet[int]()

	for i := range families {
		if slices.Contains(toCombineIds, i) {
			combined.AddSlice(families[i].List())
		} else {
			output = append(output, families[i])
		}
	}
	output = append(output, combined)

	return output
}
