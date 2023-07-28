/*
--- Day 9: Rope Bridge ---
*/

package main

import (
	"AoC2022/aoc_helpers"
	"flag"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type obj struct {
	name     string
	position point
	trail    []point // A slice of points passed by the object
}

var partFlag = flag.Int("p", 0, "part 0 (default) or part 1")
var dbgFlag = flag.Bool("d", false, "debug flag")
var stopLine = flag.Int("s", 0, "input file line to stop at, EOF if not specified or 0")

func dbgP(a ...any) {
	if *dbgFlag {
		fmt.Println("Debug:", a)
	}
}

func pointDist(a point, b point) (res point) {

	res.x = a.x - b.x
	res.y = a.y - b.y

	return res
}

func pointEqual(a point, b point) bool {
	return a.x == b.x && a.y == b.y

}

// Return slice with unique points from sorted slice sl
// Sort order not important but identical points havee to be adjacent to each other
func uniqueSlice(sl []point) []point {
	res := make([]point, 0)
	prev := sl[0]
	res = append(res, prev)
	for _, v := range sl {
		if !pointEqual(prev, v) {
			res = append(res, v)
		}
		prev = v
	}
	return res
}

func moveIncr(a, b point) point {
	res := point{} // (0,0)

	dist := pointDist(a, b)

	// Check tail too far from head
	// THIS COULD BENEFIT FROM SERIOUS REFACTORING....
	switch {
	case ((math.Abs(float64(dist.x)) > 1) && (dist.y != 0)) ||
		((dist.x != 0) && (math.Abs(float64(dist.y)) > 1)):
		// Not same column or row, move diagonally
		//dbgP("Diagonal move")
		if math.Signbit(float64(dist.x)) {
			res.x--
		} else {
			res.x++
		}
		if math.Signbit(float64(dist.y)) {
			res.y--
		} else {
			res.y++
		}
	case math.Abs(float64(dist.x)) > 1:
		// Same row, move along row
		//dbgP("Row move")
		if math.Signbit(float64(dist.x)) {
			res.x--
		} else {
			res.x++
		}
	case math.Abs(float64(dist.y)) > 1:
		// Same column, move along column
		//dbgP("Column move")
		if math.Signbit(float64(dist.y)) {
			res.y--
		} else {
			res.y++
		}
	}

	return res
}

// Move tail so it's adjacent to the head
func moveTail(h *obj, t *obj, storeTrail bool) {

	inc := moveIncr(h.position, t.position)

	t.position.x += inc.x
	t.position.y += inc.y

	if storeTrail {
		t.trail = append(t.trail, t.position)
	}

	if !pointEqual(inc, point{}) {
		moveTail(h, t, storeTrail)
	}
}

// Verify that no knot is further away than 1,1 from previous knot
func verifyKnots(k []obj) bool {
	for i := 1; i < len(k); i++ {
		d := pointDist(k[i-1].position, k[i].position)
		if d.x > 1 || d.y > 1 {
			return false
		}
	}
	return true
}

// Main function
func main() {

	flag.Parse()
	args := flag.Args()

	// Parse input file into a matrix
	lines, err := aoc_helpers.ReadLines(args[0])
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	knotCount := 2
	if *partFlag == 1 {
		knotCount = 10
	}

	knots := make([]obj, knotCount)
	tailInd := knotCount - 1

	// Initialize starting positions of knots, name them for debugging
	for n := 0; n < knotCount; n++ {
		name := strconv.Itoa(n)
		switch {
		case n == 0:
			name = "Head"
		case n == tailInd:
			name = "Tail"
		}
		knots[n].name = name
	}

	// Add initial position to trail
	knots[tailInd].trail = append(knots[tailInd].trail, knots[tailInd].position)

	// Loop through head movements, adjust tail if needed
	for n, line := range lines {
		if *stopLine > 0 && n >= *stopLine {
			fmt.Println("Stopping at line ", *stopLine)
			break
		}
		s := strings.Split(line, " ")

		steps, _ := strconv.Atoi(s[1])

		for n := 0; n < steps; n++ {
			switch {
			case s[0] == "U":
				knots[0].position.y += 1
				dbgP("Up", s[1], knots[0])
			case s[0] == "D":
				knots[0].position.y -= 1
				dbgP("Down", s[1], knots[0])
			case s[0] == "L":
				knots[0].position.x -= 1
				dbgP("Left", s[1], knots[0])
			case s[0] == "R":
				knots[0].position.x += 1
				dbgP("Right", s[1], knots[0])
			default:
				fmt.Println("Unknown input: ", line)
			}
			for i := 1; i < knotCount; i++ {
				moveTail(&knots[i-1], &knots[i], i == tailInd)
			}
		}
		if !verifyKnots(knots) {
			fmt.Println("invalid knot distances - after command ", line)
		}
		//fmt.Println(knots)
	}
	dbgP("Final position", knots[0])

	// Check unique positions by sorting the trail slice and remove duplicate values
	sort.Slice(knots[tailInd].trail, func(p, q int) bool {
		if knots[tailInd].trail[p].x < knots[tailInd].trail[q].x {
			return true
		}
		if (knots[tailInd].trail[p].x == knots[tailInd].trail[q].x) && (knots[tailInd].trail[p].y < knots[tailInd].trail[q].y) {
			return true
		}
		return false
	})

	dbgP(knots)
	u := uniqueSlice(knots[tailInd].trail)
	dbgP(u)

	fmt.Println("All tail positions:", len(knots[tailInd].trail))
	fmt.Println("Unique tail positions:", len(u))

}
