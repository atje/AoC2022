/*
--- Day 6: Tuning Trouble ---
*/

package main

import (
	"AoC2022/aoc_helpers"
	"flag"
	"fmt"
	"log"
)

var partFlag = flag.Int("p", 0, "part 0 (default) or part 1")
var dbgFlag = flag.Bool("d", false, "debug flag")

// Find first nchar unique char sequence in string
// returns it's position
func findMarker(str string, nchar int) int {

	slice := make([]rune, 0)

	for i, c := range str {
		found := -1
		for j := 0; j < len(slice); j++ {
			if c == slice[j] {
				found = j
				break
			}
		}
		if found >= 0 {
			if found == 0 {
				_, slice = slice[0], slice[1:]

			} else {
				_, slice = slice[0:found+1], slice[found+1:]
			}
			if *dbgFlag {
				fmt.Println("Repeated char found at position", i, "(", string(c), ", ", found, ")")
			}
		} //else {
		slice = append(slice, c)
		if len(slice) > nchar {
			return i
		}
		//		}
		if *dbgFlag {
			fmt.Println("Marker sequence:", string(slice))
		}
	}

	return 0
}

func main() {

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	if *partFlag > 1 {
		fmt.Println("p flag not 0 or 1, aborting!")
		return
	}

	pos := 0
	nchar := 4

	if *partFlag == 1 {
		nchar = 14
	}

	lines, err := aoc_helpers.ReadLines(args[0])
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println("*** Part", *partFlag+1, "***")
	for i, line := range lines {
		pos = findMarker(line, nchar)
		if pos > 0 {
			fmt.Println("Line", i, ", found marker at position", pos)
		}
	}
}
