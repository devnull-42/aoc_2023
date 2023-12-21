package day21

import (
	"aoc/util"
	"fmt"
)

func Run() {
	lines := util.ReadInput("day21.txt")
	result := partA(lines, 64)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines, 26501365)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string, steps int) int {
	grid, start := getGridAndStart(lines)
	result := modifiedBFS(grid, start, steps)

	return len(result)
}

func partB(lines []string, steps int) int {
	grid, start := getGridAndStart(lines)
	result := modifiedBFS_B(grid, start, steps)

	return len(result)
}

// get the grid and start location
func getGridAndStart(lines []string) ([][]string, Location) {
	var start Location
	grid := make([][]string, len(lines))
	for r, line := range lines {
		grid[r] = make([]string, len(line))
		for c, char := range line {
			grid[r][c] = string(char)
			if char == 'S' {
				start = Location{r, c}
			}
		}
	}
	return grid, start
}

type Location struct {
	row int
	col int
}

func (loc Location) String() string {
	return fmt.Sprintf("%d,%d", loc.row, loc.col)
}

func resizeLoc(loc Location, rowSize, colSize int) Location {
	newRow := loc.row % rowSize
	newCol := loc.col % colSize
	if newRow < 0 {
		newRow += rowSize
	}
	if newCol < 0 {
		newCol += colSize
	}
	return Location{newRow, newCol}
}

func modifiedBFS(g [][]string, start Location, maxSteps int) map[string]struct{} {
	type queueItem struct {
		loc   Location
		steps int
	}

	queue := []queueItem{{start, maxSteps}}
	visited := make(map[string]struct{})
	result := make(map[string]struct{})

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		loc, steps := queue[0].loc, queue[0].steps
		queue = queue[1:]

		if steps%2 == 0 {
			result[loc.String()] = struct{}{}
		}
		if steps == 0 {
			continue
		}

		for _, dir := range directions {
			nextLoc := Location{loc.row + dir[0], loc.col + dir[1]}
			if nextLoc.row < 0 || nextLoc.row >= len(g) || nextLoc.col < 0 || nextLoc.col >= len(g[0]) || g[nextLoc.row][nextLoc.col] == "#" {
				continue
			}
			if _, exists := visited[nextLoc.String()]; !exists {
				visited[nextLoc.String()] = struct{}{}
				queue = append(queue, queueItem{nextLoc, steps - 1})
			}
		}
	}
	return result
}

func modifiedBFS_B(g [][]string, start Location, maxSteps int) map[string]struct{} {
	type queueItem struct {
		loc   Location
		steps int
	}

	queue := []queueItem{{start, maxSteps}}
	visited := make(map[string]struct{})
	result := make(map[string]struct{})

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		loc, steps := queue[0].loc, queue[0].steps
		queue = queue[1:]

		if steps%2 == 0 {
			result[loc.String()] = struct{}{}
		}
		if steps == 0 {
			continue
		}

		for _, dir := range directions {
			nextLoc := Location{loc.row + dir[0], loc.col + dir[1]}
			resizedLoc := resizeLoc(nextLoc, len(g), len(g[0]))
			if g[resizedLoc.row][resizedLoc.col] == "#" {
				continue
			}
			if _, exists := visited[nextLoc.String()]; !exists {
				visited[nextLoc.String()] = struct{}{}
				queue = append(queue, queueItem{nextLoc, steps - 1})
			}
		}
	}
	return result
}
