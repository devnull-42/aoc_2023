package day08

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strings"
)

func Run() {
	lines := util.ReadInput("day08.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string) int {
	nodeMap := make(map[string][2]string)
	reNode := regexp.MustCompile(`^(\w{3}) = \((\w{3}), (\w{3})\)$`)

	for _, line := range lines {
		matches := reNode.FindStringSubmatch(line)
		if len(matches) == 4 {
			nodeMap[matches[1]] = [2]string{matches[2], matches[3]}
		}
	}

	instructions := strings.Split(lines[0], "")

	var steps int
	node := "AAA"

	for node != "ZZZ" {
		switch instructions[steps%len(instructions)] {
		case "R":
			node = nodeMap[node][1]
		case "L":
			node = nodeMap[node][0]
		default:
			panic("invalid instruction")
		}
		steps++
	}
	return steps
}

type node struct {
	name  string
	steps int
}

func partB(lines []string) int {
	nodeMap := make(map[string][2]string)
	reNode := regexp.MustCompile(`^(\w{3}) = \((\w{3}), (\w{3})\)$`)

	var nodes []node

	for _, line := range lines {
		matches := reNode.FindStringSubmatch(line)
		if len(matches) == 4 {
			nodeMap[matches[1]] = [2]string{matches[2], matches[3]}
			if matches[1][2] == 'A' {
				nodes = append(nodes, node{matches[1], 0})
			}
		}
	}

	instructions := strings.Split(lines[0], "")

	for i := range nodes {
		for nodes[i].name[2] != 'Z' {
			switch instructions[nodes[i].steps%len(instructions)] {
			case "R":
				nodes[i].name = nodeMap[nodes[i].name][1]
			case "L":
				nodes[i].name = nodeMap[nodes[i].name][0]
			default:
				panic("invalid instruction")
			}
			nodes[i].steps++
		}
	}

	var steps []int
	for _, n := range nodes {
		steps = append(steps, n.steps)
	}

	return util.LcmMultiple(steps...)
}
