package day04

import (
	"aoc/util"
	"fmt"
	"regexp"
)

func Run() {
	lines := util.ReadInput("day04.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string) int {
	var result int
	for _, line := range lines {
		var score int
		winning, owned := getNumberLists(line)
		for _, num := range owned {
			if util.Contains(winning, num) {
				if score == 0 {
					score++
				} else {
					score *= 2
				}
			}
		}
		result += score
	}

	return result
}

func partB(lines []string) int {
	scratchcards := make([]int, len(lines))

	for i, line := range lines {
		// add original card to scratchcards
		scratchcards[i]++

		// get number of matches
		var matches int
		winning, owned := getNumberLists(line)
		for _, num := range owned {
			if util.Contains(winning, num) {
				matches++
			}
		}

		// increment scratchcards for each match based on number of copies of current card
		for j := 1; j <= matches; j++ {
			scratchcards[i+j] += scratchcards[i]
		}
	}

	// sum scratchcards
	var result int
	for _, v := range scratchcards {
		result += v
	}
	return result
}

func getNumberLists(line string) ([]int, []int) {
	//split line into two lists
	reLists := regexp.MustCompile(`Card\s+\d+\:\s+(.*)\s\|\s(.*)`)
	matches := reLists.FindStringSubmatch(line)
	if len(matches) != 3 {
		panic(fmt.Sprintf("getNumberLists: unexpected number of matches: %d", len(matches)))
	}

	// split lists into numbers
	reNums := regexp.MustCompile(`(\d+)`)
	winningList := reNums.FindAllStringSubmatch(matches[1], -1)
	ownedList := reNums.FindAllStringSubmatch(matches[2], -1)

	// convert to ints
	winning := make([]int, len(winningList))
	for i, num := range winningList {
		winning[i] = util.MustAtoi(num[1])
	}
	owned := make([]int, len(ownedList))
	for i, num := range ownedList {
		owned[i] = util.MustAtoi(num[1])
	}

	return winning, owned
}
