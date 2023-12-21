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
