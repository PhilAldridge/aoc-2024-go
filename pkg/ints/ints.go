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

func Max(numbers []int) int {
	m := numbers[0]
	for _, num := range numbers {
		if num > m {
			m = num
		}
	}
	return m
}

func Mean(numbers []int) int {
	m := 0
	for _, num := range numbers {
		m+=num
	}
	return m/len(numbers)
}

func MinMax(numbers []int) (int,int) {
	m := numbers[0]
	n := numbers[0]
	for _, num := range numbers {
		if num < m {
			m = num
			continue
		}
		if num > n {
			n = num
		}
	}
	return m,n
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

func LCM(a,b int) int {
	if a< 0 {
		a *= -1
	}
	if b<0 {
		b *= -1
	}
	return (a*b)/GCD(a,b)
}

func Mod(x, d int) int {
  if d < 0 { d = -d }
  x = x % d
  if x < 0 { return x + d }
  return x
}

func GetIntsBetween(a,b int) []int {
	result:= []int{}

	if a<b {
		for i:= a+1; i<b; i++ {
			result = append(result, i)
		}
	} else if b<a {
		for i:=b+1; i<a; i++ {
			result = append(result, i)
		}
	}
	return result
}

func GetIntsBetweenInclusive(a,b int) []int {
	result:= []int{}

	if a<b {
		for i:= a; i<=b; i++ {
			result = append(result, i)
		}
	} else {
		for i:=b; i<=a; i++ {
			result = append(result, i)
		}
	}
	return result
}

func IsBetween(a,b,c int) bool {
	if a < b && a > c {
		return true
	}

	if a > b && a < c {
		return true
	}

	return false
}

func Factorial(a int) int {
	total:=1
	for i:=2; i<=a; i++ {
		total*=i
	}
	return total
}

func SumMap[T comparable](mapping map[T]int) int {
	total:= 0

	for _, val:= range mapping {
		total += val
	}

	return total
}

func ModularDifference(a,b,mod int) int {
	try1:= a-b
	if try1 < 0 {
		try1 +=mod
	}

	try2:=b-a
	if try2 < 0 {
		try2 += mod
	}

	if try1 < try2 {
		return try1
	}

	return try2
}