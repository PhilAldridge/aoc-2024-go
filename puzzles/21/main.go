package main

import (
	"fmt"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1: ", part1("input.txt"), " in:", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", part2("input.txt"), " in:", time.Since(start))
}

func part1(name string) int {
	total:=0
	// codes:= files.ReadLines(name)
	// for _,code:= range codes {
	// 	nPadInst:= nPadCode(code)
		fmt.Println(dPadAmounts())
	// }
	return total
}

func part2(name string) int {
	total:=0
	return total
}

func getPadMap(name string) map[[2]rune][2]int{
	mapping:= make(map[[2]rune][2]int)
	pad:= files.ReadLines(name)
	for i,row:= range pad {
		for j,char:= range row {
			for k,row2:= range pad {
				for l,char2:= range row2{
					y:= k-i
					x:= l-j
					val:= [2]int{0,0}
					if y<0 {
						val[0] = -y
					} else {
						val[0] = y
					}
					if x<0 {
						val[1] = -x
					} else {
						val[1] = x
					}
					mapping[[2]rune{char,char2}] = val
				}
			}
		}
	}
	return mapping
}

func dPadAmounts() ([2]int,[2]int) {
	dpadMap:= getPadMap("dPad.txt")
	thereY:= dpadMap[[2]rune{'A','^'}]
	thereY2:= dpadMap[[2]rune{'A','v'}]
	backY:= dpadMap[[2]rune{'^','A'}]
	backY2:= dpadMap[[2]rune{'v','A'}]
	thereX:= dpadMap[[2]rune{'A','<'}]
	thereX2:= dpadMap[[2]rune{'A','>'}]
	backX:= dpadMap[[2]rune{'<','A'}]
	backX2:= dpadMap[[2]rune{'>','A'}]
	totalY:= [2]int{	thereY[0]+thereY2[0]+backY[0]+backY2[0],thereY[1]+thereY2[1]+backY[1]+backY2[1]}
	totalX:= [2]int{	thereX[0]+thereX2[0]+backX[0]+backX2[0],thereX[1]+thereX2[1]+backX[1]+backX2[1]}
	return totalY,totalX
}


func nPadCode(code string) [][2]int {
	nPadMap:= getPadMap("nPad.txt")
	res:= [][2]int{}
	prev:= 'A'
	for _,button:= range code {
		val:= nPadMap[[2]rune{prev,button}]
		res = append(res, val)
		prev = button
	}
	return res
}

// func dPadCode(instructions [][2]int) [][2]int {
// 	res:= [][2]int{}
	
// 	for _, instr:= range instructions {
		
// 	}
// 	return res
// }

