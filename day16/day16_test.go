package main

import (
	"AoC2022/aoc_helpers"
	"testing"
)

var p1tests = []aoc_helpers.File_resT{
	{Args: []string{"example.txt"}, Res: 1651},
	{Args: []string{"input.txt"}, Res: 2359},
}

func TestSolveDay16Part1(t *testing.T) {
	aoc_helpers.ExecTests(t, solvePart1, p1tests)
}

/*
var p2tests = []aoc_helpers.File_resT{
	{Args: []string{"example.txt", "10", "20"}, Res: 56000011},
	{Args: []string{"input.txt", "4000000", "4000000"}, Res: 13360899249595},
}

func TestSolveDay16Part2(t *testing.T) {
	aoc_helpers.ExecTests(t, solvePart2, p2tests)
}
*/
