package day14

import (
	"aoc/util"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

func Run() {
	lines := util.ReadInput("day14.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string) int {
	var platform [][]string
	for _, line := range lines {
		platform = append(platform, strings.Split(line, ""))
	}

	// shift rocks north
	northCycle(platform)

	return getWeight(platform)
}

func partB(lines []string) int {
	var platform [][]string
	for _, line := range lines {
		platform = append(platform, strings.Split(line, ""))
	}

	var cycleCount, cycleLength int
	hashes := make(map[string]int)

	// cycle until we find a repeat hash. then calculate the cycle length
	for i := 0; i < 1000000000; i++ {
		hash := hashPlatform(platform)
		if c, ok := hashes[hash]; ok {
			cycleCount = i
			cycleLength = i - c
			break
		}
		hashes[hash] = i
		fullCycle(platform)
	}

	// calculate the number of cycles remaining
	cycles := (1000000000 - cycleCount) % cycleLength

	for i := 0; i < cycles; i++ {
		fullCycle(platform)
	}

	return getWeight(platform)
}

func hashPlatform(platform [][]string) string {
	var builder strings.Builder

	for _, row := range platform {
		builder.WriteString(strings.Join(row, ""))
	}

	hash := sha256.New()
	hash.Write([]byte(builder.String()))
	return hex.EncodeToString(hash.Sum(nil))
}

func fullCycle(platform [][]string) {
	northCycle(platform)
	westCycle(platform)
	southCycle(platform)
	eastCycle(platform)
}

func northCycle(platform [][]string) {
	for r, row := range platform {
		if r > 0 {
			for c := range row {
				shiftNorth(platform, r, c)
			}
		}
	}
}

func shiftNorth(platform [][]string, r, c int) {
	if platform[r][c] == "O" && platform[r-1][c] == "." {
		platform[r][c] = "."
		platform[r-1][c] = "O"

		if r > 1 {
			shiftNorth(platform, r-1, c)
		}
	}
}

func westCycle(platform [][]string) {
	for r, row := range platform {
		for c := range row {
			if c > 0 {
				shiftWest(platform, r, c)
			}
		}
	}
}

func shiftWest(platform [][]string, r, c int) {
	if platform[r][c] == "O" && platform[r][c-1] == "." {
		platform[r][c] = "."
		platform[r][c-1] = "O"

		if c > 1 {
			shiftWest(platform, r, c-1)
		}
	}
}

func southCycle(platform [][]string) {
	for r := len(platform) - 2; r >= 0; r-- {
		for c := range platform[r] {
			shiftSouth(platform, r, c)
		}
	}
}

func shiftSouth(platform [][]string, r, c int) {
	if platform[r][c] == "O" && platform[r+1][c] == "." {
		platform[r][c] = "."
		platform[r+1][c] = "O"

		if r < len(platform)-2 {
			shiftSouth(platform, r+1, c)
		}
	}
}

func eastCycle(platform [][]string) {
	for r, row := range platform {
		for c := len(row) - 2; c >= 0; c-- {
			shiftEast(platform, r, c)
		}
	}
}

func shiftEast(platform [][]string, r, c int) {
	if platform[r][c] == "O" && platform[r][c+1] == "." {
		platform[r][c] = "."
		platform[r][c+1] = "O"

		if c < len(platform[r])-2 {
			shiftEast(platform, r, c+1)
		}
	}
}

func getWeight(platform [][]string) int {
	var weight int
	re := regexp.MustCompile(`O`)
	for r, row := range platform {
		matches := re.FindAllString(strings.Join(row, ""), -1)
		weight += len(matches) * (len(platform) - r)
	}
	return weight
}
