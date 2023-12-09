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
		allZeros := false
		for !allZeros {
			nextDiff := make([]int, 0)
			allZeros = true
			for i := 1; i < len(numbers); i++ {
				diff := numbers[i] - numbers[i-1]
				if diff != 0 {
					allZeros = false
				}
				nextDiff = append(nextDiff, diff)
			}
			lastNumbers = append(lastNumbers, nextDiff[len(nextDiff)-1])
			numbers = nextDiff
		}
		sum := predictFirst(lastNumbers)
		predictedNums = append(predictedNums, sum)
	}
	return util.SliceSum(predictedNums)
}

func partB(lines []string) int {
	predictedNums := make([]int, 0)
	for _, line := range lines {
		numbers := parseLine(line)
		firstNumbers := make([]int, 0)
		firstNumbers = append(firstNumbers, numbers[0])
		allZeros := false
		for !allZeros {
			nextDiff := make([]int, 0)
			allZeros = true
			for i := 1; i < len(numbers); i++ {
				diff := numbers[i] - numbers[i-1]
				if diff != 0 {
					allZeros = false
				}
				nextDiff = append(nextDiff, diff)
			}
			firstNumbers = append(firstNumbers, nextDiff[0])
			numbers = nextDiff
		}
		diff := predictLast(firstNumbers)
		predictedNums = append(predictedNums, diff)
	}
	return util.SliceSum(predictedNums)
}

func parseLine(line string) []int {
	numbers := make([]int, 0)
	numStrs := strings.Split(line, " ")
	for _, numStr := range numStrs {
		numbers = append(numbers, util.MustAtoi(numStr))
	}
	return numbers
}

func predictFirst(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func predictLast(numbers []int) int {
	diff := 0
	slices.Reverse(numbers)
	for _, num := range numbers[1:] {
		diff = num - diff
	}
	return diff
}
