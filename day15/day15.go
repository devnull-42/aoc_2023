package day15

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strings"
)

func Run() {
	lines := util.ReadInput("day15.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string) int {
	var result int
	steps := strings.Split(lines[0], ",")
	for _, step := range steps {
		result += hash(step)
	}
	return result
}

func partB(lines []string) int {
	lbox := new(lightbox)

	steps := strings.Split(lines[0], ",")
	for _, step := range steps {
		parseStep(lbox, step)
	}

	return totalPower(lbox)
}

type lens struct {
	label  string
	length int
}

type lightbox [256][]lens

// parse instruction
func parseStep(lbox *lightbox, step string) {
	reStep := regexp.MustCompile(`(\w+)(-|=)([0-9])*`)
	match := reStep.FindStringSubmatch(step)
	if len(match) >= 3 {
		switch match[2] {
		case "=":
			length := util.MustAtoi(match[3])
			upsertLens(lbox, match[1], length)
		case "-":
			removeLens(lbox, match[1])
		}
	}
}

// calculate hash
func hash(input string) int {
	var result int
	for _, c := range input {
		result += int(c)
		result *= 17
		result %= 256
	}
	return result
}

// update or insert lens length
func upsertLens(lbox *lightbox, label string, length int) {
	labelHash := hash(label)
	for i, lens := range lbox[labelHash] {
		if lens.label == label {
			lbox[labelHash][i].length = length
			return
		}
	}
	lbox[labelHash] = append(lbox[labelHash], lens{label, length})
}

// remove lens
func removeLens(lbox *lightbox, label string) {
	labelHash := hash(label)
	for i, lens := range lbox[labelHash] {
		if lens.label == label {
			lbox[labelHash] = append(lbox[labelHash][:i], lbox[labelHash][i+1:]...)
			return
		}
	}
}

// calculate power for part b
func totalPower(lbox *lightbox) int {
	var result int
	for i, lenses := range lbox {
		for j, lens := range lenses {
			result += (i + 1) * (j + 1) * lens.length
		}
	}
	return result
}
