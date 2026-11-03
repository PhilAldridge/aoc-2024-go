package main

import (
	"fmt"
	"math"
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
	total := 0

	lines := files.ReadLines(name)

	for _, line := range lines {
		vals := ints.FromStringSlice(strings.Split(line, ","))

		visited := []int{0}

		for _, val := range vals {
			current := visited[len(visited)-1]
			if val < current && !slices.Contains(visited, current-val) {
				visited = append(visited, current-val)
			} else {
				visited = append(visited, current+val)
			}
		}

		total += visited[len(visited)-1]
	}

	return total
}

func part2(name string) int {
	total := 0

	lines := files.ReadLines(name)

	for _, line := range lines {
		vals := ints.FromStringSlice(strings.Split(line, ","))

		visited := []int{0}

		for _, val := range vals {
			current := visited[len(visited)-1]
			if val < current && !slices.Contains(visited, current-val) {
				visited = append(visited, current-val)
			} else {
				new := current + val
				for slices.Contains(visited, new) {
					new++
				}
				visited = append(visited, new)
			}
		}

		total += visited[len(visited)-1]
	}

	return total
}

func part3(name string) int {
	total := 0

	lines := files.ReadLines(name)

	for i, line := range lines {
		vals := ints.FromStringSlice(strings.Split(line, ","))

		visited := []int{0}

		for _, val := range vals {
			current := visited[len(visited)-1]

			new := current - val

			if new <= 0 || slices.Contains(visited, new) || crossesLoop(new, current, visited) {
				// try forward
				new = current + val

				for new <= insideLoopMax(current, visited) {
					if slices.Contains(visited, new) || crossesLoop(new, current, visited) {
						new++
						continue
					}

					visited = append(visited, new)
					break
				}

				continue
			}

			// add backward move
			visited = append(visited, new)
		}

		total += visited[len(visited)-1]

		// For displaying in Desmos
		if i == len(lines)-1 {
			// 	for j:=0; j< len(visited) - 1; j++ {
			// 		centre:= float32(visited[j] + visited[j+1])/2.0

			// 		r:= float32(visited[j]) - centre

			// 		constraint:= ` \left\{ y > 0 \right\}`
			// 		if j%2 ==0 {
			// 			constraint = ` \left\{ y < 0 \right\}`
			// 		}

			// 		fmt.Printf("(x-%f)^2+y^2=%f%s\n", centre, r*r, constraint)
			// 	}
		}
	}

	return total
}

func crossesLoop(a, b int, visited []int) bool {
	if a > b {
		a, b = b, a
	}

	startAt := 0

	if len(visited)%2 == 0 {
		startAt = 1
	}

	for startAt < len(visited)-1 {
		loopStart := visited[startAt]
		loopEnd := visited[startAt+1]

		if loopStart > loopEnd {
			loopStart, loopEnd = loopEnd, loopStart
		}

		if a < loopStart && b > loopStart && b < loopEnd {
			return true
		}

		if a < loopEnd && b > loopEnd && a > loopStart {
			return true
		}

		startAt += 2
	}

	return false
}

func insideLoopMax(a int, visited []int) int {
	max := math.MaxInt

	startAt := 0

	if len(visited)%2 == 0 {
		startAt = 1
	}

	for startAt < len(visited)-1 {
		loopStart := visited[startAt]
		loopEnd := visited[startAt+1]

		if loopStart > loopEnd {
			loopStart, loopEnd = loopEnd, loopStart
		}

		if loopStart < a && a < loopEnd && loopEnd-1 < max {
			max = loopEnd - 1
		}

		startAt += 2
	}

	return max
}
