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

// var partFlag = flag.Int("p", 0, "part 0 (default) or part 1")
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

func uniqueSlice(sl []point) []point {
	res := make([]point, 0)

	prev := point{}

	for _, v := range sl {
		if (prev == point{}) || (prev != point{} && !pointEqual(prev, v)) {
			res = append(res, v)
		}
		prev = v
	}
	return res
}

// Move tail so it's adjacent to the head
func moveTail(h *obj, t *obj) {

	dbgP("moveTail", *h, *t)
	dist := pointDist(h.position, t.position)

	// Check tail too far from head
	switch {
	case ((math.Abs(float64(dist.x)) > 1) && (dist.y != 0)) ||
		((dist.x != 0) && (math.Abs(float64(dist.y)) > 1)):
		// Not same column or row, move diagonally
		dbgP("Diagonal move")
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
		dbgP("Row move")
		if math.Signbit(float64(dist.x)) {
			t.position.x--
		} else {
			t.position.x++
		}
	case math.Abs(float64(dist.y)) > 1:
		// Same column, move along column
		dbgP("Column move")
		if math.Signbit(float64(dist.y)) {
			t.position.y--
		} else {
			t.position.y++
		}
	default:
		// No need to move tail, just return
		dbgP("No move")
		return
	}
	t.trail = append(t.trail, t.position)
	moveTail(h, t)
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

	// Initialize starting positions of head, tail
	head := obj{name: "Head"}
	tail := obj{name: "Tail"}

	// Add initial position to trail
	tail.trail = append(tail.trail, tail.position)

	// Loop through head movements, adjust tail if needed
	for _, line := range lines {
		s := strings.Split(line, " ")

		steps, _ := strconv.Atoi(s[1])
		switch {
		case s[0] == "U":
			head.position.y += steps
			dbgP("Up", s[1], head)
		case s[0] == "D":
			head.position.y -= steps
			dbgP("Down", s[1], head)
		case s[0] == "L":
			head.position.x -= steps
			dbgP("Left", s[1], head)
		case s[0] == "R":
			head.position.x += steps
			dbgP("Right", s[1], head)
		default:
			fmt.Println("Unknown input: ", line)
		}
		moveTail(&head, &tail)
	}
	dbgP("Final position", head)

	// Check unique positions by sorting the trail slice and remove duplicate values
	sort.Slice(tail.trail, func(p, q int) bool {
		if tail.trail[p].y < tail.trail[q].y {
			return true
		}
		if (tail.trail[p].y == tail.trail[q].y) && (tail.trail[p].x < tail.trail[q].x) {
			return true
		}
		return false
	})

	dbgP(tail)
	u := uniqueSlice(tail.trail)
	dbgP(u)

	fmt.Println("Unique tail positions:", len(u))

}
