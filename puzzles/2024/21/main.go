package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

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
	total:=0
	codes:= files.ReadLines(name)
	cache:= make(map[string][]string)
	dPadCode:= getPadMap("dPad.txt")
	for _,code:= range codes {
		nPadInst:= nPadCode(code)
		for i:=0;i<2;i++ {
			nPadInst = nextRobot(nPadInst,cache,dPadCode)
		}
		subTotal:= 0
		for k,v:= range nPadInst {
			subTotal+= len(k)*v
		}
		fmt.Printf("%d * %d\n",subTotal, intVal(code))
		total += subTotal * intVal(code)
	}

	return total
}

func part2(name string) int {
	total:=0
	codes:= files.ReadLines(name)
	cache:= make(map[string][]string)
	dPadCode:= getPadMap("dPad.txt")
	for _,code:= range codes {
		nPadInst:= nPadCode(code)
		for i:=0;i<25;i++ {
			nPadInst = nextRobot(nPadInst,cache,dPadCode)
		}
		subTotal:= 0
		for k,v:= range nPadInst {
			subTotal+= len(k)*v
		}
		total += subTotal * intVal(code)
	}
	return total
}

func intVal(code string) int {
	r:= regexp.MustCompile(`\d+`)
	matches:= r.FindAllString(code,-1)
	return ints.FromString(strings.Join(matches,""))
}

func getPadMap(name string) map[[2]rune]string{
	mapping:= make(map[[2]rune]string)
	pad:= files.ReadLines(name)
	nPadFlag:= false
	if name == "nPad.txt" {
		nPadFlag = true
	}
	for i,row:= range pad {
		for j,char:= range row {
			if char == ' ' {continue}
			for k,row2:= range pad {
				for l,char2:= range row2{
					if char2 == ' ' {continue}
					val:= ""
					y:= k-i
					x:= l-j
					zeroFlag:= (l==0||j==0)&&(
						(nPadFlag && (k==len(pad)-1 || i==len(pad)-1) )||
						(!nPadFlag && (k==0 || i ==0)))

					//prefer < before vertical and > after vertical, but avoid the empty space
					if x<0 {
						if y<0 {
							if nPadFlag && zeroFlag {
								val = strings.Repeat("^",-y)+strings.Repeat("<",-x) 
							} else {
								val = strings.Repeat("<",-x) +strings.Repeat("^",-y)
							}
						} else {
							if nPadFlag || !zeroFlag {
								val = strings.Repeat("<",-x) +strings.Repeat("v",y)
							} else {
								val = strings.Repeat("v",y)+strings.Repeat("<",-x) 
							}
						}
					} else {
						if y<0 {
							if zeroFlag {
								val = strings.Repeat(">",x) +strings.Repeat("^",-y)
							} else {
								val = strings.Repeat("^",-y)+strings.Repeat(">",x) 
							}
						} else {
							if zeroFlag {
								val = strings.Repeat(">",x) +strings.Repeat("v",y)
							} else {
								val = strings.Repeat("v",y)+strings.Repeat(">",x) 
							}
						}
					}
					mapping[[2]rune{char,char2}] = val+"A"
				}
			}
		}
	}
	return mapping
}

func nPadCode(code string) map[string]int {
	nPadMap:= getPadMap("nPad.txt")
	res:= make(map[string]int)
	prev:= 'A'
	for _,button:= range code {
		val:= nPadMap[[2]rune{prev,button}]
		res[val] ++
		prev = button
	}
	return res
}

func nextRobot(prev map[string]int, cache map[string][]string, dPad map[[2]rune]string) map[string]int {
	next:= make(map[string]int)
	for k,v:=range prev {
		prevChar:='A'
		if c,ok:= cache[k]; ok {
			for _,cVal:= range c {
				next[cVal] += v
			}
			continue
		}
		forCache:= []string{}
		for _,button:= range k {
			val:= dPad[[2]rune{prevChar,button}]
			next[val] +=v
			forCache = append(forCache, val)
			prevChar = button
		}
		cache[k] = forCache
	}
	return next
}