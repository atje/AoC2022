/*
--- Day 14: Regolith Reservoir ---

Some definitions from the text:
- Coordinates (x,y) shape the path
- x represents distance to the right and y represents distance down
- After the first point of each path, each point indicates the end of a
  straight horizontal or vertical line to be drawn from the previous point
- Drawing rock as #, air as ., and the source of the sand as +

Objective:
How many units of sand come to rest before sand starts flowing into the abyss below?

Rules:
- Sand is produced one unit at a time, and the next unit of sand is
  not produced until the previous unit of sand comes to rest

- A unit of sand is large enough to fill one tile of air in your scan

- A unit of sand always falls down one step if possible

- If the tile immediately below is blocked (by rock or sand),
  the unit of sand attempts to instead move diagonally one step down and to the left
- If that tile is blocked, the unit of sand attempts to instead move
  diagonally one step down and to the right
- If all three possible destinations are blocked, the unit of sand comes to rest


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

type CaveMap struct {
	x0, y0 int //Origo, corresponding to point 0, 0 in the point map
	point  [][]rune
}

const rockChar rune = '#'
const sandChar rune = 'o'

var dbgFlag = flag.Bool("d", false, "debug flag")
var traceFlag = flag.Bool("t", false, "trace flag")

var caveMap CaveMap

// Expand CaveMap to fit the new coordinates if needed
// Move origo if necessary
func expandCM(cm CaveMap, x, y int) CaveMap {
	xCap := 0
	yCap := len(cm.point)
	if yCap > 0 {
		xCap = len(cm.point[0])
	}

	dX := cm.x0 + xCap - x - 1
	dY := cm.y0 + yCap - y - 1

	log.Tracef("expandCM(%o, %d, %d), dX = %d, dY = %d", cm.point, x, y, dX, dY)
	// Within capacity, no change needed

	// Expand down
	if y > cm.y0 && dY < 0 {
		for i := 0; i > dY; i-- {
			newRow := make([]rune, xCap)
			cm.point = append(cm.point, newRow)
		}
	}
	// Expand up, move y0
	// NOT IMPLEMENTED

	if x < cm.x0 {
		// Expand left, move x0
		dX = cm.x0 - x
		for i := range cm.point {
			newCol := make([]rune, dX)
			cm.point[i] = append(newCol, cm.point[i]...)
		}
		cm.x0 = x
	} else if x > cm.x0 && dX < 0 {
		// Expand right
		for i := range cm.point {
			newCol := make([]rune, -dX)
			cm.point[i] = append(cm.point[i], newCol...)
		}
	}
	log.Tracef("expandCM result = %o", cm.point)
	return cm
}

// Add char to point in map, expand map if needed
func add2Map(x int, y int, c rune) {
	log.Tracef("add2Map(%d, %d, %c, %o)", x, y, c, caveMap)

	// Add rune at point (x,y)
	if log.GetLevel() == log.DebugLevel {
		printCaveMap(caveMap)
	}
	log.Tracef("char '%c' to be added in position [%d][%d]\nCaveMap.point = %o", c, y, x, caveMap.point)
	caveMap.point[y-caveMap.y0][x-caveMap.x0] = c
}

func printCaveMap(cm CaveMap) {
	fmt.Printf("---CaveMap:---\n(x0, y0) = (%d, %d)\n", cm.x0, cm.y0)
	for n, row := range cm.point {
		fmt.Printf("%d ", n)
		for _, v := range row {
			pointVal := '.'
			if v == rockChar || v == sandChar {
				pointVal = v
			}
			fmt.Printf("%c", pointVal)
		}
		fmt.Printf("\n")
	}
	fmt.Println("--------------")
}

// Drop of one grain of sand, return true if the sand falls into void,
// otherwise false
// Sand grain is added to map
func dropSand(x int, y int) bool {
	sandX := x - caveMap.x0
	sandY := y - caveMap.y0

	prevX := sandX
	prevY := sandY
	for sandY < len(caveMap.point)-1 {
		sandY++
		if caveMap.point[sandY][sandX] == 0 {
			// Check if current position is free
		} else if (sandX > 0) && (caveMap.point[sandY][sandX-1] == 0) {
			// Otherwise check if one left is free
			sandX--
		} else if (sandX+1 < len(caveMap.point[sandY])) && (caveMap.point[sandY][sandX+1] == 0) {
			// Otherwise check if one right is free
			sandX++
		} else if (sandX == 0) && (caveMap.point[sandY][sandX] != 0) {
			// Falling into the void
			break
		} else {
			// Otherwise it has come to rest
			add2Map(prevX+caveMap.x0, prevY+caveMap.y0, sandChar)
			//			printCaveMap(caveMap)
			return false
		}
		prevY = sandY
		prevX = sandX
	}

	return true
}

// Parse a line of coordinates describing rock structue
// line consists of <value> [--> <value>]+
// <value> consists of coordinates on the format x, y
// x is horisontal, y vertical
func parseLine(line string) {
	log.Tracef("parseLine \"%s\"", line)

	exprRE := regexp.MustCompile(`\s+->\s+`)
	coordinateStrings := exprRE.Split(line, -1)

	//log.Debugf("Matches %q", coordinateStrings)

	exprRE = regexp.MustCompile(`,`)

	var xPrev, yPrev int

	for i, cs := range coordinateStrings {
		csSplit := exprRE.Split(cs, -1)
		x, _ := strconv.Atoi(csSplit[0])
		y, _ := strconv.Atoi(csSplit[1])

		caveMap = expandCM(caveMap, x, y)
		if i != 0 {
			log.Debugf("i=%d\t(x,y) = (%d,%d)", i, x, y)
			if yPrev == y {
				// Add horisontal line
				if xPrev < x {
					// left to right
					for j := xPrev; j <= x; j++ {
						add2Map(j, y, rockChar)
					}
				} else {
					// right to left
					for j := x; j <= xPrev; j++ {
						add2Map(j, y, rockChar)
					}
				}
			} else {
				// Add vertical line
				if yPrev < y {
					// left to right
					for j := yPrev; j <= y; j++ {
						add2Map(x, j, rockChar)
					}
				} else {
					// right to left
					for j := y; j <= yPrev; j++ {
						add2Map(x, j, rockChar)
					}
				}
			}
		}
		xPrev = x
		yPrev = y
	}
}

func parseIntoCaveMap(fn string) {
	// Initialize cave map
	caveMap = CaveMap{
		x0: 500,
		y0: 0,
		point: [][]rune{
			{0},
		},
	}

	// Read file
	lines, err := aoc_helpers.ReadLines(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Go through all lines, add packets to a new slice
	for _, line := range lines {
		parseLine(line)
	}

	if log.GetLevel() == log.DebugLevel {
		printCaveMap(caveMap)
	}
}

func solvePart1(args []string) int {

	fn := args[0]
	// Parse input file into cave map
	parseIntoCaveMap(fn)

	// Then drop one grain of sand from point 500,0 into the cave until the cave is full
	// Calculate the number of sand grains in the map
	sandUnits := 0

	for !dropSand(500, 0) {
		sandUnits++
	}

	return sandUnits
}

func solvePart2(args []string) int {

	fn := args[0]
	// Parse input file into cave map
	parseIntoCaveMap(fn)

	// Then add floor which is y coord + 2 and going from 500,0 in both directions with
	// the highest y coord + 2 / 2
	newY := len(caveMap.point) - 1 + 2

	caveMap = expandCM(caveMap, 500+newY, newY)
	caveMap = expandCM(caveMap, 500-newY, newY)
	for i := 0; i < len(caveMap.point[0]); i++ {
		add2Map(i+caveMap.x0, newY, rockChar)
	}

	// Then drop one grain of sand from point 500,0 into the cave until the cave is full
	// Calculate the number of sand grains in the map
	sandUnits := 0

	for !dropSand(500, -1) {
		sandUnits++
	}

	return sandUnits
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
