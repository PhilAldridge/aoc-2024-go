package main

import (
	"fmt"
	"slices"
	"sync"
	"time"

	"github.com/kylehoehns/aoc-2023-go/pkg/files"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1: ", part1("input.txt"), " in:", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", part2("input.txt"), " in:", time.Since(start))
}

func part1(name string) int {
	lines := files.ReadLines(name)
	currentLocation, direction := getStart(lines)
	locationsVisited := [][2]int{currentLocation}
	for {
		//x,y in front of guard
		newY := currentLocation[0] + getDirectionVector(direction)[0]
		newX := currentLocation[1] + getDirectionVector(direction)[1]
		if newY < 0 || newY >= len(lines) || newX < 0 || newX >= len(lines[0]) {
			//gone out of bounds = finished
			break
		}
		switch lines[newY][newX] {
		case '#':
			//obstacle - turn right - directions in getDirectionVector are in clockwise order
			direction = (direction + 1) % 4
		default:
			//no obstacle - move forward
			currentLocation = [2]int{newY, newX}
		}
		if !slices.Contains(locationsVisited, currentLocation) {
			locationsVisited = append(locationsVisited, currentLocation)
		}
	}
	return len(locationsVisited)
}

type safeList struct {
	mu   sync.Mutex
	list [][2]int
}

func part2(name string) int {
	lines := files.ReadLines(name)
	currentLocation, direction := getStart(lines)
	obstaclesAdded := safeList{}
	locationsVisited := [][2]int{currentLocation}
	wg := &sync.WaitGroup{}
	for {
		//x,y in front of guard
		newY := currentLocation[0] + getDirectionVector(direction)[0]
		newX := currentLocation[1] + getDirectionVector(direction)[1]
		if newY < 0 || newY >= len(lines) || newX < 0 || newX >= len(lines[0]) {
			//gone out of bounds = finished
			break
		}
		switch lines[newY][newX] {
		case '#':
			//obstacle - turn right - directions in getDirectionVector are in clockwise order
			direction = (direction + 1) % 4
		default:
			//no obstacle - check if putting an obstacle in front creates a loop
			//then move forward
			inFront := [2]int{newY, newX}
			if !slices.Contains(locationsVisited, inFront) {
				//go routine not necessary. Without this, ran in 3s.
				//added only as practice. New time 0.5s
				wg.Add(1)
				go checkLoop(lines, currentLocation, direction, inFront, &obstaclesAdded, wg)
			}
			currentLocation = [2]int{newY, newX}

			if !slices.Contains(locationsVisited, currentLocation) {
				locationsVisited = append(locationsVisited, currentLocation)
			}
		}
	}
	wg.Wait()
	return len(obstaclesAdded.list)
}

func getStart(lines []string) ([2]int, int) {
	for i, line := range lines {
		for j, char := range line {
			if char == '.' || char == '#' {
				continue
			}
			direction := 0
			switch char {
			case '>':
				direction = 3
			case 'v':
				direction = 0
			case '<':
				direction = 1
			case '^':
				direction = 2
			}
			return [2]int{i, j}, direction
		}
	}
	panic("start not found")
}

func getDirectionVector(directionInt int) [2]int {
	directions := [4][2]int{
		//clockwise, so that adding 1 to directionInt = turning right
		{1, 0},  //down
		{0, -1}, //left
		{-1, 0}, //up
		{0, 1},  //right
	}
	return directions[directionInt]
}

func checkLoop(lines []string, currentLocation [2]int, direction int, obstacleAdded [2]int, obstaclesAdded *safeList, wg *sync.WaitGroup) {
	defer wg.Done()
	//list of all locations visited in the loop check and which direction they were facing
	visited := [][3]int{{currentLocation[0], currentLocation[1], direction}}
	for {
		//x,y in front of guard
		newY := currentLocation[0] + getDirectionVector(direction)[0]
		newX := currentLocation[1] + getDirectionVector(direction)[1]
		if newY < 0 || newY >= len(lines) || newX < 0 || newX >= len(lines[0]) {
			//out of bounds means the guard escapes - no loop
			return
		}
		if slices.Contains(visited, [3]int{newY, newX, direction}) {
			//we have got to the same place facing the same direction as already visited - loop
			obstaclesAdded.mu.Lock()
			defer obstaclesAdded.mu.Unlock()
			obstaclesAdded.list = append(obstaclesAdded.list, [2]int{newY, newX})
			return
		}
		visited = append(visited, [3]int{currentLocation[0], currentLocation[1], direction})
		if obstacleAdded[0] == newY && obstacleAdded[1] == newX || lines[newY][newX] == '#' {
			//Old obstacle or addedobstacle in front = turn right
			direction = (direction + 1) % 4
		} else {
			//No obstacle = move forward
			currentLocation = [2]int{newY, newX}
		}

	}
}
