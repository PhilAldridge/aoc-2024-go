package main

import (
	"fmt"
	"math"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/kylehoehns/aoc-2023-go/pkg/files"
	"github.com/kylehoehns/aoc-2023-go/pkg/ints"
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
	lines:= files.ReadLines(name)
	for _,line:=range lines{
		wg.Add(1)
		go testLine(line, &total,wg,false)
	}
	wg.Wait()
	return int(total.Load())
}

func part2(name string) int {
	var total atomic.Uint64
	wg := &sync.WaitGroup{} 
	lines:= files.ReadLines(name)
	for _,line:=range lines{
		wg.Add(1)
		go testLine(line, &total,wg,true)
	}
	wg.Wait()
	return int(total.Load())
}

func testLine(line string, total *atomic.Uint64, wg *sync.WaitGroup,incConcat bool) {
	defer wg.Done()
	split1:= strings.Split(line,": ")
	testValue := ints.FromString(split1[0])
	nums:= ints.FromStringSlice(strings.Split(split1[1]," "))
	if checkOps(testValue,nums[0],nums[1:],incConcat) {
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
	return checkOps(testValue,currentValue*nums[0], nums[1:],incConcat) ||
		checkOps(testValue, currentValue+nums[0],nums[1:],incConcat) ||
		(incConcat && checkOps(testValue, concat(currentValue,nums[0]),nums[1:],incConcat))
}

func concat(v1 int, v2 int) int {
	return v1*pow(10,countDigits(v2)) + v2
}

func pow (base int, exp int) int {
	res:=1
	for exp>0 {
		res *= base
		exp--
	}
	return res
}

func countDigits(num int) int {
    if num == 0 {
        return 1 // Special case for 0, which has 1 digit
    }
    return int(math.Log10(float64(num))) + 1
}
