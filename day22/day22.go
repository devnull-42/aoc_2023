package day22

import (
	"aoc/util"
	"cmp"
	"fmt"
	"slices"
)

func Run() {
	lines := util.ReadInput("day22.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string) int {
	bricks := getBricks(lines)
	bricks.Sort()

	dropBricks(bricks)
	bricks.Sort()

	supports, supportedBy := getSupports(bricks)

	return getRedundant(supports, supportedBy)
}

func partB(lines []string) int {
	return 0
}

type Bricks []Brick

func (b Bricks) Sort() {
	slices.SortFunc(b, func(a, b Brick) int {
		return cmp.Compare(a.Z1, b.Z1)
	})
}

func (b Bricks) Print() {
	for _, brick := range b {
		fmt.Println(brick.String())
	}
}

type Brick struct {
	BrickNum               int
	X1, Y1, X2, Y2, Z1, Z2 int
}

func (b Brick) String() string {
	return fmt.Sprintf("#%d: %d,%d,%d~%d,%d,%d", b.BrickNum, b.X1, b.Y1, b.Z1, b.X2, b.Y2, b.Z2)
}

func getBricks(lines []string) Bricks {
	bricks := make(Bricks, len(lines))
	for i, line := range lines {
		var b Brick
		fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d", &b.X1, &b.Y1, &b.Z1, &b.X2, &b.Y2, &b.Z2)
		b.BrickNum = i
		bricks[i] = b
	}
	return bricks
}

func overlaps(a, b Brick) bool {
	return slices.Max([]int{a.X1, b.X1}) <= slices.Min([]int{a.X2, b.X2}) && slices.Max([]int{a.Y1, b.Y1}) <= slices.Min([]int{a.Y2, b.Y2})
}

func dropBricks(bricks Bricks) {
	for i, brick := range bricks {
		zMax := 1
		for _, check := range bricks[:i] {
			if overlaps(brick, check) {
				zMax = slices.Max([]int{zMax, check.Z2 + 1})
			}
		}
		bricks[i].Z1 = zMax
		bricks[i].Z2 -= brick.Z1 - zMax
	}
}

type setMap map[int]map[int]struct{}

func getSupports(bricks Bricks) (setMap, setMap) {
	supports := make(setMap)
	supportedBy := make(setMap)
	for i, brick := range bricks {
		if _, ok := supports[brick.BrickNum]; !ok {
			supports[brick.BrickNum] = make(map[int]struct{})
		}
		if _, ok := supportedBy[brick.BrickNum]; !ok {
			supportedBy[brick.BrickNum] = make(map[int]struct{})
		}
		for j, check := range bricks[:i] {
			if overlaps(check, brick) && brick.Z1 == check.Z2+1 {
				supports[bricks[j].BrickNum][bricks[i].BrickNum] = struct{}{}
				supportedBy[bricks[i].BrickNum][bricks[j].BrickNum] = struct{}{}
			}
		}
	}
	return supports, supportedBy
}

func getRedundant(supports, supportedBy setMap) int {
	result := make(map[int]struct{})
	for brickNum, supts := range supports {
		for supt := range supts {
			if len(supportedBy[supt]) >= 2 {
				result[brickNum] = struct{}{}
			}
		}
		if len(supports[brickNum]) == 0 {
			result[brickNum] = struct{}{}
		}
	}
	return len(result)
}
