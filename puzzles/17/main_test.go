package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		expected := "4,6,3,5,6,3,5,2,1,0"
		actual := part1("test-input.txt")
		assert.Equal(t, expected, actual)
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		expected := 117440
		actual := part2("test-input.txt")
		assert.Equal(t, expected, actual)
	})

}
