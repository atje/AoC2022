package main

import (
	"AoC2022/aoc_helpers"
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

const DOT byte = '.'
const ANTINODE byte = '#'

var dbgFlag = flag.Bool("d", false, "debug flag")
var traceFlag = flag.Bool("t", false, "trace flag")

type coordT struct {
	x, y int
}

// abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func initializeMap(row, col int) [][]byte {
	m := make([][]byte, row)
	for i := range m {
		m[i] = make([]byte, col)
	}
	return m
}

func addAntenna(m [][]byte, ant coordT) {
	if ant.x < 0 || ant.x >= len(m[0]) || ant.y < 0 || ant.y >= len(m) {
		if *dbgFlag {
			log.Debugf("[DEBUG] Antenna out of bounds: (%d,%d)", ant.x, ant.y)
		}
		return
	}

	if *dbgFlag {
		log.Debugf("[DEBUG] Adding antenna at (%d,%d)", ant.x, ant.y)
	}
	// Add the antinode to the map
	m[ant.y][ant.x] = ANTINODE
}

func addAntiNodes(m [][]byte, antennas []coordT) [][]byte {
	for i := range antennas {
		for j := range antennas {
			if j <= i {
				// Skip the same antenna
				continue
			}

			// Calculate the distance between the antennas
			dx := antennas[j].x - antennas[i].x
			dy := antennas[j].y - antennas[i].y

			// Add the antinode to the map
			addAntenna(m, coordT{x: antennas[i].x + 2*dx, y: antennas[i].y + 2*dy})
			addAntenna(m, coordT{x: antennas[i].x - dx, y: antennas[i].y - dy})
		}
	}

	return m
}

func addAntiNodes2(m [][]byte, antennas []coordT) [][]byte {
	for i := range antennas {
		for j := range antennas {
			if j <= i {
				// Skip the same antenna
				continue
			}

			// Calculate the distance between the antennas
			dx := antennas[j].x - antennas[i].x
			dy := antennas[j].y - antennas[i].y

			k := 0
			for {
				// Add the antinode to the map
				addAntenna(m, coordT{x: antennas[i].x + k*dx, y: antennas[i].y + k*dy})
				addAntenna(m, coordT{x: antennas[i].x - k*dx, y: antennas[i].y - k*dy})
				k++
				if abs(k*dx) > len(m[0]) || abs(k*dy) > len(m) {
					break
				}
			}
		}
	}

	return m
}

func printMap(m [][]byte) {
	for i := range m {
		for j := range m[i] {
			if m[i][j] == ANTINODE {
				fmt.Print(string(ANTINODE))
			} else {
				fmt.Print(string(DOT))
			}
		}
		fmt.Println()
	}
}

func solvePart1(args []string) int {
	fn := args[0]
	res := 0

	// Parse input file
	lines, err := aoc_helpers.ReadLinesToByteSlice(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Find antennas, add them to the hashmap containing the antennas
	pairs := map[byte][]coordT{}

	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] != DOT {
				// Add antenna to the map
				pairs[lines[i][j]] = append(pairs[lines[i][j]], coordT{x: j, y: i})
				if *dbgFlag {
					log.Debugf("[DEBUG] Found antenna %c at (%d,%d)", lines[i][j], j, i)
				}
			}
		}

	}

	if *dbgFlag {
		log.Debugf("[DEBUG] Antenna pairs: %v", pairs)
	}

	// Go through all found antennas and add antinodes to the map
	m := initializeMap(len(lines), len(lines[0]))

	for _, v := range pairs {
		m = addAntiNodes(m, v)
	}

	// Print the map
	if *dbgFlag {
		log.Debugf("[DEBUG] Antenna map:")
		printMap(m)
	}

	// Count the number of unique antinode locations
	for i := range m {
		for j := range m[i] {
			if m[i][j] == ANTINODE {
				res++
			}
		}
	}

	// Print the number of unique antinode locations
	return res
}

func solvePart2(args []string) int {
	fn := args[0]
	res := 0

	// Parse input file
	lines, err := aoc_helpers.ReadLinesToByteSlice(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Find antennas, add them to the hashmap containing the antennas
	pairs := map[byte][]coordT{}

	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] != DOT {
				// Add antenna to the map
				pairs[lines[i][j]] = append(pairs[lines[i][j]], coordT{x: j, y: i})
				if *dbgFlag {
					log.Debugf("[DEBUG] Found antenna %c at (%d,%d)", lines[i][j], j, i)
				}
			}
		}

	}

	if *dbgFlag {
		log.Debugf("[DEBUG] Antenna pairs: %v", pairs)
	}

	// Go through all found antennas and add antinodes to the map
	m := initializeMap(len(lines), len(lines[0]))

	for _, v := range pairs {
		m = addAntiNodes2(m, v)
	}

	// Print the map
	if *dbgFlag {
		log.Debugf("[DEBUG] Antenna map:")
		printMap(m)
	}

	// Count the number of unique antinode locations
	for i := range m {
		for j := range m[i] {
			if m[i][j] == ANTINODE {
				res++
			}
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
