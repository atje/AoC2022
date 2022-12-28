/*
--- Day 4: Camp Cleanup ---
*/

package main

import (
	"AoC2022/aoc_helpers"
	"flag"
	"fmt"
	"log"
)

var partFlag = flag.Int("p", 0, "part 0 (default) or part 1")
var dbgFlag = flag.Bool("d", false, "debug flag")

func prettyPrintRange(start int, end int) {
	for i := 1; i <= end; i++ {
		if i < start {
			fmt.Print(".")
		} else {
			fmt.Print("-")
		}
	}
	fmt.Print("\t", start, "-", end, "\n")
}

func checkInRange(v int, start int, end int) bool {
	if v >= start && v <= end {
		return true
	}
	return false
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
	tot2 := 0
	for i := 0; i < len(lines); i++ {
		if *dbgFlag {
			fmt.Println("Row #", i, ": ", lines[i])
		}

		var e1p1, e1p2, e2p1, e2p2 int

		n, _ := fmt.Sscanf(lines[i], "%d-%d,%d-%d", &e1p1, &e1p2, &e2p1, &e2p2)
		if n != 4 {
			fmt.Println("Error scanning line #", i, ", n=", n, ":", lines[i])
			return
		}

		if *dbgFlag {
			prettyPrintRange(e1p1, e1p2)
			prettyPrintRange(e2p1, e2p2)
		}

		// Part 1
		if e1p1 >= e2p1 && e1p2 <= e2p2 {
			tot1++

			if *dbgFlag {
				fmt.Println("Found p1 within p2, row", i)
			}

		} else if e2p1 >= e1p1 && e2p2 <= e1p2 {
			tot1++

			if *dbgFlag {
				fmt.Println("Found p2 within p1, row", i)
			}
		}

		// Part 2
		if checkInRange(e1p1, e2p1, e2p2) || checkInRange(e1p2, e2p1, e2p2) ||
			checkInRange(e2p1, e1p1, e1p2) || checkInRange(e2p2, e1p1, e1p2) {
			tot2++
		}
	}

	fmt.Println("*** Part 1 ***")
	fmt.Println("Complete overlaps:", tot1)

	// Part two
	fmt.Println("*** Part 2 ***")
	fmt.Println("Partial or complete overlaps:", tot2)
}
