package day22

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartA(t *testing.T) {
	rawData := `1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`

	lines := strings.Split(rawData, "\n")
	expected := 5

	t.Run("problem test case", func(t *testing.T) {
		result := partA(lines)
		assert.Equal(t, expected, result, "PartA(%v) should equal %d", lines, expected)
	})
}

func TestPartB(t *testing.T) {
	rawData := `1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`

	lines := strings.Split(rawData, "\n")
	expected := 7

	t.Run("problem test case", func(t *testing.T) {
		result := partB(lines)
		assert.Equal(t, expected, result, "PartB(%v) should equal %d", lines, expected)
	})
}
