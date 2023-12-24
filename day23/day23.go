package day23

import (
	"aoc/util"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/emirpasic/gods/queues/arrayqueue"
)

func Run() {
	lines := util.ReadInput("day23.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string) int {
	grid, start, end := getGrid(lines)
	result := dfs(grid, start.r, start.c, end.r, end.c)
	// result := parallelDFS(grid, start, end)

	return result
}

func partB(lines []string) int {
	// grid, start, end := getGrid(lines)
	result := SolvePart2(lines)

	return result
}

func getGrid(lines []string) ([][]rune, Point, Point) {
	grid := make([][]rune, len(lines))
	var start, end Point

	for i, line := range lines {
		grid[i] = make([]rune, len(line))
		for j, ch := range line {
			if i == 0 && ch == '.' {
				start = Point{i, j}
			} else if i == len(lines)-1 && ch == '.' {
				end = Point{i, j}
			}
			grid[i][j] = ch
		}
	}
	return grid, start, end
}

type Point struct {
	r, c int
}

var directions = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // Right, Down, Left, Up

func isValid(grid [][]rune, r, c int) bool {
	return r >= 0 && c >= 0 && r < len(grid) && c < len(grid[0]) && grid[r][c] != '#'
}

func dfs(grid [][]rune, r, c, endR, endC int) int {
	if r == endR && c == endC {
		return 0 // Reached the end point
	}

	original := grid[r][c]
	grid[r][c] = '#' // Mark as visited
	maxPath := -1

	for _, dir := range directions {
		nx, ny := r+dir[0], c+dir[1]
		if isValid(grid, nx, ny) {
			pathLength := -1
			switch original {
			case '.':
				pathLength = dfs(grid, nx, ny, endR, endC)
			case '>':
				right := [2]int{0, 1}
				if dir == right {
					pathLength = dfs(grid, nx, ny, endR, endC)
				}
			case '<':
				left := [2]int{0, -1}
				if dir == left {
					pathLength = dfs(grid, nx, ny, endR, endC)
				}
			case '^':
				up := [2]int{-1, 0}
				if dir == up {
					pathLength = dfs(grid, nx, ny, endR, endC)
				}
			case 'v':
				down := [2]int{1, 0}
				if dir == down {
					pathLength = dfs(grid, nx, ny, endR, endC)
				}
			}

			if pathLength >= 0 {
				pathLength++ // Increment path length
				if pathLength > maxPath {
					maxPath = pathLength // Update max path if longer path found
				}
			}
		}
	}

	grid[r][c] = original // Restore original value

	return maxPath
}

// this solution is based on https://github.com/vipul0092/advent-of-code-2023/blob/main/day23/day23.go
// fought with this one for too long. This is a good solution to study and uses the packages
// gods/queues/arrayqueue and deckarep/golang-set/v2
// gods/queues/arrayqueue provides mplementations of various data structures and algorithms
// deckarep/golang-set/v2 provides a set implementation using generics that functions like python
// this one hard codes assumptions about the input data and used some global variables.

type DP [142][142][142][142]int32
type Graph [142][142]int

type Pwl struct {
	p   Point
	len int
}

var graph Graph

func SolvePart2(lines []string) int {
	rows, cols, id := len(lines), len(lines[0]), 3
	var maxid int

	adjacencyList := make(map[Point][]Point)

	compressedGraphNodes := make(map[Point]int)
	compressedGraphNodes[Point{0, 1}] = 1
	compressedGraphNodes[Point{rows - 1, cols - 2}] = 2
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !valid(i, j, lines) {
				continue
			}
			pt := Point{i, j}
			adjacencyList[pt] = make([]Point, 0)
			for _, df := range directions {
				ni, nj := i+df[0], j+df[1]
				if valid(ni, nj, lines) {
					adjacencyList[pt] = append(adjacencyList[pt], Point{ni, nj})
				}
			}

			if len(adjacencyList[pt]) > 2 {
				compressedGraphNodes[pt] = id
				id++
			}
		}
	}
	maxid = id

	graph = Graph{}
	for pt := range compressedGraphNodes {
		populateGraphForNode(pt, compressedGraphNodes, adjacencyList)
	}
	return dfs2(1, 1, 2, 0, maxid)
}

func dfs2(current, prev, end int, visited int, maxid int) int {
	if current == end {
		return 0
	}
	maxi := -1
	for j := 1; j < maxid; j++ {
		if graph[current][j] != 0 && j != prev && !isBitSet(visited, j) {
			d := dfs2(j, current, end, setBit(visited, j), maxid)
			if d != -1 {
				maxi = max(maxi, d+graph[current][j])
			}
		}
	}
	return maxi
}

func isBitSet(bitSet, bitPos int) bool {
	return (bitSet & (1 << bitPos)) != 0
}

func setBit(bitSet, bitPos int) int {
	return bitSet | (1 << bitPos)
}

func populateGraphForNode(point Point, compressedGraphNodes map[Point]int, adjacencyList map[Point][]Point) {
	id, queue, visited := compressedGraphNodes[point], arrayqueue.New(), mapset.NewSet[Point]()
	queue.Enqueue(Pwl{point, 1})
	visited.Add(point)

	for !queue.Empty() {
		pwl, _ := queue.Dequeue()
		p, distance := pwl.(Pwl).p, pwl.(Pwl).len
		for _, neighbor := range adjacencyList[p] {
			if !visited.Contains(neighbor) {
				if nid, has := compressedGraphNodes[neighbor]; has {
					graph[id][nid] = distance
					graph[nid][id] = distance
				} else {
					queue.Enqueue(Pwl{neighbor, distance + 1})
				}
				visited.Add(neighbor)
			}
		}
	}
}

func valid(i, j int, lines []string) bool {
	return i >= 0 && j >= 0 && i < len(lines) && j < len(lines[0]) && lines[i][j] != '#'
}
