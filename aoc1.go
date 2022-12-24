package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}

	lines, err := readLines(os.Args[1])
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
			calories = 0
			elf_count++
			fmt.Println("line", i, "\nCounting new Elf", elf_count)
		}
		cal, _ := strconv.Atoi(line)
		calories += cal
	}
	fmt.Println("Elf #", max_elf, "calories", max_calories)
}
