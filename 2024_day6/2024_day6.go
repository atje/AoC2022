package main

import (
	"AoC2022/aoc_helpers"
	"flag"
	"fmt"
	"os"
	"regexp"

	log "github.com/sirupsen/logrus"
)

// Guard movement direction
const G_U = 0
const G_D = 1
const G_R = 2
const G_L = 3
const OBSTACLE = 0x23
const BREADCRUMB = 0x58

var dbgFlag = flag.Bool("d", false, "debug flag")
var traceFlag = flag.Bool("t", false, "trace flag")

func solvePart1(args []string) int {
	fn := args[0]
	res := 0
	regex1 := regexp.MustCompile(`\^`)

	// Parse input file
	lines, err := aoc_helpers.ReadLinesToByteSlice(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	guardX := -1
	guardY := -1
	guardPos := G_U

	// Find starting position of guard
	for i, l := range lines {
		loc := regex1.FindIndex(l)
		if loc != nil {
			guardX = i
			guardY = loc[0]
			//fmt.Println("Found guard in position ", guardX, guardY)
			break
		}
	}

	// Simulate guard walking until she leaves map
	for {
		newX := guardX
		newY := guardY

		lines[guardX][guardY] = BREADCRUMB

		// Check movement direction
		switch guardPos {
		case G_U:
			newX = newX - 1
		case G_D:
			newX = newX + 1
		case G_R:
			newY = newY + 1
		case G_L:
			newY = newY - 1
		}

		// Check if guard leaves the matrix
		if newX < 0 || newY < 0 || newX >= len(lines) || newY >= len(lines[0]) {
			break
		}

		// Check for obstacles in new position, turn right 90 degrees
		if lines[newX][newY] == OBSTACLE {
			switch guardPos {
			case G_U:
				newX = newX + 1
				newY = newY + 1
				guardPos = G_R
			case G_D:
				newX = newX - 1
				newY = newY - 1
				guardPos = G_L
			case G_R:
				newX = newX + 1
				newY = newY - 1
				guardPos = G_D
			case G_L:
				newX = newX - 1
				newY = newY + 1
				guardPos = G_U
			}
		}

		// Check if guard leaves the matrix
		if newX < 0 || newY < 0 || newX >= len(lines) || newY >= len(lines[0]) {
			break
		}

		guardX = newX
		guardY = newY
	}

	// Count breadcrumbs on the map
	for _, l := range lines {
		for _, b := range l {
			if b == BREADCRUMB {
				res = res + 1
			}
		}
	}

	return res
}

func solvePart2(args []string) int {
	fn := args[0]
	res := 0

	// Parse input file
	_, err := aoc_helpers.ReadLinesToByteSlice(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
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
