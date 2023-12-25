package day24

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
)

func Run() {
	lines := util.ReadInput("day24.txt")
	result := partA(lines, 200000000000000, 400000000000000)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines, 200000000000000, 400000000000000)
	fmt.Printf("partB: %d\n", result)
}

func partA(input []string, min, max float64) int {
	lines := parseInput(input)
	result := 0
	for i, line1 := range lines {
		for _, line2 := range lines[i+1:] {
			intersection, ok := findIntersection(line1, line2)
			if ok && checkIntersection(intersection, min, max) &&
				isAfterFirstPoint(intersection, line1) &&
				isAfterFirstPoint(intersection, line2) {
				result++
			}
		}
	}

	return result
}

// SageMath output from https://sagecell.sagemath.org/ using the output from this program
const SageOutput = `[
[x == 131246724405205, y == 399310844858926, z == 277550172142625, vx == 279, vy == -184, vz == 16, t1 == 130621773037, t2 == 423178590960, t3 == 631793973864]
]`

func partB(input []string, min, max float64) int {
	partBSolution(input, min, max)
	return 0
}

type Point struct {
	X, Y, Z float64
}

type Line struct {
	P1, P2 Point
}

func parseInput(input []string) []Line {
	re := regexp.MustCompile(`^(\d+),\s+(\d+),\s+(\d+)\s+@\s+(-?\d+),\s+(-?\d+),\s+(-?\d+)$`)
	var lines []Line
	for _, i := range input {
		matches := re.FindStringSubmatch(i)
		if matches == nil {
			panic(fmt.Sprintf("no match for %s", i))
		}
		x := parseFloat(matches[1])
		y := parseFloat(matches[2])
		z := parseFloat(matches[3])
		dx := parseFloat(matches[4])
		dy := parseFloat(matches[5])
		dz := parseFloat(matches[6])

		lines = append(lines, Line{
			P1: Point{X: x, Y: y, Z: z},
			P2: Point{X: x + dx, Y: y + dy, Z: z + dz},
		})
	}
	return lines
}

func parseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}

// slopeInterceptForm calculates the slope and y-intercept of a line
func slopeInterceptForm(line Line) (float64, float64) {
	m := (line.P2.Y - line.P1.Y) / (line.P2.X - line.P1.X)
	b := line.P1.Y - m*line.P1.X
	return m, b
}

// findIntersection calculates the intersection point of two lines
func findIntersection(line1, line2 Line) (Point, bool) {
	m1, b1 := slopeInterceptForm(line1)
	m2, b2 := slopeInterceptForm(line2)

	// check if lines are parallel
	if m1 == m2 {
		return Point{}, false
	}

	// calculate intersection point
	x := (b2 - b1) / (m1 - m2)
	y := m1*x + b1

	return Point{X: x, Y: y}, true
}

func checkIntersection(point Point, min, max float64) bool {
	return min <= point.X && point.X <= max && min <= point.Y && point.Y <= max
}

func isAfterFirstPoint(intersection Point, line Line) bool {
	dirVector := Point{X: line.P2.X - line.P1.X, Y: line.P2.Y - line.P1.Y}
	intersectVector := Point{X: intersection.X - line.P1.X, Y: intersection.Y - line.P1.Y}

	// check if the dot product of direction vector and intersection vector is positive
	dotProduct := dirVector.X*intersectVector.X + dirVector.Y*intersectVector.Y
	return dotProduct > 0
}

// part b from https://github.com/vipul0092/advent-of-code-2023/blob/main/day24/day24.go
// this solution uses an external program called SageMath to solve the problem. All the solutions I saw
// were using an external package or program. Al lpython solutions were using sympy.
// Putting it here for future reference and for SageMath reference. It is modified to work along side
// my existing code.
type LineB struct {
	a float64
	b float64
	c float64
}

type Stone struct {
	x    int
	y    int
	z    int
	vx   int
	vy   int
	vz   int
	line LineB
}

