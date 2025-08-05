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
	"sort"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

var dbgFlag = flag.Bool("d", false, "debug flag")
var traceFlag = flag.Bool("t", false, "trace flag")

var leftList []int
var rightList []int

func parseFile(fn string) {

	// Read file
	lines, err := aoc_helpers.ReadLines(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Go through all lines, add sensor readings to slice
	leftList = make([]int, 0)
	rightList = make([]int, 0)

	for i, line := range lines {

		values := strings.Fields(line)

		if len(values) != 2 {
			fmt.Println("Error scanning line #", i, ":", line)
			return
		}
		v, _ := strconv.Atoi(values[0])
		leftList = append(leftList, v)
		v, _ = strconv.Atoi(values[1])
		rightList = append(rightList, v)
	}
	sort.Ints(leftList)
	sort.Ints(rightList)
}

func solvePart1(args []string) int {

	fn := args[0]
	totalDist := 0

	// Parse input file
	parseFile(fn)

	for i := range leftList {
		dist := leftList[i] - rightList[i]
		log.Debug("dist = ", dist)
		if dist < 0 {
			dist = -dist
		}
		totalDist += dist
	}
	return totalDist
}

func solvePart2(args []string) int {
	fn := args[0]

	var similarMap = make(map[int]int)
	score := 0

	// Parse input file
	leftList = nil
	rightList = nil
	parseFile(fn)

	for _, v := range leftList {
		for _, k := range rightList {
			if v < k {
				break
			}

			if v == k {
				similarMap[v] = similarMap[v] + v
			}
		}
	}
	for _, m := range similarMap {
		score = score + m
	}

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
