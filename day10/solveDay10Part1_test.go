package main

import (
	"AoC2022/aoc_helpers"
	"testing"

	log "github.com/sirupsen/logrus"
)

type input struct {
	filename string
	answer   int
}

var inputs = []input{
	{"simpletest.txt", 0},
	{"largetest.txt", 13140},
}

func TestSolveDay10Part1(t *testing.T) {

	for _, test := range inputs {
		lines, err := aoc_helpers.ReadLines(test.filename)
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}

		output := solveDay10Part1(lines)

		if output != test.answer {
			t.Errorf("Output %v not equal to expected %v", output, test.answer)
		}
	}

}
