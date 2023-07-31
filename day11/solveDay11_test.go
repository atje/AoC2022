package main

import (
	"testing"
)

type input1 struct {
	filename string
	rounds   int
	answer   int
}

var inputs1 = []input1{
	{"example.txt", 20, 10605},
}

func TestSolveDay10Part1(t *testing.T) {

	for _, test := range inputs1 {
		output := solveDay11Part1(test.filename, test.rounds)

		if output != test.answer {
			t.Errorf("Output %v not equal to expected %v", output, test.answer)
		}
	}

}
