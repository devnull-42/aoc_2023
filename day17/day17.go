package day17

import (
	"aoc/util"
	"container/heap"
	"fmt"
)

func Run() {
	lines := util.ReadInput("day17.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string) int {
	// build grid from input
	grid := buildGrid(lines)

	// niitialize priority queue
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// set initial value
	heap.Push(&pq, &Node{0, 0, 0, 0, 0, 0, 0})

	// initialize visited map
	visited := make(map[string]struct{})

	// traverse grid
	for pq.Len() > 0 {
		// pop node from queue
		node := heap.Pop(&pq).(*Node)

		// check if we've reached the bottom right corner
		if node.posRow == len(grid)-1 && node.posCol == len(grid[0])-1 {
			return node.heatLoss
		}

		// check if we've visited this node before
		if _, exists := visited[node.String()]; exists {
			continue
		}

		// add node to visited map
		visited[node.String()] = struct{}{}

		// check if we can move in any direction
		for _, dir := range directrions {
			// calculate new position
			newRow := node.posRow + dir[0]
			newCol := node.posCol + dir[1]

			// check if new position is valid
			if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) {
				continue
			}

			// cannot go in the same direction more than 3 times in a row
			if node.consecutiveSteps == 3 && (node.dRow == dir[0] || node.dCol == dir[1]) {
				continue
			}

			// cannot go in reverse direction
			if node.dRow == -dir[0] && node.dCol == -dir[1] {
				continue
			}

			// calculate new heat loss
			newHeatLoss := node.heatLoss + grid[newRow][newCol]

			// reset or increment consecutive steps
			newConsecutiveSteps := node.consecutiveSteps
			if node.dRow == dir[0] && node.dCol == dir[1] {
				newConsecutiveSteps++
			} else {
				newConsecutiveSteps = 1
			}

			// add new node to queue
			heap.Push(&pq, &Node{
				posRow:           newRow,
				posCol:           newCol,
				dRow:             dir[0],
				dCol:             dir[1],
				consecutiveSteps: newConsecutiveSteps,
				heatLoss:         newHeatLoss,
			})
		}

	}
	return 0
}

func partB(lines []string) int {
	// build grid from input
	grid := buildGrid(lines)

	// niitialize priority queue
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// set initial value
	heap.Push(&pq, &Node{0, 0, 0, 0, 0, 0, 0})

	// initialize visited map
	visited := make(map[string]struct{})

	// traverse grid
	for pq.Len() > 0 {
		// pop node from queue
		node := heap.Pop(&pq).(*Node)

		// check if we've reached the bottom right corner
		if node.posRow == len(grid)-1 && node.posCol == len(grid[0])-1 && node.consecutiveSteps >= 4 {
			return node.heatLoss
		}

		// check if we've visited this node before
		if _, exists := visited[node.String()]; exists {
			continue
		}

		// add node to visited map
		visited[node.String()] = struct{}{}

		// check if we can move in any direction
		for _, dir := range directrions {
			// calculate new position
			newRow := node.posRow + dir[0]
			newCol := node.posCol + dir[1]

			// check if new position is valid
			if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) {
				continue
			}

			// must go in same direction at least 4 times in a row except for the starting position
			if node.consecutiveSteps < 4 && (node.dRow != dir[0] || node.dCol != dir[1]) && !(node.dRow == 0 && node.dCol == 0) {
				continue
			}

			// cannot go in the same direction more than 10 times in a row
			if node.consecutiveSteps >= 10 && (node.dRow == dir[0] || node.dCol == dir[1]) {
				continue
			}

			// cannot go in reverse direction
			if node.dRow == -dir[0] && node.dCol == -dir[1] {
				continue
			}

			// calculate new heat loss
			newHeatLoss := node.heatLoss + grid[newRow][newCol]

			// reset or increment consecutive steps
			newConsecutiveSteps := node.consecutiveSteps
			if node.dRow == dir[0] && node.dCol == dir[1] {
				newConsecutiveSteps++
			} else {
				newConsecutiveSteps = 1
			}

			// add new node to queue
			heap.Push(&pq, &Node{
				posRow:           newRow,
				posCol:           newCol,
				dRow:             dir[0],
				dCol:             dir[1],
				consecutiveSteps: newConsecutiveSteps,
				heatLoss:         newHeatLoss,
			})
		}

	}
	return 0
}

func buildGrid(lines []string) [][]int {
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, c := range line {
			grid[i][j] = util.MustAtoi(string(c))
		}
	}
	return grid
}

var directrions = [4][2]int{
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
	{-1, 0}, // up
}

type Node struct {
	posRow           int
	posCol           int
	dRow             int
	dCol             int
	consecutiveSteps int
	heatLoss         int
	index            int
}

func (n *Node) String() string {
	return fmt.Sprintf("%d-%d-%d-%d-%d",
		n.posRow, n.posCol, n.dRow, n.dCol, n.consecutiveSteps)
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].heatLoss < pq[j].heatLoss
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	node := old[n-1]
	node.index = -1
	*pq = old[0 : n-1]
	return node
}
