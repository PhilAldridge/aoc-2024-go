package main

import (
	"fmt"
	"strconv"
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
	dial := 50
	MAX_VALUE:= 99
	zeroCount:=0
	lines := files.ReadLines(name)
	for _, l:= range lines {
		amount,err := strconv.Atoi(l[1:])
		if err != nil {
			panic(err)
		}

		if l[0] == 'L' {
			dial -= amount
		} else {
			dial += amount
		}

		dial = ints.Mod(dial, MAX_VALUE+1)

		if dial ==0 {
			zeroCount ++
		}
	}
	return zeroCount
}

func part2(name string) int {
	dial := 50
	MAX_VALUE:= 99
	modulo:= MAX_VALUE + 1
	zeroCount:=0
	lines := files.ReadLines(name)
	for _, l:= range lines {
		amount,err := strconv.Atoi(l[1:])
		if err != nil {
			panic(err)
		}

		zeroCount += amount/(modulo)
		amount %= (modulo)

		if l[0] == 'L' {
			wasZero:= dial==0
			dial -= amount
			if dial < 0 {
				dial += modulo
				if !wasZero {
					zeroCount ++
				}
			} else if dial == 0 {
				zeroCount ++
			}
		} else {			
			dial += amount
			if dial > MAX_VALUE {
				dial -= modulo
				zeroCount ++
			}
		}
	}
	return zeroCount
}
