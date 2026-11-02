package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		expected := 2024
		actual := part1("test-input.txt")
		assert.Equal(t, expected, actual)
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		expected := "aaa,aoc,bbb,ccc,eee,ooo,z24,z99"
		actual := part2("test-input.txt")
		assert.Equal(t, expected, actual)
	})

}
