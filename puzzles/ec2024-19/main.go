package main

import (
	"fmt"
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

func part1(name string) string {
	maze, key := parseInput(name)

	width := len(maze[0])
	height := len(maze)

	index := 0

	for i, row := range maze {
		for j := range row {
			pos := coords.NewCoord(i, j)
			if len(getNeighbours(pos, width, height)) != 8 {
				continue
			}

			if key[index] == 'L' {
				maze = rotateLeft(maze, pos)
			} else {
				maze = rotateRight(maze, pos)
			}
			index = (index + 1) % len(key)
		}
	}

	return getDecodedString(maze)
}

func part2(name string) string {
	maze, key := parseInput(name)

	rounds:=100

	width := len(maze[0])
	height := len(maze)

	changeMap := getRoundChangeMap(height, width, key)

	doubleUp(changeMap, rounds)

	finalMaze:= mapToFinalPosition(maze,changeMap,rounds)

	return getDecodedString(finalMaze)
}

func part3(name string) string {
	maze, key := parseInput(name)
	rounds := 1048576000

	width := len(maze[0])
	height := len(maze)

	changeMap := getRoundChangeMap(height, width, key)

	doubleUp(changeMap, rounds)

	finalMaze:= mapToFinalPosition(maze,changeMap,rounds)

	printMaze(finalMaze)

	return getDecodedString(finalMaze)
}

func rotateRight[T any](maze [][]T, locus coords.Coord) [][]T {
	neighbours := getNeighbours(locus, len(maze[0]), len(maze))
	prevPos := neighbours[len(neighbours)-1]
	prevChar := maze[prevPos.I][prevPos.J]

	for _, pos := range neighbours {
		nextChar := maze[pos.I][pos.J]
		maze[pos.I][pos.J] = prevChar
		prevChar = nextChar
	}

	return maze
}

func rotateLeft[T any](maze [][]T, locus coords.Coord) [][]T {
	neighbours := getNeighbours(locus, len(maze[0]), len(maze))
	prevPos := neighbours[0]
	prevChar := maze[prevPos.I][prevPos.J]

	for i := len(neighbours) - 1; i >= 0; i-- {
		pos := neighbours[i]
		nextChar := maze[pos.I][pos.J]
		maze[pos.I][pos.J] = prevChar
		prevChar = nextChar
	}

	return maze
}

func getNeighbours(locus coords.Coord, width, height int) []coords.Coord {
	adjacents := locus.GetAdjacentIncludingDiagonals()

	var toMove []coords.Coord

	for _, adj := range adjacents {
		if adj.I < 0 || adj.J < 0 || adj.I >= height || adj.J >= width {
			continue
		}
		toMove = append(toMove, adj)
	}

	return toMove
}

func parseInput(name string) ([][]rune, []rune) {
	input := files.ReadLinesAsRunes(name)

	key := input[0]

	maze := input[2:]

	return maze, key
}

func printMaze(maze [][]rune) {
	for _, row := range maze {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func getDecodedString(maze [][]rune) string {
	answer := ""

	for _, row := range maze {
		startFound := false
		for _, char := range row {
			if !startFound {
				if char == '>' {
					startFound = true
				}
				continue
			}

			if char == '<' {
				return answer
			}

			answer += string(char)
		}
	}
	printMaze(maze)
	panic("output incorrect")
}

type state struct {
	position coords.Coord
	rounds   int
}

type stateToAdd struct {
	state
	newPosition coords.Coord
}

func getRoundChangeMap(height, width int, key []rune) map[state]coords.Coord {
	answer := make(map[state]coords.Coord)

	coordMaze := make([][]coords.Coord, height)

	for i := range height {
		coordMaze[i] = make([]coords.Coord, width)
		for j := range width {
			coordMaze[i][j] = coords.NewCoord(i, j)
		}
	}

	index := 0
	for i := range height {
		for j := range width {
			pos := coords.NewCoord(i, j)
			if len(getNeighbours(pos, width, height)) != 8 {
				continue
			}

			if key[index] == 'L' {
				coordMaze = rotateLeft(coordMaze, pos)
			} else {
				coordMaze = rotateRight(coordMaze, pos)
			}
			index = (index + 1) % len(key)
		}
	}

	for i := range height {
		for j := range width {
			answer[state{position: coordMaze[i][j], rounds: 1}] = coords.NewCoord(i, j)
		}
	}

	return answer
}

func doubleUp(changeMap map[state]coords.Coord, rounds int) {
	statesToCheck := []stateToAdd{}
	power:=1

	for k, v := range changeMap {
		statesToCheck = append(statesToCheck, stateToAdd{
			state:       k,
			newPosition: v,
		})
	}

	for rounds%2 ==0 {
		toAdd := []stateToAdd{}
		rounds /=2
		for _, stateToCheck := range statesToCheck {
			toAdd = append(toAdd, stateToAdd{
				state: state{
					position: stateToCheck.state.position,
					rounds:   power * 2,
				},
				newPosition: changeMap[state{position: stateToCheck.newPosition, rounds: power}],
			})
		}

		for _, newState := range toAdd {
			changeMap[newState.state] = newState.newPosition
		}
		power *=2
		statesToCheck = toAdd
	}

	for rounds%5 == 0 {
		toAdd := []stateToAdd{}
		rounds /=5
		for _, stateToCheck := range statesToCheck {
			position := coords.NewCoord(stateToCheck.state.position.I,stateToCheck.state.position.J)
			for range 5 {
				position = changeMap[state{
					position: position,
					rounds: power,
				}]
			}
			toAdd = append(toAdd, stateToAdd{
				state: state{
					position: stateToCheck.state.position,
					rounds: power * 5,
				},
				newPosition: position,
			})
		}

		for _, newState := range toAdd {
			changeMap[newState.state] = newState.newPosition
		}
		power *=5
		statesToCheck = toAdd
	}
}

func mapToFinalPosition(maze [][]rune, changeMap map[state]coords.Coord, rounds int) [][]rune {
	width := len(maze[0])
	height := len(maze)

	finalMaze:= make([][]rune,height)
	for i:=range height {
		finalMaze[i] = make([]rune, width)
	}

	for i:=range height {
		for j:= range width {
			position := changeMap[state{position: coords.NewCoord(i,j), rounds: rounds}]

			finalMaze[position.I][position.J] = maze[i][j]
		}
	}

	return finalMaze
}