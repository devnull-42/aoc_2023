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
	default:
		fmt.Println("Day not recognized")
	}
}
