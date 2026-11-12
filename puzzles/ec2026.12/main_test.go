package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		expected := 7
		actual := part1("test-input.txt")
		assert.Equal(t, expected, actual)
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		expected := 6
		actual := part2("test-input2.txt")
		assert.Equal(t, expected, actual)
	})

	t.Run("Part 2b", func(t *testing.T) {
		expected := 23
		actual := part2("test-input2b.txt")
		assert.Equal(t, expected, actual)
	})

}

func TestPart3(t *testing.T) {

	t.Run("Part 3", func(t *testing.T) {
		expected := 8
		actual := part3("test-input3.txt")
		assert.Equal(t, expected, actual)
	})

	t.Run("Part 3b", func(t *testing.T) {
		expected := 29
		actual := part3("test-input3b.txt")
		assert.Equal(t, expected, actual)
	})

	t.Run("Part 3c", func(t *testing.T) {
		expected := 16
		actual := part3("test-input3c.txt")
		assert.Equal(t, expected, actual)
	})

	// t.Run("Part 3actual", func(t *testing.T) {
	// 	expected := 16
	// 	actual := part3("input3.txt")
	// 	assert.Equal(t, expected, actual)
	// })

}
