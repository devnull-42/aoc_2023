package util

import (
	"bufio"
	"os"
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
