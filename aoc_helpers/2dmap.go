package aoc_helpers

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type TwoDimMap struct {
	x0, y0 int //Origo, corresponding to point 0, 0 in the point map
	point  [][]rune
}

// Expand 2D Map to fit the new coordinates if needed
// Move origo if necessary
func Expand(m TwoDimMap, x, y int) TwoDimMap {
	xCap := 0
	yCap := len(m.point)
	if yCap > 0 {
		xCap = len(m.point[0])
	}

	dX := m.x0 + xCap - x - 1
	dY := m.y0 + yCap - y - 1

	log.Tracef("Expand(%o, %d, %d), dX = %d, dY = %d", m.point, x, y, dX, dY)
	// Within capacity, no change needed

	// Expand down
	if y > m.y0 && dY < 0 {
		for i := 0; i > dY; i-- {
			newRow := make([]rune, xCap)
			m.point = append(m.point, newRow)
		}
	}
	// Expand up, move y0
	// NOT IMPLEMENTED

	if x < m.x0 {
		// Expand left, move x0
		dX = m.x0 - x
		for i := range m.point {
			newCol := make([]rune, dX)
			m.point[i] = append(newCol, m.point[i]...)
		}
		m.x0 = x
	} else if x > m.x0 && dX < 0 {
		// Expand right
		for i := range m.point {
			newCol := make([]rune, -dX)
			m.point[i] = append(m.point[i], newCol...)
		}
	}
	log.Tracef("Expand result = %o", m.point)
	return m
}

// Add char to point in map, expand map if needed
func Add2Map(dm TwoDimMap, x int, y int, c rune) TwoDimMap {
	log.Tracef("Add2Map(%d, %d, %c, %o)", x, y, c, dm)

	// Add rune at point (x,y)
	if log.GetLevel() == log.DebugLevel {
		PrintTwoDimMap(dm)
	}
	log.Tracef("char '%c' to be added in position [%d][%d]\nTwoDimMap.point = %o", c, y, x, dm.point)
	dm.point[y-dm.y0][x-dm.x0] = c

	return dm
}

func PrintTwoDimMap(dm TwoDimMap) {
	fmt.Printf("---CaveMap:---\n(x0, y0) = (%d, %d)\n", dm.x0, dm.y0)
	for n, row := range dm.point {
		fmt.Printf("%d ", n)
		for _, v := range row {
			pointVal := '.'
			if v != 0 {
				pointVal = v
			}
			fmt.Printf("%c", pointVal)
		}
		fmt.Printf("\n")
	}
	fmt.Println("--------------")
}
