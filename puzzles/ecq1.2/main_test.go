package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		expected := "CFGNLK"
		actual := part1("test-input.txt")
		assert.Equal(t, expected, actual)
	})

	t.Run("Part 1a", func(t *testing.T) {
		expected := "EVERYBODYCODES"
		actual := part1("test-input1a.txt")
		assert.Equal(t, expected, actual)
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		expected := "MGFLNK"
		actual := part2("test-input2.txt")
		assert.Equal(t, expected, actual)
	})

}

func TestPart3(t *testing.T) {

	t.Run("Part 3", func(t *testing.T) {
		expected := "DJMGL"
		actual := part3("test-input3.txt")
		assert.Equal(t, expected, actual)
	})

	t.Run("Part 3", func(t *testing.T) {
		expected := "DJCGL"
		actual := part3("test-input3a.txt")
		assert.Equal(t, expected, actual)
	})

}
