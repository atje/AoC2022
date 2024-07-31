/*
--- Day 15: Beacon Exclusion Zone ---

Some definitions from the text:
-

Objective:
Consult the report from the sensors you just deployed.
In the row where y=2000000, how many positions cannot contain a beacon?

Rules:
-

Approach:
-

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

	log "github.com/sirupsen/logrus"
)

type coordType struct {
	x, y int
}

type sensorType struct {
	pos    coordType
	beacon coordType
	tcDist int
}

var dbgFlag = flag.Bool("d", false, "debug flag")
var traceFlag = flag.Bool("t", false, "trace flag")
var nFlag = flag.Int("n", 26, "row to calculate beacon positions")

// abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// calcTCDist returns the calculated TrafficCab distance between two cartesian coordinates
func calcTCDist(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func calcRadiusOnRow(s sensorType, row int) int {
	if abs(s.pos.y-row) > s.tcDist {
		return 0
	}

	return s.tcDist - abs(s.pos.y-row)

}

// Parse a line into sensor struct
// Also calculates Taxicab distance
func parseLine(line string) sensorType {
	log.Tracef("parseLine '%s'", line)

	exprRE := regexp.MustCompile(`x=(?P<sensor_x>\-*\d+), y=(?P<sensor_y>-*\d+):[\w|\s]+x=(?P<beacon_x>\-*\d+), y=(?P<beacon_y>-*\d+)`)
	m := exprRE.FindStringSubmatch(line)
	if m == nil {
		log.Fatalf("Line '%s' not matching format!", line)
	}

	var res sensorType
	for i, name := range exprRE.SubexpNames() {
		switch name {
		case "sensor_x":
			res.pos.x, _ = strconv.Atoi(m[i])
		case "sensor_y":
			res.pos.y, _ = strconv.Atoi(m[i])
		case "beacon_x":
			res.beacon.x, _ = strconv.Atoi(m[i])
		case "beacon_y":
			res.beacon.y, _ = strconv.Atoi(m[i])
		}
	}

	res.tcDist = calcTCDist(res.pos.x, res.pos.y, res.beacon.x, res.beacon.y)

	return res
}

func parseFile(fn string) []sensorType {

	// Read file
	lines, err := aoc_helpers.ReadLines(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Go through all lines, add sensor readings to slice
	sensors := make([]sensorType, 0)

	for _, line := range lines {
		sensors = append(sensors, parseLine(line))
	}

	return sensors
}

func solvePart1(args []string) int {

	fn := args[0]
	row, _ := strconv.Atoi(args[1])

	// Parse input file
	sensors := parseFile(fn)

	set := make(map[int]bool)

	// Loop through sensors and check sensor coverage on row 10 = Y-coord
	for i, s := range sensors {
		res := calcRadiusOnRow(s, row)
		log.Tracef("Sensor %d: x=%d, y=%d, tcDist=%d, res=%d", i, s.pos.x, s.pos.y, s.tcDist, res)
		if res > 0 {
			for n := s.pos.x - res; n <= s.pos.x+res; n++ {
				set[n] = true
			}
		}
	}

	return len(set) - 1
}

/*
	func solvePart2(fn string) int {
		// Parse input file
		parseFile(fn)

		return -1
	}
*/
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
	//fmt.Println("part 2:", solvePart2(args[0]))
}
