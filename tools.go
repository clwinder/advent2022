package tools

import (
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

func Max[T constraints.Integer | constraints.Float](in []T) T {
	var max T
	for _, v := range in {
		if v > max {
			max = v
		}
	}
	return max
}

func Sum[T constraints.Integer | constraints.Float](in []T) T {
	var t T
	for _, v := range in {
		t += v
	}
	return t
}

func SliceData[T any](data string, separator string, conv func(string) T) []T {
	rows := strings.Split(data, separator)
	var output []T
	for _, r := range rows {
		output = append(output, conv(r))
	}
	return output
}

func Intersection[T comparable](a []T, b []T) []T {
	intersectionMap := make(map[T]struct{})
	for _, v := range a {
		for _, w := range b {
			if v == w {
				intersectionMap[v] = struct{}{}
			}
		}
	}
	var intersection []T
	for v := range intersectionMap {
		intersection = append(intersection, v)
	}
	return intersection
}

func Set[T comparable](in []T) []T {
	s := make(map[T]struct{})
	for _, v := range in {
		s[v] = struct{}{}
	}

	set := make([]T, 0, len(s))
	for k := range s {
		set = append(set, k)
	}
	return set
}

func ReadFile(name string) []byte {
	b, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return b
}

func StringToInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}
