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
	default:
		fmt.Println("Day not recognized")
	}
}
