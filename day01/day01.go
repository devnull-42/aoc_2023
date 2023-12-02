package day01

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
)

func Run() {
	lines := util.ReadInput("day01.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string) int {
	var sum int
	re := regexp.MustCompile(`\d`)

	for _, line := range lines {
		matches := re.FindAllString(line, -1)

		num := string(matches[0]) + string(matches[len(matches)-1])
		inum, _ := strconv.Atoi(num)
		sum += inum
	}
	return sum
}

var numberStrings = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

var reverseNumberStrings = map[string]string{
	"orez":  "0",
	"eno":   "1",
	"owt":   "2",
	"eerht": "3",
	"ruof":  "4",
	"evif":  "5",
	"xis":   "6",
	"neves": "7",
	"thgie": "8",
	"enin":  "9",
}

func partB(lines []string) int {
	var sum int

	re := regexp.MustCompile(`(zero|one|two|three|four|five|six|seven|eight|nine)|\d`)
	reverseRE := regexp.MustCompile(`(orez|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)|\d`)

	for _, line := range lines {
		revLine := util.ReverseString(line)

		match := re.FindString(line)
		revMatch := reverseRE.FindString(revLine)

		var numString string

		// check forward match
		if val, ok := numberStrings[match]; ok {
			numString += val
		} else {
			numString += match
		}

		// check reverse match
		if val, ok := reverseNumberStrings[revMatch]; ok {
			numString += val
		} else {
			numString += revMatch
		}

		inum, _ := strconv.Atoi(numString)
		sum += inum
	}
	return sum
}
