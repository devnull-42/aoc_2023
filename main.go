package main

import (
	"fmt"
	"os"

	"aoc/day01"
	"aoc/day02"
	"aoc/day03"
	"aoc/day04"
	"aoc/day05"
	"aoc/day06"
	"aoc/day07"
	"aoc/day08"
	"aoc/day09"
	"aoc/day10"
	"aoc/day11"
	"aoc/day12"
	"aoc/day13"
	"aoc/day14"
	"aoc/day15"
	"aoc/day16"
	"aoc/day17"
	"aoc/day18"
	"aoc/day19"
	"aoc/day20"
	"aoc/day21"
	"aoc/day22"
	"aoc/day23"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("specify a day")
		return
	}

	day := os.Args[1]

	switch day {
	case "1":
		day01.Run()
	case "2":
		day02.Run()
	case "3":
		day03.Run()
	case "4":
		day04.Run()
	case "5":
		day05.Run()
	case "6":
		day06.Run()
	case "7":
		day07.Run()
	case "8":
		day08.Run()
	case "9":
		day09.Run()
	case "10":
		day10.Run()
	case "11":
		day11.Run()
	case "12":
		day12.Run()
	case "13":
		day13.Run()
	case "14":
		day14.Run()
	case "15":
		day15.Run()
	case "16":
		day16.Run()
	case "17":
		day17.Run()
	case "18":
		day18.Run()
	case "19":
		day19.Run()
	case "20":
		day20.Run()
	case "21":
		day21.Run()
	case "22":
		day22.Run()
	case "23":
		day23.Run()
	default:
		fmt.Println("Day not recognized")
	}
}
