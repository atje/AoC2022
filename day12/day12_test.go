package main

import (
	"AoC2022/aoc_helpers"
	"testing"
)

var day12p1tests = []aoc_helpers.File_resT{
	{Fname: "example.txt", Res: 31},
	{Fname: "input.txt", Res: 420},
}

func TestSolveDay12Part1(t *testing.T) {
	aoc_helpers.ExecTests(t, solveDay12Part1, day12p1tests)
}
