package day25

import (
	"aoc/util"
	"fmt"
	"math/rand"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
)

func Run() {
	lines := util.ReadInput("day25.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string) int {
	V, E := ParseInput(lines)
	return findSections(V, E)
}

func partB(lines []string) int {
	return 0
}

type Edge struct {
	U, V string
}

func ParseInput(lines []string) (mapset.Set[string], mapset.Set[Edge]) {
	V := mapset.NewSet[string]()
	E := mapset.NewSet[Edge]()

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		v := parts[0]
		V.Add(v)
		connectedNodes := strings.Split(parts[1], " ")
		for _, n := range connectedNodes {
			V.Add(n)
			E.Add(Edge{U: v, V: n})
		}
	}
	return V, E
}

// This solution uses Kruskal's algorithm to find the number of connected components: https://en.wikipedia.org/wiki/Karger%27s_algorithm
// It was implemented in this python solution for the same problem:
// https://www.reddit.com/r/adventofcode/comments/18qbsxs/comment/ketzp94/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
// I ported it to Go and used the golang-set library to make it easier to work with sets.
func findSections(V mapset.Set[string], E mapset.Set[Edge]) int {
	rand.Seed(time.Now().UnixNano())

	ss := func(v string, subsets []mapset.Set[string]) mapset.Set[string] {
		for _, s := range subsets {
			if s.Contains(v) {
				return s
			}
		}
		return nil
	}

	var subsets []mapset.Set[string]
	for {
		subsets = make([]mapset.Set[string], 0)
		for v := range V.Iterator().C {
			subset := mapset.NewSet[string]()
			subset.Add(v)
			subsets = append(subsets, subset)
		}

		for len(subsets) > 2 {
			var chosenEdge Edge
			for e := range E.Iterator().C {
				chosenEdge = e
				break
			}
			s1, s2 := ss(chosenEdge.U, subsets), ss(chosenEdge.V, subsets)
			if !s1.Equal(s2) {
				merged := s1.Union(s2)
				// Remove s1 and s2 from subsets and add merged
				newSubsets := make([]mapset.Set[string], 0)
				for _, s := range subsets {
					if !s.Equal(s1) && !s.Equal(s2) {
						newSubsets = append(newSubsets, s)
					}
				}
				newSubsets = append(newSubsets, merged)
				subsets = newSubsets
			}
		}

		edgeCount := 0
		for edge := range E.Iterator().C {
			if ss(edge.U, subsets) != ss(edge.V, subsets) {
				edgeCount++
			}
		}
		if edgeCount < 4 {
			break
		}
	}

	result := 1
	for _, s := range subsets {
		result *= s.Cardinality()
	}
	return result
}
