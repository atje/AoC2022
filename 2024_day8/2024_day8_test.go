package main

import (
	"AoC2022/aoc_helpers"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	tests := []aoc_helpers.File_resT{
		{Args: []string{"debug.txt"}, Res: 2},
		{Args: []string{"debug2.txt"}, Res: 4},
		{Args: []string{"example.txt"}, Res: 14},
		{Args: []string{"input.txt"}, Res: 367},
	}

	aoc_helpers.ExecTests(t, solvePart1, tests)
}

func TestSolvePart2(t *testing.T) {
	tests := []aoc_helpers.File_resT{
		{Args: []string{"debug3.txt"}, Res: 9},
		{Args: []string{"example.txt"}, Res: 34},
		{Args: []string{"input.txt"}, Res: 1285},
	}

	aoc_helpers.ExecTests(t, solvePart2, tests)
}
