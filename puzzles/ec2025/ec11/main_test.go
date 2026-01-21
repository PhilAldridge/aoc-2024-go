package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		expected := 109
		actual := part1("test-input.txt")
		assert.Equal(t, expected, actual)
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		expected := 11
		actual := part2("test-input.txt")
		assert.Equal(t, expected, actual)
	})

	t.Run("Part 2", func(t *testing.T) {
		expected := 1579
		actual := part2("test-input2.txt")
		assert.Equal(t, expected, actual)
	})

	t.Run("Part 2 actual", func(t *testing.T) {
		expected := 3586398
		actual := part2("input2.txt")
		assert.Equal(t, expected, actual)
	})

}
