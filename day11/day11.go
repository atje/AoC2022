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
	"sort"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type fnT func(int) int

type monkeyT struct {
	id         int
	operation  fnT
	test       int
	throwTrue  int
	throwFalse int
	finspCnt   int
	tinspCnt   int
	items      []int
}

var monkeys []monkeyT
var dbgFlag = flag.Bool("d", false, "debug flag")

func parseOperation(bytes []byte, div int) fnT {

	exprRE := regexp.MustCompile(`new = old\s+([\*\+\-])\s+(\d+|old)`)

	matches := exprRE.FindAllSubmatch(bytes, -1)
	log.Traceln(matches)
	op, val := string(matches[0][1]), string(matches[0][2])
	v, old := strconv.Atoi(val)
	vali := int(v)
	log.Tracef("n = o %v %v", op, val)

	switch {
	case op == "*":
		if old != nil {
			return func(a int) int { log.Tracef("%d * %d / %d", a, a, div); return a * a / div }
		}
		return func(a int) int { log.Tracef("%d * %d / %d", a, vali, div); return a * vali / div }
	case op == "+":
		if old != nil {
			return func(a int) int { log.Tracef("(%d + %d) / %d", a, a, div); return (a + a) / div }
		}
		return func(a int) int { log.Tracef("(%d + %d) / %d", a, vali, div); return (a + vali) / div }
	default:
		log.Fatalln("could not parse operator", op)
	}

	return func(a int) int { return a }
}

// Load initial monkeys
func loadMonkeys(file string, div int) []monkeyT {

	m := make([]monkeyT, 0)

	fp, err := os.Open(file)
	if err != nil {
		return nil
	}
	defer fp.Close()

	newMonkeyRE := regexp.MustCompile(`Monkey\s+(\d):`)
	itemsRE := regexp.MustCompile(`\s+Starting items:\s+([\s,\d]+)`)
	itemRE := regexp.MustCompile(`(\d+)`)
	opRE := regexp.MustCompile(`Operation:\s+(.+)`)
	testRE := regexp.MustCompile(`Test:\s+divisible by\s+(\d+)`)
	trueRE := regexp.MustCompile(`\s+If true: throw to monkey\s+(\d+)`)
	falseRE := regexp.MustCompile(`\s+If false: throw to monkey\s+(\d+)`)

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
				m[len(m)-1].items = append(m[len(m)-1].items, int(item))
			}
		case opRE.FindSubmatch(bytes) != nil:
			matches := opRE.FindSubmatch(bytes)
			log.Tracef("Found operation: %v", string(matches[1]))

			m[len(m)-1].operation = parseOperation(matches[1], div)

		case testRE.FindSubmatch(bytes) != nil:
			matches := testRE.FindSubmatch(bytes)
			val, _ := strconv.Atoi(string(matches[1]))
			m[len(m)-1].test = int(val)

		case trueRE.FindSubmatch(bytes) != nil:
			matches := trueRE.FindSubmatch(bytes)
			m[len(m)-1].throwTrue, _ = strconv.Atoi(string(matches[1]))

		case falseRE.FindSubmatch(bytes) != nil:
			matches := falseRE.FindSubmatch(bytes)
			m[len(m)-1].throwFalse, _ = strconv.Atoi(string(matches[1]))
		}
	}
	log.Tracef("Monkeys:\n%v", m)
	return m
}

func calcLCD() int {
	denoms := make(map[int]int)

	for _, m := range monkeys {
		denoms[m.test] = m.test
	}

	for {
		all_equal := true
		least_ind := -1
		for k, v := range denoms {
			if (least_ind != -1) && (denoms[least_ind] != v) {
				all_equal = false
			}
			if least_ind == -1 {
				least_ind = k
			} else if v < denoms[least_ind] {
				least_ind = k
			}
		}
		if all_equal {
			return denoms[least_ind]
		}
		denoms[least_ind] += least_ind
	}
}

// Play one round
func playRound(lcd int) {
	for i, m := range monkeys {
		for _, item := range m.items {
			worryLvl := m.operation(item)
			throwMonkey := m.throwFalse
			if worryLvl%m.test == int(0) {
				log.Traceln("True", worryLvl)
				throwMonkey = m.throwTrue
				monkeys[i].tinspCnt++
			} else {
				monkeys[i].finspCnt++
			}
			log.Tracef("Throwing %d to monkey %d\n", worryLvl, throwMonkey)
			monkeys[throwMonkey].items = append(monkeys[throwMonkey].items, worryLvl%lcd)
		}
		monkeys[i].items = []int{}
	}
}

// Dump monkeys to stdout
func dumpMonkeys() {
	for _, n := range monkeys {
		log.Debugf("Monkey %d: ", n.id)
		log.Debugf("truecnt(%d), falsecnt(%d) %v", n.tinspCnt, n.finspCnt, n.items)
	}
}

// Multiply the two most active monkey inspection count
func calcMonkeyBusiness() int {
	res := -1
	activecnt := make([]int, 0)

	for _, m := range monkeys {
		activecnt = append(activecnt, m.finspCnt+m.tinspCnt)
	}

	sort.Slice(activecnt, func(i, j int) bool {
		return activecnt[i] > activecnt[j]
	})

	res = activecnt[0] * activecnt[1]
	return res
}

func solveDay11Part1(file string, rounds int) int {
	monkeys = loadMonkeys(file, 3)

	lcd := calcLCD()
	for i := 0; i < rounds; i++ {
		log.Debugln("Round", i)
		playRound(lcd)
		dumpMonkeys()
	}

	return calcMonkeyBusiness()
}

func solveDay11Part2(file string, rounds int) int {
	monkeys = loadMonkeys(file, 1)

	lcd := calcLCD()
	for i := 0; i < rounds; i++ {
		log.Debugln("Round", i)
		playRound(lcd)
		dumpMonkeys()
	}

	return calcMonkeyBusiness()
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

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

	fmt.Println("Day 11, part 1 answer:", solveDay11Part1(args[0], 20))
	fmt.Println("Day 11, part 2 answer:", solveDay11Part2(args[0], 10000))
}
