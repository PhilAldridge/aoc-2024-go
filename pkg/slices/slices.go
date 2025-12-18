package slices

import (
	"strings"

	"github.com/PhilAldridge/aoc-2024-go/pkg/ints"
)

func SlidingWindow(slice []int, size int) [][]int {
	var result [][]int = make([][]int, 0)
	for i := 0; i <= len(slice)-size; i++ {
		window := slice[i : i+size]
		result = append(result, window)
	}
	return result
}

func StringSliceToStringMapStringSlice(input []string, separator1 string, separator2 string) map[string][]string {
	result := make(map[string][]string)

	for _,line:= range input {
		split1:= strings.Split(line,separator1)
		split2:= strings.Split(split1[1],separator2)

		result[split1[0]] = split2
	}

	return result
}

func StringSliceToStringMapIntSlice(input []string, separator1 string, separator2 string) map[string][]int {
	result := make(map[string][]int)

	for _,line:= range input {
		split1:= strings.Split(line,separator1)
		split2:= strings.Split(split1[1],separator2)

		result[split1[0]] = ints.FromStringSlice(split2)
	}

	return result
}

func StringSliceToIntMapIntSlice(input []string, separator1 string, separator2 string) map[int][]int {
	result := make(map[int][]int)

	for _,line:= range input {
		split1:= strings.Split(line,separator1)
		split2:= strings.Split(split1[1],separator2)

		result[ints.FromString(split1[0])] = ints.FromStringSlice(split2)
	}

	return result
}
