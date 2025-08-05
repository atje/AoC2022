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

type Bits uint64

func Set(b Bits, flag int) Bits    { return b | (1 << flag) }        // Set the bit
func Clear(b Bits, flag int) Bits  { return b &^ (1 << flag) }       // Clear the bit
func Toggle(b Bits, flag int) Bits { return b ^ (1 << flag) }        // Toggle the bit
func Has(b Bits, flag int) bool    { return (b & (1 << flag)) != 0 } // Check if the bit is set

// The graph to hold the valves and their connections
// The valveFlowMap to hold the flow rates of the valves
var graph *dijkstra.Graph
var valveFlowMap map[int]int
var numValves int

// memo to hold the calculated pressure for each state
var memo map[string]int
var distArr [][]int // Distance from 1st valve to 2nd valve

// Helper to create a unique key for memoization
func stateKey(curValve int, closedValves Bits, minutes int) string {
	key := fmt.Sprintf("%d-%d-%d", curValve, minutes, closedValves)

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
	id, _ := graph.GetMapping(m[1])
	valveFlowMap[id] = flowRate
}

func parseFile(fn string) {

	// Read file into lines
	lines, err := aoc_helpers.ReadLines(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Initialize the graph and valveFlowMap
	graph = dijkstra.NewGraph()
	valveFlowMap = make(map[int]int, 0)

	// Parse each line in the file
	for _, line := range lines {
		parseLine(line)
	}
}

// Calculate maximum released pressure using recursion to test all possible path
// within given time limits
// Works on n number of people & elephants
// Returns total released pressure over time
func calcReleasedPressure(indent string, curValve int, closedValves Bits, minutes int) int {
	nextValve := -1

	if *dbgFlag {
		fmt.Printf("%stime %d min - Current valves: %v, closed valves: %b\n", indent, minutes, curValve, closedValves)
	}
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
	for closedValve := 0; closedValve < numValves; closedValve++ {
		min := minutes

		if !Has(closedValves, closedValve) || closedValve == curValve {
			if *dbgFlag {
				fmt.Printf("Skipping valve %d, closedValves = %b\n", closedValve, closedValves)
			}
			continue // Skip the current valve
		}
		/*
			path, err := graph.Shortest(curValve, closedValve)

			if err != nil {
				log.Fatalf("Error finding shortest path from %d to %d: %s", curValve, closedValve, err)
			}
		*/
		min -= distArr[curValve][closedValve] + 1
		//min -= (int)(path.Distance) + 1
		pressure := 0
		if min > 0 {
			// Calculate the pressure released by opening the valve
			pressure = valveFlowMap[closedValve] * min
		} else {
			if *dbgFlag {
				fmt.Printf("Not enough time to open valve %d, time left: %d\n", closedValve, min)
			}
		}

		cv := Clear(closedValves, closedValve) // Remove the valve from the closed valves

		added := calcReleasedPressure(indent+"  ", closedValve, cv, min)
		if added+pressure > maxPressure {
			maxPressure = added + pressure
			nextValve = closedValve
		}
	}
	memo[key] = maxPressure
	if *dbgFlag {
		fmt.Printf("%stime %d min - Max pressure from %d is %d, next valve: %d\n", indent, minutes, curValve, maxPressure, nextValve)
	}

	return maxPressure
}

func solvePart1(args []string) int {

	fn := args[0]
	minutes := 30
	cur := "AA"

	// Parse input file
	parseFile(fn)

	// Find all valves with flow rate > 0, create bit array to represent them
	var valves Bits
	for valve, flow := range valveFlowMap {
		if flow > 0 {
			valves = Set(valves, valve)
		}
		if *dbgFlag {
			fmt.Printf("Valve %d has flow rate %d, bit: %64b\n", valve, flow, valves)
		}
	}
	numValves = len(valveFlowMap)

	memo = make(map[string]int)
	curID, _ := graph.GetMapping(cur)

	// Pre-calcluate shortest path from starting valve to any closed valve
	distArr = make([][]int, numValves)
	// Pre-calculate shortest path from a valve to any other
	// Set a ridiculusly high distance value if they are not reachable
	// Store distance in 2D array
	for i := 0; i < numValves; i++ {
		distArr[i] = make([]int, numValves)
		for j := 0; j < numValves; j++ {
			path, err := graph.Shortest(i, j)
			if err == nil {
				distArr[i][j] = (int)(path.Distance)
			} else {
				distArr[i][j] = 50000000
			}
		}
	}
	return calcReleasedPressure("", curID, valves, minutes)
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
