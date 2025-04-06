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

// blinkStones blinks the stones according to the rules
// Returns the number of stones after the blinks
// Uses memorization to avoid recalculating the same stone
// The memo map stores the number of stones for each stone and count
// The key is the stone and count, and the value is the number of stones
// The function is recursive
func blinkStone(stone string, count int, memo map[string]int) int {

	if count == 0 {
		return 1
	}

	// Check if the result is already in the memo
	key := fmt.Sprintf("%s,%d", stone, count)
	if val, exists := memo[key]; exists {
		return val
	}

	//	for i := 0; i < len(stones); i++ {
	n, err := strconv.Atoi(stone)
	if err != nil {
		log.Fatalf("Error converting string to int: %s", err)
	}

	res := 0
	if n == 0 {
		res = blinkStone("1", count-1, memo)
	} else {
		if len(stone)%2 == 0 {
			mid := len(stone) / 2
			left := strings.TrimLeft(stone[:mid], "0")
			right := strings.TrimLeft(stone[mid:], "0")

			if left == "" {
				left = "0"
			}
			if right == "" {
				right = "0"
			}
			res = blinkStone(left, count-1, memo) + blinkStone(right, count-1, memo)
		} else {
			res = blinkStone(strconv.Itoa(n*2024), count-1, memo)
		}
	}
	//	}
	memo[key] = res
	return res
}

func blinkStones(stones []string, count int) int {
	memo := make(map[string]int) // Initialize the memoization map
	res := 0
	for i := 0; i < len(stones); i++ {
		res += blinkStone(stones[i], count, memo)
	}
	return res
}

func solvePart1(args []string) int {
	fn := args[0]

	// Parse input file
	lines, err := aoc_helpers.ReadLines(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Split the first line into a slice of strings
	stones := strings.Split(lines[0], " ")

	return blinkStones(stones, 25)
}

func solvePart2(args []string) int {
	fn := args[0]

	// Parse input file
	lines, err := aoc_helpers.ReadLines(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Split the first line into a slice of strings
	stones := strings.Split(lines[0], " ")

	return blinkStones(stones, 75)
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
