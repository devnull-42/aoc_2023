package day15

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartA(t *testing.T) {
	rawData := `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

	lines := strings.Split(rawData, "\n")
	expected := 1320

	t.Run("problem test case", func(t *testing.T) {
		result := partA(lines)
		assert.Equal(t, expected, result, "PartA(%v) should equal %d", lines, expected)
	})
}

func TestPartB(t *testing.T) {
	rawData := `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

	lines := strings.Split(rawData, "\n")
	expected := 145

	t.Run("problem test case", func(t *testing.T) {
		result := partB(lines)
		assert.Equal(t, expected, result, "PartB(%v) should equal %d", lines, expected)
	})
}
