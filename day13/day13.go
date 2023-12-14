package day13

import (
	"aoc/util"
	"fmt"
	"slices"
	"strings"

	"golang.org/x/exp/maps"
)

func Run() {
	lines := util.ReadInput("day13.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string) int {
	patterns := getPatterns(lines)
	var sum int

	for _, pattern := range patterns {
		// check for vertical mirror
		match, idx := checkMirror(pattern)
		if match {
			sum += idx[0]
		}
		// check for horizontal mirror
		match, idx = checkMirror(invertGrid(pattern))
		if match {
			sum += idx[0] * 100
		}

	}

	return sum
}

func partB(lines []string) int {
	patterns := getPatterns(lines)
	var sum int

	for _, pattern := range patterns {
		vertReflect := make(map[int]struct{})
		horizReflect := make(map[int]struct{})
		originalReflect := make([]int, 2)
		// check for original vertical mirror
		match, idx := checkMirror(pattern)
		if match {
			originalReflect[0] = idx[0]
		}
		// check for original horizontal mirror
		match, idx = checkMirror(invertGrid(pattern))
		if match {
			originalReflect[1] = idx[0]
		}
		for r, row := range pattern {
			for c := range row {
				swapChars(pattern, r, c)
				// check for vertical mirror
				match, idx := checkMirror(pattern)
				if match {
					for i := range idx {
						if idx[i] != originalReflect[0] {
							vertReflect[idx[i]] = struct{}{}
						}
					}
				}
				// check for horizontal mirror
				match, idx = checkMirror(invertGrid(pattern))
				if match {
					for i := range idx {
						if idx[i] != originalReflect[1] {
							horizReflect[idx[i]] = struct{}{}
						}
					}
				}
				swapChars(pattern, r, c)
			}
		}
		if len(vertReflect) > 0 {
			sum += maps.Keys(vertReflect)[0]
		}
		if len(horizReflect) > 0 {
			sum += maps.Keys(horizReflect)[0] * 100
		}
	}

	return sum
}

func getPatterns(lines []string) [][][]string {
	var patterns [][][]string
	var pattern [][]string
	for _, line := range lines {
		if line == "" {
			patterns = append(patterns, pattern)
			pattern = nil
		} else {
			pattern = append(pattern, strings.Split(line, ""))
		}
	}
	patterns = append(patterns, pattern)
	return patterns
}

func checkMirror(pattern [][]string) (bool, []int) {
	columns := make(map[int]struct{})
	for i := 1; i < len(pattern[0]); i++ {
		columns[i] = struct{}{}
	}

	for _, line := range pattern {
		for i := range columns {
			if i <= len(line)/2 {
				l := slices.Clone(line[:i])
				r := line[i : i*2]

				slices.Reverse(l)

				if !slices.Equal(l, r) {
					delete(columns, i)
				}
			} else {
				l := slices.Clone(line[i-(len(line)-i) : i])
				r := line[i:]

				slices.Reverse(l)

				if !slices.Equal(l, r) {
					delete(columns, i)
				}
			}
		}
		if len(columns) == 0 {
			return false, nil
		}
	}

	return true, maps.Keys(columns)
}

func invertGrid(pattern [][]string) [][]string {
	inverted := make([][]string, len(pattern[0]))

	for i := 0; i < len(pattern[0]); i++ {
		inverted[i] = make([]string, len(pattern))
	}

	for i, line := range pattern {
		for j, char := range line {
			inverted[j][i] = char
		}
	}

	return inverted
}

func swapChars(pattern [][]string, r, c int) {
	switch pattern[r][c] {
	case "#":
		pattern[r][c] = "."
	case ".":
		pattern[r][c] = "#"
	}
}
