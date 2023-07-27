/*
--- Day 3: Rucksack Reorganization ---
One Elf has the important job of loading all of the rucksacks with supplies for the jungle journey. Unfortunately, that Elf didn't quite follow the packing instructions, and so a few items now need to be rearranged.

Each rucksack has two large compartments. All items of a given type are meant to go into exactly one of the two compartments. The Elf that did the packing failed to follow this rule for exactly one item type per rucksack.

The Elves have made a list of all of the items currently in each rucksack (your puzzle input), but they need your help finding the errors. Every item type is identified by a single lowercase or uppercase letter (that is, a and A refer to different types of items).

The list of items for each rucksack is given as characters all on a single line. A given rucksack always has the same number of items in each of its two compartments, so the first half of the characters represent items in the first compartment, while the second half of the characters represent items in the second compartment.

For example, suppose you have the following list of contents from six rucksacks:

vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
The first rucksack contains the items vJrwpWtwJgWrhcsFMMfFFhFp, which means its first compartment contains the items vJrwpWtwJgWr, while the second compartment contains the items hcsFMMfFFhFp. The only item type that appears in both compartments is lowercase p.
The second rucksack's compartments contain jqHRNqRjqzjGDLGL and rsFMfFZSrLrFZsSL. The only item type that appears in both compartments is uppercase L.
The third rucksack's compartments contain PmmdzqPrV and vPwwTWBwg; the only common item type is uppercase P.
The fourth rucksack's compartments only share item type v.
The fifth rucksack's compartments only share item type t.
The sixth rucksack's compartments only share item type s.
To help prioritize item rearrangement, every item type can be converted to a priority:

Lowercase item types a through z have priorities 1 through 26.
Uppercase item types A through Z have priorities 27 through 52.
In the above example, the priority of the item type that appears in both compartments of each rucksack is 16 (p), 38 (L), 42 (P), 22 (v), 20 (t), and 19 (s); the sum of these is 157.

Find the item type that appears in both compartments of each rucksack. What is the sum of the priorities of those item types?

------------
Noted:
- 2 compartments in each Elf rucksack
- case determines item type
- Same number of items in each compartment, first half of chars for first compartment, second half for second
- Find items that appear in both compartments
- Add common items according to prio number [a-z] 1-26, [A-Z] 27-52

------------

--- Part Two ---
As you finish identifying the misplaced items, the Elves come to you with another issue.

For safety, the Elves are divided into groups of three. Every Elf carries a badge that identifies their group. For efficiency, within each group of three Elves, the badge is the only item type carried by all three Elves. That is, if a group's badge is item type B, then all three Elves will have item type B somewhere in their rucksack, and at most two of the Elves will be carrying any other item type.

The problem is that someone forgot to put this year's updated authenticity sticker on the badges. All of the badges need to be pulled out of the rucksacks so the new authenticity stickers can be attached.

Additionally, nobody wrote down which item type corresponds to each group's badges. The only way to tell which item type is the right one is by finding the one item type that is common between all three Elves in each group.

Every set of three lines in your list corresponds to a single group, but each group can have a different badge item type. So, in the above example, the first group's rucksacks are the first three lines:

vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
And the second group's rucksacks are the next three lines:

wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
In the first group, the only item type that appears in all three rucksacks is lowercase r; this must be their badges. In the second group, their badge item type must be Z.

Priorities for these items must still be found to organize the sticker attachment efforts: here, they are 18 (r) for the first group and 52 (Z) for the second group. The sum of these is 70.

Find the item type that corresponds to the badges of each three-Elf group. What is the sum of the priorities of those item types?
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
