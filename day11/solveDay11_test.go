package main

import (
	"AoC2022/aoc_helpers"
	"testing"

	log "github.com/sirupsen/logrus"
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
		lines, err := aoc_helpers.ReadLines(test.filename)
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}

		output := solveDay11Part1(lines, 20)

		if output != test.answer {
			t.Errorf("Output %v not equal to expected %v", output, test.answer)
		}
	}

}
