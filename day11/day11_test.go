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

	lines := strings.Split(rawData, "\n")
	expected := 374

	t.Run("problem test case", func(t *testing.T) {
		result := partA(lines)
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

	larger := 100
	lines := strings.Split(rawData, "\n")
	expected := 8410

	t.Run("problem test case", func(t *testing.T) {
		result := partB(lines, larger)
		assert.Equal(t, expected, result, "PartB(%v) should equal %d", lines, expected)
	})
}
