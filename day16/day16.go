/*
--- Day 16: Proboscidea Volcanium ---

Some definitions from the text:
- There's even a valve in the room you and the elephants are currently standing in labeled AA.
You estimate it will take you one minute to open a single valve and one minute to follow any tunnel from one valve to another.
What is the most pressure you could release?


Objective:
Work out the steps to release the most pressure in 30 minutes.
What is the most pressure you can release?

Rules:
- 1min moving between valves
- 1min opening a valve

Approach:
- build a directed graph using the provided valves and tunnels
- use Dijkstra's algorithm to find the shortest path between valves wherr the flow rate is > 0
- use a priority queue to find the best path to follow
- start at the AA valve, open it and follow the path with the highest flow rate
- keep track of the time left and the pressure released
- stop when time is 0 or all valves are open
- return the total pressure released

-- Part two --

Approach:
-

*/

package main

import (
	"AoC2022/aoc_helpers"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/RyanCarrier/dijkstra"

	log "github.com/sirupsen/logrus"
)

var dbgFlag = flag.Bool("d", false, "debug flag")
var traceFlag = flag.Bool("t", false, "trace flag")

// The graph to hold the valves and their connections
// The valveFlowMap to hold the flow rates of the valves
var graph *dijkstra.Graph
var valveFlowMap map[string]int

// Add this at the top of the file
var memo map[string]int

// Helper to create a unique key for memoization
func stateKey(curValve string, closedValves []string, minutes int) string {
	key := curValve + fmt.Sprintf("-%d-", minutes)
	for _, v := range closedValves {
		key += v + ","
	}
	return key
}

// Parse a line into graph and valveMap
func parseLine(line string) {
	log.Tracef("parseLine '%s'", line)

	// Find valve, flow rate and list of tunnels to other valves
	exprRE := regexp.MustCompile("(?i)Valve ([A-Za-z]+) has flow rate=([0-9]+); tunnel[s]* lead[s]* to valve[s]* (.*)")
	m := exprRE.FindStringSubmatch(line)
	if m == nil {
		log.Fatalf("Line '%s' not matching format!", line)
	}

	if *dbgFlag {
		fmt.Printf("%q\n", m)
	}

	re := regexp.MustCompile("(?i)[A-Za-z]+")
	tunnels := re.FindAllString(string(m[3]), -1)

	flowRate, _ := strconv.Atoi(m[2])

	if graph == nil {
		graph = dijkstra.NewGraph()
		valveFlowMap = make(map[string]int, 0)
	}
	// Add valve to graph
	graph.AddMappedVertex(m[1])

	for _, tunnel := range tunnels {
		if *dbgFlag {
			fmt.Printf("Adding tunnel from %s to %s\n", m[1], tunnel)
		}
		// Add tunnel to graph
		graph.AddMappedVertex(tunnel)
		graph.AddMappedArc(m[1], tunnel, 1) // 1 minute to move between valves
	}
	// Add flow rate to valveFlowMap
	valveFlowMap[m[1]] = flowRate
}

func parseFile(fn string) {

	// Read file
	lines, err := aoc_helpers.ReadLines(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for _, line := range lines {
		parseLine(line)
	}
}

// Remove a string from a slice of strings
// Returns a new slice with the string removed
// If the string is not found, the original slice is returned
func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func calcReleasedPressure(indent, curValve string, closedValves []string, minutes int) int {
	nextValve := ""

	if minutes <= 0 {
		return 0
	}

	// Check if we have already calculated the pressure for this state
	key := stateKey(curValve, closedValves, minutes)
	if val, ok := memo[key]; ok {
		return val
	}

	maxPressure := 0
	// For each valve in the list of closed valves, find the shortest path to it from the current valve
	for _, valve := range closedValves {
		min := minutes

		if valve == curValve {
			continue // Skip the current valve
		}

		curID, _ := graph.GetMapping(curValve)
		valveID, _ := graph.GetMapping(valve)
		path, err := graph.Shortest(curID, valveID)

		if err != nil {
			log.Fatalf("Error finding shortest path from %s to %s: %s", curValve, valve, err)
		}

		min -= (int)(path.Distance) + 1
		pressure := 0
		if min > 0 {
			// Calculate the pressure released by opening the valve
			pressure = valveFlowMap[valve] * min
		} else {
			if *dbgFlag {
				fmt.Printf("Not enough time to open valve %s, time left: %d\n", valve, min)
			}
		}

		cv := make([]string, len(closedValves))
		_ = copy(cv, closedValves)

		added := calcReleasedPressure(indent+"  ", valve, remove(cv, valve), min)
		if added+pressure > maxPressure {
			maxPressure = added + pressure
			nextValve = valve
		}
	}
	if *dbgFlag {
		fmt.Printf("%stime %d min - Max pressure from %s is %d, next valve: %s\n", indent, minutes, curValve, maxPressure, nextValve)
	}

	memo[key] = maxPressure
	return maxPressure
}

func solvePart1(args []string) int {

	fn := args[0]
	//minutes, _ := strconv.Atoi(args[1])
	minutes := 30
	cur := "AA"

	// Parse input file
	parseFile(fn)

	// Find all valves with flow rate > 0
	valves := make([]string, 0)
	for valve, flow := range valveFlowMap {
		if flow > 0 {
			valves = append(valves, valve)
		}
	}

	memo = make(map[string]int)
	return calcReleasedPressure("", cur, valves, minutes)
}

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

// Expected input args on commandline <filename> <part1_row> <part2_max>
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
	// fmt.Println("part 2:", solvePart2(args))
}
