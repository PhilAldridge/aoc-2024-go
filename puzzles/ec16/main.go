package main

import (
	"fmt"
	"math"
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
	spell:= ints.FromStringSlice(strings.Split(files.Read(name),","))
	
	return countBlocks(spell,90)
}

func part2(name string) int {
	columns:= ints.FromStringSlice(strings.Split(files.Read(name),","))
	total:=1

	spell:=getSpell(columns)

	for _,val:= range spell {
		total*=val
	}

	return total
}


func part3(name string) int {
	columns:= ints.FromStringSlice(strings.Split(files.Read(name),","))

	totalBlocks:= 202520252025000

	spell:=getSpell(columns)

	min:=0
	max:=math.MaxInt
	

	for min + 1< max {
		mid:= (max+min)/2

		blocks:= countBlocks(spell,mid)
		
		if totalBlocks == blocks {
			return mid
		}

		if totalBlocks > blocks {
			min = mid
		} else {
			max = mid
		}
	}

	return min
}

func getSpell(columns []int) []int {
	spell:=[]int{}

	for i,column:= range columns {
		count:=0

		for _,val:= range spell {
			if (i+1)%val ==0 {
				count++
			}
		}

		if column-count>0 {
			spell = append(spell, i+1)
		}
	}

	return spell
}

func countBlocks(spell []int, columnCount int) int {
	total:= 0

	for _,val:= range spell {
		total+=columnCount/val
	}

	return total	
}