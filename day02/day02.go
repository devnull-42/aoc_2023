package day02

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
)

func Run() {
	lines := util.ReadInput("day02.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string) int {
	gameNumRegex := regexp.MustCompile(`Game (\d+):`)
	colorRegex := regexp.MustCompile(`(\d+)\s*(red|blue|green)`)
	var result int
	const (
		maxRed   = 12
		maxGreen = 13
		maxBlue  = 14
	)

	for _, line := range lines {
		gameMatch := gameNumRegex.FindStringSubmatch(line)
		gameNum, _ := strconv.Atoi(gameMatch[1])

		maxColorValues := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}

		matches := colorRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			count, _ := strconv.Atoi(match[1])
			if count > maxColorValues[match[2]] {
				maxColorValues[match[2]] = count
			}
		}

		if maxColorValues["red"] <= maxRed && maxColorValues["green"] <= maxGreen && maxColorValues["blue"] <= maxBlue {
			result += gameNum
		}

	}
	return result
}

func partB(lines []string) int {
	var result int
	colorRegex := regexp.MustCompile(`(\d+)\s*(red|blue|green)`)

	for _, line := range lines {
		maxColorValues := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}

		matches := colorRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			count, _ := strconv.Atoi(match[1])
			if count > maxColorValues[match[2]] {
				maxColorValues[match[2]] = count
			}
		}

		result += maxColorValues["red"] * maxColorValues["green"] * maxColorValues["blue"]
	}
	return result
}
