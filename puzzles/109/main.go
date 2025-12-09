package main

import (
	"fmt"
	"math"
	"strings"
	"sync"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/ints"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1 answer: ", part1("input.txt"))
	split := time.Now()
	fmt.Println("Part 2 answer: ", part2("input.txt"))
	
	fmt.Println()
	fmt.Println("Part 1: ", split.Sub(start))
	fmt.Println("Part 2: ", time.Since(split))
}

func part1(name string) int {
	coords,_,_ := parseCoords(name)

	max:=0

	for i:=0;i<len(coords); i++ {
		for j:=i+1; j<len(coords); j++ {
			area:= (ints.Abs(coords[i][0] - coords[j][0])+1) *
				 	(ints.Abs(coords[i][1] - coords[j][1])+1) 
			
			if area > max {
				max = area
			}
		}
	}

	return max
}

func part2(name string) int {
	coords, _, _:= parseCoords(name)

	grid:= make(map[[2]int]bool)
	
	lastCoord:=coords[len(coords)-1]
	grid[lastCoord] = true
	for i:=0; i<len(coords); i++ {
		grid[coords[i]] = true

		if coords[i][0] != lastCoord[0] {
			yFill:= ints.GetIntsBetween(coords[i][0], lastCoord[0])
			for _,y:= range yFill {
				grid[[2]int{y,lastCoord[1]}] = true
			}
		} else {
			xFill:= ints.GetIntsBetween(coords[i][1], lastCoord[1])
			for _,x:= range xFill {
				grid[[2]int{lastCoord[0],x}] = true
			}
		}

		lastCoord = coords[i]
	}

	IsValid := func (i,j int) bool {
		lastCoord = coords[len(coords)-1]
		for k:=0; k<len(coords); k++ {
			if rectIntersection(coords[i],coords[j],coords[k], lastCoord) {
				return false
			}
			lastCoord = coords[k]
		}

		count:=0
		midY:= (coords[i][0]+coords[j][0])/2
		midX:= (coords[i][1]+coords[j][1])/2
		if coords[i][0] != coords[j][0] {
			for k:=0; k<=midY; k++ {
				if grid[[2]int{k,midX}] {
					count ++
				}
			}
		} else {
			for k:=0; k<=midX; k++ {
				if grid[[2]int{midY,k}] {
					count ++
				}
			}
		}
		if count %2 == 0 {
			return false
		}		

		return true
	}

	channel := make(chan int)

	var wg sync.WaitGroup

	getArea:= func (i,j int) {
		defer wg.Done()
		if IsValid(i,j) {
			channel <- (ints.Abs(coords[i][0] - coords[j][0])+1) *
				(ints.Abs(coords[i][1] - coords[j][1])+1)
		}
	}

	for i:=0;i<len(coords); i++ {
		for j:=i+1; j<len(coords); j++ {
			wg.Add(1)
			go getArea(i,j)
		}
	}

	// goroutine that closes channel after work is done
	go func() {
		wg.Wait()
		close(channel)
	}()

	return MaxFromChan(channel)
}

func parseCoords(name string) ([][2]int, int, int) {
	lines:=files.ReadLines(name)
	coords:= [][2]int{}
	maxX,maxY:=0,0

	for _,line:= range lines {
		sliced:= strings.Split(line, ",")
		vals:= ints.FromStringSlice(sliced)
		if vals[0] > maxY {
			maxY = vals[0]
		}
		if vals[1] > maxX {
			maxX = vals[1]
		}
		coords = append(coords, [2]int{vals[0],vals[1]})
	}

	return coords, maxY, maxX
}

func rectIntersection(a,b,c,d [2]int) bool {
	topAB, bottomAB:= ints.MinMax([]int{a[0],b[0]})
	topCD, bottomCD:= ints.MinMax([]int{c[0],d[0]})
	if topAB >=bottomCD || topCD >= bottomAB {
		return false
	}

	leftAB, rightAB:= ints.MinMax([]int{a[1],b[1]})
	leftCD, rightCD:= ints.MinMax([]int{c[1],d[1]})
	if leftAB >= rightCD || leftCD >= rightAB {
		return false
	}

	return true
}

func MaxFromChan(ch <-chan int) int {
    max := math.MinInt

    for v := range ch {
        if v > max {
            max = v
        }
    }
    return max
}