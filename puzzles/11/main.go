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
	stonesMap := getStonesMap(name)
	linkedListOfStoneBlink := make(map[uint64][]uint64)
	return blinkRecursion(stonesMap,linkedListOfStoneBlink,25)
}

func part2(name string) int {
	stonesMap := getStonesMap(name)
	linkedListOfStoneBlink := make(map[uint64][]uint64)
	return blinkRecursion(stonesMap,linkedListOfStoneBlink,75)
}

//stones stored as a map so that stones of the same value can be grouped
//saving calculation time and array size issues
//key is the stone value, value is the number of those stones
func getStonesMap(name string) map[uint64]int {
	stones := ints.FromStringSlice(strings.Split(files.Read(name)," "))
	stoneMap := make(map[uint64]int)
	for _,stone:= range stones {
		stoneMap[uint64(stone)] = stoneMap[uint64(stone)]+1
	}
	return stoneMap
}

func blinkRecursion(stoneMap map[uint64]int, linkedListOfBlink map[uint64][]uint64, blinksLeft int) int {
	//Base case - blinks done. Count up all stones
	if blinksLeft == 0 {
		total:=0
		for _,count := range stoneMap {
			total+= count
		}
		return total
	}
	nextStoneMap := make(map[uint64]int)
	for stone,count := range stoneMap {
		//Map of pre-calculated stone numbers
		//if in map, use the map val as the list of stones this stone turns into
		if len(linkedListOfBlink[stone])>0 {
			for _,stoneVal := range linkedListOfBlink[stone] {
				nextStoneMap[stoneVal] = nextStoneMap[stoneVal] + count
			}
			continue
		}
		//else, apply rules manually and add results to both nextStoneMap and linkedListOfBlink
		applyRules(stone,count,linkedListOfBlink,nextStoneMap)
		
	}
	return blinkRecursion(nextStoneMap,linkedListOfBlink,blinksLeft-1)
}

func applyRules(stone uint64, count int, linkedListOfBlink map[uint64][]uint64, nextStoneMap map[uint64]int) {
	stoneStr:= strconv.FormatUint(stone,10)
	if stone == 0 {
		//All 0 stones turn into 1 stones
		nextStoneMap[1] = nextStoneMap[1] + count
		linkedListOfBlink[stone] = []uint64{1}
	} else if len(stoneStr)%2 ==0 {
		//All stones with even no of digits are split into two stones
		left,_ := strconv.ParseUint(stoneStr[:len(stoneStr)/2],10,64)
		right,_ := strconv.ParseUint(stoneStr[len(stoneStr)/2:],10,64)
		nextStoneMap[left] = nextStoneMap[left] + count
		nextStoneMap[right] = nextStoneMap[right] + count
		linkedListOfBlink[stone] = []uint64{left,right}
	} else {
		//All other stones have their value multiplied by 2024
		newVal := stone*2024
		nextStoneMap[newVal] = nextStoneMap[newVal] + count
		linkedListOfBlink[stone] = []uint64{newVal}
	}
}
