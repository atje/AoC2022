package main

import (
	"AoC2022/aoc_helpers"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	tests := []aoc_helpers.File_resT{
		{Args: []string{"example.txt"}, Res: 55312},
		{Args: []string{"input.txt"}, Res: 190865},
	}

	aoc_helpers.ExecTests(t, solvePart1, tests)
}

func TestSolvePart2(t *testing.T) {
	tests := []aoc_helpers.File_resT{
		//{Args: []string{"debug.txt"}, Res: 2},
		//{Args: []string{"debug3.txt"}, Res: 13},
		//{Args: []string{"example.txt"}, Res: 81},
		{Args: []string{"input.txt"}, Res: 225404711855335},
	}

	aoc_helpers.ExecTests(t, solvePart2, tests)
}
