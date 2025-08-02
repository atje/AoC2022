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

var p2tests = []aoc_helpers.File_resT{
	{Args: []string{"example.txt"}, Res: 11387},
	//{Args: []string{"input.txt"}, Res: 4998764814652},
}

func TestSolvePart2(t *testing.T) {
	aoc_helpers.ExecTests(t, solvePart2, p2tests)
}
func TestConcatenate(t *testing.T) {
	tests := []struct {
		name     string
		v1       int
		v2       int
		expected int
	}{
		{
			name:     "Concatenate single-digit numbers",
			v1:       1,
			v2:       2,
			expected: 12,
		},
		{
			name:     "Concatenate multi-digit numbers",
			v1:       12,
			v2:       34,
			expected: 1234,
		},
		{
			name:     "Concatenate number with zero",
			v1:       0,
			v2:       5,
			expected: 5,
		},
		{
			name:     "Concatenate zero with number",
			v1:       7,
			v2:       0,
			expected: 70,
		},
		{
			name:     "Concatenate two zeros",
			v1:       0,
			v2:       0,
			expected: 0,
		},
		{
			name:     "Concatenate negative and positive number",
			v1:       -1,
			v2:       2,
			expected: -12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := concatenate(tt.v1, tt.v2)
			if result != tt.expected {
				t.Errorf("concatenate(%d, %d) = %d; want %d", tt.v1, tt.v2, result, tt.expected)
			}
		})
	}
}
