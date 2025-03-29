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
				fmt.Println("Found guard in position ", guardX, guardY)
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

// Walk guard in direction heading from position x, y
// Return true if guard leaves the map, false otherwise
// Return true if falling out of area, false otherwise, and number of steps taken
func walkGuard(heading byte, x, y int, areamap [][]byte, loopcheck bool) (bool, [][]byte) {
	nx := x
	ny := y
	lheading := heading

	if *dbgFlag {
		fmt.Printf("walkGuard(%s, %d, %d)\n", dirStr(heading), x, y)
	}

	trailmap := make([][]byte, len(areamap))

	for i := range trailmap {
		trailmap[i] = make([]byte, len(areamap[0]))
	}

	maxstep := len(areamap) * len(areamap[0])
	rotcount := 0

	for {
		px, py := nx, ny

		// Mark trail
		trailmap[nx][ny] = genBC(trailmap[nx][ny], heading)

		switch {
		case heading&(BC_DU) > 0:
			nx = nx - 1
		case heading&(BC_UD) > 0:
			nx = nx + 1
		case heading&(BC_LR) > 0:
			ny = ny + 1
		case heading&(BC_RL) > 0:
			ny = ny - 1
		}

		// Check if guard is out of bounds
		// If so, return true
		if outOfMatrix(areamap, nx, ny) {
			return true, trailmap
		}

		// Check if guard is back to start position heading as started, and have not gotten there by rotating
		// If so, return false
		if loopcheck && nx == x && ny == y && (lheading == heading) && rotcount < 2 { // Check for loop
			return false, trailmap
		}

		if areamap[nx][ny] == OBSTACLE {
			if *dbgFlag {
				//	fmt.Printf("Obstacle at (%d, %d)\n", nx, ny)
			}
			heading = rotClockwise(heading)
			rotcount = rotcount + 1
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
			fmt.Printf("Moving to (%d, %d) %s\n", nx, ny, dirStr(heading))
		}

	}
}

func genBC(cur byte, heading byte) byte {

	return cur | heading
}

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

func leftPos(dir byte, x, y int) (int, int) {
	switch {
	case dir&(BC_DU) > 0:
		return x, y - 1
	case dir&(BC_UD) > 0:
		return x, y + 1
	case dir&(BC_LR) > 0:
		return x - 1, y
	case dir&(BC_RL) > 0:
		return x + 1, y
	}
	return x, y
}

// Rotate clockwise from current direction
func rotClockwise(dir byte) byte {
	return ((dir << 1) | (dir >> 3)) & 0b1111
}

// Rotate counterclockwise from current direction
func rotCounterClockwise(dir byte) byte {
	return ((dir >> 1) | (dir << 3)) & 0b1111
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
				steps = steps + 1
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
	//trailmap[startX][startY] = BC_GUARD_STR

	// Check all trail positions for potential loops
	for i := range trailmap {
		for j := range trailmap[i] {
			if *dbgFlag {
				fmt.Printf("Checking position %d, %d - %b\n", i, j, trailmap[i][j])
			}
			if trailmap[i][j] != 0 {
				//d := trailmap[i][j]

				// Check all directions
				//for d := BC_LR; d <= BC_DU; d = d << 1 {
				//	if *dbgFlag {
				//		fmt.Printf("Checking (%d, %d) %s\n", i, j, dirStr(d))
				//	}

				//	if trailmap[i][j]&d > 0 {
				// Check for loop 90 degrees to the right
				//ldir := rotClockwise(d)
				//obsX, obsY := leftPos(ldir, i, j)

				//if !outOfMatrix(lines, obsX, obsY) {
				// Put a temporary obstacle to the left of the current position
				// Check if the position is already an obstacle, or if this is the guard starting position
				// If so, skip this position
				if (lines[i][j] != OBSTACLE) && !(i == startX && j == startY) {
					if *dbgFlag {
						fmt.Printf("Testing obstacle at (%d, %d)\n", i, j)
					}
					tmp := lines[i][j]
					lines[i][j] = OBSTACLE
					out, _ := walkGuard(guardPos, startX, startY, lines, true)
					lines[i][j] = tmp
					if !out {
						// Put a new obstacle in front of current position
						lines[i][j] = NEW_OBSTACLE
						res = res + 1
						if *dbgFlag {
							fmt.Printf("O #%d (%d, %d)\n", res, i, j)
						}
					}
				}
				//}
				//	}
				//}
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
