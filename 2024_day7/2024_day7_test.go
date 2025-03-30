package main

import (
	"AoC2022/aoc_helpers"
	"testing"
)

func TestIsEqual(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		values   []int
		expected bool
	}{
		{
			name:     "Single value equal to target",
			target:   5,
			values:   []int{5},
			expected: true,
		},
		{
			name:     "Single value not equal to target",
			target:   5,
			values:   []int{3},
			expected: false,
		},
		{
			name:     "Two values can sum to target",
			target:   7,
			values:   []int{3, 4},
			expected: true,
		},
		{
			name:     "Two values can multiply to target",
			target:   12,
			values:   []int{3, 4},
			expected: true,
		},
		{
			name:     "Two values cannot sum or multiply to target",
			target:   10,
			values:   []int{3, 4},
			expected: false,
		},
		{
			name:     "Multiple values can form target through addition",
			target:   10,
			values:   []int{2, 3, 5},
			expected: true,
		},
		{
			name:     "Multiple values can form target through multiplication",
			target:   30,
			values:   []int{2, 3, 5},
			expected: true,
		},
		{
			name:     "Multiple values cannot form target",
			target:   20,
			values:   []int{2, 3, 5},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isEqual(tt.target, tt.values)
			if result != tt.expected {
				t.Errorf("isEqual(%d, %v) = %v; want %v", tt.target, tt.values, result, tt.expected)
			}
		})
	}
}

var p1tests = []aoc_helpers.File_resT{
	{Args: []string{"example.txt"}, Res: 3749},
	{Args: []string{"input.txt"}, Res: 4998764814652},
}

func TestSolvePart1(t *testing.T) {
	aoc_helpers.ExecTests(t, solvePart1, p1tests)
}
