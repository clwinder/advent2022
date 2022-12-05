package main

import (
	"fmt"
	"log"
	"strings"

	tools "github.com/clwinder/advent2022"
)

func main() {
	b := tools.ReadFile("input.txt")

	pairs := tools.SliceData(string(b), "\n", func(s string) pair {
		p := strings.Split(s, ",")
		if len(p) != 2 {
			panic(fmt.Sprintf("Expected two pairs, got %v", p))
		}
		return pair{
			first:  stringToVals(p[0]),
			second: stringToVals(p[1]),
		}
	})

	log.Printf("Part 1 answer: %d", part1(pairs))
	log.Printf("Part 2 answer: %d", part2(pairs))
}

func part1(pairs []pair) int {
	var count int
	for _, p := range pairs {
		if p.first.min <= p.second.min && p.first.max >= p.second.max {
			count++
			continue
		}
		if p.second.min <= p.first.min && p.second.max >= p.first.max {
			count++
			continue
		}
	}
	return count
}

func part2(pairs []pair) int {
	var count int
	for _, p := range pairs {
		if p.second.min > p.first.max {
			continue
		}
		if p.second.max < p.first.min {
			continue
		}
		count++
	}
	return count
}

type pair struct {
	first  vals
	second vals
}

type vals struct {
	min int
	max int
}

func stringToVals(s string) vals {
	v := strings.Split(s, "-")
	if len(v) != 2 {
		panic(fmt.Sprintf("Expected two values, got %v", v))
	}
	return vals{
		min: tools.StringToInt(v[0]),
		max: tools.StringToInt(v[1]),
	}
}
