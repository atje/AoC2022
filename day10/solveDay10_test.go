package main

import (
	"AoC2022/aoc_helpers"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
)

type input1 struct {
	filename string
	answer   int
}

type input2 struct {
	filename string
	answer   string
}

var inputs1 = []input1{
	{"simpletest.txt", 0},
	{"largetest.txt", 13140},
}

var inputs2 = []input2{
	{"largetest.txt",
		`##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....
`},
}

func TestSolveDay10Part1(t *testing.T) {

	for _, test := range inputs1 {
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

func TestSolveDay10Part2(t *testing.T) {

	for _, test := range inputs2 {
		lines, err := aoc_helpers.ReadLines(test.filename)
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}

		output := solveDay10Part2(lines)

		if strings.Compare(output, test.answer) != 0 {
			t.Errorf("Day10, part 2: Output \n%s not equal to expected \n%s", output, test.answer)
		}

	}
}
