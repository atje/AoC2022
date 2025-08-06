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
var dumpNOKReport = flag.Bool("n", false, "dump invalid reports")

func isAdjacent(v1, v2 int) bool {
	diff := v1 - v2
	if diff < 0 {
		diff = -diff
	}
	if (diff > 0) && (diff < 4) {
		return true
	}
	return false
}

func deleteElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func checkReport(n int, values []int) bool {
	prev := -1
	dirInc := false
	safe := false

	for i, v := range values {
		if i > 0 {
			if i == 1 {
				if prev < v {
					dirInc = true
				} else if prev == v {
					safe = false
					break
				}
			}
			if isAdjacent(prev, v) && dirInc && (prev < v) {
				safe = true
			} else if isAdjacent(prev, v) && !dirInc && (prev > v) {
				safe = true
			} else {
				safe = false
				break
			}
		}
		prev = v

	}

	if safe {
		log.Printf("%d: Report OK\n", n+1)
	} else {
		log.Printf("%d: Report NOK\n", n+1)
		if *dumpNOKReport {
			fmt.Println(values)
		}
	}

	return safe
}

func solvePart1(args []string) int {

	fn := args[0]
	safeReports := 0

	// Parse input file
	lines, err := aoc_helpers.ReadLines(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for n, line := range lines {

		values := strings.Fields(line)
		var valuesInt []int

		for _, k := range values {
			v, _ := strconv.Atoi(k)
			valuesInt = append(valuesInt, v)
		}
		if checkReport(n+1, valuesInt) {
			safeReports++
		}
	}
	return safeReports
}

func solvePart2(args []string) int {
	fn := args[0]

	safeReports := 0

	// Parse input file
	lines, err := aoc_helpers.ReadLines(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for n, line := range lines {

		values := strings.Fields(line)
		var valuesInt []int

		for _, k := range values {
			v, _ := strconv.Atoi(k)
			valuesInt = append(valuesInt, v)
		}
		res := checkReport(n+1, valuesInt)
		if res {
			log.Printf("%d: Report OK\n", n+1)
			safeReports++
		} else {
			for i := range valuesInt {
				s1 := make([]int, len(valuesInt))
				copy(s1, valuesInt)
				res = checkReport(n+1, deleteElement(s1, i))
				if res {
					break
				}
			}
			if res {
				safeReports++
			}
		}
	}

	return safeReports
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
