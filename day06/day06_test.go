package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := map[string]struct {
		in       []byte
		expected int
	}{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb": {
			in:       []byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb"),
			expected: 7,
		},
		"bvwbjplbgvbhsrlpgdmjqwftvncz": {
			in:       []byte("bvwbjplbgvbhsrlpgdmjqwftvncz"),
			expected: 5,
		},
		"nppdvjthqldpwncqszvftbrmjlhg": {
			in:       []byte("nppdvjthqldpwncqszvftbrmjlhg"),
			expected: 6,
		},
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": {
			in:       []byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"),
			expected: 10,
		},
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw": {
			in:       []byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"),
			expected: 11,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := part1(test.in)
			if actual != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, actual)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := map[string]struct {
		in       []byte
		expected int
	}{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb": {
			in:       []byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb"),
			expected: 19,
		},
		"bvwbjplbgvbhsrlpgdmjqwftvncz": {
			in:       []byte("bvwbjplbgvbhsrlpgdmjqwftvncz"),
			expected: 23,
		},
		"nppdvjthqldpwncqszvftbrmjlhg": {
			in:       []byte("nppdvjthqldpwncqszvftbrmjlhg"),
			expected: 23,
		},
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": {
			in:       []byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"),
			expected: 29,
		},
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw": {
			in:       []byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"),
			expected: 26,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := part2(test.in)
			if actual != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, actual)
			}
		})
	}
}
