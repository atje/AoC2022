package main

import (
	"AoC2022/aoc_helpers"
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

//const DOT byte = '.'
//const ANTINODE byte = '#'

var dbgFlag = flag.Bool("d", false, "debug flag")
var traceFlag = flag.Bool("t", false, "trace flag")

// File format:
// Even positions (starting with first position being 0) are file
// Odd positions are space
// The file ID is the file number position divided by 2

type blockT struct {
	id    int // File ID
	start int // Start position
	len   int // Length of the file
}

func string2Blocks(s string) []blockT {
	blocks := make([]blockT, 0)
	id, pos := 0, 0

	// Split the string into blocks
	for i := 0; i < len(s); i += 2 {
		if s[i] == ' ' {
			continue
		}

		length := int(s[i] - '0')

		block := blockT{
			id:    id,
			start: pos,
			len:   length,
		}
		blocks = append(blocks, block)
		id++
		padding := 0
		if i+1 < len(s) {
			padding = int(s[i+1]) - '0'
		}
		pos += length + padding
	}

	return blocks
}

func printBlocks(blocks []blockT) {
	for _, block := range blocks {
		fmt.Printf("Block ID: %d, Start: %d, Length: %d\n", block.id, block.start, block.len)
	}
}

func compactBlocks(blocks []blockT) []blockT {
	compacted := make([]blockT, 0)

	// p = pointer to current position in the file
	p := 0
	for n := 0; n < len(blocks); n++ {
		if blocks[n].start > p {
			// If the start of the block is greater than the current position
			// Add block from right until the space is filled
			if blocks[n].len <= 0 {
				break
			}
			for blocks[n].start > p {
				last := len(blocks) - 1
				cnt := blocks[n].start - p
				if cnt > blocks[last].len {
					cnt = blocks[last].len
				}
				newBlock := blockT{
					id:    blocks[last].id,
					start: p,
					len:   cnt,
				}
				compacted = append(compacted, newBlock)
				p += cnt
				blocks[last].len -= cnt

				if blocks[last].len == 0 {
					blocks = blocks[:last]
					if len(blocks) <= n {
						return compacted
					}
				}
			}
		}
		if len(blocks) > n-1 {
			compacted = append(compacted, blocks[n])
			p += blocks[n].len
		}
	}
	return compacted
}

func checkSum(blocks []blockT) int {
	sum, pos := 0, 0

	for _, block := range blocks {
		for i := block.len; i > 0; i-- {
			sum += pos * block.id
			pos++
		}
	}
	return sum
}

func solvePart1(args []string) int {
	fn := args[0]

	// Parse input file
	lines, err := aoc_helpers.ReadLines(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// We're only interested in the first line
	// Parse the first line
	blocks := string2Blocks(lines[0])
	// Print the blocks
	if *dbgFlag {
		log.Debugf("[DEBUG] Blocks:")
		printBlocks(blocks)
	}

	// Compact the blocks
	blocks = compactBlocks(blocks)
	// Print the compacted blocks
	if *dbgFlag {
		log.Debugf("[DEBUG] Blocks:")
		printBlocks(blocks)
	}

	return checkSum(blocks)
}

func solvePart2(args []string) int {
	/*fn := args[0]

	// Parse input file
	lines, err := aoc_helpers.ReadLinesToByteSlice(fn)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Find antennas, add them to the hashmap containing the antennas
	pairs := parseAntennas(lines)

	// Go through all found antennas and add antinodes to the map
	m := initializeAndFillMap(lines, pairs, true)
	*/
	return 0
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

	// Set up logging
	aoc_helpers.SetupLogging(dbgFlag, traceFlag)

	if len(args) == 0 {
		log.Fatalln("Please provide input file!")
	}

	fmt.Println("part 1:", solvePart1(args))
	fmt.Println("part 2:", solvePart2(args))
}
