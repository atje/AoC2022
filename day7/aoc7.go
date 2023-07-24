/*
--- Day 6: Tuning Trouble ---
*/

package main

import (
	"AoC2022/aoc_helpers"
	"flag"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

var partFlag = flag.Int("p", 0, "part 0 (default) or part 1")
var dbgFlag = flag.Bool("d", false, "debug flag")

var total_size uint64

type dirs []*dir

// First, define directory tree struct with methods
type dir struct {
	name      string
	size      uint64
	parentdir *dir
	subdirs   dirs
}

func (slice dirs) Len() int {
	return len(slice)
}

func (slice dirs) Less(i, j int) bool {
	return slice[i].size > slice[j].size
}

func (slice dirs) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (d *dir) add_subdir(name string) {
	if *dbgFlag {
		fmt.Println("Adding subdirectory", name, "to directory", d.name)
	}
	d.subdirs = append(d.subdirs, &dir{name: name, size: 0, parentdir: d})
}

func (d *dir) find_subdir(name string) *dir {
	for _, subdir := range d.subdirs {
		if subdir.name == name {
			return subdir
		}
	}
	return nil
}

/*
* Add filesize to directory and all it's parent directories
 */
func (d *dir) add_size(n uint64) {
	d.size = d.size + n
	if *dbgFlag {
		fmt.Println("dir", d.name, "new size ", d.size)
	}
	if d.parentdir != nil {
		d.parentdir.add_size(n)
	}
}

/*
* Print a tree of directories, with subdirectories indented
 */
func (d *dir) tree(prefix string) {
	fmt.Println(prefix, "- ", d.name, "(dir, size=", d.size)
	for _, subdir := range d.subdirs {
		subdir.tree(prefix + "  ")
	}
}

/*
* Create a flat list of directories, sorted by size (largest first)
 */
func flatten(d *dir) []*dir {

	df := []*dir{d}

	for _, subdir := range d.subdirs {
		df = append(df, flatten(subdir)...)
	}
	sort.Sort(dirs(df))

	return df
}

func find_smallest(dirs []*dir, tresh_size uint64) dir {
	prev := *dirs[0]
	for _, d := range dirs {
		if d.size < tresh_size {
			return prev
		}
		prev = *d
	}
	return prev
}

// Main function
func main() {

	flag.Parse()
	args := flag.Args()

	total_size = 0

	if len(args) == 0 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	if *partFlag > 1 {
		fmt.Println("p flag not 0 or 1, aborting!")
		return
	}

	root := dir{name: "/", size: uint64(0)}

	dirp := &root

	lines, err := aoc_helpers.ReadLines(args[0])
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println("*** Part", *partFlag+1, "***")

	// Parse commandline commands
	for _, line := range lines {
		w := strings.Split(line, " ")

		switch w[0] {
		case "$":
			// A command
			switch w[1] {
			case "ls":
				// ls, not useful for this exercise
			case "cd":
				// Change of directory, move current directory pointer to new directory (if found)
				switch w[2] {
				case "/":
					dirp = &root
				case "..":
					dirp = dirp.parentdir
				default:
					// Assuming directory change
					d := dirp.find_subdir(w[2])
					if d != nil {
						dirp = d
					} else {
						log.Fatalln("Directrory not found!")
						dirp = dirp.find_subdir(w[2])
					}
				}
			}
		case "dir":
			// A listed directory
			dirp.add_subdir(w[1])
		}
		if i, err := strconv.ParseUint(w[0], 10, 64); err == nil {
			dirp.add_size(i)
		}
	}

	// Print tree from root
	if *dbgFlag {
		root.tree("")
	}

	flatDirs := flatten(&root)

	// Calculate total size of directories smaler than a certain number
	total_size = 0
	for _, d := range flatDirs {
		if *dbgFlag {
			fmt.Println(d.name, d.size)
		}
		if d.size <= 100000 {
			total_size += d.size
		}
	}
	fmt.Println("Sum of sizes =", total_size)

	fmt.Println("*** Part 2 ***")
	fmt.Println("root.size = ", root.size)

	// Calulate required size to achieve
	delta := uint64(30000000)
	if root.size > uint64(70000000) {
		delta += root.size - uint64(70000000)
	} else {
		delta -= uint64(70000000) - root.size
	}

	if *dbgFlag {
		fmt.Println("delta =", delta)
	}
	fmt.Println("Smallest dir size = ", find_smallest(flatDirs, delta).size)
}
