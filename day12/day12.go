package day12

import (
	"aoc/util"
	"fmt"
	"slices"
	"strings"
)

func Run() {
	lines := util.ReadInput("day12.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string) int {
	var sum int
	for _, line := range lines {
		springs, pattern := parseLine(line)
		combinations := generateCombinations(springs, pattern)
		sum += combinations
	}

	return sum
}

func partB(lines []string) int {
	var sum int
	cache = make(map[string]int)
	for _, line := range lines {
		springs, pattern := parseLine(line)
		springs, pattern = unfold(springs, pattern)
		combinations := combinationsB(springs, pattern)
		sum += combinations
	}

	return sum
}

// parseLine takes a line of input and returns a string and slice of ints
func parseLine(line string) (string, []int) {
	var s string
	var ints []int
	input := strings.Split(line, " ")
	s = input[0]
	intStrings := strings.Split(input[1], ",")
	for _, intString := range intStrings {
		ints = append(ints, util.MustAtoi(intString))
	}
	return s, ints
}

// generateCombinations generates all possible combinations by replacing '?' with '#' and '.'
//
// this is a bad solution because it is exponential time, but i'm keeping it here
// because it was my first solution
func generateCombinations(s string, pattern []int) int {
	// find the first occurrence of '?'
	hashPattern := make([]int, 0)
	hashCount := 0
	index := -1

FindIndex:
	for i, char := range s {
		switch char {
		case '?':
			index = i
			if hashCount > 0 {
				hashPattern = append(hashPattern, hashCount)
				hashCount = 0
			}
			break FindIndex
		case '#':
			hashCount++
		case '.':
			if hashCount > 0 {
				hashPattern = append(hashPattern, hashCount)
				hashCount = 0
			}
		}
	}
	// if the last character is a '#', add it to the hash pattern
	if hashCount > 0 {
		hashPattern = append(hashPattern, hashCount)
	}

	// if the current hash pattern is longer than the pattern, return an empty slice
	if len(hashPattern) > len(pattern) {
		return 0
	}

	// if the current hash pattern is not a match, return an empty slice
	for i, hash := range hashPattern {
		if hash > pattern[i] {
			return 0
		}
	}

	// base case: no '?' found, return the string as it is
	if index == -1 {
		// check the has pattern to make sure it is an exact match
		if slices.Compare(hashPattern, pattern) != 0 {
			return 0
		}
		return 1
	}

	// recursive case: replace '?' with '#' and '.' and recurse
	replaceWithHash := s[:index] + "#" + s[index+1:]
	replaceWithDot := s[:index] + "." + s[index+1:]

	return generateCombinations(replaceWithHash, pattern) + generateCombinations(replaceWithDot, pattern)
}

func unfold(springs string, pattern []int) (string, []int) {
	unfoldedSprings := springs
	for i := 0; i < 4; i++ {
		unfoldedSprings = fmt.Sprintf("%s?%s", unfoldedSprings, springs)
	}

	unfoldedPattern := make([]int, 0)
	for i := 0; i < 5; i++ {
		unfoldedPattern = append(unfoldedPattern, pattern...)
	}
	return unfoldedSprings, unfoldedPattern
}

var cache map[string]int

func makeKey(springs string, pattern []int) string {
	return fmt.Sprintf("%s:%v", springs, pattern)
}

func combinationsB(springs string, pattern []int) int {
	// base case if the springs string is empty and the pattern is empty
	// then we have a match
	if springs == "" {
		if len(pattern) == 0 {
			return 1
		}
		return 0
	}

	// base case if the pattern is empty and the springs string does not
	// contain another #, then we have a match
	if len(pattern) == 0 {
		if !strings.Contains(springs, "#") {
			return 1
		}
		return 0
	}

	// create the cache key
	cacheKey := makeKey(springs, pattern)

	// check the cache
	if val, ok := cache[cacheKey]; ok {
		return val
	}

	// initialize result
	var result int

	// split cases based on the first character of the pattern
	if springs[0] == '.' || springs[0] == '?' {
		result += combinationsB(springs[1:], pattern)
	}

	if springs[0] == '#' || springs[0] == '?' {
		if pattern[0] <= len(springs) &&
			!strings.Contains(springs[:pattern[0]], ".") {
			if pattern[0] == len(springs) {
				result += combinationsB("", pattern[1:])
			} else if springs[pattern[0]] != '#' {
				result += combinationsB(springs[pattern[0]+1:], pattern[1:])
			}
		}
	}

	// add the result to the cache
	cache[cacheKey] = result
	return result
}
