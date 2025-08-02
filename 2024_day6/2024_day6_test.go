package main

import (
	"AoC2022/aoc_helpers"
	"testing"
)

var p1tests = []aoc_helpers.File_resT{
	{Args: []string{"example.txt"}, Res: 41},
	{Args: []string{"input.txt"}, Res: 5129},
}

func TestSolvePart1(t *testing.T) {
	aoc_helpers.ExecTests(t, solvePart1, p1tests)
}

var p2tests = []aoc_helpers.File_resT{
	{Args: []string{"example.txt"}, Res: 6},
	{Args: []string{"debug2.txt"}, Res: 1},
	{Args: []string{"debug3.txt"}, Res: 1},
	{Args: []string{"debug4.txt"}, Res: 3},
	{Args: []string{"input.txt"}, Res: 1888},
}

func TestSolvePart2(t *testing.T) {
	aoc_helpers.ExecTests(t, solvePart2, p2tests)
}

func TestRotClockwise(t *testing.T) {
	tests := []struct {
		input    byte
		expected byte
	}{
		{BC_DU, BC_LR},
		{BC_LR, BC_UD},
		{BC_UD, BC_RL},
		{BC_RL, BC_DU},
	}

	for _, test := range tests {
		result := rotClockwise(test.input)
		if result != test.expected {
			t.Errorf("rotClockwise(%b) = %b; want %b", test.input, result, test.expected)
		}
	}
}
