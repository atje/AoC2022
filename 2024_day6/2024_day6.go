package main

import (
	"AoC2022/aoc_helpers"
	"flag"
	"fmt"
	"os"
	"regexp"

	log "github.com/sirupsen/logrus"
)

// Guard movement direction
const OBSTACLE byte = '#'
const BREADCRUMB byte = 'X'
const NEW_OBSTACLE byte = 'O'
const BC_UD_STR byte = '|'
const BC_LR_STR byte = '-'
const BC_X_STR byte = '+'
const BC_GUARD_STR byte = '^'

// Use bitpositions to indicate earlier point passing movement direction of guard
// There are four possible ways the guard could pass through a point
// 1. Up to Down (UD)
// 2. Down to Up (DU)
// 3. Left to Right (LR)
// 4. Right to Left (RL)
// These are encoded in a 4 bit byte so we can use bitwise operations to check for earlier point passing, and
// to combine them with the current direction, and
// rotate by using barrel shifting
const EMPTY byte = 0
const BC_LR byte = 0b1
const BC_UD byte = 0b10
const BC_RL byte = 0b100
const BC_DU byte = 0b1000

var dbgFlag = flag.Bool("d", false, "debug flag")
var traceFlag = flag.Bool("t", false, "trace flag")

func findGuardPos(m [][]byte) (int, int) {
	regex1 := regexp.MustCompile(`\^`)

	guardX := -1
	guardY := -1

	for i, l := range m {
		loc := regex1.FindIndex(l)
		if loc != nil {
			guardX = i
			guardY = loc[0]
			if *traceFlag {
				fmt.Println("[TRACE] Found guard in position ", guardX, guardY)
			}
			break
		}
	}

	return guardX, guardY
}

func dirStr(dir byte) string {
	switch {
	case dir&(BC_DU) > 0:
		return "Up"
	case dir&(BC_UD) > 0:
		return "Dn"
	case dir&(BC_LR) > 0:
		return "Right"
	case dir&(BC_RL) > 0:
		return "Left"
	}
	return "Err"
}

func initializeMap(row, col int) [][]byte {
	m := make([][]byte, row)
	for i := range m {
		m[i] = make([]byte, col)
	}
	return m
}

func stepGuard(x, y int, heading byte) (int, int) {
	switch {
	case heading&(BC_DU) > 0:
		return x - 1, y
	case heading&(BC_UD) > 0:
		return x + 1, y
	case heading&(BC_LR) > 0:
		return x, y + 1
	case heading&(BC_RL) > 0:
		return x, y - 1
	}
	return x, y
}

func isBackToStart(nx, ny, x, y int, lheading, heading byte, rotcount int) bool {
	return nx == x && ny == y && lheading == heading && rotcount < 2
}

// Walk guard in direction heading from position x, y
// Return true if guard leaves the map, false otherwise
// Return true if falling out of area, false otherwise, and number of steps taken
func walkGuard(heading byte, x, y int, areamap [][]byte, loopcheck bool) (bool, [][]byte) {
	nx, ny := x, y
	lheading := heading

	if *dbgFlag {
		fmt.Printf("[DEBUG] walkGuard(%s, %d, %d)\n", dirStr(heading), x, y)
	}

	trailmap := initializeMap(len(areamap), len(areamap[0]))
	maxstep := len(areamap) * len(areamap[0])
	rotcount := 0

	for {
		px, py := nx, ny

		// Mark trail
		trailmap[nx][ny] = genBC(trailmap[nx][ny], heading)

		nx, ny = stepGuard(nx, ny, heading)

		// Check if guard is out of bounds
		// If so, return true
		if outOfMatrix(areamap, nx, ny) {
			return true, trailmap
		}

		// Check if guard is back to start position heading as started, and have not gotten there by rotating
		// If so, return false
		if loopcheck && isBackToStart(nx, ny, x, y, lheading, heading, rotcount) { // Check for loop
			return false, trailmap
		}

		if areamap[nx][ny] == OBSTACLE {
			if *dbgFlag {
				//	fmt.Printf("Obstacle at (%d, %d)\n", nx, ny)
			}
			heading = rotClockwise(heading)
			rotcount++
			nx, ny = px, py
		} else {
			rotcount = 0
		}

		// Check if guard is stuck in one place just rotating
		// If so, return false
		if rotcount > 3 {
			return false, trailmap
		}

		maxstep = maxstep - 1
		// Check if guard has moved too many steps to avoid infinite loop
		// If so, return false
		if maxstep < 0 {
			return false, trailmap
		}

		if *dbgFlag {
			fmt.Printf("[DEBUG] Moving to (%d, %d) %s\n", nx, ny, dirStr(heading))
		}

	}
}

