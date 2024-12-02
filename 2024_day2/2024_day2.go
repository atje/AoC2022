/*
--- Day 15: Beacon Exclusion Zone ---

Some definitions from the text:
-

Objective:
Consult the report from the sensors you just deployed.
In the row where y=2000000, how many positions cannot contain a beacon?

Rules:
-

Approach:
-

-- Part two --

Approach:
-

*/

package main

import (
	"AoC2022/aoc_helpers"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

var dbgFlag = flag.Bool("d", false, "debug flag")
var traceFlag = flag.Bool("t", false, "trace flag")

func solvePart1(args []string) int {

	fn := args[0]
	safeReports := 0

	// Parse input file
	lines, err := aoc_helpers.ReadLines(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for _, line := range lines {

		values := strings.Fields(line)
		prev := 0
		dirInc := false
		safe := false

		for i, k := range values {
			v, _ := strconv.Atoi(k)
			if i > 0 {
				if i == 1 {
					if prev < v {
						dirInc = true
					} else if prev == v {
						safe = false
						break
					}
				}
				diff := prev - v
				if diff < 0 {
					diff = -diff
				}
				if dirInc && (diff > 0) && (diff < 4) && (prev < v) {
					safe = true
				} else if !dirInc && (diff > 0) && (diff < 4) && (prev > v) {
					safe = true
				} else {
					safe = false
					break
				}
			}
			prev = v

		}
		if safe {
			safeReports++
		}
	}

	return safeReports
}

func solvePart2(args []string) int {
	score := 0

	// Parse input file

	return score
}

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func main() {

	flag.Parse()
	args := flag.Args()

	if *dbgFlag {
		log.SetLevel(log.DebugLevel)
	} else if *traceFlag {
		log.SetLevel(log.TraceLevel)
	}

	if len(args) == 0 {
		log.Fatalln("Please provide input file!")
	}

	fmt.Println("part 1:", solvePart1(args))
	fmt.Println("part 2:", solvePart2(args))
}
