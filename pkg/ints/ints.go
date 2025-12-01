package ints

import (
	"math"
	"strconv"
)

func Sum(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func FromString(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return val
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// FromStringSlice converts a slice of strings to a slice of ints
func FromStringSlice(input []string) []int {
	output := make([]int, 0)
	for _, str := range input {
		output = append(output, FromString(str))
	}
	return output
}

func Min(numbers []int) int {
	m := numbers[0]
	for _, num := range numbers {
		if num < m {
			m = num
		}
	}
	return m
}

func AllSame(numbers []int) bool {
	for i := 1; i < len(numbers); i++ {
		if numbers[i] != numbers[i-1] {
			return false
		}
	}
	return true
}

func Pow (base int, exp int) int {
	res:=1
	for exp>0 {
		res *= base
		exp--
	}
	return res
}

func CountDigits(num int) int {
    if num == 0 {
        return 1 // Special case for 0, which has 1 digit
    }
    return int(math.Log10(float64(num))) + 1
}

func GCD(a, b int) int {
	if a< 0 {
		a *= -1
	}
	if b<0 {
		b *= -1
	}
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func Mod(x, d int) int {
  if d < 0 { d = -d }
  x = x % d
  if x < 0 { return x + d }
  return x
}