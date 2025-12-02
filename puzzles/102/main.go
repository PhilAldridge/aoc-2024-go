package main

import (
	"fmt"
	"strconv"
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
	codeRanges := strings.Split(files.Read(name), ",")
	invalidCodeSet:= make(map[int]bool)

	for _, codeRange:= range codeRanges {
		rangeStrings := strings.Split(codeRange, "-")
		rangeInts := ints.FromStringSlice(rangeStrings)

		for totalDigits:= len(rangeStrings[0]); totalDigits <= len(rangeStrings[1]); totalDigits++ {
			if totalDigits % 2 ==0 {
				recursiveCreateDoubleCode("",rangeInts[0],rangeInts[1],totalDigits,invalidCodeSet)
			}
			
		}
	}

	total:= 0
	for code:= range invalidCodeSet {
		total += code
	}

	return total
}

func part2(name string) int {
	codeRanges := strings.Split(files.Read(name), ",")
	invalidCodeSet:= make(map[int]bool)

	for _, codeRange:= range codeRanges {
		rangeStrings := strings.Split(codeRange, "-")
		rangeInts := ints.FromStringSlice(rangeStrings)

		for totalDigits:= len(rangeStrings[0]); totalDigits <= len(rangeStrings[1]); totalDigits++ {
			if totalDigits%2 ==0 {
				recursiveNCode("",rangeInts[0],rangeInts[1],totalDigits,invalidCodeSet)
			}
			
		}
	}

	total:= 0
	for code:= range invalidCodeSet {
		total += code
	}

	return total
}

func recursiveCreateDoubleCode(halfCodeSoFar string, min int, max int, totalDigits int, codeSet map[int]bool) {
	if len(halfCodeSoFar) > totalDigits/2 {
		return
	}

	minCode:= ints.FromString(halfCodeSoFar + strings.Repeat("0", totalDigits-len(halfCodeSoFar)))
	if minCode > max {
		return
	}

	maxCode:= ints.FromString(halfCodeSoFar + strings.Repeat("9", totalDigits-len(halfCodeSoFar)))
	if maxCode < min {
		return
	}

	if len(halfCodeSoFar)>0 && len(halfCodeSoFar) == totalDigits / 2 {
		code:= ints.FromString(strings.Repeat(halfCodeSoFar, 2))
		if code >= min && code <= max {
			codeSet[code] = true
		}
	}

	i:=0
	if len(halfCodeSoFar)==0 {
		i=1
	}

	for i<10 {
		recursiveCreateDoubleCode(halfCodeSoFar + strconv.Itoa(i),min,max,totalDigits,codeSet)
		i++
	}
}

func recursiveNCode(halfCodeSoFar string, min int, max int, totalDigits int, codeSet map[int]bool) {
	if len(halfCodeSoFar) > totalDigits/2 {
		return
	}

	minCode:= ints.FromString(halfCodeSoFar + strings.Repeat("0", totalDigits-len(halfCodeSoFar)))
	if minCode > max {
		return
	}

	maxCode:= ints.FromString(halfCodeSoFar + strings.Repeat("9", totalDigits-len(halfCodeSoFar)))
	if maxCode < min {
		return
	}

	if len(halfCodeSoFar) > 0 && totalDigits % len(halfCodeSoFar) == 0 {
		code:= ints.FromString(strings.Repeat(halfCodeSoFar, totalDigits/len(halfCodeSoFar)))
		if code >= min && code <= max {
			codeSet[code] = true
		}
	}

	i:=0
	if len(halfCodeSoFar)==0 {
		i=1
	}

	for i<10 {
		recursiveNCode(halfCodeSoFar + strconv.Itoa(i),min,max,totalDigits,codeSet)
		i++
	}
}