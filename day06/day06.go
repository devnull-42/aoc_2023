package day06

import (
	"aoc/util"
	"fmt"
	"strings"
)

func Run() {
	lines := util.ReadInput("day06.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string) int {
	result := 1

	timeStrings := strings.Fields(lines[0])
	distanceStrings := strings.Fields(lines[1])

	records := make([][2]int, len(timeStrings)-1)

	for i := 1; i < len(timeStrings); i++ {
		records[i-1] = [2]int{util.MustAtoi(timeStrings[i]), util.MustAtoi(distanceStrings[i])}
	}

	for _, record := range records {
		var validTimes int

		for i := 0; i <= record[0]; i++ {
			if i*(record[0]-i) > record[1] {
				validTimes++
			}
		}
		result *= validTimes
	}

	return result
}

func partB(lines []string) int {
	timeStrings := strings.Fields(lines[0])
	distanceStrings := strings.Fields(lines[1])

	timeString := strings.Join(timeStrings[1:], "")
	distanceString := strings.Join(distanceStrings[1:], "")

	time := util.MustAtoi(timeString)
	distance := util.MustAtoi(distanceString)

	var validTimes int

	for i := 0; i <= time; i++ {
		if i*(time-i) > distance {
			validTimes++
		}
	}

	return validTimes
}
