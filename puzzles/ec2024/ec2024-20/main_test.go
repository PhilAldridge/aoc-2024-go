package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		expected := 1045
		actual := part1("test-input.txt")
		assert.Equal(t, expected, actual)
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		expected := 24
		actual := part2("test-input2.txt")
		assert.Equal(t, expected, actual)
	})

	t.Run("Part 2a", func(t *testing.T) {
		expected := 78
		actual := part2("test-input2a.txt")
		assert.Equal(t, expected, actual)
	})

	t.Run("Part 2b", func(t *testing.T) {
		expected := 206
		actual := part2("test-input2b.txt")
		assert.Equal(t, expected, actual)
	})

}

func TestPart3(t *testing.T) {

	t.Run("Part 3", func(t *testing.T) {
		expected := 768790
		actual := part3("test-input3.txt")
		assert.Equal(t, expected, actual)
	})

}
