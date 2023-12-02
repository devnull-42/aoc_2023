package day01

import (
	"strings"
	"testing"
)

func TestPartA(t *testing.T) {
	rawData := `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`

	lines := strings.Split(rawData, "\n")
	expected := 142

	t.Run("problem test case", func(t *testing.T) {
		result := partA(lines)
		if result != expected {
			t.Errorf("PartA(%v) = %d; want %d", lines, result, expected)
		}
	})
}

func TestPartB(t *testing.T) {
	rawData := `two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`

	lines := strings.Split(rawData, "\n")
	expected := 281

	t.Run("problem test case", func(t *testing.T) {
		result := partB(lines)
		if result != expected {
			t.Errorf("PartB(%v) = %d; want %d", lines, result, expected)
		}
	})
}
