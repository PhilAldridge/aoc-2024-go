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

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		expected := 11
		actual := part2("test-input.txt")
		assert.Equal(t, expected, actual)
	})

}

func TestPart3(t *testing.T) {

	t.Run("Part 3", func(t *testing.T) {
		expected := 34
		actual := part3("test-input3.txt",1,10)
		assert.Equal(t, expected, actual)
	})

	t.Run("Part 3", func(t *testing.T) {
		expected := 72
		actual := part3("test-input3.txt",2,10)
		assert.Equal(t, expected, actual)
	})

	t.Run("Part 3", func(t *testing.T) {
		expected := 3442321
		actual := part3("test-input3.txt",1000,1000)
		assert.Equal(t, expected, actual)
	})

}
