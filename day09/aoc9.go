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

// Move tail so it's adjacent to the head
func moveTail(h *obj, t *obj, storeTrail bool) {

	//dbgP("moveTail", *h, *t)
	dist := pointDist(h.position, t.position)

	// Check tail too far from head
	switch {
	case ((math.Abs(float64(dist.x)) > 1) && (dist.y != 0)) ||
		((dist.x != 0) && (math.Abs(float64(dist.y)) > 1)):
		// Not same column or row, move diagonally
		//dbgP("Diagonal move")
		if math.Signbit(float64(dist.x)) {
			t.position.x--
		} else {
			t.position.x++
		}
		if math.Signbit(float64(dist.y)) {
			t.position.y--
		} else {
			t.position.y++
		}
	case math.Abs(float64(dist.x)) > 1:
		// Same row, move along row
		//dbgP("Row move")
		if math.Signbit(float64(dist.x)) {
			t.position.x--
		} else {
			t.position.x++
		}
	case math.Abs(float64(dist.y)) > 1:
		// Same column, move along column
		//dbgP("Column move")
		if math.Signbit(float64(dist.y)) {
			t.position.y--
		} else {
			t.position.y++
		}
	default:
		// No need to move tail, just return
		//dbgP("No move")
		return
	}
	if storeTrail {
		t.trail = append(t.trail, t.position)
	}

	moveTail(h, t, storeTrail)
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
	for _, line := range lines {
		s := strings.Split(line, " ")

		steps, _ := strconv.Atoi(s[1])
		switch {
		case s[0] == "U":
			knots[0].position.y += steps
			dbgP("Up", s[1], knots[0])
		case s[0] == "D":
			knots[0].position.y -= steps
			dbgP("Down", s[1], knots[0])
		case s[0] == "L":
			knots[0].position.x -= steps
			dbgP("Left", s[1], knots[0])
		case s[0] == "R":
			knots[0].position.x += steps
			dbgP("Right", s[1], knots[0])
		default:
			fmt.Println("Unknown input: ", line)
		}
		for i := 1; i < knotCount; i++ {
			moveTail(&knots[i-1], &knots[i], i == tailInd)
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

	dbgP(knots[tailInd])
	u := uniqueSlice(knots[tailInd].trail)
	dbgP(u)

	fmt.Println("Unique tail positions:", len(u))

}
