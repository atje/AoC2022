package main

import (
	"AoC2022/aoc_helpers"
	"testing"
)

var mergeSortTests = []aoc_helpers.File_resT{
	{Fname: "sort_test_1.txt", Res: 13},
}

func TestMergeSort(t *testing.T) {
	aoc_helpers.ExecTests(MergeSort, MergeSortTests)
}
