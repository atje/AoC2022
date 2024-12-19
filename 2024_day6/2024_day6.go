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

// Use bitpositions to indicate earlier point passing movement direction of guard
// There are eigth possible ways the guard could pass through a point
const EMPTY byte = 0
const BC_LR byte = 0b01
const BC_RL byte = 0b10
const BC_UD byte = 0b100
const BC_DU byte = 0b1000

/*
const BC_LD byte = 0b10000
const BC_DR byte = 0b100000
const BC_RU byte = 0b1000000
const BC_UL byte = 0b10000000
*/
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
			//fmt.Println("Found guard in position ", guardX, guardY)
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

func moveGuard(dir byte, x, y int, rotate90 bool, reverse ...bool) (byte, int, int) {
	nx := x
	ny := y
	ndir := dir
	rev := false

	if len(reverse) > 0 {
		rev = reverse[0]
	}

	if rotate90 {
		switch {
		case dir&BC_DU > 0:
			nx = nx + 1
			ny = ny + 1
			ndir = BC_LR
		case dir&BC_UD > 0:
			nx = nx - 1
			ny = ny - 1
			ndir = BC_RL
		case dir&BC_LR > 0:
			nx = nx + 1
			ny = ny - 1
			ndir = BC_UD
		case dir&BC_RL > 0:
			nx = nx - 1
			ny = ny + 1
			ndir = BC_DU
		}
	} else {
		v := 1
		if rev {
			v = -1
		}
		switch {
		case dir&(BC_DU) > 0:
			nx = nx - v
		case dir&(BC_UD) > 0:
			nx = nx + v
		case dir&(BC_LR) > 0:
			ny = ny + v
		case dir&(BC_RL) > 0:
			ny = ny - v
		}
	}

	//	fmt.Printf("moveGuard(%b, %d,%d,%t, %t) --> %b, %d, %d\n", dir, x, y, rotate90, rev, ndir, nx, ny)
	return ndir, nx, ny
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
	switch {
	case dir&(BC_DU) > 0:
		return BC_LR
	case dir&(BC_UD) > 0:
		return BC_RL
	case dir&(BC_LR) > 0:
		return BC_UD
	case dir&(BC_RL) > 0:
		return BC_DU
	}
	return dir
}

// Check if there is a somewhere from the startx, starty position moving in direction dir
// A path has to contain an obstacle preceeded by a tail in the same direction as dir
func checkPath(obsmap, trailmap [][]byte, dir byte, startx, starty int) bool {
	nx, ny := 0, 0
	ndir := dir

	px := startx
	py := starty
	for {
		switch {
		case ndir&(BC_DU) > 0:
			nx = -1
		case ndir&(BC_UD) > 0:
			nx = 1
		case ndir&(BC_LR) > 0:
			ny = 1
		case ndir&(BC_RL) > 0:
			ny = -1
		}

		px = px + nx
		py = py + ny

		fmt.Println("Moving ", dirStr(ndir), px, py)

		if px == startx && py == starty {
			fmt.Println("Loop!")
			// We're in a loop if we're back where started
			return true
		}

		if outOfMatrix(trailmap, px, py) {
			fmt.Println("Out")
			return false
		} else if obsmap[px][py] == OBSTACLE {
			fmt.Println("Turning right")
			ndir = rotClockwise(ndir)
		}
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

	// Find starting position of guard
	guardX, guardY := findGuardPos(lines)
	guardPos := BC_DU

	// Simulate guard walking until she leaves map
	for {
		newX, newY := 0, 0
		lines[guardX][guardY] = BREADCRUMB

		guardPos, newX, newY = moveGuard(guardPos, guardX, guardY, false)

		// Check if guard leaves the matrix
		if outOfMatrix(lines, newX, newY) {
			break
		}

		// Check for obstacles in new position, turn right 90 degrees
		if lines[newX][newY] == OBSTACLE {
			guardPos, newX, newY = moveGuard(guardPos, newX, newY, true)

		}

		// Check if guard leaves the matrix
		if outOfMatrix(lines, newX, newY) {
			break
		}

		guardX = newX
		guardY = newY
	}

	// Count breadcrumbs on the map
	for _, l := range lines {
		for _, b := range l {
			if b == BREADCRUMB {
				res = res + 1
			}
		}
	}

	return res
}

func solvePart2(args []string) int {
	fn := args[0]
	res := 0

	// Parse input file
	obsmap, err := aoc_helpers.ReadLinesToByteSlice(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Create an empty trail map
	trailmap := make([][]byte, len(obsmap))

	for i := range trailmap {
		trailmap[i] = make([]byte, len(obsmap[0]))
	}

	// Find starting position of guard
	startX, startY := findGuardPos(obsmap)
	guardX, guardY := startX, startY
	guardPos := BC_DU

	// Simulate guard walking until she leaves map
	for {
		newX, newY := 0, 0

		trailmap[guardX][guardY] = genBC(trailmap[guardX][guardY], guardPos)

		guardPos, newX, newY = moveGuard(guardPos, guardX, guardY, false)

		// Check if guard leaves the matrix
		if outOfMatrix(obsmap, newX, newY) {
			break
		}

		// Check for obstacles in new position, turn right 90 degrees
		if obsmap[newX][newY] == OBSTACLE {
			guardPos, newX, newY = moveGuard(guardPos, newX, newY, true)
			trailmap[guardX][guardY] = genBC(trailmap[guardX][guardY], guardPos)

			// Check if guard leaves the matrix
			if outOfMatrix(obsmap, newX, newY) {
				break
			}

		}

		// Move guard to new position
		guardX = newX
		guardY = newY
	}

	// Check trail map for potential loops by walking through it again
	guardX, guardY = startX, startY
	guardPos = BC_DU

	for {
		newX, newY := 0, 0

		guardPos, newX, newY = moveGuard(guardPos, guardX, guardY, false)

		if outOfMatrix(obsmap, newX, newY) {
			break
		}

		// Check for obstacles in new position, turn right 90 degrees
		if obsmap[newX][newY] == OBSTACLE {
			guardPos, newX, newY = moveGuard(guardPos, newX, newY, true)

			// Check if guard leaves the matrix
			if outOfMatrix(obsmap, newX, newY) {
				break
			}

		}
		fmt.Println("Guard move ", newX, newY)

		newDir := rotClockwise(guardPos)
		if checkPath(obsmap, trailmap, newDir, guardX, guardY) {
			obsmap[newX][newY] = NEW_OBSTACLE
			res = res + 1
			fmt.Printf("O #%d (%d, %d)\tguardPos %s\n", res, newX, newY, dirStr(guardPos))
		}

		guardX = newX
		guardY = newY
	}

	printMap(obsmap, trailmap)

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
