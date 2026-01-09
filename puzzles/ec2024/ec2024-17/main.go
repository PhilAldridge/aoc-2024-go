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
	stars:= getStars(name)
	constellations:= getConstellations(stars, math.MaxInt)

	return constellations[0].Score()
}

func part2(name string) int {
	return part1(name)
}


func part3(name string) int {
	stars:= getStars(name)
	constellations:= getConstellations(stars, 5)

	slices.SortFunc(constellations, func(a,b constellationType) int {
		return b.Score() - a.Score()
	})

	return constellations[0].Score() * constellations[1].Score() * constellations[2].Score()
}

func getStars(name string) []coords.Coord {
	lines:= files.ReadLines(name)
	var stars []coords.Coord
	for i,row:= range lines{
		for j,char:= range row {
			if char == '*' {
				stars = append(stars, coords.NewCoord(i,j))
			}
		}
	}

	return stars
}

func getDistances(stars []coords.Coord) []distance {
	var distances []distance
	for i:= range stars {
		for j:=i+1; j<len(stars); j++ {
			distances = append(distances, distance{
				a: stars[i],
				b: stars[j],
				distance: coords.ManhattanDistance(stars[i],stars[j]),
			})
		}
	}

	return distances
}

func getConstellations(stars []coords.Coord, maxDistance int) []constellationType {
	distances:= getDistances(stars)
	slices.SortFunc(distances, func(a,b distance) int {
		return a.distance - b.distance
	})

	var constellations []constellationType
	for _,star:= range stars {
		constellations = append(constellations, constellationType{stars: []coords.Coord{star}})
	}

	for len(constellations)>1 {
		nextConnection:= distances[0]
		distances = distances[1:]

		if nextConnection.distance > maxDistance {
			break
		}

		aIndex := slices.IndexFunc(constellations,func(constellation constellationType) bool {
			return slices.ContainsFunc(constellation.stars, func(star coords.Coord) bool {
				return star.Equals(nextConnection.a)
			})
		})

		bIndex := slices.IndexFunc(constellations,func(constellation constellationType) bool {
			return slices.ContainsFunc(constellation.stars, func(star coords.Coord) bool {
				return star.Equals(nextConnection.b)
			})
		})

		if aIndex == -1 || bIndex == -1 {
			panic("star lost!")
		}

		if aIndex == bIndex {
			continue
		}

		newConstellations := []constellationType{{
			stars: append(constellations[aIndex].stars,constellations[bIndex].stars...),
			totalDistances: constellations[aIndex].totalDistances + constellations[bIndex].totalDistances + nextConnection.distance,
		}}

		for i:= range constellations {
			if i == aIndex || i ==bIndex {
				continue
			}

			newConstellations = append(newConstellations, constellations[i])
		}

		constellations = newConstellations
	}

	return constellations
}

type distance struct {
	a,b coords.Coord
	distance int
}

type constellationType struct {
	stars []coords.Coord
	totalDistances int
}

func (c constellationType) Score() int {
	return c.totalDistances + len(c.stars)
}