package day07

import (
	"aoc/util"
	"fmt"
	"sort"
	"strings"
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

	switch len(charCount) {
	case 1:
		// 5 of a kind
		return getCardValues(hand, 7)
	case 2:
		for _, count := range charCount {
			if count == 4 {
				// 4 of a kind
				return getCardValues(hand, 6)
			}
		}
		// full house
		return getCardValues(hand, 5)
	case 3:
		for _, count := range charCount {
			if count == 3 {
				// 3 of a kind
				return getCardValues(hand, 4)
			}
		}
		// two pair
		return getCardValues(hand, 3)
	case 4:
		// 1 pair
		return getCardValues(hand, 2)
	case 5:
		// high card
		return getCardValues(hand, 1)
	default:
		fmt.Println(hand, charCount)
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
		// 5 of a kind
		return getCardValuesB(hand, 7)
	}

	if jokerCount > 0 {
		var maxChar string
		var maxCount int
		for char, count := range charCount {
			if count > maxCount {
				maxChar = char
				maxCount = count
			}
		}

		charCount[maxChar] += jokerCount
	}

	switch len(charCount) {
	case 1:
		// 5 of a kind
		return getCardValuesB(hand, 7)
	case 2:
		for _, count := range charCount {
			if count == 4 {
				// 4 of a kind
				return getCardValuesB(hand, 6)
			}
		}
		// full house
		return getCardValuesB(hand, 5)
	case 3:
		for _, count := range charCount {
			if count == 3 {
				// 3 of a kind
				return getCardValuesB(hand, 4)
			}
		}
		// two pair
		return getCardValuesB(hand, 3)
	case 4:
		// 1 pair
		return getCardValuesB(hand, 2)
	case 5:
		// high card
		return getCardValuesB(hand, 1)
	default:
		fmt.Println(hand, charCount)
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
