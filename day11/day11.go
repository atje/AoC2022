/*
--- Day 11: Monkey in the Middle ---
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type monkeyT struct {
	id        int
	operation func(int) int
	test      func(int) int
	items     []int
}

var monkeys []monkeyT

// Load initial monkeys
func loadMonkeys(file string) []monkeyT {

	m := make([]monkeyT, 0)

	fp, err := os.Open(file)
	if err != nil {
		return nil
	}
	defer fp.Close()

	newMonkeyRE := regexp.MustCompile(`Monkey (\d):`)
	itemsRE := regexp.MustCompile(`\s+Starting items:\s+([\s,\d]+)`)
	itemRE := regexp.MustCompile(`(\d+)`)

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		bytes := []byte(scanner.Text())

		switch {
		case newMonkeyRE.FindSubmatch(bytes) != nil:
			matches := newMonkeyRE.FindSubmatch(bytes)
			//log.Traceln("found Monkey", string(matches[1]))

			id, _ := strconv.Atoi(string(matches[1]))
			m = append(m, monkeyT{id: id})

		case itemsRE.Find(bytes) != nil:
			matches := itemsRE.Find(bytes)

			itemsFound := itemRE.FindAll(matches, -1)
			//log.Tracef("found starting items, %v", itemsFound)
			m[len(m)-1].items = make([]int, 0)
			for i := 0; i < len(itemsFound); i++ {
				item, _ := strconv.Atoi(string(itemsFound[i]))
				m[len(m)-1].items = append(m[len(m)-1].items, item)
			}
		}
	}
	log.Tracef("Monkeys:\n%v", m)
	return m
}

func solveDay11Part1(file string, rounds int) int {
	monkeys = loadMonkeys(file)

	return -1
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.TraceLevel)
}

func main() {

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		log.Fatalln("Please provide input file!")
	}

	fmt.Println("Day 11, part 1 answer:", solveDay11Part1(args[0], 20))
}
