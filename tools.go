package tools

import (
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
