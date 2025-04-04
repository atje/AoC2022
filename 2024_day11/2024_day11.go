package main

import (
	"AoC2022/aoc_helpers"
	"flag"
	"fmt"
	"os"
	"slices"
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

	for i := 0; i < len(stones); i++ {
		n, err := strconv.Atoi(stones[i])
		if err != nil {
			log.Fatalf("Error converting string to int: %s", err)
		}
		if n == 0 {
			//If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
			stones[i] = "1"
		} else if len(stones[i])%2 == 0 {
			//If the stone is engraved with a number that has an even number of digits, it is replaced by two stones.
			// The left half of the digits are engraved on the new left stone, and the right half of the digits
			// are engraved on the new right stone.
			// (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
			// Split the number into two halves
			numStr := strconv.Itoa(n)
			mid := len(numStr) / 2
			left := numStr[:mid]
			right := numStr[mid:]
			// Remove leading zeros
			left = strings.TrimLeft(left, "0")
			right = strings.TrimLeft(right, "0")
			if left == "" {
				left = "0"
			}
			if right == "" {
				right = "0"
			}

			stones[i] = left
			stones = slices.Insert(stones, i+1, right)
			i++
		} else {
			//If none of the other rules apply, the stone is replaced by a new stone;
			// the old stone's number multiplied by 2024 is engraved on the new stone.
			stones[i] = strconv.Itoa(n * 2024)

		}
	}
	return stones

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
		stones = blinkOnce(stones)
		if *dbgFlag {
			fmt.Printf("After %d blinks: ", i+1)
			printStones(stones)
		}
	}

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
