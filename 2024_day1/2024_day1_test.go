package main

import (
	"AoC2022/aoc_helpers"
	"testing"
)

var p1tests = []aoc_helpers.File_resT{
	{Args: []string{"example.txt"}, Res: 11},
	{Args: []string{"input.txt"}, Res: 1223326},
}

func TestSolvePart1(t *testing.T) {
	aoc_helpers.ExecTests(t, solvePart1, p1tests)
}

var p2tests = []aoc_helpers.File_resT{
	{Args: []string{"example.txt"}, Res: 31},
	{Args: []string{"input.txt"}, Res: 21070419},
}

func TestSolvePart2(t *testing.T) {
	aoc_helpers.ExecTests(t, solvePart2, p2tests)
}
