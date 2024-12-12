package main

import (
	"fmt"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/bools"
	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1: ", part1("input.txt"), " in:", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", part2("input.txt"), " in:", time.Since(start))
}

type plot struct {
	perimeter int
	area int
}

func part1(name string) int {
	garden := files.ReadLines(name)
	plotMap:= make(map[[2]int]int)
	plots:= []plot{}
	for i, row := range garden {
		for j, _ := range row {
			pos := [2]int{i,j}
			//If not in a plot already, create a new plot and add all tiles to it
			_, plotInMap := plotMap[pos]
			if !plotInMap {
				plotIndex := len(plots)
				plots = append(plots, plot{
					area:0,
					perimeter: 0,
				})
				
				//Recursive function that spreads through adjacent tiles to find all tiles in same plot as the current tile
				//Adds each tile to the plotMap
				updateMap(plotMap, plotIndex, i,j, garden)
			} 

			//Each tile in plot counts as 1 area
			plots[plotMap[pos]].area += 1
			//Function checks around current tile in each of the four directions
			//On each side, if there is a border, add 1 to perimeter
			plots[plotMap[pos]].perimeter += getPerimeter(i,j, garden)
		}
	}

	//Sum up area*perimeter
	return getTotal(plots)
}

func part2(name string) int {
	garden := files.ReadLines(name)
	plotMap:= make(map[[2]int]int)
	plots:= []plot{}
	for i, row := range garden {
		for j, _ := range row {
			pos := [2]int{i,j}
			//If not in a plot already, create a new plot and add all tiles to it
			_, plotInMap := plotMap[pos]
			if !plotInMap {
				plotIndex := len(plots)
				plots = append(plots, plot{
					area:0,
					perimeter: 0,
				})
				
				//Recursive function that spreads through adjacent tiles to find all tiles in same plot as the current tile
				//Adds each tile to the plotMap
				updateMap(plotMap, plotIndex, i,j, garden)
			} 
			
			//Each tile in plot counts as 1 area
			plots[plotMap[pos]].area += 1
			//Function checks around current tile in each corner
			//Number of corners in polygon = number of sides
			plots[plotMap[pos]].perimeter += getCorners(i,j, garden)
		}
	}

	//Sum up area*perimeter
	return getTotal(plots)
}

func updateMap(plotMap map[[2]int]int, plotIndex int, i int, j int, garden []string) {
	_,okIMinus := plotMap[[2]int{i-1,j}]
	_,okIPlus := plotMap[[2]int{i+1,j}]
	_,okJMinus := plotMap[[2]int{i,j-1}]
	_,okJPlus := plotMap[[2]int{i,j+1}]
	plotMap[[2]int{i,j}] = plotIndex
	//If adjacent matching tile that is not already added to a plot, add it and check around it
	if i>0 && garden[i-1][j]==garden[i][j] && !okIMinus {
		updateMap(plotMap,plotIndex,i-1,j,garden)
	}
	if i<len(garden)-1 && garden[i+1][j]==garden[i][j] && !okIPlus {
		updateMap(plotMap,plotIndex,i+1,j,garden)
	}
	if j>0 && garden[i][j-1]==garden[i][j] && !okJMinus{
		updateMap(plotMap,plotIndex,i,j-1,garden)
	}
	if j<len(garden[0])-1 && garden[i][j+1]==garden[i][j] && !okJPlus {
		updateMap(plotMap,plotIndex,i,j+1,garden)
	}
}

func getPerimeter(i int, j int, garden []string) int {
	current := garden[i][j]
	//If on edge or bordered by different plot, add 1 to perimeter
	//Repeat for each side
	return bools.CountTrues(
		i==0 || current != garden[i-1][j],
		i == len(garden)-1 || current != garden[i+1][j],
		j == 0 || current != garden[i][j-1],
		j == len(garden[0])-1 || current != garden[i][j+1],
	)
}

func getCorners(i int, j int, garden []string) int {
	corners:=0
	current := garden[i][j]
	//Convex corners = Adjacent tile on two side do not match current tile            
	if (i==0 || current != garden[i-1][j]) {
		corners += bools.CountTrues(
			j == 0 || current != garden[i][j-1],
			j == len(garden[0])-1 || current != garden[i][j+1],
		)
	} 
	if i == len(garden)-1 || current != garden[i+1][j] {
		corners += bools.CountTrues(
			j == 0 || current != garden[i][j-1],
			j == len(garden[0])-1 || current != garden[i][j+1],
		)
	}
	
	//Concave corners = Adjacent tile of two sides match, but diagonal tile does not match
	if i>0 && current == garden[i-1][j] {
		corners += bools.CountTrues(
			j>0 && current == garden[i][j-1] && current != garden[i-1][j-1],
			j<len(garden[0])-1 && current == garden[i][j+1] && current != garden[i-1][j+1],
		)
	}
	if i<len(garden)-1 && current == garden[i+1][j] {
		corners += bools.CountTrues(
			j>0 && current == garden[i][j-1] && current != garden[i+1][j-1],
			j<len(garden[0])-1 && current == garden[i][j+1] && current != garden[i+1][j+1],
		)
	}
	
	return corners
}

func getTotal(plots []plot) int {
	total:= 0
	for _,plot:=range plots {
		total += plot.area*plot.perimeter
	}
	return total
}

