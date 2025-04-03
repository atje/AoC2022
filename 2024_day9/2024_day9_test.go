package main

import (
	"AoC2022/aoc_helpers"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	tests := []aoc_helpers.File_resT{
		{Args: []string{"debug.txt"}, Res: 60},
		//	{Args: []string{"debug2.txt"}, Res: 4},
		{Args: []string{"example.txt"}, Res: 1928},
		{Args: []string{"input.txt"}, Res: 6395800119709},
	}

	aoc_helpers.ExecTests(t, solvePart1, tests)
}
