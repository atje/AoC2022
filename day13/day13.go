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
-

*/

package main

import (
	"AoC2022/aoc_helpers"
	"flag"
	"fmt"
	"os"

	"github.com/RyanCarrier/dijkstra"
	log "github.com/sirupsen/logrus"
)

var dbgFlag = flag.Bool("d", false, "debug flag")

// Generate a unique node ID, based on x/y coordinates
func generateID(x, mult, y int) int {
	return x*mult + y
}

func genHeightgrid(file string) (grid [][]int, startPos int, endPos int) {
	grid = make([][]int, 0)
	endPos = -1
	startPos = -1

	// Read file into [][]grid
	lines, err := aoc_helpers.ReadLines(file)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for x, line := range lines {
		grid = append(grid, make([]int, 0))
		for y, height := range line {
			name := x*len(line) + y
			if height == 'E' {
				endPos = name
				height = 'z'
			}
			if height == 'S' {
				startPos = name
				height = 'a'
			}
			grid[x] = append(grid[x], int(height-rune('a')))
		}
	}

	log.Debugf("S at %d, E at %d", startPos, endPos)
	log.Debugf("Grid: %v", grid)

	return grid, startPos, endPos

}

func genGraph(grid [][]int) *dijkstra.Graph {
	// Create Graph
	g := dijkstra.NewGraph()

	// Add vertices
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			vertexID := generateID(x, len(grid[x]), y)
			g.AddVertex(vertexID) //
			log.Debugf("Adding vertex x%d y%d with ID %d", x, y, vertexID)
		}
	}

	// Add arcs
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			vertexID := generateID(x, len(grid[x]), y)

			// Check neighbouring squares for possible paths //

			// Down
			if x < (len(grid) - 1) {
				if grid[x+1][y] <= grid[x][y]+1 {
					// Down move possible, add path
					destID := generateID(x+1, len(grid[x]), y)
					log.Debugf("Adding arc (ID %d) x%d y%d --> x%d y%d (ID %d)", vertexID, x, y, x+1, y, destID)
					g.AddArc(vertexID, destID, 1)
				}
			}
			// Up
			if x > 0 {
				if grid[x-1][y] <= grid[x][y]+1 {
					// Up move possible, add path
					log.Debugf("ID %d: Adding arc x%d y%d --> x%d y%d", vertexID, x, y, x-1, y)
					g.AddArc(vertexID, generateID(x-1, len(grid[x]), y), 1)
				}
			}
			// Left
			if y > 0 {
				if grid[x][y-1] <= grid[x][y]+1 {
					// Left move possible, add path
					log.Debugf("ID %d: Adding arc x%d y%d --> x%d y%d", vertexID, x, y, x, y-1)
					g.AddArc(vertexID, generateID(x, len(grid[x]), y-1), 1)
				}
			}

			// Right
			if y < (len(grid[x]) - 1) {
				if grid[x][y+1] <= grid[x][y]+1 {
					// Right move possible, add path
					log.Debugf("ID %d: Adding arc x%d y%d --> x%d y%d", vertexID, x, y, x, y+1)
					g.AddArc(vertexID, generateID(x, len(grid[x]), y+1), 1)
				}
			}
		}
	}
	log.Tracef("graph: %+v", g)

	return g
}

func findAll(grid [][]int, i int) []int {
	res := make([]int, 0)

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			log.Debugf("findAll: grid[%d][%d] = %d", x, y, grid[x][y])
			if grid[x][y] == i {
				vertexID := generateID(x, len(grid[x]), y)
				log.Debugf("findAll: Appending %d", vertexID)
				res = append(res, vertexID)
			}
		}
	}

	return res
}

func solveDay12Part1(file string) int {

	// Generate heatmap as a grid
	heightGrid, startPos, endPos := genHeightgrid(file)

	// Create Graph
	g := genGraph(heightGrid)

	// Find shortest path from S to E
	path, err := g.Shortest(startPos, endPos)
	if err != nil {
		log.Fatal(err)
	}

	return int(path.Distance)
}

func solveDay12Part2(file string) int {
	// Generate heatmap as a grid
	heightGrid, _, endPos := genHeightgrid(file)

	// Create Graph
	g := genGraph(heightGrid)

	// Find all points with height 'a', which is 0 in the heightmap
	startPositions := findAll(heightGrid, 0)

	// Calculate distance from all starting positions
	dist := int64(1 << 62)
	for _, startPos := range startPositions {
		path, err := g.Shortest(startPos, endPos)

		if err == nil && path.Distance < dist {
			log.Debugf("New distance %d is smaller than old %d, startPos %d", path.Distance, dist, startPos)
			dist = path.Distance
		} else {
			log.Debugf("No path found from %d to %d, err = %v, dist %d", startPos, endPos, err, dist)
		}
	}
	// Find & return minimum distance
	return int(dist)
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

	fmt.Println("Day 12, part 1:", solveDay12Part1(args[0]))
	fmt.Println("Day 12, part 2:", solveDay12Part2(args[0]))
}
