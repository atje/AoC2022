package main

import (
	"testing"
)

type uniqueTest struct {
	arg []point
	res []point
}

var uniqueTests = []uniqueTest{
	{[]point{{0, 0}}, []point{{0, 0}}},
	{[]point{{0, 0}, {0, 0}}, []point{{0, 0}}},
	{[]point{{0, 0}, {0, 0}, {0, 0}}, []point{{0, 0}}},
	{[]point{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 1}, {1, 2}, {2, 2}, {3, 2}, {4, 2}, {3, 3}, {4, 3}, {4, 3}, {2, 4}, {3, 4}},
		[]point{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 1}, {1, 2}, {2, 2}, {3, 2}, {4, 2}, {3, 3}, {4, 3}, {2, 4}, {3, 4}}},
}

func testEq(a, b []point) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestUniqueSlice(t *testing.T) {

	for _, test := range uniqueTests {
		output := uniqueSlice(test.arg)

		if !testEq(output, test.res) {
			t.Errorf("Output %q not equal to expected %q", output, test.res)
		}
	}

}
