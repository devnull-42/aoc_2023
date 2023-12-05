package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadInput(filename string) []string {
	lines := make([]string, 0)
	filepath := "input/" + filename
	file, _ := os.Open(filepath)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		tmp := runes[i]
		runes[i] = runes[j]
		runes[j] = tmp
	}
	return string(runes)
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("MustAtoi: %v", err))
	}
	return i
}

func Contains[T comparable](slice []T, val T) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
