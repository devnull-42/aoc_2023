package day13

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartA(t *testing.T) {
	rawData := `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

	lines := strings.Split(rawData, "\n")
	expected := 405

	t.Run("problem test case", func(t *testing.T) {
		result := partA(lines)
		assert.Equal(t, expected, result, "PartA(%v) should equal %d", lines, expected)
	})
}

func TestPartB(t *testing.T) {
	rawData := `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

	lines := strings.Split(rawData, "\n")
	expected := 400

	t.Run("problem test case", func(t *testing.T) {
		result := partB(lines)
		assert.Equal(t, expected, result, "PartB(%v) should equal %d", lines, expected)
	})
}
