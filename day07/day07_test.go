package day07

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartA(t *testing.T) {
	rawData := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

	lines := strings.Split(rawData, "\n")
	expected := 6440

	t.Run("problem test case", func(t *testing.T) {
		result := partA(lines)
		assert.Equal(t, expected, result, "PartA(%v) should equal %d", lines, expected)
	})
}

func TestPartB(t *testing.T) {
	rawData := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

	lines := strings.Split(rawData, "\n")
	expected := 5905

	t.Run("problem test case", func(t *testing.T) {
		result := partB(lines)
		assert.Equal(t, expected, result, "PartB(%v) should equal %d", lines, expected)
	})
}
