package main

import (
	"fmt"
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
	plants,_,_:= parseInput(name)
	return getFinalEnergy(plants,[]int{})
}

func part2(name string) int {
	plants, testCases,_:= parseInput(name)
	total:=0
	for _, testCase:= range testCases {
		total+= getFinalEnergy(plants, testCase)
	}

	return total
}

type state struct {
	testCase []int
	energy int
}


func part3(name string) int {
	plants, testCases, maxID:= parseInput(name)
	energies:= []state{}
	for _, testCase:= range testCases {
		energy := getFinalEnergy(plants, testCase)
		energies = append(energies, state{
			testCase: testCase,
			energy: energy,
		})
	}

	slices.SortFunc(energies, func(a,b state) int {
		return b.energy - a.energy
	})

	max:= getMaxEnergy(plants,energies[0].testCase, maxID)

	total:=0

	for _,energy:= range energies {
		if energy.energy == 0 {
			continue
		}

		total += max - energy.energy
	}

	return total
}

type plant struct {
	id int
	branchedFrom []branch
	thickness int
}

type branch struct {
	source int
	thickness int
}

func parseInput(name string) ([]plant, [][]int, int) {
	output:= []plant{}
	testCases:= [][]int{}
	input:= files.ReadParagraphs(name)
	maxID:=0

	words:= []string{
		"Plant",
		"with",
		"thickness",
		"-",
		"free",
		"branch",
		"to",
	}

	for _, plantInput:= range input {
		if len(plantInput) == 0 {
			continue
		}

		if plantInput[0][0] != 'P' {
			for _,row:= range plantInput {
				testCase:= []int{}
				split:= strings.Split(row," ")

				for j,str:= range split {
					if str == "0" {
						testCase = append(testCase, j+1)
					}
				}
				testCases = append(testCases, testCase)
			}
			continue
		}

		cleanStr:= strings.ReplaceAll(plantInput[0],":","")
		split := strings.Split(cleanStr," ")
		vals:= excludeStrings(split,words)
		branches:= parseBranches(plantInput[1:],words)

		output = append(output, plant{
			id: vals[0],
			thickness: vals[1],
			branchedFrom: branches,
		})

		if len(branches) == 1 && branches[0].source==0 && vals[0] > maxID {
			maxID = vals[0]
		}
	}

	return output,testCases,maxID
}

func excludeStrings(src, exclusions []string) []int {
	output:= []int{}

	for _, str:= range src {
		if !slices.Contains(exclusions,str) {
			output = append(output, ints.FromString(str))
		}
	}

	return output
}

func parseBranches(input []string, words []string) []branch {
	output:= []branch{}

	for _,str:= range input {
		split := strings.Split(str," ")
		vals:= excludeStrings(split,words)

		if len(vals) == 1 {
			output = append(output, branch{
				thickness:vals[0],
			})
		} else {
			output = append(output, branch{
				thickness: vals[1],
				source: vals[0],
			})
		}
	}

	return output
}

func getPlantEnergy(plant plant, mappedPlants map[int]int) (int,bool) {
	total:= 0

	for _, branch := range plant.branchedFrom {
		energy,ok:= mappedPlants[branch.source]

		if !ok {
			return 0,false
		}

		total += energy*branch.thickness
	}

	if total < plant.thickness {
		return 0, true
	}

	return total, true
}

func getFinalEnergy(plants []plant, excludedPlants []int) int {
	mappedPlants:= map[int]int{
		0:1,
	}

	for _,p:= range excludedPlants {
		mappedPlants[p] = 0
	}

	for {
		newMapping:= false
		for _, plant:= range plants {
			if _,ok:= mappedPlants[plant.id]; ok {
				continue
			}

			energy,ok := getPlantEnergy(plant, mappedPlants)

			if ok {
				mappedPlants[plant.id] = energy
				newMapping = true
			}
		}

		if !newMapping {
			break
		}
	}	

	return mappedPlants[plants[len(plants)-1].id]
}

func getMaxEnergy(plants []plant, seed []int, maxId int) int {
	max:= getFinalEnergy(plants, seed)

	for {
		maxIncreased := false
		var maxSeed []int
		
		for i:= range maxId {
			newSeed:= applySeedChange(seed,i+1)
			energy:= getFinalEnergy(plants,newSeed)

			if energy > max {
				max = energy
				maxIncreased = true
				maxSeed = newSeed
			}
		}

		if !maxIncreased {
			return max
		}

		seed = maxSeed
	}
}

func applySeedChange(seed []int, id int) []int {
	if slices.Contains(seed,id) {
		newSeed:= []int{}

		for _,i:= range seed {
			if i == id {
				continue
			}

			newSeed = append(newSeed, i)
		}

		return newSeed
	}

	return append(seed,id)
}