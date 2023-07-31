package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	calories := readFile("input_test.txt")
	if part1(calories) != 24000 {
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	calories := readFile("input_test.txt")
	if part2(calories) != 45000 {
		t.Fail()
	}
}
