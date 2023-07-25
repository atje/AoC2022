/*
--- Day 8: Treetop Tree House ---
*/

package main

import (
	"AoC2022/aoc_helpers"
	"flag"
	"fmt"
	"log"

	"golang.org/x/exp/constraints"
)

// var partFlag = flag.Int("p", 0, "part 0 (default) or part 1")
var dbgFlag = flag.Bool("d", false, "debug flag")

// transpose rows to columns
func transpose[E constraints.Integer](matrix [][]E) [][]E {
	res := make([][]E, len(matrix))

	for _, row := range matrix {
		for j, v := range row {
			res[j] = append(res[j], v)
		}
	}

	return res
}

// Mirror matrix left-right
func mirror[E constraints.Integer](matrix [][]E) {
	for _, row := range matrix {

		for i, j := 0, len(row)-1; i < j; i, j = i+1, j-1 {
			row[i], row[j] = row[j], row[i]
		}
	}
}

// check rows, build map of visible trees from left
func checkRows(treeMap [][]rune, visible [][]int) (res [][]int) {
	if visible == nil {
		res = make([][]int, len(treeMap))
	} else {
		res = visible
	}

	for r, row := range treeMap {
		if res[r] == nil {
			res[r] = make([]int, len(row)-2)
		}
		maxVal := row[0]
		for i := 1; i < len(row)-1; i++ {
			if row[i] > maxVal {
				res[r][i-1] = 1
				maxVal = row[i]
			}
		}
	}
	if *dbgFlag {
		fmt.Println(res)
	}

	return res
}

func countOnes(matrix [][]int) int {
	res := 0

	for _, row := range matrix {
		for _, v := range row {
			if v == 1 {
				res++
			}
		}
	}

	return res
}

// Main function
func main() {

	flag.Parse()
	args := flag.Args()

	// Parse input file into a matrix
	lines, err := aoc_helpers.ReadLines(args[0])
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	m := make([][]rune, len(lines))

	for i, line := range lines {
		m[i] = []rune(line)
	}

	if *dbgFlag {
		fmt.Println(m)
	}
	fmt.Println("*** Part 1 ***")

	treeCount := 2*len(m) + 2*len(m[0]) - 4

	// Check from left
	f := checkRows(m[1:len(m)-1], nil)

	// Check from right
	mirror(m)
	mirror(f)
	f = checkRows(m[1:len(m)-1], f)

	// Check from bottom
	m = transpose(m)
	f = transpose(f)
	f = checkRows(m[1:len(m)-1], f)

	// Check from top
	mirror(m)
	mirror(f)
	f = checkRows(m[1:len(m)-1], f)

	treeCount += countOnes(f)

	fmt.Println("treeCount", treeCount)
}
