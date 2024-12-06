package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/kylehoehns/aoc-2023-go/pkg/files"
	"github.com/kylehoehns/aoc-2023-go/pkg/ints"
)

func main() {
	start:= time.Now()
	fmt.Println("Part 1: ", part1("input.txt"), " in:", time.Since(start))
	start= time.Now()
	fmt.Println("Part 2: ", part2("input.txt"), " in:", time.Since(start))
}

func part1(name string) int {
	sections := files.ReadParagraphs(name)
	manuals:= sections[1]
	mapping:= createMapping(sections[0])
	total:=0
	for _,manual:= range manuals {
		pages:= ints.FromStringSlice(strings.Split(manual,","))
		if checkValid(pages,mapping) {
			total += pages[(len(pages)-1)/2]
		}
	}
	return total
}

func part2(name string) int {
	sections := files.ReadParagraphs(name)
	manuals:= sections[1]
	mapping:= createMapping(sections[0])
	total:=0
	
	for _,manual:= range manuals {
		pages:= ints.FromStringSlice(strings.Split(manual,","))
		if !checkValid(pages,mapping) {
			total += getMiddleOrdered(pages,mapping)
		}
	}
	return total
}

func createMapping (mappings []string) map[int][]int {
	mapping:= make(map[int][]int)
	for _,m:= range mappings {
		split := ints.FromStringSlice(strings.Split(m,"|"))
		mapping[split[0]] = append(mapping[split[0]], split[1])
	}
	return mapping
}

func checkValid (pages []int, mapping map[int][]int) bool {
		for j:=1; j<len(pages);j++ {
			for k:=0;k<j;k++ {
				if slices.Contains(mapping[pages[j]],pages[k]){
					return false
				}
			}
		}
		return true
}

func getMiddleOrdered(pages []int, mapping map[int][]int) int {
	//bubble sort
	errorsFound:= true
	for errorsFound {
		errorsFound = false
		for j:=1; j<len(pages);j++ {
			for k:=0;k<j;k++ {
				if slices.Contains(mapping[pages[j]],pages[k]){
					errorsFound = true
					pages[j], pages[k] = pages[k],pages[j]
					break
				}
			}
		}
	}
	return pages[(len(pages)-1)/2]
}
