package day11

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartA(t *testing.T) {
	rawData := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

	growthMultiple := 2
	lines := strings.Split(rawData, "\n")
	expected := 374

	t.Run("problem test case", func(t *testing.T) {
		result := partA(lines, growthMultiple)
		assert.Equal(t, expected, result, "PartA(%v) should equal %d", lines, expected)
	})
}

func TestPartB(t *testing.T) {
	rawData := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

	growthMultiple := 100
	lines := strings.Split(rawData, "\n")
	expected := 8410

	t.Run("problem test case", func(t *testing.T) {
		result := partB(lines, growthMultiple)
		assert.Equal(t, expected, result, "PartB(%v) should equal %d", lines, expected)
	})
}
