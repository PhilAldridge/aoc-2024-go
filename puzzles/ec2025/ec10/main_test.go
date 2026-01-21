package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		expected := 27
		actual := part1("test-input.txt",3)
		assert.Equal(t, expected, actual)
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		expected := 27
		actual := part2("test-input2.txt",3)
		assert.Equal(t, expected, actual)
	})

}

func TestPart3(t *testing.T) {

	t.Run("Part 3", func(t *testing.T) {
		expected := 15
		actual := part3("test-input3.txt")
		assert.Equal(t, expected, actual)
	})

	t.Run("Part 3", func(t *testing.T) {
		expected := 8
		actual := part3("test-input3a.txt")
		assert.Equal(t, expected, actual)
	})

	t.Run("Part 3", func(t *testing.T) {
		expected := 44
		actual := part3("test-input3b.txt")
		assert.Equal(t, expected, actual)
	})

	t.Run("Part 3", func(t *testing.T) {
		expected := 4406
		actual := part3("test-input3c.txt")
		assert.Equal(t, expected, actual)
	})

	t.Run("Part 3", func(t *testing.T) {
		expected := 13033988838
		actual := part3("test-input3d.txt")
		assert.Equal(t, expected, actual)
	})

}
