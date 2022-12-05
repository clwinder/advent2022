package main

import (
	"log"
	"sort"
	"strings"

	tools "github.com/clwinder/advent2022"
)

func main() {
	b := tools.ReadFile("input.txt")

	elves := tools.SliceData(string(b), "\n\n", func(s string) string { return s })

	calTots := totalsForElves(elves)

	log.Printf("Part 1 answer: %d", tools.Max(calTots))

	sort.Ints(calTots)
	topThree := calTots[len(calTots)-3:]

	log.Printf("Part 2 answer: %d", tools.Sum(topThree))
}

func totalsForElves(elves []string) []int {
	calTots := make([]int, len(elves))
	for i, e := range elves {
		rows := strings.Split(e, "\n")

		cals := 0
		for _, r := range rows {
			cals += tools.StringToInt(r)
		}

		calTots[i] = cals
	}
	return calTots
}
