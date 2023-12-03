package day03

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
)

func Run() {
	lines := util.ReadInput("day03.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string) int {
	var result int

	reNum := regexp.MustCompile(`\d+`)
	reSymbol := regexp.MustCompile(`[^\d.]`)

	for i, line := range lines {
		var (
			prevLineSymbols [][]int
			curLineSymbols  [][]int
			nextLineSymbols [][]int
		)

		if i > 0 {
			prevLineSymbols = reSymbol.FindAllStringIndex(lines[i-1], -1)
		}
		curLineSymbols = reSymbol.FindAllStringIndex(line, -1)
		if i < len(lines)-1 {
			nextLineSymbols = reSymbol.FindAllStringIndex(lines[i+1], -1)
		}

		numberMatches := reNum.FindAllStringIndex(line, -1)

		for _, match := range numberMatches {
			start, end := match[0], match[1]
			if isSymbolAdjacent(prevLineSymbols, start, end) || isSymbolAdjacent(curLineSymbols, start, end) || isSymbolAdjacent(nextLineSymbols, start, end) {
				num, _ := strconv.Atoi(line[start:end])
				result += num
				continue
			}
		}
	}
	return result
}

func partB(lines []string) int {
	var result int

	reGear := regexp.MustCompile(`\*`)

	for i := range lines {
		gearMatches := reGear.FindAllStringIndex(lines[i], -1)

		for _, match := range gearMatches {
			var numMatches int
			var matchingNumbers []int

			if i > 0 {
				count, nums := adjacentNums(lines[i-1], match[0])
				numMatches += count
				matchingNumbers = append(matchingNumbers, nums...)
			}

			count, nums := adjacentNums(lines[i], match[0])
			numMatches += count
			matchingNumbers = append(matchingNumbers, nums...)

			if i < len(lines)-1 {
				count, nums := adjacentNums(lines[i+1], match[0])
				numMatches += count
				matchingNumbers = append(matchingNumbers, nums...)
			}

			if numMatches == 2 {
				result += matchingNumbers[0] * matchingNumbers[1]
			}

		}
	}

	return result
}

func isSymbolAdjacent(symbolMatches [][]int, start, end int) bool {
	for _, match := range symbolMatches {
		if match[0] >= start-1 && match[0] <= end {
			return true
		}
	}
	return false
}

func adjacentNums(line string, gearIdx int) (int, []int) {
	var result int
	var adjacentMatches []int

	reNum := regexp.MustCompile(`\d+`)
	numMatches := reNum.FindAllStringIndex(line, -1)

	for _, match := range numMatches {
		if gearIdx >= match[0]-1 && gearIdx <= match[1] {
			result++
			num, _ := strconv.Atoi(line[match[0]:match[1]])
			adjacentMatches = append(adjacentMatches, num)
		}
	}
	return result, adjacentMatches
}
