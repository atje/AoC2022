package main

import (
	"fmt"
	"testing"
)

var moveIncrTests = []pointTest{
	// Zero-move cases, there are nine of them
	{point{0, 0}, point{0, 0}, point{0, 0}},
	{point{1, 0}, point{0, 0}, point{0, 0}},
	{point{0, 1}, point{0, 0}, point{0, 0}},
	{point{-1, 0}, point{0, 0}, point{0, 0}},
	{point{0, -1}, point{0, 0}, point{0, 0}},
	{point{1, 1}, point{0, 0}, point{0, 0}},
	{point{-1, -1}, point{0, 0}, point{0, 0}},
	{point{-1, 1}, point{0, 0}, point{0, 0}},
	{point{1, -1}, point{0, 0}, point{0, 0}},

	// x-axis move cases
	{point{-2, 0}, point{0, 0}, point{-1, 0}},
	{point{2, 0}, point{0, 0}, point{1, 0}},
	{point{-3, 0}, point{-1, 0}, point{-1, 0}},
	{point{1, 0}, point{-1, 0}, point{1, 0}},

	// y-axis move cases
	{point{-2, 1}, point{0, 0}, point{-1, 1}},
	{point{-2, -1}, point{0, 0}, point{-1, -1}},

	// diagonal moves
	{point{-1, 1}, point{1, 0}, point{-1, 1}},
	{point{-1, 1}, point{1, 2}, point{-1, -1}},

	{point{0, 0}, point{1, 2}, point{-1, -1}},
	{point{-2, 3}, point{1, 2}, point{-1, 1}},
}

func TestMoveIncr(t *testing.T) {

	for _, test := range moveIncrTests {
		output := moveIncr(test.arg0, test.arg1)

		fmt.Println(output, test.res)
		if output != test.res {
			t.Errorf("Output %v not equal to expected %v", output, test.res)
		}
	}

}
