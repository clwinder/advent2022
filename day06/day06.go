package main

import (
	"log"

	tools "github.com/clwinder/advent2022"
)

func main() {
	b := tools.ReadFile("input.txt")

	log.Printf("Answer for part 1: %d", part1(b))
	log.Printf("Answer for part 2: %d", part2(b))
}

func part1(s []byte) int {
	for i := 0; i < len(s)-3; i++ {
		set := tools.Set(s[i : i+4])
		if len(set) == 4 {
			return i + 4
		}
	}
	return 0
}

func part2(s []byte) int {
	for i := 0; i < len(s)-13; i++ {
		set := tools.Set(s[i : i+14])
		if len(set) == 14 {
			return i + 14
		}
	}
	return 0
}
