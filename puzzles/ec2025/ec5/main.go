package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
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

type vertebrae struct {
	val         int
	left, right *int
}

func part1(name string) int {
	input := files.Read(name)
	split1 := strings.Split(input, ":")
	split2 := strings.Split(split1[1], ",")
	vals := ints.FromStringSlice(split2)

	quality, _ := calculateQuality(vals)
	return quality
}

func part2(name string) int {
	input := files.ReadLines(name)
	max := 0
	min := math.MaxInt

	for _, sword := range input {
		split1 := strings.Split(sword, ":")
		split2 := strings.Split(split1[1], ",")
		vals := ints.FromStringSlice(split2)
		quality,_ := calculateQuality(vals)
		if quality < min {
			min = quality
		}
		if quality > max {
			max = quality
		}
	}

	return max-min
}

func part3(name string) int {
	input := files.ReadLines(name)
	swords:= []swordType{}

	for _, sword := range input {
		split1 := strings.Split(sword, ":")
		split2 := strings.Split(split1[1], ",")
		vals := ints.FromStringSlice(split2)
		quality,fishbone := calculateQuality(vals)
		swords = append(swords, swordType{
			id: ints.FromString(split1[0]),
			quality:quality,
			fishbone:fishbone,
		})
	}

	slices.SortFunc(swords, func(a,b swordType) int {
		// Test quality first
		if a.quality != b.quality {
			return b.quality - a.quality
		}

		// Then value of each vertebrae
		for i:= range a.fishbone {
			if i>= len(b.fishbone) {
				return -1
			}

			aVal:= a.fishbone[i].value()
			bVal:= b.fishbone[i].value()

			if aVal != bVal {
				return bVal - aVal
			}
		}

		if len(b.fishbone) > len(a.fishbone) {
			return 1
		}

		// Identical swords, sort by id
		return b.id - a.id
	})

	total:=0
	for i,sword:= range swords {
		total += (i+1)*sword.id
	}

	return total
}

type swordType struct {
	id,quality int
	fishbone []vertebrae
}

func intPtrOrEmpty(p *int) string {
	if p == nil {
		return ""
	}
	return fmt.Sprintf("%d", *p)
}

func calculateQuality(vals []int) (int,[]vertebrae) {
	currentVertebae := vertebrae{val: vals[0]}

	fishbone := []vertebrae{currentVertebae}

	for i := 1; i < len(vals); i++ {
		placed := false
		for j := range fishbone {
			if vals[i] < fishbone[j].val && fishbone[j].left == nil {
				fishbone[j].left = &vals[i]
				placed = true
				break
			} else if vals[i] > fishbone[j].val && fishbone[j].right == nil {
				fishbone[j].right = &vals[i]
				placed = true
				break
			}
		}
		if placed {
			continue
		}
		currentVertebae = vertebrae{val: vals[i]}
		fishbone = append(fishbone, currentVertebae)
	}

	result := ""
	for _, vertebrae := range fishbone {
		result += strconv.Itoa(vertebrae.val)
	}

	return ints.FromString(result),fishbone
}

func (v vertebrae) value() int {
	answer:= intPtrOrEmpty(v.left) + strconv.Itoa(v.val) + intPtrOrEmpty(v.right)
	return ints.FromString(answer)
}