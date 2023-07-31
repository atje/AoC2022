package main

import (
	"fmt"
	"testing"
)

type pointTest struct {
	arg0 point
	arg1 point
	res  point
}

var pointTests = []pointTest{
	{point{0, 0}, point{0, 0}, point{0, 0}},
	{point{-123, 0}, point{12, 0}, point{-135, 0}},
	{point{12, 0}, point{-123, 0}, point{135, 0}},
	{point{0, 12}, point{0, -123}, point{0, 135}},
	{point{1, 0}, point{-1, 0}, point{2, 0}},
	{point{-1, 0}, point{1, 0}, point{-2, 0}},
	{point{0, 1}, point{0, -1}, point{0, 2}},
	{point{-1, 0}, point{1, 0}, point{-2, 0}},
}

func TestPointDist(t *testing.T) {

	for _, test := range pointTests {
		output := pointDist(test.arg0, test.arg1)

		fmt.Println(output, test.res)
		if output != test.res {
			t.Errorf("Output %v not equal to expected %v", output, test.res)
		}
	}

}
