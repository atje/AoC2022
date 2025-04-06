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

func printStones(stones []string) {
	for i := 0; i < len(stones); i++ {
		fmt.Printf("%s ", stones[i])
	}
	fmt.Println()
}

func blinkOnce(stones []string) []string {
	// Create a new slice to store the updated stones
	newStones := make([]string, 0, len(stones)*2) // Preallocate capacity to avoid frequent resizing

	for _, stone := range stones {
		n, err := strconv.Atoi(stone)
		if err != nil {
			log.Fatalf("Error converting string to int: %s", err)
		}

		if n == 0 {
			// If the stone is engraved with the number 0, replace it with "1"
			newStones = append(newStones, "1")
		} else {
			numStr := strconv.Itoa(n)
			if len(numStr)%2 == 0 {
				// If the number has an even number of digits, split it into two halves
				mid := len(numStr) / 2
				left := strings.TrimLeft(numStr[:mid], "0")
				right := strings.TrimLeft(numStr[mid:], "0")

				// Handle empty strings after trimming leading zeros
				if left == "" {
					left = "0"
				}
				if right == "" {
					right = "0"
				}

				newStones = append(newStones, left, right)
			} else {
				// Otherwise, multiply the number by 2024 and replace the stone
				newStones = append(newStones, strconv.Itoa(n*2024))
			}
		}
	}

	return newStones
}

func solvePart1(args []string) int {
	fn := args[0]
	n := 1 // Default number of blinks

	if len(args) >= 2 {
		t, err := strconv.Atoi(args[1])

		if err != nil {
			log.Fatalf("Error converting string to int: %s", err)
		}
		n = t
	}

	// Parse input file
	lines, err := aoc_helpers.ReadLines(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Split the first line into a slice of strings
	stones := strings.Split(lines[0], " ")

	for i := 0; i < n; i++ {
		fmt.Printf(".")
		stones = blinkOnce(stones)
		if *dbgFlag {
			fmt.Printf("After %d blinks: %d stones\n", i+1, len(stones))
		}
	}
	fmt.Println()

	return len(stones)
}

func solvePart2(args []string) int {
	return 0
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
