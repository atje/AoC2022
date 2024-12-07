package main

import (
	"AoC2022/aoc_helpers"
	"testing"
)

var p1tests = []aoc_helpers.File_resT{
	{Args: []string{"example.txt"}, Res: 18},
	{Args: []string{"input.txt"}, Res: 2591},
}

func TestSolvePart1(t *testing.T) {
	aoc_helpers.ExecTests(t, solvePart1, p1tests)
}

var p2tests = []aoc_helpers.File_resT{
	{Args: []string{"example2.txt"}, Res: 48},
	// {Args: []string{"input.txt"}, Res: 88802350},
}

func TestSolvePart2(t *testing.T) {
	aoc_helpers.ExecTests(t, solvePart2, p2tests)
}
