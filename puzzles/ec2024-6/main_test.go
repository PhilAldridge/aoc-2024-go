package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		expected := "RRB@"
		actual := part1("test-input.txt")
		assert.Equal(t, expected, actual)
	})

}
