package main

import (
	"log"
	"os"

	tools "github.com/clwinder/advent2022"
)

const allItems = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	itemPriorities := make(map[string]int)
	for i, s := range allItems {
		itemPriorities[string(s)] = i + 1
	}

	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	rucksacks := tools.SliceData(string(b), "\n", func(s string) string { return s })

	part1Answer := part1(rucksacks, itemPriorities)
	log.Printf("Answer to part 1: %d", part1Answer)

	part2Answer := part2(rucksacks, itemPriorities)
	log.Printf("Answer to part 2: %d", part2Answer)
}

func part1(rucksacks []string, itemPriorities map[string]int) int {
	var total int
	for _, r := range rucksacks {
		items := tools.SliceData(r, "", func(s string) string { return s })
		size := len(items)

		intersection := tools.Intersection(items[:size/2], items[size/2:])
		if len(intersection) != 1 {
			panic("didn't get only one item in the intersection")
		}

		total += itemPriorities[intersection[0]]
	}
	return total
}

func part2(rucksacks []string, itemPriorities map[string]int) int {
	var total int
	for i := 0; i <= len(rucksacks)-3; i += 3 {
		items1 := tools.SliceData(rucksacks[i], "", func(s string) string { return s })
		items2 := tools.SliceData(rucksacks[i+1], "", func(s string) string { return s })
		items3 := tools.SliceData(rucksacks[i+2], "", func(s string) string { return s })

		intersect1 := tools.Intersection(items1, items2)
		intersect2 := tools.Intersection(intersect1, items3)

		if len(intersect2) != 1 {
			panic("didn't get only one item in the intersection")
		}

		total += itemPriorities[intersect2[0]]
	}

	return total
}
