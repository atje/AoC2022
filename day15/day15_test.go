package main

import (
	"AoC2022/aoc_helpers"
	"testing"
)

var p1tests = []aoc_helpers.File_resT{
	{Args: []string{"example.txt", "10"}, Res: 26},
	{Args: []string{"input.txt", "2000000"}, Res: 5832528},
}

func TestSolveDay15Part1(t *testing.T) {
	aoc_helpers.ExecTests(t, solvePart1, p1tests)
}

var p2tests = []aoc_helpers.File_resT{
	{Args: []string{"example.txt", "20"}, Res: 56000011},
	//{Fname: "input.txt", Res: 28594},
}

func TestSolveDay15Part2(t *testing.T) {
	aoc_helpers.ExecTests(t, solvePart2, p2tests)
}
