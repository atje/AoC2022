/*
--- Day 10: Cathode-Ray Tube ---
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

type opFnT func(int, int) int

type opT struct {
	o string // operation id
	c int    // cycles to execute
	f opFnT  // operation function
}

type cpuT struct {
	cycle int // CPU cycle
	x     int // Register 'x' value
	op    opT // Current operation
	arg   int // Operation argument
	opcnt int // Cycle count remaining for current operation
}

func (c *cpuT) tick() {
	log.Tracef("tick() called, starting cycle %d, opcnt %d", c.cycle, c.opcnt)

	if calcSignalStrenghtAt[c.cycle] {
		signalStrength += c.cycle * c.x
		log.Tracef("Time to count signal strength, x = %d, cycle = %d, signalStrength = %d",
			cpu.x, cpu.cycle, signalStrength)
	}

	if c.opcnt > 0 {
		c.opcnt--
	}

	if c.opcnt == 0 {
		c.x = c.op.f(c.x, c.arg)
		log.Tracef("tick() executing operation %s", c.op.o)
	}
	c.cycle++
}

var cpu cpuT

var ops = map[string]opT{
	"noop": {"noop", 1, func(reg, i int) int { return reg }},
	"addx": {"addx", 2, func(reg, i int) int { return reg + i }},
}

var signalStrength int = 0
var calcSignalStrenghtAt = map[int]bool{
	20: true, 60: true, 100: true, 140: true, 180: true, 220: true,
}

func solveDay10Part1(lines []string) int {

	// initialize cpu
	cpu.x = 1
	cpu.cycle = 1

	for _, line := range lines {
		s := strings.Split(line, " ")

		op := ops[s[0]]
		arg := 0

		if len(s) > 1 {
			arg, _ = strconv.Atoi(s[1])
		}

		log.Tracef("op %s, arg %d", op.o, arg)
		cpu.op, cpu.arg, cpu.opcnt = op, arg, op.c

		log.Tracef("CPU %v", cpu)
		for cpu.opcnt > 0 {
			cpu.tick()
		}
	}

	return signalStrength
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
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

	fmt.Println("Day 10, part 1 answer:", solveDay10Part1(lines))
}
