/*
--- Day 5: Supply Stacks ---
*/

package main

import (
	"AoC2022/aoc_helpers"
	"container/list"
	"flag"
	"fmt"
	"log"
	"regexp"
)

var partFlag = flag.Int("p", 0, "part 0 (default) or part 1")
var dbgFlag = flag.Bool("d", false, "debug flag")

func loadCrates(lines []string) (int, []list.List) {
	l := []list.List{}
	n := 0

	stop := false
	re := regexp.MustCompile(`\s{4}|\[[A-Z]\]`)
	for !stop && n < len(lines) {

		sarr := re.FindAllString(lines[n], -1)
		fmt.Printf("%q\n", sarr)
		if sarr == nil {
			stop = true
		}
		n++

	}
	return n, l

}

func moveCrates(l []list.List, spec string) {
	fmt.Println("Moving :", spec)
}

func topCrates([]list.List) string {
	return "empty"
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

	lines, err := aoc_helpers.ReadLines(args[0])
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Load initial crate cfg
	n, stacks := loadCrates(lines)
	if n <= 0 {
		fmt.Println("Failed loading crates")
		return
	}

	// Move crates according to spec
	for i := n + 1; i < len(lines); i++ {
		moveCrates(stacks, lines[i])
	}

	top := topCrates(stacks)

	fmt.Println("*** Part 1 ***")
	fmt.Println("Top crates per row:", top)

	// Part two
	fmt.Println("*** Part 2 ***")
	// fmt.Println("Partial or complete overlaps:", tot2)
}
