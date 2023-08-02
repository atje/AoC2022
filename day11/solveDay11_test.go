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
	{"input.txt", 20, 55944},
}

var inputs2 = []input1{
	{"example.txt", 1, 6 * 4},
	{"example.txt", 20, 103 * 99},
	{"example.txt", 1000, 5204 * 5192},
	{"example.txt", 2000, 10419 * 10391},
	{"example.txt", 10000, 52166 * 52013},
}

func TestSolveDay10Part1(t *testing.T) {

	for i, test := range inputs1 {
		output := solveDay11Part1(test.filename, test.rounds)

		if output != test.answer {
			t.Errorf("Test %d: Output %v != expected %v", i, output, test.answer)
		}
	}

}

func TestSolveDay10Part2(t *testing.T) {

	for i, test := range inputs2 {
		output := solveDay11Part2(test.filename, test.rounds)

		if output != test.answer {
			t.Errorf("Test %d: Out(%v) != Exp(%v)", i, output, test.answer)
		}
	}

}
