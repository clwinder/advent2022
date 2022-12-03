package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	tools "github.com/clwinder/advent2022"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
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
			c, err := strconv.Atoi(r)
			if err != nil {
				panic(err)
			}
			cals += c
		}

		calTots[i] = cals
	}
	return calTots
}
