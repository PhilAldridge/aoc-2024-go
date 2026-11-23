package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		expected := 5
		actual := part1("test-input.txt")
		assert.Equal(t, expected, actual)
	})

	t.Run("Part 1b", func(t *testing.T) {
		expected := 7
		actual := part1("test-inputb.txt")
		assert.Equal(t, expected, actual)
	})

	t.Run("Part 1c", func(t *testing.T) {
		expected := 9
		actual := part1("test-inputc.txt")
		assert.Equal(t, expected, actual)
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		expected := 21
		actual := part2("test-input2.txt")
		assert.Equal(t, expected, actual)
	})

}

func TestPart3(t *testing.T) {

	t.Run("Part 3", func(t *testing.T) {
		expected := 984329851511
		actual := part3("test-input3.txt")
		assert.Equal(t, expected, actual)
	})

}
