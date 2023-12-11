package day10

import (
	"aoc/util"
	"fmt"
	"slices"
	"sync"
	"sync/atomic"

	"golang.org/x/exp/maps"
)

func Run() {
	lines := util.ReadInput("day10.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string) int {
	grid, start := createGrid(lines)
	g := createGraphFromGrid(grid)

	visited := util.BFS(g, start)
	return len(visited) / 2
}

func partB(lines []string) int {
	grid, start := createGrid(lines)
	g := createGraphFromGrid(grid)

	visited := util.BFS(g, start)
	visitedNodes := maps.Keys(visited)
	count := raycast(grid, visitedNodes)
	return int(count)
}

var pipes = map[rune]map[string][]rune{
	'S': {
		"up":    []rune{'|', '-', 'L', 'J', 'F', '7'},
		"down":  []rune{'|', '-', 'L', 'J', 'F', '7'},
		"right": []rune{'|', '-', 'L', 'J', 'F', '7'},
		"left":  []rune{'|', '-', 'L', 'J', 'F', '7'},
	},
	'|': {
		"up":    []rune{'|', 'F', '7'},
		"down":  []rune{'|', 'L', 'J'},
		"right": []rune{},
		"left":  []rune{},
	},
	'-': {
		"up":    []rune{},
		"down":  []rune{},
		"right": []rune{'-', 'J', '7'},
		"left":  []rune{'-', 'L', 'F'},
	},
	'L': {
		"up":    []rune{'|', 'F', '7'},
		"down":  []rune{},
		"right": []rune{'-', 'J', '7'},
		"left":  []rune{},
	},
	'J': {
		"up":    []rune{'|', 'F', '7'},
		"down":  []rune{},
		"right": []rune{},
		"left":  []rune{'-', 'L', 'F'},
	},
	'F': {
		"up":    []rune{},
		"down":  []rune{'|', 'L', 'J'},
		"right": []rune{'-', 'J', '7'},
		"left":  []rune{},
	},
	'7': {
		"up":    []rune{},
		"down":  []rune{'|', 'L', 'J'},
		"right": []rune{},
		"left":  []rune{'-', 'L', 'F'},
	},
}

// createGrid creates a grid from a slice of strings
func createGrid(lines []string) ([][]rune, string) {
	grid := make([][]rune, len(lines))
	var start string

	for row, line := range lines {
		grid[row] = make([]rune, len(line))
		for col, char := range line {
			switch char {
			case 'S':
				start = fmt.Sprintf("%d,%d", row, col)
				grid[row][col] = char
			default:
				grid[row][col] = char
			}
		}
	}
	return grid, start
}

// createGraphFromGrid creates a graph structure from a grid
func createGraphFromGrid(grid [][]rune) *util.Graph {
	g := util.NewGraph()

	for row, gridRow := range grid {
		for col, char := range gridRow {
			// Create and add a node for each character
			nodeName := fmt.Sprintf("%d,%d", row, col)
			currentNode := &util.Node{Name: nodeName, Value: int(char)}
			if g.GetNode(currentNode.Name) == nil {
				g.AddNode(currentNode)
			}

			type direction struct {
				dx, dy int
			}

			// directions
			directions := make(map[string]direction)
			directions["right"] = direction{0, 1} // Right
			directions["left"] = direction{0, -1} // Left
			directions["down"] = direction{1, 0}  // Down
			directions["up"] = direction{-1, 0}   // Up

			for dir, coord := range directions {
				newRow, newCol := row+coord.dx, col+coord.dy
				if char != '.' && newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(gridRow) && pipeIsValid(char, grid[newRow][newCol], dir) {
					adjacentNodeName := fmt.Sprintf("%d,%d", newRow, newCol)
					adjacentNode := &util.Node{Name: adjacentNodeName, Value: int(grid[newRow][newCol])}
					if g.GetNode(adjacentNodeName) == nil {
						g.AddNode(adjacentNode)
					}
					g.AddDirectedEdge(currentNode, adjacentNode, 0)
				}
			}
		}

	}

	return g
}

// pipeIsValid is a helper function for createGraphFromGrid that checks
// if a pipe is valid in a given direction
func pipeIsValid(pipe, adjacent rune, direction string) bool {
	for _, validPipe := range pipes[pipe][direction] {
		if validPipe == adjacent {
			return true
		}
	}
	return false
}

// raycast is a helper function for partB that counts the number of
// points in a grid that are inside the loop
func raycast(grid [][]rune, visited []string) int32 {
	var counter int32
	var wg sync.WaitGroup

	for row := range grid {
		wg.Add(1)
		// go routine to chese performance
		go func(row int) {
			defer wg.Done()
			count := raycastRow(grid, row, visited)
			atomic.AddInt32(&counter, count)
		}(row)
	}
	wg.Wait()
	return counter
}

// raycastRow is a helper function for raycast that counts the number of
// intersections in a given row and returns the count of how many are odd
func raycastRow(grid [][]rune, row int, visited []string) int32 {
	var count int32
	gridRow := grid[row]
	for col := range gridRow {
		var crosses int
		nodeName := fmt.Sprintf("%d,%d", row, col)
		if slices.Contains(visited, nodeName) {
			continue
		}
		for i := col + 1; i < len(gridRow); i++ {
			testNode := fmt.Sprintf("%d,%d", row, i)
			if slices.Contains(visited, testNode) && slices.Contains([]rune{'|', 'F', '7'}, grid[row][i]) {
				crosses++
			}
		}
		if crosses == 0 {
			break
		}
		if crosses%2 == 1 {
			count++
		}
	}
	return count
}
