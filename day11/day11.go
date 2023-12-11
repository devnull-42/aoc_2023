package day11

import (
	"aoc/util"
	"fmt"
	"slices"
	"strings"
)

func Run() {
	lines := util.ReadInput("day11.txt")
	result := partA(lines, 2)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines, 1000000)
	fmt.Printf("partB: %d\n", result)
}

type galaxy struct {
	row int
	col int
}

func partA(lines []string, growthMultiple int) int {
	var pathSum int

	galaxies, emptyRows, emptyCols := parseLines(lines)

	for i, g1 := range galaxies {
		dist := evaluateGalaxy(g1, galaxies[i+1:], emptyRows, emptyCols, growthMultiple)
		pathSum += dist
	}
	return pathSum
}

func partB(lines []string, growthMultiple int) int {
	var pathSum int

	galaxies, emptyRows, emptyCols := parseLines(lines)

	for i, g1 := range galaxies {
		dist := evaluateGalaxy(g1, galaxies[i+1:], emptyRows, emptyCols, growthMultiple)
		pathSum += dist
	}
	return pathSum
}

// parseLines creates a slice of galaxies and a map of empty rows and columns
func parseLines(lines []string) ([]galaxy, map[int]struct{}, map[int]struct{}) {
	emptyRows := make(map[int]struct{})
	emptyCols := make(map[int]struct{})
	for i := 0; i < len(lines[0]); i++ {
		emptyCols[i] = struct{}{}
	}
	galaxies := make([]galaxy, 0)

	for i, line := range lines {
		// if there are no '#'s in this row, add another row of all 0's
		if !slices.Contains(strings.Split(line, ""), "#") {
			emptyRows[i] = struct{}{}
		}
		for j, char := range line {
			switch char {
			case '#':
				galaxies = append(galaxies, galaxy{i, j})
				// if there is a '#' in this column, remove it from emptyCols
				delete(emptyCols, j)
			case '.':
				continue
			default:
				panic(fmt.Sprintf("unexpected character: %s", string(char)))
			}
		}
	}
	return galaxies, emptyRows, emptyCols
}

// evaluateGalaxy calculates the distance between a galaxy and all galaxies after it
// it uses the growth multiple to increase the distance for each empty row or column
func evaluateGalaxy(g1 galaxy, galaxies []galaxy, emptyRows, emptyCols map[int]struct{}, growthMultiple int) int {
	pathSum := 0
	for _, g2 := range galaxies {
		extraDist := 0
		// check if there are any emptyRows between g1 and g2
		rows := []int{g1.row, g2.row}
		slices.Sort(rows)
		for r := rows[0] + 1; r < rows[1]; r++ {
			if _, ok := emptyRows[r]; ok {
				extraDist++
			}
		}
		// check if there are any emptyCols between g1 and g2
		cols := []int{g1.col, g2.col}
		slices.Sort(cols)
		for c := cols[0] + 1; c < cols[1]; c++ {
			if _, ok := emptyCols[c]; ok {
				extraDist++
			}
		}
		pathSum += (rows[1] - rows[0]) + (cols[1] - cols[0]) + extraDist*growthMultiple - extraDist
	}
	return pathSum
}
