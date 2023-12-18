package day18

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
)

func Run() {
	lines := util.ReadInput("day18.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string) int {
	// get verticies and perimeter
	verticies, perimeter := getVerticiesAndPerimiter(lines)

	// get area
	area := getArea(verticies)
	area = getAreaUsingPicksTheorem(area, perimeter)

	// add perimeter
	area += perimeter

	return area
}

func partB(lines []string) int {
	// get verticies and perimeter
	verticies, perimeter := getVerticiesAndPerimiterFromHex(lines)

	// get area
	area := getArea(verticies)
	area = getAreaUsingPicksTheorem(area, perimeter)

	// add perimeter
	area += perimeter

	return area
}

func parseLine(line string) (string, int) {
	re := regexp.MustCompile(`^([UDLR]) (\d+) (\(#\w+\))$`)
	matches := re.FindStringSubmatch(line)

	return matches[1], util.MustAtoi(matches[2])
}

func parseHexLine(line string) (string, int) {
	re := regexp.MustCompile(`^([UDLR]) (\d+) \((#\w+)\)$`)
	matches := re.FindStringSubmatch(line)
	dir := string(matches[3][6])
	hex := matches[3][1:6]

	switch dir {
	case "0":
		dir = "R"
	case "1":
		dir = "D"
	case "2":
		dir = "L"
	case "3":
		dir = "U"
	}

	distance, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		panic(err)
	}
	return dir, int(distance)
}

func getVerticiesAndPerimiter(lines []string) ([][2]int, int) {
	verticies := make([][2]int, len(lines)+1)
	verticies[0] = [2]int{0, 0}
	perimeter := 0
	r, c := 0, 0

	for i, line := range lines {
		dir, dist := parseLine(line)
		perimeter += dist
		switch dir {
		case "U":
			r -= dist
			verticies[i+1] = [2]int{r, c}
		case "D":
			r += dist
			verticies[i+1] = [2]int{r, c}
		case "L":
			c -= dist
			verticies[i+1] = [2]int{r, c}
		case "R":
			c += dist
			verticies[i+1] = [2]int{r, c}
		}
	}
	return verticies, perimeter
}

func getVerticiesAndPerimiterFromHex(lines []string) ([][2]int, int) {
	verticies := make([][2]int, len(lines)+1)
	verticies[0] = [2]int{0, 0}
	perimeter := 0
	r, c := 0, 0

	for i, line := range lines {
		dir, dist := parseHexLine(line)
		perimeter += dist
		switch dir {
		case "U":
			r -= dist
			verticies[i+1] = [2]int{r, c}
		case "D":
			r += dist
			verticies[i+1] = [2]int{r, c}
		case "L":
			c -= dist
			verticies[i+1] = [2]int{r, c}
		case "R":
			c += dist
			verticies[i+1] = [2]int{r, c}
		}
	}
	return verticies, perimeter
}

// get area using shoelace theorem https://en.wikipedia.org/wiki/Shoelace_formula
func getArea(verticies [][2]int) int {
	n := len(verticies)
	area := 0
	for i := 0; i < n-1; i++ {
		j := (i + 1) % n
		area += verticies[i][0]*verticies[j][1] - verticies[j][0]*verticies[i][1]
	}

	if area < 0 {
		area *= -1
	}
	return area / 2
}

// use pick's theorem https://en.wikipedia.org/wiki/Pick%27s_theorem
// this is because the shoelace theorem overestimates the area since the
// verticies are in the middle of the squares
func getAreaUsingPicksTheorem(area int, perimeter int) int {
	return area - perimeter/2 + 1
}
