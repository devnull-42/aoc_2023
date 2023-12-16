package day16

import (
	"aoc/util"
	"fmt"
)

func Run() {
	lines := util.ReadInput("day16.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string) int {
	grid := buildGrid(lines)
	startingPos := pos{0, -1, 0, 1}

	energized := evalGrid(startingPos, grid)

	return energized
}

func partB(lines []string) int {
	grid := buildGrid(lines)
	startingPositions := make([]pos, 0)

	// create positions for top and bottom rows
	for i := range grid[0] {
		startingPositions = append(startingPositions, pos{-1, i, 1, 0})
		startingPositions = append(startingPositions, pos{len(grid) + 1, i, -1, 0})
	}
	// create positions for left and right columns
	for i := range grid {
		startingPositions = append(startingPositions, pos{i, -1, 0, 1})
		startingPositions = append(startingPositions, pos{i, len(grid[0]) + 1, 0, -1})
	}

	maxEnergized := 0
	for _, p := range startingPositions {
		energized := evalGrid(p, grid)
		if energized > maxEnergized {
			maxEnergized = energized
		}
	}

	return maxEnergized
}

func buildGrid(lines []string) [][]rune {
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = make([]rune, len(line))
		for j, char := range line {
			grid[i][j] = char
		}
	}
	return grid
}

type pos struct {
	r  int
	c  int
	dr int
	dc int
}

func (p pos) String() string {
	return fmt.Sprintf("(%d, %d, %d, %d)", p.r, p.c, p.dr, p.dc)
}

func (p pos) coords() string {
	return fmt.Sprintf("(%d, %d)", p.r, p.c)
}

type queue []pos

func (q *queue) push(p pos) {
	*q = append(*q, p)
}

func (q *queue) pop() pos {
	p := (*q)[0]
	*q = (*q)[1:]
	return p
}

func evalGrid(start pos, grid [][]rune) int {
	visited := make(map[string]struct{})
	coords := make(map[string]int)

	// initial position
	q := make(queue, 0)
	q.push(start)

	for len(q) > 0 {
		p := q.pop()

		// move position
		p.r += p.dr
		p.c += p.dc

		// check if position is valid
		if p.c < 0 || p.c >= len(grid[0]) || p.r < 0 || p.r >= len(grid) {
			continue
		}

		// get grid rune
		r := grid[p.r][p.c]

		// evaluate position
		switch {
		case r == '/':
			p.dr, p.dc = -p.dc, -p.dr
			if _, exists := visited[p.String()]; !exists {
				visited[p.String()] = struct{}{}
				coords[p.coords()]++
				q.push(p)
			}
		case r == '\\':
			p.dr, p.dc = p.dc, p.dr
			if _, exists := visited[p.String()]; !exists {
				visited[p.String()] = struct{}{}
				coords[p.coords()]++
				q.push(p)
			}
		case r == '|' && p.dc != 0:
			p1 := pos{p.r, p.c, p.dc, p.dr}
			p2 := pos{p.r, p.c, -p.dc, p.dr}
			if _, exists := visited[p1.String()]; !exists {
				visited[p.String()] = struct{}{}
				coords[p.coords()]++
				q.push(p1)
			}
			if _, exists := visited[p2.String()]; !exists {
				visited[p.String()] = struct{}{}
				coords[p.coords()]++
				q.push(p2)
			}
		case r == '-' && p.dr != 0:
			p1 := pos{p.r, p.c, p.dc, p.dr}
			p2 := pos{p.r, p.c, p.dc, -p.dr}
			if _, exists := visited[p1.String()]; !exists {
				visited[p.String()] = struct{}{}
				coords[p.coords()]++
				q.push(p1)
			}
			if _, exists := visited[p2.String()]; !exists {
				visited[p.String()] = struct{}{}
				coords[p.coords()]++
				q.push(p2)
			}
		default:
			if _, exists := visited[p.String()]; !exists {
				visited[p.String()] = struct{}{}
				coords[p.coords()]++
				q.push(pos{p.r, p.c, p.dr, p.dc})
			}
		}
	}
	return len(coords)
}
