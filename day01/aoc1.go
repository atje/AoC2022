package main

import (
	"AoC2022/aoc_helpers"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// Return list of hashmap keys sorted by value in descending order
func sortByCalories(hm map[int]int) []int {
	keys := make([]int, 0, len(hm))

	for key := range hm {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return hm[keys[i]] > hm[keys[j]]
	})
	return keys
}

func sumCalories(keys []int, hm map[int]int, len int) int {
	sum := 0
	for i := 0; i < len; i++ {
		sum += hm[keys[i]]
		fmt.Println("value", hm[keys[i]], "sum", sum)
	}
	return sum
}

func main() {
	const sumElfs = 3

	m := make(map[int]int)

	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}

	lines, err := aoc_helpers.ReadLines(os.Args[1])
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	calories := 0
	max_calories := 0
	max_elf := 1
	elf_count := 1
	for i, line := range lines {
		if line == "" {
			if calories > max_calories {
				max_calories = calories
				max_elf = elf_count
			}
			m[elf_count] = calories
			elf_count++
			calories = 0

			fmt.Println("line", i, "\nCounting new Elf", elf_count)
		}
		cal, _ := strconv.Atoi(line)
		calories += cal
	}
	fmt.Println("Elf #", max_elf, "calories", max_calories)

	cal_count := sumCalories(sortByCalories(m), m, sumElfs)
	fmt.Println("Calorie sum for top", sumElfs, "Elfs:", cal_count)

}
