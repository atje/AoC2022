package main

import (
	"AoC2022/aoc_helpers"
	"reflect"
	"testing"
)

var p1tests = []aoc_helpers.File_resT{
	{Args: []string{"example.txt"}, Res: 24},
	{Args: []string{"input.txt"}, Res: 698},
}

func TestSolveDay14Part1(t *testing.T) {
	aoc_helpers.ExecTests(t, solvePart1, p1tests)
}

var p2tests = []aoc_helpers.File_resT{
	{Args: []string{"example.txt"}, Res: 93},
	{Args: []string{"input.txt"}, Res: 28594},
}

func TestSolveDay14Part2(t *testing.T) {
	aoc_helpers.ExecTests(t, solvePart2, p2tests)
}

func TestExpandCM(t *testing.T) {
	cm := CaveMap{
		x0: 500,
		y0: 0,
		point: [][]rune{
			{0}},
	}

	expectedCM := []CaveMap{
		{
			x0: 499,
			y0: 0,
			point: [][]rune{
				{0, 0},
				{0, 0}},
		},
		{
			x0: 498,
			y0: 0,
			point: [][]rune{
				{0, 0, 0},
				{0, 0, 0}},
		},
		{
			x0: 494,
			y0: 0,
			point: [][]rune{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0}},
		},
	}

	inputCoords := [][]int{
		{499, 1},
		{498, 1},
		{494, 9},
	}

	for i, expected := range expectedCM {

		t.Logf("Running expandCM test #%d", i)
		updatedCM := expandCM(cm, inputCoords[i][0], inputCoords[i][1])

		if !reflect.DeepEqual(expected, updatedCM) {
			t.Fatalf("expandCM test #%d failed!", i)
		}

	}
}
