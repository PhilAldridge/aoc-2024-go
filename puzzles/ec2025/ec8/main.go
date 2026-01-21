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
	fmt.Println("Part 1 answer: ", part1("input.txt",32))
	split := time.Now()
	fmt.Println("Part 2 answer: ", part2("input2.txt",256))
	split2 := time.Now()
	fmt.Println("Part 3 answer: ", part3("input3.txt",256))

	
	fmt.Println()
	fmt.Println("Part 1: ", split.Sub(start))
	fmt.Println("Part 2: ", split2.Sub(split))
	fmt.Println("Part 2: ", time.Since(split2))
}

func part1(name string,nails int) int {
	input:= files.Read(name)
	steps:= ints.FromStringSlice( strings.Split(input,","))
	total:= 0

	for i:=1; i<len(steps); i++ {
		if ints.Mod(steps[i]-steps[i-1],nails) == nails/2 {
			total++
		}
	}

	return total
}

func part2(name string,nails int) int {
	input:= files.Read(name)
	steps:= ints.FromStringSlice( strings.Split(input,","))
	total:=0
	strings:= [][]int{}

	for i:=1; i<len(steps); i++ {
		newString:= []int{steps[i-1],steps[i]}
		slices.Sort(newString) 

		strings = append(strings, newString)
	}

	for i:= range strings {
		for j:=0;j<i;j++ {
			if checkIntersection(strings[i],strings[j],nails) {
				total++
			}
		}
	}

	return total
}


func part3(name string,nails int) int {
	input:= files.Read(name)
	steps:= ints.FromStringSlice( strings.Split(input,","))
	max:=0
	strings:= [][]int{}

	for i:=1; i<len(steps); i++ {
		newString:= []int{ints.Mod(steps[i-1],nails),ints.Mod(steps[i],nails)}
		slices.Sort(newString) 

		strings = append(strings, newString)
	}

	for i:=range nails {
		for j:=i+1;j<nails;j++ {
			newSlice:=[]int{i,j}
			total:=0
			for _,stringObject:= range strings {
				if checkIntersection(newSlice,stringObject,nails) {
					total++
				}
			}
			
			if total>max {
				max = total
			}
		}
	}

	return max
}

func checkIntersection(a,b []int, nails int) bool {
	betweenCount:= 0

	if a[0] == b[0] && a[1] == b[1] {
		return true
	}

	for _, val:= range b {
		if val == a[0] || val == a[1] {
			return false
		}
		if val>a[0] && val<a[1] {
			betweenCount++
		}
	}

	return betweenCount == 1
}