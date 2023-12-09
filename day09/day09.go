package day09

import (
	"aoc/util"
	"fmt"
	"slices"
	"strings"
)

func Run() {
	lines := util.ReadInput("day09.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string) int {
	predictedNums := make([]int, 0)
	for _, line := range lines {
		numbers := parseLine(line)
		lastNumbers := make([]int, 0)
		lastNumbers = append(lastNumbers, numbers[len(numbers)-1])
		lastNumbers = append(lastNumbers, calculateDiff(numbers)...)
		sum := util.SliceSum(lastNumbers)
		predictedNums = append(predictedNums, sum)
	}
	return util.SliceSum(predictedNums)
}

func partB(lines []string) int {
	predictedNums := make([]int, 0)
	for _, line := range lines {
		numbers := parseLine(line)
		slices.Reverse(numbers)
		lastNumbers := make([]int, 0)
		lastNumbers = append(lastNumbers, numbers[len(numbers)-1])
		lastNumbers = append(lastNumbers, calculateDiff(numbers)...)
		sum := util.SliceSum(lastNumbers)
		predictedNums = append(predictedNums, sum)
	}
	return util.SliceSum(predictedNums)
}

func calculateDiff(numbers []int) []int {
	nextDiff := make([]int, len(numbers)-1)
	allZeros := true
	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]
		if diff != 0 {
			allZeros = false
		}
		nextDiff[i-1] = diff
	}
	if allZeros {
		return []int{0}
	}

	return append([]int{nextDiff[len(nextDiff)-1]}, calculateDiff(nextDiff)...)
}

func parseLine(line string) []int {
	numbers := make([]int, 0)
	numStrs := strings.Split(line, " ")
	for _, numStr := range numStrs {
		numbers = append(numbers, util.MustAtoi(numStr))
	}
	return numbers
}
