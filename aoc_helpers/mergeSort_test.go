package aoc_helpers

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {

	// Test ascending order.
	sorted_slice := MergeSort([]int{38, 27, 43, 10},
		func(left int, right int) bool { return left < right })

	if !reflect.DeepEqual(sorted_slice, []int{10, 27, 38, 43}) {
		t.Fatalf("Sort failed")
	}

	// Test descending order.
	desc_sorted_slice := MergeSort([]int{11, 23, 17, 3, 5}, func(left int, right int) bool {
		return left > right
	})

	if !reflect.DeepEqual(desc_sorted_slice, []int{23, 17, 11, 5, 3}) {
		t.Fatalf("Sort failed")
	}
}
