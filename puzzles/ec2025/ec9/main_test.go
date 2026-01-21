package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		expected := 414
		actual := part1("test-input.txt")
		assert.Equal(t, expected, actual)
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		expected := 1245
		actual := part2("test-input2.txt")
		assert.Equal(t, expected, actual)
	})

}

func TestPart3(t *testing.T) {

	t.Run("Part 3", func(t *testing.T) {
		expected := 12
		actual := part3("test-input3.txt")
		assert.Equal(t, expected, actual)
	})

	t.Run("Part 3a", func(t *testing.T) {
		expected := 36
		actual := part3("test-input3a.txt")
		assert.Equal(t, expected, actual)
	})

}
