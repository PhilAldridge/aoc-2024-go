package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
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

func part1(name string) int {
	mapping:= parseMapping(name)

	ingredientsUsed:= make(map[string]bool)

	collectIngredients("NULL", mapping, ingredientsUsed)

	delete(ingredientsUsed, "NULL")

	return len(ingredientsUsed)
}

func part2(name string) int {
	mapping:= reverseMapping(name)

	itemsUsingNull:= make(map[string]bool)

	collectIngredients("NULL", mapping, itemsUsingNull)

	delete(itemsUsingNull, "NULL")

	return len(itemsUsingNull)
}

func part3(name string) int {
	mapping:= reverseMapping(name)

	itemsUsingNullV1:= make(map[string]bool)

	collectIngredients("NULLv1", mapping, itemsUsingNullV1)

	itemsUsingNullV2:= make(map[string]bool)

	collectIngredients("NULLv2", mapping, itemsUsingNullV2)

	itemsUsingNullV3:= make(map[string]bool)

	collectIngredients("NULLv2", mapping, itemsUsingNullV3)

	counter:= make(map[string]int)

	for item:= range itemsUsingNullV1 {
		counter[item]++
	}

	for item:= range itemsUsingNullV2 {
		counter[item]++
	}

	for item:= range itemsUsingNullV3 {
		counter[item]++
	}

	total:= 0

	for _, count:= range counter {
		if count >1 {
			total ++
		}
	}

	return total
}

func parseMapping(name string) map[string][]string {
	out:= make(map[string][]string)

	for _, line:= range files.ReadLines(name) {
		split:= strings.Split(line, ":")
		out[split[0]] = strings.Split(split[1],",")
	}
	
	return out
}

func reverseMapping(name string) map[string][]string {
	out:= make(map[string][]string)

	for _, line:= range files.ReadLines(name) {
		split:= strings.Split(line, ":")

		for _, item := range strings.Split(split[1], ",") {
			current:= out[item]
			if !slices.Contains(current, split[0]) {
				out[item] = append(out[item], split[0])
			}
		}
	}

	return out
}

func collectIngredients(item string, mapping map[string][]string, out map[string]bool) {
	for _, ingredient:= range mapping[item] {
		if out[ingredient] { continue}
		out[ingredient] = true
		collectIngredients(ingredient, mapping, out)
	}
}