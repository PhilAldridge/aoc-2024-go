package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/complexn"
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

func part1(name string) string {
	input:= files.Read(name)
	a,_:= parseComplex(input)

	divisor:= complexn.NewComplex(10,10)

	r := performSteps(a, divisor, 3)

	return complexToString(r)
}

func part2(name string) int {
	input:= files.Read(name)
	a,_:= parseComplex(input)

	divisor:= complexn.NewComplex(100000,100000)

	count:=0

	for i:= range 101 {
		rowString:=""
		for j:= range 101 {
			point := complexn.Add(a,complexn.NewComplex(j*10,i*10))
			result:= performSteps(point,divisor, 100)
			if checkValid(result) {
				rowString += "#"
				count ++
			} else {
				rowString += " "
			}
		}
		fmt.Println(rowString)
	}


	return count
}


func part3(name string) int {
	input:= files.Read(name)
	a,_:= parseComplex(input)

	divisor:= complexn.NewComplex(100000,100000)

	count:=0
	
	for i:= range 1001 {
		rowString:=""
		for j:= range 1001 {
			point := complexn.Add(a,complexn.NewComplex(j*1,i*1))
			result:= performSteps(point,divisor, 100)
			if checkValid(result) {
				rowString += "#"
				count ++
			} else {
				rowString += " "
			}
		}
	}


	return count
}

func parseComplex(input string) (complexn.Complex, string) {
	split:= strings.Split(input,"=")
	name:= split[0]
	nums:= ints.FromStringSlice(strings.Split(split[1][1:len(split[1])-1],","))
	comp:= complexn.NewComplex(nums[0],nums[1])

	return comp, name
}

func complexToString(input complexn.Complex) string {
	return fmt.Sprintf("[%d,%d]",input.R,input.I)
}

func performSteps(input, divisor complexn.Complex, repetitions int) complexn.Complex {
	r:= complexn.NewComplex(0,0)
	for range repetitions {
		r = complexn.Multiply(r,r)
		r = complexn.DivideBasic(r,divisor)
		r = complexn.Add(r,input)
	}

	return r
} 

func checkValid(result complexn.Complex) bool {
	limit:= 1000000
	if result.R < -limit || result.R > limit || result.I < -limit || result.I > limit {
		return false
	}

	return true
}