package aoc_helpers

// Generic function to sort any list using provided function
// The provided function should return true if left should be sorted before right in the resulting list
func merge[T comparable](left []T, right []T, f func(T, T) bool) []T {
	merged := make([]T, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if f(left[i], right[j]) {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return merged
}

// Use merge sort to sort the packets
func MergeSort[T comparable](packets []T, f func(T, T) bool) []T {

	length := len(packets)
	if length <= 1 {
		return packets
	}

	mid := length / 2
	left := MergeSort(packets[:mid], f)
	right := MergeSort(packets[mid:], f)

	return merge(left, right, f)

}
