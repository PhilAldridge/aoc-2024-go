package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/ints"
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
	total := 0

	trees:= parseInput(name)
	for _, tree:= range trees {
		seconds, size := tree.growToSizeLimit()
		fmt.Println(tree.id, seconds, size)
		total += seconds
	}

	return total
}

func part2(name string) int {
	total := 0

	trees:= parseInput(name)
	for _, tree:= range trees {
		seconds, leaves := tree.growToLeafLimit()
		fmt.Println(tree.id, seconds, leaves)
		total += seconds
	}

	return total
}

func part3(name string) int {
	return part2(name)
}

type tree struct {
	id, growth,limit int
	root branch
	previousResults map[int]int
}

type branch struct {
	size, lifeTime int
	left, right *branch
}

func (b *branch) grow(growth int) (int, int) {
	total:= 0
	leaves:=0
	if b.left != nil {
		subTotal, subLeaves:= b.left.grow(growth)
		total += subTotal
		leaves += subLeaves
	}

	if b.right != nil {
		subTotal, subLeaves:= b.right.grow(growth)
		total += subTotal
		leaves += subLeaves
	}

	b.size ++
	total += b.size

	if b.left == nil && b.right == nil {
		leaves++
	}
	
	if b.size >= growth {
		b.size /=2
		b.left = &branch{
			size: b.size,
			left: b.left,
			right: b.right,
		}
		b.right = &branch{
			size: 1,
		}
		leaves++ 
	}

	if b.left != nil || b.right != nil {
		total++
	}

	return total,leaves
}

func (b *branch) growLeavesWithPreviousResults(growth int, previousResults map[int]int) (int) {
	b.lifeTime ++
	if result,ok:= previousResults[b.lifeTime]; ok {
		return result
	}

	leaves:=0

	if b.left != nil {
		_, subLeaves:= b.left.grow(growth)
		leaves += subLeaves
	}

	if b.right != nil {
		_, subLeaves:= b.right.grow(growth)
		leaves += subLeaves
	}

	b.size ++

	if b.left == nil && b.right == nil {
		leaves++
	}
	
	if b.size >= growth {
		b.size /=2
		b.left = &branch{
			size: b.size,
			left: b.left,
			right: b.right,
		}
		b.right = &branch{
			size: 1,
		}
		leaves++ 
	}

	previousResults[b.lifeTime]= leaves

	return leaves
}

// return seconds, total elements (including branches)
func (t tree) growToSizeLimit() (int,int) {
	seconds:=0
	size:=0
	for {
		seconds++
		size, _ = t.root.grow(t.growth)

		if size >= t.limit {
			break
		}
	}

	return seconds, size
}

// return seconds, leaves
func (t tree) growToLeafLimit() (int,int) {
	seconds:=0
	leaf:=0
	for {
		seconds++
		leaf = t.root.growLeavesWithPreviousResults(t.growth, t.previousResults)

		if leaf >= t.limit {
			break
		}
	}

	return seconds, leaf
}

func parseInput(name string) []tree {
	out:= []tree{}

	for _, row:= range files.ReadLines(name) {
		split:= strings.Split(row, ", ")
		out = append(out, tree{
			id: ints.FromString(split[0][3:]),
			growth: ints.FromString(split[1][7:]),
			limit: ints.FromString(split[2][6:]),
			previousResults: make(map[int]int),
		})
	}

	return out
}
