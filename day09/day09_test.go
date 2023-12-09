package day09

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartA(t *testing.T) {
	rawData := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

	lines := strings.Split(rawData, "\n")
	expected := 114

	t.Run("problem test case", func(t *testing.T) {
		result := partA(lines)
		assert.Equal(t, expected, result, "PartA(%v) should equal %d", lines, expected)
	})
}

func TestPartB(t *testing.T) {
	rawData := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

	lines := strings.Split(rawData, "\n")
	expected := 2

	t.Run("problem test case", func(t *testing.T) {
		result := partB(lines)
		assert.Equal(t, expected, result, "PartB(%v) should equal %d", lines, expected)
	})
}
