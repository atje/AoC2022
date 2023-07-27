/*
--- Day 3: Rucksack Reorganization ---
*/

package main

import (
	"AoC2022/aoc_helpers"
	"bytes"
	"flag"
	"fmt"
	"log"
	"strings"
)

var partFlag = flag.Int("p", 0, "part 0 (default) or part 1")
var dbgFlag = flag.Bool("d", false, "debug flag")

func findDuplicates(list1 []rune, list2 []rune) bytes.Buffer {
	buf := bytes.Buffer{}

	if *dbgFlag {
		fmt.Println("list1: ", string(list1), ", list2=", string(list2))
	}

	for i := 0; i < len(list1); i++ {
		for j := 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				if *dbgFlag {
					fmt.Println("Found duplicate: ", string(list2[j]), "i=", i, "j=", j)
				}
				if !strings.Contains(buf.String(), string(list2[j])) {
					buf.WriteRune(list2[j])
					if *dbgFlag {
						fmt.Println("New duplicate: ", string(list2[j]))
					}
				}
			}
		}
	}
	return buf

}

func calcScore1(str string) int {
	score := 0

	chars := []rune(str)
	for i := 0; i < len(chars); i++ {
		c := int(chars[i])
		if c < 96 {
			// Captial letters, subtract 38
			score += c - 38
		} else if c > 96 {
			// Small caps, subtract 96
			score += c - 96
		}
	}
	return score
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

	btot := bytes.Buffer{}
	for i, line := range lines {
		if *dbgFlag {
			fmt.Println("Rucksack #", i, ": ", line)
		}

		chars := []rune(line)
		mid := len(chars) / 2

		buf := findDuplicates([]rune(line[0:mid]), []rune(line[mid:]))
		btot.Write(buf.Bytes())
	}

	fmt.Println("*** Part 1 ***")
	fmt.Println("Found duplicates:", btot.String(), ", total score ", calcScore1(btot.String()))

	// Part two
	gtot := bytes.Buffer{}
	for i := 0; i < len(lines); i++ {
		// Compare first two rucksacks in the group
		d1 := findDuplicates([]rune(lines[i]), []rune(lines[i+1]))
		if *dbgFlag {
			fmt.Println("gtot =", gtot.String())
		}
		// Compare result from first two rucksacks with the third
		d2 := findDuplicates([]rune(d1.String()), []rune(lines[i+2]))
		gtot.Write(d2.Bytes())
		if *dbgFlag {
			fmt.Println("gtot =", gtot.String())
		}
		i += 2
	}
	fmt.Println("*** Part 2 ***")
	fmt.Println("Found group badges:", gtot.String(), ", total score ", calcScore1(gtot.String()))
}
