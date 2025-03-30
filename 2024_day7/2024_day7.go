package main

import (
	"AoC2022/aoc_helpers"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var dbgFlag = flag.Bool("d", false, "debug flag")
var traceFlag = flag.Bool("t", false, "trace flag")

// parseNumbers extracts all integers from a given string and returns them as a slice of ints.
func parseNumbers(input string) ([]int, error) {
	// Regular expression to match integers (positive and negative)
	re := regexp.MustCompile(`-?\d+`)

	// Find all matches in the input string
	matches := re.FindAllString(input, -1)

	// Convert matches to integers
	numbers := make([]int, len(matches))
	for i, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			return nil, fmt.Errorf("error converting %s to int: %v", match, err)
		}
		numbers[i] = num
	}

	return numbers, nil
}

func isEqual(t int, v []int) bool {

	if len(v) == 1 {
		return t == v[0]
	}

	plus := []int{v[0] + v[1]}
	mult := []int{v[0] * v[1]}

	return isEqual(t, append(plus, v[2:]...)) || isEqual(t, append(mult, v[2:]...))
}

func solvePart1(args []string) int {
	lines, err := aoc_helpers.ReadLines(args[0])
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	sum := 0
	for n, line := range lines {
		numbers, _ := parseNumbers(line)
		if isEqual(numbers[0], numbers[1:]) {
			if *dbgFlag {
				fmt.Printf("[DEBUG] line %d: calibration line ok, value %d\n", n+1, numbers[0])
			}
			sum = sum + numbers[0]
		}
	}
	return sum
}

func solvePart2(args []string) int {
	return -1
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
