package day05

import (
	"aoc/util"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

func Run() {
	lines := util.ReadInput("day05.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string) int {
	seeds := getSeeds(lines[0])
	reSeedMap := regexp.MustCompile(`^(\d+ )+\d+$`)

	changes := make(map[int]int)
	for _, line := range lines {
		if reSeedMap.MatchString(line) {
			destStart, sourceStart, rangeLen := seedMap(line)
			for i := range seeds {
				if seeds[i] >= sourceStart && seeds[i] < sourceStart+rangeLen {
					changes[i] = destStart + (seeds[i] - sourceStart)
				}
			}
		}
		if line == "" {
			for k, v := range changes {
				seeds[k] = v
			}
		}
	}
	return slices.Min(seeds)
}

func partB(lines []string) int {
	seedRanges := make([][2]int, 0)
	input := getSeeds(lines[0])
	for i := 0; i < len(input); i += 2 {
		seedRanges = append(seedRanges, [2]int{input[i], input[i] + input[i+1]})
	}
	// get mapping groups
	groups, labels := getMapGroups(lines[2:])
	newRanges := make([][2]int, 0)

	// loop through each label in order from input
	for _, mapLabel := range labels {
		mapGroup := groups[mapLabel]

	seedGroupLoop:
		// loop through each seed range. pop the first seed range off the list and compare it to each map range for
		// the current label. if any part of the range matches the map range, modify the overlaping part of the range
		// based on the map and save it to newRanges. the leftover range or ranges that did not match are appended to
		// seedRanges.
		for {
			if len(seedRanges) == 0 {
				seedRanges = append(seedRanges, newRanges...)
				newRanges = make([][2]int, 0)
				break
			}
			// pop off the first range
			seedRange := seedRanges[0]
			seedRanges = seedRanges[1:]

			// look through each map range for the current label
			for _, mapRange := range mapGroup {
				destStart, sourceStart, rangeLen := mapRange[0], mapRange[1], mapRange[2]
				overlapStart := util.Max(seedRange[0], sourceStart)
				overlapEnd := util.Min(seedRange[1], sourceStart+rangeLen)
				if overlapStart < overlapEnd {
					// if there is an overlap creta e new range for the overlap and modify it based on the map
					// then add it to newRanges
					newRanges = append(newRanges, [2]int{overlapStart - sourceStart + destStart, overlapEnd - sourceStart + destStart})
					if overlapStart > seedRange[0] {
						// if there is a leftover range before the overlap, add it to seedRanges
						seedRanges = append(seedRanges, [2]int{seedRange[0], overlapStart})
					}
					if overlapEnd < seedRange[1] {
						// if there is a leftover range after the overlap, add it to seedRanges
						seedRanges = append(seedRanges, [2]int{overlapEnd, seedRange[1]})
					}
					// break out of the map range loop and start over with the next seed range so the current seed range
					// is not compared to more map ranges
					continue seedGroupLoop
				}
			}
			// if there was no overlap after checking all map ranges, add the seed range to newRanges
			newRanges = append(newRanges, seedRange)
		}
	}

	// find the lowest seed value
	min := seedRanges[0][0]
	for _, seedRange := range seedRanges {
		if seedRange[0] < min {
			min = seedRange[0]
		}
	}
	return min
}

func getSeeds(line string) []int {
	seedNums := strings.TrimPrefix(line, "seeds: ")
	seeds := strings.Split(seedNums, " ")
	result := make([]int, len(seeds))
	for i, num := range seeds {
		result[i] = util.MustAtoi(num)
	}
	return result
}

// takes seed map string and returns destination start, source start and range
func seedMap(numString string) (int, int, int) {
	nums := strings.Split(numString, " ")
	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = util.MustAtoi(num)
	}

	return result[0], result[1], result[2]
}

// returns map of label to seed map groups and list of labels in order
func getMapGroups(lines []string) (map[string][][3]int, []string) {
	groups := make(map[string][][3]int, 0)
	reLabel := regexp.MustCompile(`^.*:$`)
	reSeedMap := regexp.MustCompile(`^(\d+ )+\d+$`)
	labels := make([]string, 0)
	currentLabel := ""
	for _, line := range lines {
		if reLabel.MatchString(line) {
			currentLabel = strings.TrimSuffix(line, ":")
			labels = append(labels, currentLabel)
		}
		if reSeedMap.MatchString(line) {
			destStart, sourceStart, rangeLen := seedMap(line)
			groups[currentLabel] = append(groups[currentLabel], [3]int{destStart, sourceStart, rangeLen})
		}
	}
	return groups, labels
}
