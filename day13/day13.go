/*
--- Day 13: Distress Signal ---

Some definitions from the text:
- Your list consists of pairs of packets
- Pairs are separated by a blank line
- Packet data consists of lists and integers
- Each list starts with [, ends with ]
- Lists contains zero or more comma-separated values (either integers or other lists)
- Each packet is always a list and appears on its own line
- When comparing two values, the first value is called left and the second value is called right

Objective:
You need to identify how many pairs of packets are in the right order

Rules:
- If both values are integers, the lower integer should come first

- If the left integer is lower than the right integer, the inputs are in the right order
- If the left integer is higher than the right integer, the inputs are not in the right order
- Otherwise, the inputs are the same integer; continue checking the next part of the input

- If both values are lists, compare the first value of each list, then the second value, and so on
- If the left list runs out of items first, the inputs are in the right order
- If the right list runs out of items first, the inputs are not in the right order
- If the lists are the same length and no comparison makes a decision about the order, continue checking the next part of the input

- If exactly one value is an integer, convert the integer to a list which contains that integer as its only value, then retry the comparison


Approach:
- Read input lines
- for each pair, run comparison adding output

*/

package main

import (
	"AoC2022/aoc_helpers"
	"flag"
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var dbgFlag = flag.Bool("d", false, "debug flag")

// Parse the packet into a list of packet data which consists of lists and ints
func parsePacket(p string) []string {
	l := 0

	var res []string

	// Check for empty packet
	if len(p) == 0 {
		log.Debugln("Return nil")
		return nil
	}

	// Check for int
	_, perr := strconv.Atoi(p)
	if perr == nil {
		res = append(res, p)
		return res
	}

	// Check for list
	for l < len(p) {
		ptr := l

		if p[l] == '[' {
			cnt := 1
			// find matching closing bracket
			for {
				ptr++
				if ptr > len(p) {
					log.Fatalln("incorrect packet format, could not find list closing")
				} else if p[ptr] == '[' {
					cnt++
				} else if p[ptr] == ']' {
					cnt--
					if cnt == 0 {
						break
					}
				}
			}
			res = append(res, p[l:ptr+1])
			l = ptr
		} else if p[l] >= 48 && p[l] <= 57 {
			// Found a number
			for {
				ptr++
				if ptr >= len(p) {
					res = append(res, p[l:ptr])
					l = ptr
					break
				}

				if p[ptr] == ',' {
					res = append(res, p[l:ptr])
					l = ptr
					break
				}
			}

		}
		l++
	}
	return res
}

// Returns 0 for wrong order, 1 right order, 2 for equal/undecisive
func isOrderedPair(left string, right string, prefixStr string) int {
	var res int = 2

	log.Debugf("%sCompare %s vs %s", prefixStr, left, right)

	prefixStr = "  " + prefixStr
	// check if both are empty
	if left == "" && right == "" {
		log.Debugf("%sBoth are empty", prefixStr)
		return 2
	}

	// check if left is empty while right is not
	if left == "" && right != "" {
		log.Debugf("%sLeft side ran out of items, so input is in the right order", prefixStr)
		return 1
	}

	// check if both are empty lists
	if left == "[]" && right == "[]" {
		log.Debugf("%sBoth are empty lists", prefixStr)
		return 2
	}

	// Try to convert to int
	lint, lerr := strconv.Atoi(left)
	rint, rerr := strconv.Atoi(right)

	// check if both are ints
	if lerr == nil && rerr == nil {
		v := lint - rint
		if v < 0 {
			log.Debugf("%sLeft side is smaller, so input is in the right order", prefixStr)
			return 1
		} else if v == 0 {
			return 2
		}
		log.Debugf("%sRight side is smaller, so input is NOT in the right order", prefixStr)
		return 0
	}

	// Check if one is int & the other is a list, convert to list and re-run check
	if lerr == nil && right[0] == '[' {
		left = "[" + left + "]"
	}
	if rerr == nil && left[0] == '[' {
		right = "[" + right + "]"
	}

	// Parse packets
	leftVals := parsePacket(left[1 : len(left)-1])
	rightVals := parsePacket(right[1 : len(right)-1])

	for n := 0; n < len(leftVals); n++ {
		// If the right list runs out of items first, the inputs are not in the right order
		if n >= len(rightVals) {
			log.Debugf("%sRight side ran out of items, so input is NOT in the right order", prefixStr)
			res = 0
			break
		}

		// Check the pair, exit loop if answer found
		res = isOrderedPair(leftVals[n], rightVals[n], prefixStr)
		if res != 2 {
			break
		}
	}

	if (len(leftVals) < len(rightVals)) && res == 2 {
		res = 1
	}

	return res
}

func solvePart1(file string) int {
	var result int = 0
	var pairCnt int = 1
	var left, right string

	// Read file
	lines, err := aoc_helpers.ReadLines(file)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Go through all lines, find pairs of packets
	for _, line := range lines {
		if line == "" {
			log.Debugln("")
			pairCnt++
			left = ""
			right = ""
		} else {
			if left == "" {
				left = line
			} else {
				// Foudn a pair, compare them
				right = line
				log.Debugf("== Pair %d ==", pairCnt)
				if isOrderedPair(left, right, "- ") == 1 {
					// Pair is ordered, add pair index to sum
					log.Debugln("inputs are in the right order")
					result += pairCnt
				} else {
					log.Debugln("inputs are NOT in the right order")

				}
			}
		}
	}
	return result
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
	}

	if len(args) == 0 {
		log.Fatalln("Please provide input file!")
	}

	fmt.Println("part 1:", solvePart1(args[0]))
}
