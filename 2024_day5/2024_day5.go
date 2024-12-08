package main

import (
	"AoC2022/aoc_helpers"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var dbgFlag = flag.Bool("d", false, "debug flag")
var traceFlag = flag.Bool("t", false, "trace flag")

type RuleT struct {
	rules map[int]map[int]bool // rules is a map of maps
}

func str2ints(strs []string) []int {
	res := make([]int, len(strs))

	for i, s := range strs {
		res[i], _ = strconv.Atoi(s)
	}
	return res
}

func addRule(r RuleT, p1, p2 int) RuleT {

	if r.rules == nil {
		r.rules = make(map[int]map[int]bool)
	}
	_, ok := r.rules[p1]
	if !ok {
		// Create new slice & add page
		r.rules[p1] = make(map[int]bool)
	}
	r.rules[p1][p2] = true

	return r
}

func checkUpdate(r RuleT, pages []int) bool {

	for i, p := range pages {
		rules := r.rules[p]

		if rules != nil {
			for j := 0; j < i; j++ {
				if rules[pages[j]] {
					return false
				}
			}
		}

	}
	return true
}

func orderPages(r RuleT, pages []int) []int {
	res := pages

	for i, p := range res {
		rules := r.rules[p]

		if rules != nil {
			for j := 0; j < i; j++ {
				if rules[res[j]] {
					tmp := res[i]
					res[i] = res[j]
					res[j] = tmp
					return orderPages(r, res)
				}
			}
		}

	}

	return res
}

func midPageNo(l []int) int {

	return l[len(l)/2]
}

func solvePart1(args []string) int {
	fn := args[0]
	res := 0
	regex1 := regexp.MustCompile(`(\d+)\|(\d+)`)
	regex2 := regexp.MustCompile(`,`)
	ruleSet := RuleT{}

	// Parse input file
	lines, err := aoc_helpers.ReadLinesToByteSlice(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	//Load page ordering rules
	updateStart := 0
	for i, l := range lines {
		if len(l) == 0 {
			updateStart = i + 1
			break
		}
		m := regex1.FindSubmatch(l)
		//		fmt.Printf("%d: %s\t%s\n", i+1, string(m[1]), string(m[2]))
		p1, _ := strconv.Atoi(string(m[1]))
		p2, _ := strconv.Atoi(string(m[2]))

		ruleSet = addRule(ruleSet, p1, p2)
	}
	// Check updates
	for i := updateStart; i < len(lines); i++ {
		m := str2ints(regex2.Split(string(lines[i]), -1))
		if checkUpdate(ruleSet, m) {
			res = res + midPageNo(m)
		}
	}

	return res
}

func solvePart2(args []string) int {
	fn := args[0]
	res := 0
	regex1 := regexp.MustCompile(`(\d+)\|(\d+)`)
	regex2 := regexp.MustCompile(`,`)
	ruleSet := RuleT{}

	// Parse input file
	lines, err := aoc_helpers.ReadLinesToByteSlice(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	//Load page ordering rules
	updateStart := 0
	for i, l := range lines {
		if len(l) == 0 {
			updateStart = i + 1
			break
		}
		m := regex1.FindSubmatch(l)
		//		fmt.Printf("%d: %s\t%s\n", i+1, string(m[1]), string(m[2]))
		p1, _ := strconv.Atoi(string(m[1]))
		p2, _ := strconv.Atoi(string(m[2]))

		ruleSet = addRule(ruleSet, p1, p2)
	}
	// Check updates, correct page ordering
	for i := updateStart; i < len(lines); i++ {
		m := str2ints(regex2.Split(string(lines[i]), -1))
		if !checkUpdate(ruleSet, m) {
			m = orderPages(ruleSet, m)
			res = res + midPageNo(m)
		}
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
