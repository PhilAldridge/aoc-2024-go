package main

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
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
	var total atomic.Uint64
	wg := &sync.WaitGroup{}
	lines := files.ReadLines(name)
	for _, line := range lines {
		wg.Add(1)
		go testLine(line, &total, wg, false)
	}
	wg.Wait()
	return int(total.Load())
}

func part2(name string) int {
	var total atomic.Uint64
	wg := &sync.WaitGroup{}
	lines := files.ReadLines(name)
	for _, line := range lines {
		wg.Add(1)
		go testLine(line, &total, wg, true)
	}
	wg.Wait()
	return int(total.Load())
}

func testLine(line string, total *atomic.Uint64, wg *sync.WaitGroup, incConcat bool) {
	defer wg.Done()
	split1 := strings.Split(line, ": ")
	testValue := ints.FromString(split1[0])
	nums := ints.FromStringSlice(strings.Split(split1[1], " "))
	if checkOps(testValue, nums[0], nums[1:], incConcat) {
		total.Add(uint64(testValue))
	}
}

func checkOps(testValue int, currentValue int, nums []int, incConcat bool) bool {
	if currentValue > testValue {
		return false
	}
	if len(nums) == 0 {
		return testValue == currentValue
	}
	return checkOps(testValue, currentValue*nums[0], nums[1:], incConcat) ||
		checkOps(testValue, currentValue+nums[0], nums[1:], incConcat) ||
		(incConcat && checkOps(testValue, concatInts(currentValue, nums[0]), nums[1:], incConcat))
}

func concatInts(v1 int, v2 int) int {
	return v1*ints.Pow(10, ints.CountDigits(v2)) + v2
}
