package main

import (
	"AoC2022/aoc_helpers"
	"testing"
)

var day13p1tests = []aoc_helpers.File_resT{
	{Fname: "example.txt", Res: 13},
	{Fname: "input.txt", Res: 6395},
}

func TestSolveDay13Part1(t *testing.T) {
	aoc_helpers.ExecTests(t, solvePart1, day13p1tests)
}

var day12p2tests = []aoc_helpers.File_resT{
	{Fname: "example.txt", Res: 140},
	{Fname: "input.txt", Res: 24921},
}

func TestSolveDay12Part2(t *testing.T) {
	aoc_helpers.ExecTests(t, solvePart2, day12p2tests)
}
