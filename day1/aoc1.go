package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// Read input file and return a slice containing the number
// of total calories per elf.
func readFile(name string) []int {
	input, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	elfCalories := []int{0}
	elfIndex := 0
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			calories, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			elfCalories[elfIndex] += calories
		} else {
			elfCalories = append(elfCalories, 0)
			elfIndex++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return elfCalories
}

// Return the sum of calories for the elfs with highest number of
// calories.
func sumTopElfs(calories []int, elfs int) int {
	sort.Ints(calories)
	sum := 0
	for i := 1; i <= elfs; i++ {
		sum += calories[len(calories)-i]
	}
	return sum
}

func part1(calories []int) int {
	return sumTopElfs(calories, 1)
}

func part2(calories []int) int {
	return sumTopElfs(calories, 3)
}

func main() {
	calories := readFile("input.txt")
	fmt.Println("Part 1: ", part1(calories))
	fmt.Println("Part 2: ", part2(calories))
}
