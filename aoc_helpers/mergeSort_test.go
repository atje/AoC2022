package aoc_helpers

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	// Test ascending order.
	fn := func(left int, right int) bool {
		return left < right
	}
	sorted_slice := MergeSort([]int{38, 27, 43, 10}, fn)
	sorted := [4]int{}
	copy(sorted[:], sorted_slice)
	if sorted != [4]int{10, 27, 38, 43} {
		t.Fatalf("Sort failed")
	}
	// Test descending order.
	fn = func(left int, right int) bool {
		return left > right
	}
	desc_sorted_slice := MergeSort([]int{11, 23, 17, 3, 5}, fn)
	desc_sorted := [5]int{}
	copy(desc_sorted[:], desc_sorted_slice)
	if desc_sorted != [5]int{23, 17, 11, 5, 3} {
		t.Fatalf("Sort failed")
	}
}
