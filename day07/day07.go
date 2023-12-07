package day07

import (
	"aoc/util"
	"fmt"
	"sort"
	"strings"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func Run() {
	lines := util.ReadInput("day07.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string) int {
	hands, handsList := getHands(lines)
	sort.Sort(camelCards(handsList))

	var result int
	for i, hand := range handsList {
		result += hands[hand] * (i + 1)
	}

	return result
}

func partB(lines []string) int {
	hands, handsList := getHands(lines)
	sort.Sort(camelCardsB(handsList))

	var result int
	for i, hand := range handsList {
		result += hands[hand] * (i + 1)
	}

	return result
}

// Part A
type camelCards []string

func (c camelCards) Len() int {
	return len(c)
}

func (c camelCards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c camelCards) Less(i, j int) bool {
	a := getHandValue(c[i])
	b := getHandValue(c[j])

	for k := 0; k < 6; k++ {
		if a[k] != b[k] {
			return a[k] < b[k]
		}
	}
	return false
}

func getHands(lines []string) (map[string]int, []string) {
	hands := make(map[string]int)
	handsList := make([]string, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		hands[parts[0]] = util.MustAtoi(parts[1])
		handsList[i] = parts[0]
	}
	return hands, handsList
}

func getHandValue(hand string) [6]int {
	charCount := make(map[string]int)
	for _, char := range hand {
		charCount[string(char)]++
	}

	counts := maps.Values(charCount)
	slices.Sort(counts)

	switch {
	case slices.Equal(counts, []int{5}):
		return getCardValues(hand, 7)
	case slices.Equal(counts, []int{1, 4}):
		return getCardValues(hand, 6)
	case slices.Equal(counts, []int{2, 3}):
		return getCardValues(hand, 5)
	case slices.Equal(counts, []int{1, 1, 3}):
		return getCardValues(hand, 4)
	case slices.Equal(counts, []int{1, 2, 2}):
		return getCardValues(hand, 3)
	case slices.Equal(counts, []int{1, 1, 1, 2}):
		return getCardValues(hand, 2)
	case slices.Equal(counts, []int{1, 1, 1, 1, 1}):
		return getCardValues(hand, 1)
	default:
		panic(fmt.Sprintf("unexpected hand: %s", hand))
	}
}

func getCardValues(hand string, handType int) [6]int {
	cards := strings.Split(hand, "")
	cardValues := [6]int{}
	cardValues[0] = handType
	for i, card := range cards {
		switch card {
		case "A":
			cardValues[i+1] = 14
		case "K":
			cardValues[i+1] = 13
		case "Q":
			cardValues[i+1] = 12
		case "J":
			cardValues[i+1] = 11
		case "T":
			cardValues[i+1] = 10
		default:
			cardValues[i+1] = util.MustAtoi(card)
		}
	}
	return cardValues
}

// Part B
func getHandValueB(hand string) [6]int {
	charCount := make(map[string]int)
	var jokerCount int
	for _, char := range hand {
		if string(char) == "J" {
			jokerCount++
		} else {
			charCount[string(char)]++
		}
	}

	if jokerCount == 5 {
		return getCardValuesB(hand, 7)
	}

	counts := maps.Values(charCount)
	slices.Sort(counts)
	counts[len(counts)-1] = counts[len(counts)-1] + jokerCount

	switch {
	case slices.Equal(counts, []int{5}):
		return getCardValuesB(hand, 7)
	case slices.Equal(counts, []int{1, 4}):
		return getCardValuesB(hand, 6)
	case slices.Equal(counts, []int{2, 3}):
		return getCardValuesB(hand, 5)
	case slices.Equal(counts, []int{1, 1, 3}):
		return getCardValuesB(hand, 4)
	case slices.Equal(counts, []int{1, 2, 2}):
		return getCardValuesB(hand, 3)
	case slices.Equal(counts, []int{1, 1, 1, 2}):
		return getCardValuesB(hand, 2)
	case slices.Equal(counts, []int{1, 1, 1, 1, 1}):
		return getCardValuesB(hand, 1)
	default:
		panic(fmt.Sprintf("unexpected hand: %s", hand))
	}
}

func getCardValuesB(hand string, handType int) [6]int {
	cards := strings.Split(hand, "")
	cardValues := [6]int{}
	cardValues[0] = handType
	for i, card := range cards {
		switch card {
		case "A":
			cardValues[i+1] = 14
		case "K":
			cardValues[i+1] = 13
		case "Q":
			cardValues[i+1] = 12
		case "J":
			cardValues[i+1] = 1
		case "T":
			cardValues[i+1] = 10
		default:
			cardValues[i+1] = util.MustAtoi(card)
		}
	}
	return cardValues
}

type camelCardsB []string

func (c camelCardsB) Len() int {
	return len(c)
}

func (c camelCardsB) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c camelCardsB) Less(i, j int) bool {
	a := getHandValueB(c[i])
	b := getHandValueB(c[j])

	for k := 0; k < 6; k++ {
		if a[k] != b[k] {
			return a[k] < b[k]
		}
	}
	return false
}
