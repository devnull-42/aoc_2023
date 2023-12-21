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

	return result
}

func partB(lines []string, steps int) int {
	grid, start := getGridAndStart(lines)
	result := partBSolution(grid, start, steps)

	return result
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

func modifiedBFS(g [][]string, start Location, maxSteps int) int {
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
	return len(result)
}

// This is a solution by HyperNeutrino. It is what I was trying to do but couldn't get there.
// It doesn't work with the test data, but it does work with the real data. So I'm sure there is a better solution.
// There is a solution that uses the quadratic formula to solve the problem, but it's
// not what I was trying to do. I modified it to use my functions and structs since his solution
// was in python.
func partBSolution(grid [][]string, start Location, maxSteps int) int {
	gridSize := len(grid)
	gridWidth := maxSteps/gridSize - 1

	odd := (gridWidth/2*2 + 1) * (gridWidth/2*2 + 1)
	even := ((gridWidth + 1) / 2 * 2) * ((gridWidth + 1) / 2 * 2)

	oddResults := modifiedBFS(grid, start, gridSize*2+1)
	evenResults := modifiedBFS(grid, start, gridSize*2)

	topCorner := modifiedBFS(grid, Location{gridSize - 1, start.col}, gridSize-1)
	bottomCorner := modifiedBFS(grid, Location{0, start.col}, gridSize-1)
	rightCorner := modifiedBFS(grid, Location{start.row, 0}, gridSize-1)
	leftCorner := modifiedBFS(grid, Location{start.row, gridSize - 1}, gridSize-1)

	smallEdgeTopRight := modifiedBFS(grid, Location{gridSize - 1, 0}, gridSize/2-1)
	smallEdgeBottomRight := modifiedBFS(grid, Location{0, 0}, gridSize/2-1)
	smallEdgeTopLeft := modifiedBFS(grid, Location{gridSize - 1, gridSize - 1}, gridSize/2-1)
	smallEdgeBottomLeft := modifiedBFS(grid, Location{0, gridSize - 1}, gridSize/2-1)

	bigEdgeTopRight := modifiedBFS(grid, Location{gridSize - 1, 0}, gridSize*3/2-1)
	bigEdgeBottomRight := modifiedBFS(grid, Location{0, 0}, gridSize*3/2-1)
	bigEdgeTopLeft := modifiedBFS(grid, Location{gridSize - 1, gridSize - 1}, gridSize*3/2-1)
	bigEdgeBottomLeft := modifiedBFS(grid, Location{0, gridSize - 1}, gridSize*3/2-1)

	result := odd*oddResults +
		even*evenResults +
		topCorner + bottomCorner + rightCorner + leftCorner +
		(gridWidth+1)*(smallEdgeBottomLeft+smallEdgeBottomRight+smallEdgeTopLeft+smallEdgeTopRight) +
		gridWidth*(bigEdgeBottomLeft+bigEdgeBottomRight+bigEdgeTopLeft+bigEdgeTopRight)

	return result
}
