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
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

var dbgFlag = flag.Bool("d", false, "debug flag")
var traceFlag = flag.Bool("t", false, "trace flag")

func solvePart1(args []string) int {
	fn := args[0]
	res := 0
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	// Parse input file
	lines, err := aoc_helpers.ReadLines(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for _, line := range lines {
		m := regex.FindAllSubmatch([]byte(line), -1)
		//fmt.Printf("%q\n", m)
		for _, mul := range m {
			v1, _ := strconv.Atoi(string(mul[1]))
			v2, _ := strconv.Atoi(string(mul[2]))

			res = res + v1*v2
		}
	}
	return res
}

func solvePart2(args []string) int {
	fn := args[0]
	res := 0
	regex := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)

	// Parse input file
	lines, err := aoc_helpers.ReadLines(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	enabled := true
	for _, line := range lines {
		m := regex.FindAllSubmatch([]byte(line), -1)
		//		fmt.Printf("%q\n", m)
		for _, mul := range m {
			if strings.Contains(string(mul[0]), "don") {
				enabled = false
			} else if strings.Contains(string(mul[0]), "do") {
				enabled = true
			} else if enabled {
				v1, _ := strconv.Atoi(string(mul[1]))
				v2, _ := strconv.Atoi(string(mul[2]))

				res = res + v1*v2
			}
		}
	}
	return res
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
