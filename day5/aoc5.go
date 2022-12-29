/*
--- Day 5: Supply Stacks ---
*/

package main

import (
	"AoC2022/aoc_helpers"
	"bytes"
	"flag"
	"fmt"
	"log"
	"regexp"
)

var partFlag = flag.Int("p", 0, "part 0 (default) or part 1")
var dbgFlag = flag.Bool("d", false, "debug flag")

func loadCrates(lines []string) (int, [][]rune) {
	l := make([][]rune, 10)
	n := 0

	stop := false
	re := regexp.MustCompile(`\s{4}|\[([A-Z])\]`)
	for !stop && n < len(lines) {

		sarr := re.FindAllStringSubmatch(lines[n], -1)
		//fmt.Printf("%q\n", sarr)
		if sarr == nil {
			stop = true
		} else {
			for i, str := range sarr {
				if str[1] != "" {
					chars := []rune(str[1])
					l[i+1] = append(l[i+1], chars[0])
				}
			}
		}
		n++

	}
	if *dbgFlag {
		fmt.Printf("Initial crates: \n%q\n", l)
	}
	return n, l

}

func moveCrates(l [][]rune, spec string, part int) [][]rune {
	if *dbgFlag {
		fmt.Println("Moving :", spec)
	}
	var from, to, quant int
	n, _ := fmt.Sscanf(spec, `move %d from %d to %d`, &quant, &from, &to)
	if n != 3 {
		fmt.Println("Failed to read spec: ", spec)
		return l
	}
	if part == 0 {
		for d := 0; d < quant; d++ {
			x, a := l[from][0], l[from][1:]
			l[from] = a
			//fmt.Println("crate ", x)

			l[to] = append([]rune{x}, l[to]...)
		}
	} else {
		if quant == 1 {
			x, a := l[from][0], l[from][quant:]
			l[from] = a
			l[to] = append([]rune{x}, l[to]...)
		} else {
			x, a := l[from][0:quant], l[from][quant:]
			l[from] = a
			l[to] = append(x, l[to]...)
		}
	}
	if *dbgFlag {
		fmt.Printf("%q\n", l)
	}
	return l
}

func printTopCrates(l [][]rune) {
	chars := bytes.Buffer{}

	for i := 1; i < len(l); i++ {
		if len(l[i]) > 0 {
			chars.WriteRune(l[i][0])
		}
	}
	fmt.Println("Top crates per row:", chars.String())
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
		moveCrates(stacks, lines[i], *partFlag)
	}

	fmt.Println("*** Part 1 ***")
	printTopCrates(stacks)

	// Part two
	fmt.Println("*** Part 2 ***")
	// fmt.Println("Partial or complete overlaps:", tot2)
}
