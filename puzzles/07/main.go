package main

import (
	"fmt"
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
		go testLine(line, &total,wg)
	}
	wg.Wait()
	return int(total.Load())
}

func part2(name string) int {
	return 0
}

func testLine(line string, total *atomic.Uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	split1:= strings.Split(line,": ")
	testValue := ints.FromString(split1[0])
	nums:= ints.FromStringSlice(strings.Split(split1[1]," "))
	for i:=0; i< (1<<(len(nums)-1)); i++ {
		base:= nums[0]
		for j:=1; j<len(nums);j++ {
			if (i>>(j-1))&1 == 1 {
				base += nums[j]
			} else {
				base *= nums[j]
			}
			if base > testValue {
				break
			}
		}
		if base == testValue {
			total.Add(uint64(testValue))
			return
		}
	}
}