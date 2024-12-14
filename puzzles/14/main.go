package main

import (
	"fmt"
	"regexp"
	"sort"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/coords"
	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/ints"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1: ", part1("input.txt"), " in:", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", part2("input.txt"), " in:", time.Since(start))
}

func part1(name string) int {
	height:= 103
	width:= 101
	robotLines:= getRobotLines(name)
	finalPositions:= getPositions(robotLines,100,width,height)
	return multiplyQuadrants(finalPositions, width,height)
}

func part2(name string) int {
	height:= 103
	width:= 101
	robotLines:= getRobotLines(name)
	distances := make(map[int]int)
	//For times up to 10000, find the total (manhattan) distance between all the robots. 
	//In the position of the tree, the points will be closer together than usual
	for i:=0; i<=10000; i++ {
		distances[i] = getTotalDistance(getPositions(robotLines,i,width,height))
	}
	distancesOrdered := make([]int, 0, len(distances))
	for d := range distances {
        distancesOrdered = append(distancesOrdered, d)
    }
	sort.SliceStable(distancesOrdered, func(i, j int) bool{
        return distances[distancesOrdered[i]] < distances[distancesOrdered[j]]
    })
	//Print position where total distance is shortest to check that it forms a tree
	printPositions(robotLines,distancesOrdered[0],width,height)
	return 0
}

func getRobotLines(name string) []coords.Line {
	lines:= []coords.Line{}
	lineStr := files.ReadLines(name)
	findInts := regexp.MustCompile(`(-?[0-9])+`)
	for _,str:= range lineStr {
		lineInts := ints.FromStringSlice(findInts.FindAllString(str,-1))
		m:= coords.NewCoord(lineInts[3],lineInts[2])
		c:= coords.NewCoord(lineInts[1],lineInts[0])
		lines = append(lines, coords.NewLine(m,c))
	}
	return lines
}

func multiplyQuadrants(positions []coords.Coord, width int, height int) int {
	quadrants:= [4]int{0,0,0,0}
	for _, position:= range positions {
		i:= position.I
		j:= position.J

		middleH:= (height-1)/2
		middleW:= (width-1)/2
		if i< middleH {
			if j< middleW {
				quadrants[0] ++
			} else if j> middleW {
				quadrants[1] ++
			}
		} else if i>middleH {
			if j< middleW {
				quadrants[2] ++
			} else if j> middleW {
				quadrants[3] ++
			}
		}
	}
	return quadrants[0]*quadrants[1]*quadrants[2]*quadrants[3]
}

func getPositions(robotLines []coords.Line, time int, width int, height int) []coords.Coord {
	positions:= []coords.Coord{}
	for _, robot:= range robotLines {
		pos:= robot.C.Add(robot.M.Multiply(time))
		pos.I %= height
		pos.J %= width
		if pos.I<0 {pos.I += height}
		if pos.J<0 {pos.J+= width}
		positions = append(positions, pos)
	}
	return positions
}

func getTotalDistance(positions []coords.Coord) int {
	total:=0
	for i,pos1 := range positions {
		for j,pos2:= range positions {
			if j<=i {continue}
			iDist:= pos1.I-pos2.I
			if iDist<0 {iDist *= -1}
			
			jDist:= pos1.J-pos2.J
			if jDist<0 {jDist *= -1}
			total += iDist + jDist
		}
	}
	return total
}

func printPositions(robotLines []coords.Line,time int, width int, height int) {
	finalPositions:= getPositions(robotLines,time,width,height)
	for i:=0; i< height; i++ {
		fmt.Println(" ")
		for j:=0; j< width; j++{
			robotHere:=false
			for _,pos:= range finalPositions {
				if pos.I == i && pos.J == j {
					fmt.Print("*")
					robotHere = true
					break
				}
			}
			if !robotHere {
				fmt.Print(" ")
			}
		}
	}
	fmt.Print(time)
	fmt.Println(" ")

}