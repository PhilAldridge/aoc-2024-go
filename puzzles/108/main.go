package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/ints"
	"github.com/PhilAldridge/aoc-2024-go/pkg/sets"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1 answer: ", part1("input.txt",1000))
	split := time.Now()
	fmt.Println("Part 2 answer: ", part2("input.txt"))
	
	fmt.Println()
	fmt.Println("Part 1: ", split.Sub(start))
	fmt.Println("Part 2: ", time.Since(split))
}

func part1(name string, iterations int) int {
	lines:= files.ReadLines(name)
	distanceMap, distancesSlice:= getDistances(lines)

	connected:= []sets.Set[[3]int]{}

	connectionsCount:=0
	for connectionsCount<iterations {
		distance:= distancesSlice[connectionsCount]
		connected = makeShortestConnection(connected, distanceMap[distance])
		connectionsCount++
	}

	lengths:= []int{}

	for _,group:= range connected {
		lengths = append(lengths, group.Size())
	}

	slices.Sort(lengths)

	return lengths[len(lengths)-1]*lengths[len(lengths)-2]*lengths[len(lengths)-3]
}

func part2(name string) int {
	lines:= files.ReadLines(name)
	
	distanceMap, distancesSlice:= getDistances(lines)

	connected:= []sets.Set[[3]int]{}

	connectionsCount:=0
	for len(connected)!=1 || connected[0].Size()!=len(lines) {
		distance:= distancesSlice[connectionsCount]
		connected = makeShortestConnection(connected, distanceMap[distance])
		connectionsCount++
	}

	lastCoords:= distanceMap[distancesSlice[connectionsCount-1]]

	return lastCoords[0][0] * lastCoords[1][0]
}

func calculateDistance(a [3]int, b [3]int) float64 {
	x:= a[0]-b[0]
	y:= a[1]-b[1]
	z:= a[2]-b[2]

	return float64(math.Sqrt(float64(x*x+y*y+z*z)))
}

func makeShortestConnection(connected []sets.Set[[3]int], coords [2][3]int) ([]sets.Set[[3]int]) {
	foundIndex:= []int{}
	newConnected:= []sets.Set[[3]int]{}

	for i, group:= range connected {
		foundOne:=false
		if group.Contains(coords[0]) {
			foundIndex = append(foundIndex, i)
			foundOne=true
		}

		if group.Contains(coords[1]) {
			if foundOne {
				return connected
			}
			foundIndex = append(foundIndex, i)
		}

		if len(foundIndex)==2 {
			break
		}
	}

	newGroup := sets.NewSet[[3]int]()
	for _,i:= range foundIndex {
		for _,coord := range connected[i].List() {
			newGroup.Add(coord)
		}
	}
	newGroup.Add(coords[0])
	newGroup.Add(coords[1])

	newConnected = append(newConnected,*newGroup)


	for i,group:= range connected {
		if slices.Contains(foundIndex,i) {
			continue
		}
		newConnected = append(newConnected, group)
	}
	
	return newConnected
}

func getDistances(lines []string) (map[float64][2][3]int, []float64) {
	coords:= [][3]int{}
	for _,line:= range lines {
		ints:= ints.FromStringSlice(strings.Split(line,","))
		coords = append(coords, [3]int(ints[0:3]))
	}

	distanceMap:= make(map[float64][2][3]int)
	ditancesSlice := []float64{}

	for i,coord:= range coords {
		for j:= i+1; j<len(coords); j++ {
			distance:= calculateDistance(coord, coords[j])
			ditancesSlice = append(ditancesSlice, distance)
			distanceMap[distance] = [2][3]int{coord,coords[j]}
		}
	}

	slices.Sort(ditancesSlice)

	return distanceMap,ditancesSlice
}