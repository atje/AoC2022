package main

import (
	"AoC2022/aoc_helpers"
	"flag"
	"fmt"
	"os"

	"golang.org/x/exp/slices"

	slices0 "slices"

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
	id    int  // File ID
	start int  // Start position
	len   int  // Length of the file
	moved bool // Moved flag
}

func newBlock(id, start, len int) blockT {
	return blockT{
		id:    id,
		start: start,
		len:   len,
		moved: false,
	}
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

		block := newBlock(id, pos, length)
		blocks = append(blocks, block)

		id++
		padding := 0
		if i+1 < len(s) {
			// Update the position to account for the space
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

func blocksAsString(blocks []blockT) string {
	s := ""
	for n, block := range blocks {
		for i := 0; i < block.len; i++ {
			s += fmt.Sprintf("%d", block.id)
		}
		if len(blocks) > n+1 {
			for i := blocks[n+1].start; i > block.start+block.len; i-- {
				s += "."
			}
		}
	}
	return s
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
				cnt := min(blocks[n].start-p, blocks[last].len)
				newBlock := newBlock(blocks[last].id, p, cnt)
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

func validateBlocks(blocks []blockT) {
	for i := 1; i < len(blocks); i++ {
		if blocks[i].start < blocks[i-1].start+blocks[i-1].len {
			log.Fatalf("Blocks overlap: Block %d overlaps with Block %d", blocks[i-1].id, blocks[i].id)
		}
	}
}

// Find the first gap in the blocks which is at least the size of the given size
// If no gap is found, returns -1 and 0
func findFirstGap(blocks []blockT, size int) int {
	// Find the first gap in the blocks
	gap := 0
	for i := 0; i < len(blocks); i++ {
		if blocks[i].start-gap >= size {
			return i - 1
		}
		gap = blocks[i].start + blocks[i].len
	}
	return -1
}
func compactBlocks2(blocks []blockT) []blockT {
	// p = pointer to current position in the file

	for x := len(blocks) - 1; x >= 0; x-- {
		if blocks[x].moved {
			// current block has been moved, skip it
			continue
		}

		// Search for the first gap from left which will fit the block
		gapPos := findFirstGap(blocks, blocks[x].len)
		if gapPos == -1 {
			// No gap found, break
			continue
		}
		if gapPos >= x {
			// don't move a block if there is no space to the left of it
			continue
		}

		b := blocks[x]
		b.moved = true
		b.start = blocks[gapPos].start + blocks[gapPos].len
		tmp := slices0.Delete(blocks, x, x+1)
		tmp = slices.Insert(tmp, gapPos+1, b)
		blocks = tmp
		x++

		if *dbgFlag {
			log.Debugf("[DEBUG] Blocks: \n%s\n", blocksAsString(blocks))
		}
	}

	validateBlocks(blocks)
	return blocks
}

func checkSum(blocks []blockT) int {
	pos := 0
	sum := 0

	for _, block := range blocks {
		pos = block.start
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
	blocks = compactBlocks2(blocks)
	// Print the compacted blocks
	if *dbgFlag {
		log.Debugf("[DEBUG] Blocks: \n%s\n", blocksAsString(blocks))
	}

	return checkSum(blocks)
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
