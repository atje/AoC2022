package main

import (
	"AoC2022/aoc_helpers"
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

const X = 0x58
const M = 0x4D
const A = 0x41
const S = 0x53

var dbgFlag = flag.Bool("d", false, "debug flag")
var traceFlag = flag.Bool("t", false, "trace flag")

// Return number of instances found of XMAS | SAMX
// horisontal, vertical, or diagonal
func wordSearch(matrix [][]byte, x, y int) int {
	res := 0

	// Look left
	if x-3 >= 0 {
		if matrix[x-1][y] == M && matrix[x-2][y] == A && matrix[x-3][y] == S {
			res++
		}
	}
	// Look right
	if x+3 < len(matrix[x]) {
		if matrix[x+1][y] == M && matrix[x+2][y] == A && matrix[x+3][y] == S {
			res++
		}
	}
	// Look up
	if y-3 >= 0 {
		if matrix[x][y-1] == M && matrix[x][y-2] == A && matrix[x][y-3] == S {
			res++
		}
	}
	// Look down
	if y+3 < len(matrix) {
		if matrix[x][y+1] == M && matrix[x][y+2] == A && matrix[x][y+3] == S {
			res++
		}
	}
	// Diag down left
	if x-3 >= 0 && y+3 < len(matrix) {
		if matrix[x-1][y+1] == M && matrix[x-2][y+2] == A && matrix[x-3][y+3] == S {
			res++
		}
	}

	// Diag down right
	if x+3 < len(matrix[x]) && y+3 < len(matrix) {
		if matrix[x+1][y+1] == M && matrix[x+2][y+2] == A && matrix[x+3][y+3] == S {
			res++
		}
	}
	// Diag up left
	if x-3 >= 0 && y-3 >= 0 {
		if matrix[x-1][y-1] == M && matrix[x-2][y-2] == A && matrix[x-3][y-3] == S {
			res++
		}
	}
	// Diag up right
	if x+3 < len(matrix[x]) && y-3 >= 0 {
		if matrix[x+1][y-1] == M && matrix[x+2][y-2] == A && matrix[x+3][y-3] == S {
			res++
		}
	}

	return res
}

func solvePart1(args []string) int {
	fn := args[0]
	res := 0

	// Parse input file
	lines, err := aoc_helpers.ReadLinesToByteSlice(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// First, find all positions where X is
	for i := range lines {
		for j := range lines[0] {
			if lines[i][j] == X {
				res = res + wordSearch(lines, i, j)
			}
		}
	}

	return res
}

func solvePart2(args []string) int {
	fn := args[0]
	res := 0
	//	regex := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)

	// Parse input file
	_, err := aoc_helpers.ReadLines(fn)
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