// Combine current direction with earlier point passing direction
func genBC(cur byte, heading byte) byte {

	return cur | heading
}

// Check if guard is out of map
func outOfMatrix(m [][]byte, x, y int) bool {
	return x < 0 || y < 0 || x >= len(m) || y >= len(m[0])
}

// Print a combination of obstacle map & trail map
// Assumes both maps are of equal size
func printMap(obsmap, trailmap [][]byte) {
	for i := range obsmap {
		for j := range obsmap[i] {
			b := byte('.')
			if obsmap[i][j] == OBSTACLE {
				b = OBSTACLE
			} else if obsmap[i][j] == NEW_OBSTACLE {
				b = NEW_OBSTACLE
			} else if ((trailmap[i][j] & (BC_DU | BC_UD)) > 0) && (trailmap[i][j]&(BC_LR|BC_RL) > 0) {
				b = byte('+')
			} else if trailmap[i][j]&(BC_DU|BC_UD) > 0 {
				b = byte('|')
			} else if trailmap[i][j]&(BC_LR|BC_RL) > 0 {
				b = byte('-')
			} else if trailmap[i][j] == BC_GUARD_STR {
				b = BC_GUARD_STR
			} else {
				b = obsmap[i][j]
			}
			fmt.Printf("%s", string(b))
		}
		fmt.Println()
	}
}

// Rotate clockwise from current direction
func rotClockwise(dir byte) byte {
	// Shift left by 1 and wrap around the overflow bit to the right
	return ((dir << 1) | (dir >> 3)) & 0b1111
}

func solvePart1(args []string) int {
	fn := args[0]

	// Parse input file
	lines, err := aoc_helpers.ReadLinesToByteSlice(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Find starting position of guard
	guardX, guardY := findGuardPos(lines)
	guardPos := BC_DU

	_, trailmap := walkGuard(guardPos, guardX, guardY, lines, false)

	steps := 0
	for i := range trailmap {
		for j := range trailmap[i] {
			if trailmap[i][j] > 0 {
				steps++
			}
		}
	}
	return steps
}

func solvePart2(args []string) int {
	fn := args[0]
	res := 0

	// Parse input file
	lines, err := aoc_helpers.ReadLinesToByteSlice(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Find starting position of guard
	startX, startY := findGuardPos(lines)
	guardX, guardY := startX, startY
	guardPos := BC_DU

	_, trailmap := walkGuard(guardPos, guardX, guardY, lines, false)

	// Check all trail positions for potential loops
	for i := range trailmap {
		for j := range trailmap[i] {
			// Check if the position is already an obstacle, or if this is the guard starting position
			// If so, skip this position
			if trailmap[i][j] == 0 || (lines[i][j] == OBSTACLE) || (i == startX && j == startY) {
				continue
			}

			if *dbgFlag {
				fmt.Printf("[DEBUG] Checking position %d, %d - %b\n", i, j, trailmap[i][j])
			}

			if *dbgFlag {
				fmt.Printf("[DEBUG] Testing obstacle at (%d, %d)\n", i, j)
			}
			tmp := lines[i][j]
			lines[i][j] = OBSTACLE
			out, _ := walkGuard(guardPos, startX, startY, lines, true)
			lines[i][j] = tmp

			if !out {
				// Put a new obstacle in front of current position
				lines[i][j] = NEW_OBSTACLE
				res++
				if *dbgFlag {
					fmt.Printf("O #%d (%d, %d)\n", res, i, j)
				}
			}
		}
	}

	if *dbgFlag {
		printMap(lines, trailmap)
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
