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

func getStonesMap(name string) map[uint64]int {
	stones := ints.FromStringSlice(strings.Split(files.Read(name)," "))
	stoneMap := make(map[uint64]int)
	for _,stone:= range stones {
		stoneMap[uint64(stone)] = stoneMap[uint64(stone)]+1
	}
	return stoneMap
}

func blinkRecursion(stoneMap map[uint64]int, linkedListOfBlink map[uint64][]uint64, times int) int {
	// fmt.Println(stones)
	if times == 0 {
		total:=0
		for _,count := range stoneMap {
			total+= count
		}
		return total
	}
	nextStoneMap := make(map[uint64]int)
	for stone,count := range stoneMap {
		if len(linkedListOfBlink[stone])>0 {
			for _,stoneVal := range linkedListOfBlink[stone] {
				nextStoneMap[stoneVal] = nextStoneMap[stoneVal] + count
			}
			continue
		}
		stoneStr:= strconv.FormatUint(stone,10)
		if stone == 0 {
			nextStoneMap[1] = nextStoneMap[1] + count
			linkedListOfBlink[stone] = []uint64{1}
		} else if len(stoneStr)%2 ==0 {
			left,_ := strconv.ParseUint(stoneStr[:len(stoneStr)/2],10,64)
			right,_ := strconv.ParseUint(stoneStr[len(stoneStr)/2:],10,64)
			nextStoneMap[left] = nextStoneMap[left] + count
			nextStoneMap[right] = nextStoneMap[right] + count
			linkedListOfBlink[stone] = []uint64{left,right}
		} else {
			newVal := stone*2024
			nextStoneMap[newVal] = nextStoneMap[newVal] + count
			linkedListOfBlink[stone] = []uint64{newVal}
		}
		
	}
	return blinkRecursion(nextStoneMap,linkedListOfBlink,times-1)
}
