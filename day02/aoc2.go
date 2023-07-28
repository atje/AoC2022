/*
--- Day 2: Rock Paper Scissors ---
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"AoC2022/aoc_helpers"
)

func decodeMove(m string) int {
	var ret int = -1

	switch {
	case m == "A":
		ret = 0
	case m == "X":
		ret = 0
	case m == "B":
		ret = 1
	case m == "Y":
		ret = 1
	case m == "C":
		ret = 2
	case m == "Z":
		ret = 2
	default:
		fmt.Println("Unknown input: ", m)
	}
	return ret
}

var partFlag = flag.Int("p", 0, "part 0 (default) or part 1")
var dbgFlag = flag.Bool("d", false, "debug flag")

var costArr = [2][3][3]int{{
	{4, 8, 3}, {1, 5, 9}, {7, 2, 6}},
	{{3, 4, 8}, {1, 5, 9}, {2, 6, 7}},
}

func main() {

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	if *partFlag > 1 {
		fmt.Println("p flag not 0 or 1, aborting!")
		return
	}

	lines, err := aoc_helpers.ReadLines(args[0])
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	score := 0
	for i, line := range lines {
		if *dbgFlag {
			fmt.Print("round ", i, ": ", line)
		}
		s := strings.Split(line, " ")

		// Add points for win or draw
		score += costArr[*partFlag][decodeMove(s[0])][decodeMove(s[1])]
		if *dbgFlag {
			fmt.Println("--> score", score)
		}
	}

	fmt.Println("Total score:", score)
}
