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

	before, after, found := strings.Cut(string(b), "\n\n")
	if !found {
		panic("Arrgggghhhh!!!!")
	}

	moves := tools.SliceData(after, "\n", func(s string) move {
		// move 1 from 8 to 7
		d := strings.Fields(s)
		if len(d) != 6 {
			log.Printf("Expected 6 fields, got %v", d)
		}
		return move{
			count: tools.StringToInt(d[1]),
			from:  tools.StringToInt(d[3]),
			to:    tools.StringToInt(d[5]),
		}
	})

	log.Printf("Part 1: %s", part1(before, moves))
	log.Printf("Part 2: %s", part2(before, moves))
}

func part1(stackString string, moves []move) string {
	stacks := makeStacks(stackString)
	for _, m := range moves {
		var crates []string
		for i := 0; i < m.count; i++ {
			crates = append(crates, stacks[m.from].pop())
		}
		stacks[m.to].add(crates...)
	}

	topCrates := make([]string, 0, len(stacks))
	for _, s := range stacks {
		topCrates = append(topCrates, s.pop())
	}
	return strings.Join(topCrates, "")
}

func part2(stackString string, moves []move) string {
	stacks := makeStacks(stackString)
	for _, m := range moves {
		crates := make([]string, m.count)
		for i := m.count - 1; i >= 0; i-- {
			crates[i] = stacks[m.from].pop()
		}
		stacks[m.to].add(crates...)
	}

	topCrates := make([]string, 0, len(stacks))
	for _, s := range stacks {
		topCrates = append(topCrates, s.pop())
	}
	return strings.Join(topCrates, "")
}

type move struct {
	count int
	from  int
	to    int
}

func makeStacks(s string) []stack {
	stacks := make([]stack, 10)
	for i := 0; i < 10; i++ {
		stacks[i] = newStack()
	}
	rows := strings.Split(s, "\n")
	for i := len(rows) - 1; i >= 0; i-- {
		for j := 1; j <= len(rows[i]); j += 4 {
			if i == len(rows)-1 {
				continue
			}
			if string(rows[i][j]) == " " {
				continue
			}
			stacks[(j/4)+1].add(string(rows[i][j]))
		}
	}
	return stacks
}

type stack []string

func newStack() stack {
	return make(stack, 0)
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) add(in ...string) {
	for _, v := range in {
		*s = append(*s, v)
	}
}

func (s *stack) pop() string {
	if s.isEmpty() {
		return ""
	}
	var v string
	v, *s = (*s)[len(*s)-1], (*s)[:(len(*s)-1)]
	return v
}
