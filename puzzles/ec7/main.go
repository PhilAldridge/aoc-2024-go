package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	myslices "github.com/PhilAldridge/aoc-2024-go/pkg/slices"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1 answer: ", part1("input.txt"))
	split := time.Now()
	fmt.Println("Part 2 answer: ", part2("input2.txt"))
	split2 := time.Now()
	fmt.Println("Part 3 answer: ", part3("input3.txt"))

	fmt.Println()
	fmt.Println("Part 1: ", split.Sub(start))
	fmt.Println("Part 2: ", split2.Sub(split))
	fmt.Println("Part 2: ", time.Since(split2))
}

func part1(name string) string {
	input := files.ReadParagraphs(name)
	names := strings.Split(input[0][0], ",")
	ruleMap := myslices.StringSliceToStringMapStringSlice(input[1], " > ", ",")

	for _, name := range names {
		if checkName(name, ruleMap) {
			return name
		}

	}

	panic("something went wrong")
}

func part2(name string) int {
	input := files.ReadParagraphs(name)
	names := strings.Split(input[0][0], ",")
	ruleMap := myslices.StringSliceToStringMapStringSlice(input[1], " > ", ",")
	total:=0

	for i, name := range names {
		if checkName(name, ruleMap) {
			total += i+1
		}

	}

	return total
}

func part3(name string) int {
	input := files.ReadParagraphs(name)
	prefixes := strings.Split(input[0][0], ",")
	ruleMap := myslices.StringSliceToStringMapStringSlice(input[1], " > ", ",")
	names:= make(map[string]bool)

	for _, prefix := range prefixes {
		if checkName(prefix, ruleMap) {
			getNames(prefix, ruleMap,names)
		}

	}

	return len(names)
}

func checkName(name string, ruleMap map[string][]string) bool {
	for i := 1; i < len(name); i++ {
		if !slices.Contains(ruleMap[string(name[i-1])], string(name[i])) {
			return false
		}
	}
	return true
}

func getNames(prefix string, ruleMap map[string][]string, names map[string]bool) {
	if len(prefix) > 11 || names[prefix] {
		return
	}

	nextChars:= ruleMap[prefix[len(prefix)-1:]]

	for _, next:= range nextChars {
		getNames(prefix+next,ruleMap,names)
	}

	if len(prefix)>=7 {
		names[prefix] = true
	}
}