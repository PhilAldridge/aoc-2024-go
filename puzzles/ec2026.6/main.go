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
	mapping := parseMapping(name)

	ingredientsUsed := collectIngredients("NULL", mapping)

	return len(ingredientsUsed)
}

func part2(name string) int {
	mapping := reverseMapping(name)

	itemsUsingNull := collectIngredients("NULL", mapping)

	return len(itemsUsingNull)
}

func part3(name string) int {
	mapping := reverseMapping(name)

	itemsUsingNullV1 := collectIngredients("NULLv1", mapping)

	itemsUsingNullV2 := collectIngredients("NULLv2", mapping)

	itemsUsingNullV3 := collectIngredients("NULLv3", mapping)

	counter := make(map[string]int)

	for item := range itemsUsingNullV1 {
		counter[item]++
	}

	for item := range itemsUsingNullV2 {
		counter[item]++
	}

	for item := range itemsUsingNullV3 {
		counter[item]++
	}

	total := 0

	for _, count := range counter {
		if count > 1 {
			total++
		}
	}

	return total
}

// parseMapping maps an item to a slice of its direct ingredients - used in part 1
func parseMapping(name string) map[string][]string {
	out := make(map[string][]string)

	for _, line := range files.ReadLines(name) {
		split := strings.Split(line, ":")
		out[split[0]] = strings.Split(split[1], ",")
	}

	return out
}

// reverseMapping maps each ingredient to all the items that directly use it - used in parts 2 and 3
func reverseMapping(name string) map[string][]string {
	out := make(map[string][]string)

	for _, line := range files.ReadLines(name) {
		split := strings.Split(line, ":")

		for _, item := range strings.Split(split[1], ",") {
			current := out[item]
			if !slices.Contains(current, split[0]) {
				out[item] = append(out[item], split[0])
			}
		}
	}

	return out
}

// collectIngredientsRecursion recursively searches the mapping, to find all unique nodes that can be reached by going forward through the graph, starting at item
func collectIngredientsRecursion(item string, mapping map[string][]string, out map[string]bool) {
	for _, ingredient := range mapping[item] {
		// skip if ingredient already added - avoids infinite recursion
		if out[ingredient] {
			continue
		}

		// add ingredient
		out[ingredient] = true

		// add ingredients used to make ingredient
		collectIngredientsRecursion(ingredient, mapping, out)
	}
}

func collectIngredients(item string, mapping map[string][]string) map[string]bool {
	ingredientsUsed := make(map[string]bool)

	// populate ingredientsUsed
	collectIngredientsRecursion(item, mapping, ingredientsUsed)

	// Remove the initial node, as it is not an ingredient used to make itself
	delete(ingredientsUsed, item)

	return ingredientsUsed
}
