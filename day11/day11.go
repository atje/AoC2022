/*
--- Day 11: Monkey in the Middle ---
*/

package main

import (
	"AoC2022/aoc_helpers"
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

type monkeyT struct {
	id        int
	operation func(int) int
	test      func(int) int
	items     []int
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func solveDay11Part1(lines []string, rounds int) int {
	return -1
}

func main() {

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		log.Fatalln("Please provide input file!")
	}
	// Parse input file into a matrix
	lines, err := aoc_helpers.ReadLines(args[0])
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println("Day 11, part 1 answer:", solveDay11Part1(lines, 20))
}
