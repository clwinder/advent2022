package main

import (
	"log"
	"os"
	"strings"

	tools "github.com/clwinder/advent2022"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	rows := tools.SliceData(string(b), "\n", func(s string) string { return s })

	totalScorePart1 := 0
	totalScorePart2 := 0
	for _, r := range rows {
		shapes := strings.Fields(r)
		if len(shapes) != 2 {
			panic("length of shapes does not equal 2!!")
		}

		totalScorePart1 += gameScorePart1(shapes[0], shapes[1])
		totalScorePart2 += gameScorePart2(shapes[0], shapes[1])
	}

	log.Printf("Part 1 answer: %d", totalScorePart1)
	log.Printf("Part 2 answer: %d", totalScorePart2)
}

func gameScorePart2(them string, you string) int {
	theirScore := shapeScore(them)
	switch you {
	case "X":
		// loss
		if theirScore == 1 {
			return 3
		}
		return theirScore - 1
	case "Y":
		// draw
		return 3 + theirScore
	case "Z":
		// win
		score := 6
		if theirScore == 3 {
			return score + 1
		}
		return score + theirScore + 1
	default:
		panic("unknown you: " + you)
	}
}

func gameScorePart1(them string, you string) int {
	score := shapeScore(you)
	switch them {
	case "A":
		if you == "X" {
			return score + 3
		}
		if you == "Y" {
			return score + 6
		}
	case "B":
		if you == "Y" {
			return score + 3
		}
		if you == "Z" {
			return score + 6
		}
	case "C":
		if you == "Z" {
			return score + 3
		}
		if you == "X" {
			return score + 6
		}
	default:
		panic("unknown them: " + them)
	}
	return score
}

func shapeScore(shape string) int {
	switch shape {
	case "A", "X":
		return 1
	case "B", "Y":
		return 2
	case "C", "Z":
		return 3
	default:
		panic("unknown shape: " + shape)
	}
}
