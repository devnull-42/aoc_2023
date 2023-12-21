package day21

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartA(t *testing.T) {
	rawData := `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

	lines := strings.Split(rawData, "\n")
	steps := 6
	expected := 16

	t.Run("problem test case", func(t *testing.T) {
		result := partA(lines, steps)
		assert.Equal(t, expected, result, "PartA(%v) should equal %d", lines, expected)
	})
}

func TestPartB(t *testing.T) {
	rawData := `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`
	lines := strings.Split(rawData, "\n")

	var testData = []struct {
		steps    int
		expected int
	}{
		{
			steps:    6,
			expected: 16,
		},
		{
			steps:    10,
			expected: 50,
		},
		{
			steps:    50,
			expected: 1594,
		},
		{
			steps:    100,
			expected: 6536,
		},
		{
			steps:    500,
			expected: 167004,
		},
		{
			steps:    1000,
			expected: 668697,
		},
		{
			steps:    5000,
			expected: 16733044,
		},
	}

	for _, tt := range testData {
		t.Run("problem test case", func(t *testing.T) {
			result := partB(lines, tt.steps)
			assert.Equal(t, tt.expected, result, "PartB(%v) should equal %d", lines, tt.expected)
		})
	}
}