func partBSolution(input []string, LEAST, MOST float64) {
	re := regexp.MustCompile(`^(\d+),\s+(\d+),\s+(\d+)\s+@\s+(-?\d+),\s+(-?\d+),\s+(-?\d+)$`)
	stones := make([]Stone, len(input))
	for i, ipt := range input {
		fmt.Printf("input: %s\n", ipt)
		matches := re.FindStringSubmatch(ipt)
		fmt.Printf("matches: %v\n", matches)
		x := util.MustAtoi(matches[1])
		y := util.MustAtoi(matches[2])
		z := util.MustAtoi(matches[3])
		vx := util.MustAtoi(matches[4])
		vy := util.MustAtoi(matches[5])
		vz := util.MustAtoi(matches[6])

		m := float64(vy) / float64(vx)     // y2-y1 / x2-x1
		c := float64(y) - (m * float64(x)) // y = mx + c => c = y - mx
		// y = mx + c => mx - y + c = 0
		stones[i] = Stone{x, y, z, vx, vy, vz, LineB{m, -1, c}}
	}

	count := 0
	for i := 0; i < len(stones)-1; i++ {
		for j := i + 1; j < len(stones); j++ {
			s1, s2 := stones[i], stones[j]
			// intersection of two lines a1x + b1y + c1 = 0 & a2x + b2y + c2 = 0 is:
			// b1c2 - b2c1 / a1b2 - a2b1, c1a2 - c2a1 / a1b2 - a2b1
			a1, b1, c1, a2, b2, c2 := s1.line.a, s1.line.b, s1.line.c, s2.line.a, s2.line.b, s2.line.c
			a1b2_a2b1 := (a1 * b2) - (a2 * b1)
			if a1b2_a2b1 == 0 { // parallel or same
				continue
			}

			ix, iy := (b1*c2-b2*c1)/a1b2_a2b1, (c1*a2-c2*a1)/a1b2_a2b1

			// time should be > 0
			// x(t) = x0 + vt => t = x(t) - x0 / v
			t1, t2 := (ix-float64(s1.x))/float64(s1.vx), (ix-float64(s2.x))/float64(s2.vx)

			if t1 > 0 && t2 > 0 && ix >= LEAST && ix <= MOST && iy >= LEAST && iy <= MOST {
				count++
			}
		}
	}

	//fmt.Println("Alt Part 1: ", count) // 13892

	// Generate SageMath script. Copy the output from these lines into https://sagecell.sagemath.org/
	// and then copy the output from that into the SageOutput constant above.
	fmt.Println()
	fmt.Println("var('x y z vx vy vz t1 t2 t3')")
	fmt.Println("eq1 = x + (vx * t1) == ", stones[0].x, " + (", stones[0].vx, " * t1)")
	fmt.Println("eq2 = y + (vy * t1) == ", stones[0].y, " + (", stones[0].vy, " * t1)")
	fmt.Println("eq3 = z + (vz * t1) == ", stones[0].z, " + (", stones[0].vz, " * t1)")
	fmt.Println("eq4 = x + (vx * t2) == ", stones[1].x, " + (", stones[1].vx, " * t2)")
	fmt.Println("eq5 = y + (vy * t2) == ", stones[1].y, " + (", stones[1].vy, " * t2)")
	fmt.Println("eq6 = z + (vz * t2) == ", stones[1].z, " + (", stones[1].vz, " * t2)")
	fmt.Println("eq7 = x + (vx * t3) == ", stones[2].x, " + (", stones[2].vx, " * t3)")
	fmt.Println("eq8 = y + (vy * t3) == ", stones[2].y, " + (", stones[2].vy, " * t3)")
	fmt.Println("eq9 = z + (vz * t3) == ", stones[2].z, " + (", stones[2].vz, " * t3)")
	fmt.Println("print(solve([eq1,eq2,eq3,eq4,eq5,eq6,eq7,eq8,eq9],x,y,z,vx,vy,vz,t1,t2,t3))")
	fmt.Println()

	reSage := regexp.MustCompile(`x == (\d+), y == (\d+), z == (\d+)`)
	matches := reSage.FindStringSubmatch(SageOutput)
	if matches != nil {
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		z, _ := strconv.Atoi(matches[3])
		fmt.Println("Part 2: ", x+y+z)
	}
}
