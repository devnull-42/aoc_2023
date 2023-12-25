package day24

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartA(t *testing.T) {
	rawData := `19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3`

	lines := strings.Split(rawData, "\n")
	expected := 2

	t.Run("problem test case", func(t *testing.T) {
		result := partA(lines, 7, 27)
		assert.Equal(t, expected, result, "PartA(%v) should equal %d", lines, expected)
	})
}

func TestPartB(t *testing.T) {
	rawData := `19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3`

	lines := strings.Split(rawData, "\n")
	expected := 0

	t.Run("problem test case", func(t *testing.T) {
		result := partB(lines, 7, 27)
		assert.Equal(t, expected, result, "PartB(%v) should equal %d", lines, expected)
	})
}
