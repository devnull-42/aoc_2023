package main

import (
	"fmt"
	"os"

	"aoc/day01"
	"aoc/day02"
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
	default:
		fmt.Println("Day not recognized")
	}
}