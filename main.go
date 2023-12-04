package main

import (
	"fmt"
	"os"

	"aoc/day01"
	"aoc/day02"
	"aoc/day03"
	"aoc/day04"
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
	default:
		fmt.Println("Day not recognized")
	}
}
