/*
--- Day 4: Camp Cleanup ---
*/

package main

import (
	"AoC2022/aoc_helpers"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var partFlag = flag.Int("p", 0, "part 0 (default) or part 1")
var dbgFlag = flag.Bool("d", false, "debug flag")

func prettyPrintRange(startStr string, endStr string) {
	start, _ := strconv.Atoi(startStr)
	end, _ := strconv.Atoi(endStr)
	for i := 1; i < end; i++ {
		if i < start {
			fmt.Print(".")
		} else {
			fmt.Print("-")
		}
	}
	fmt.Print("\t", start, "-", end, "\n")
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

	tot1 := 0
	for i := 0; i < len(lines); i++ {
		if *dbgFlag {
			fmt.Println("Row #", i, ": ", lines[i])
		}

		e1, e2, found := strings.Cut(lines[i], ",")
		if !found {
			fmt.Println("Error splitting string into pairs: ", lines[i])
			return
		}
		e1p1, e1p2, found := strings.Cut(e1, "-")
		if !found {
			fmt.Println("Error splitting string into numbers: ", lines[i])
			return
		}
		if *dbgFlag {
			prettyPrintRange(e1p1, e1p2)
		}

		e2p1, e2p2, found := strings.Cut(e2, "-")
		if !found {
			fmt.Println("Error splitting string into numbers: ", lines[i])
			return
		}
		if *dbgFlag {
			prettyPrintRange(e2p1, e2p2)
		}

		if e1p1 >= e2p1 && e1p2 <= e2p2 {
			tot1++

			if *dbgFlag {
				fmt.Println("Found p1 within p2, row", i, ": ", e1, ",", e2)
			}

		} else if e2p1 >= e1p1 && e2p2 <= e1p2 {
			tot1++

			if *dbgFlag {
				fmt.Println("Found p2 within p1, row", i, ": ", e1, ",", e2)
			}
		}

	}

	fmt.Println("*** Part 1 ***")
	fmt.Println("Complete overlaps:", tot1)

	// Part two
}
