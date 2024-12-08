package aoc_helpers

import (
	"bufio"
	"os"
)

// Read all lines from a file into an array of strings
func ReadLines(path string) ([]string, error) {
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

// Read all lines from a file into an array of byte arrays
func ReadLinesToByteSlice(path string) ([][]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, []byte(scanner.Text()))
	}
	return lines, scanner.Err()
}
