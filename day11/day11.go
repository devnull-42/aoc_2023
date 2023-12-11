package day11

import (
	"aoc/util"
	"fmt"
	"slices"
	"strings"
)

func Run() {
	lines := util.ReadInput("day11.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines, 1000000)
	fmt.Printf("partB: %d\n", result)
}

type galaxy struct {
	row int
	col int
}

func partA(lines []string) int {
	pathSum := 0

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

	for i, g1 := range galaxies {
		for _, g2 := range galaxies[i+1:] {
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
			pathSum += (rows[1] - rows[0]) + (cols[1] - cols[0]) + extraDist
		}
	}

	return pathSum
}

func partB(lines []string, larger int) int {
	pathSum := 0

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

	for i, g1 := range galaxies {
		for _, g2 := range galaxies[i+1:] {
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
			pathSum += (rows[1] - rows[0]) + (cols[1] - cols[0]) + extraDist*larger - extraDist
		}
	}

	return pathSum
}
