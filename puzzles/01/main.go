package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/kylehoehns/aoc-2023-go/pkg/files"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1: ", part1("input.txt"), " in:", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", part2("input.txt"), " in:", time.Since(start))
}

func part1(name string) int {
	list1, list2 := getLists(name)
	slices.Sort(list1)
	slices.Sort(list2)
	total := 0
	for i, v := range list1 {
		w := list2[i]
		if v > w {
			total += v
			total -= w
		} else {
			total -= v
			total += w
		}
	}
	return total
}

func part2(name string) int {
	list1, list2 := getLists(name)
	count := make(map[int]int)
	for _, v := range list2 {
		count[v] = count[v] + 1
	}
	total := 0
	for _, v := range list1 {
		total += v * count[v]
	}
	return total
}

func getLists(name string) ([]int, []int) {
	lines := files.ReadLines(name)
	list1 := []int{}
	list2 := []int{}
	for _, line := range lines {
		vals := strings.Split(line, "   ")
		if i, err := strconv.Atoi(vals[0]); err == nil {
			list1 = append(list1, i)
		} else {
			log.Fatal(err)
		}
		if i, err := strconv.Atoi(vals[1]); err == nil {
			list2 = append(list2, i)
		} else {
			log.Fatal(err)
		}

	}
	return list1, list2
}
