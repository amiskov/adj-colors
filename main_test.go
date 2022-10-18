package main

import (
	"testing"
)

type suite struct {
	colors   [][]color
	expected int
}

func TestColorMatrix(t *testing.T) {
	cases := []suite{
		{
			colors:   [][]color{},
			expected: 0,
		},
		{
			colors: [][]color{
				{"🔴", "🔴", "🔴"},
				{"🔴", "🔴", "🔴"},
				{"🔴", "🔴", "🔴"},
			},
			expected: 9,
		},
		{
			colors: [][]color{
				{"🔴", "🔴", "🔴"},
				{"🔴", "🔴", "🔴"},
				{"🔴", "⚫", "🔴"},
				{"🔴", "🔴", "🔴"},
				{"🔴", "🔴", "🔴"},
			},
			expected: 14,
		},
		{
			colors: [][]color{
				{"🔴", "🟢", "🟢"},
				{"🔴", "🔴", "🔴"},
				{"🔴", "🟢", "🟢"},
			},
			expected: 5,
		},
		{
			colors: [][]color{
				{"🔴", "🟢", "🟢"},
				{"🔴", "🟢", "🟢"},
				{"🔴", "🔴", "⚫️"},
			},
			expected: 4,
		},
		{
			colors: [][]color{
				{"🔴", "🟢", "⚫"},
				{"⚫", "🔴", "🟢️"},
			},
			expected: 1,
		},
	}

	for _, c := range cases {
		got := longestAdjColors(c.colors)
		if got != c.expected {
			t.Error("Got:", got, "Expected:", c.expected)
		}
	}
}
