package main

import (
	"AoC2022/aoc_helpers"
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

var dbgFlag = flag.Bool("d", false, "debug flag")
var traceFlag = flag.Bool("t", false, "trace flag")

func createMap(m [][]byte) [][]byte {
	// Create a new map with the same dimensions
	newMap := make([][]byte, len(m))
	for i := range m {
		newMap[i] = make([]byte, len(m[i]))
	}
	return newMap
}

func printMap(m [][]byte) {
	for _, row := range m {
		for _, c := range row {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

func walkTrails(m [][]byte, t [][]byte, i, j int, start byte) [][]byte {
	//fmt.Printf("[DEBUG] Checking (%d, %d) with start %c\n", i, j, start)

	// Check if we are out of bounds
	if i < 0 || j < 0 || i >= len(m) || j >= len(m[i]) {
		//fmt.Printf("[DEBUG] Out of bounds: (%d, %d)\n", i, j)
		return t
	}

	if m[i][j] == '9' && start == '9' {
		//fmt.Printf("[DEBUG] Found a trailhead at (%d, %d)\n", i, j)
		t[i][j] = t[i][j] + 1
		return t
	}

	if m[i][j] == start {
		start++
		// Recursively check the adjacent positions
		t = walkTrails(m, t, i-1, j, start)
		t = walkTrails(m, t, i+1, j, start)
		t = walkTrails(m, t, i, j-1, start)
		t = walkTrails(m, t, i, j+1, start)

		return t
	}

	return t
}

func findTrailheads(m [][]byte, useRating bool) int {
	res := 0

	// Go through the map and find trailheads
	for i, row := range m {
		for j, c := range row {
			if c == '0' {
				// Found a trailhead
				t := createMap(m)
				t = walkTrails(m, t, i, j, c)

				// Print the trail
				if *dbgFlag {
					log.Debugf("[DEBUG] Trail:")
					printMap(t)
				}
				// Count the number of '1's in the trail
				count := 0
				for _, row := range t {
					for _, c := range row {
						if c > 0 {
							if useRating {
								count += int(c)
							} else {
								count++
							}
						}
					}
				}
				// Add the count to the result
				res += count
			}
		}
	}
	return res
}

func solvePart1(args []string) int {
	fn := args[0]

	// Parse input file
	m, err := aoc_helpers.ReadLinesToByteSlice(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Print the map
	if *dbgFlag {
		log.Debugf("[DEBUG] Map:")
		printMap(m)
	}

	return findTrailheads(m, false)
}

func solvePart2(args []string) int {
	fn := args[0]

	// Parse input file
	m, err := aoc_helpers.ReadLinesToByteSlice(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Print the map
	if *dbgFlag {
		log.Debugf("[DEBUG] Map:")
		printMap(m)
	}

	return findTrailheads(m, true)
}

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func setupLogging() {
	aoc_helpers.SetupLogging(dbgFlag, traceFlag)
}

func validateArgs(args []string) {
	if len(args) == 0 {
		log.Fatalln("Please provide input file!")
	}
}

func main() {

	flag.Parse()
	args := flag.Args()

	setupLogging()
	validateArgs(args)

	fmt.Println("part 1:", solvePart1(args))
	fmt.Println("part 2:", solvePart2(args))
}
