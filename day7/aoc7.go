/*
--- Day 7: No Space Left On Device ---
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

// First, define directory tree struct with methods
type dir struct {
	name      string
	size      uint64
	parentdir *dir
	subdirs   []*dir
}

func (d *dir) add_subdir(name string) {
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

func (d *dir) add_size(n uint64) {
	d.size = d.size + n
	//fmt.Println("new size ", d.size)
}

func (d *dir) tree(prefix string, max_size uint64) uint64 {
	size := d.size
	for _, subdir := range d.subdirs {
		size += subdir.tree(prefix+"  ", max_size)
	}
	if size <= max_size {
		fmt.Println(prefix, "- ", d.name, "(dir, size=", size)
		total_size += size
	}
	return size
}

type bySize []dir

func (a bySize) Len() int           { return len(a) }
func (a bySize) Less(i, j int) bool { return a[i].size > a[j].size }
func (a bySize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (d *dir) find_smallest2delete(delta uint64) *dir {
	if d.size > delta {
		sort.Sort(bySize(d.subdirs))

		for i, subdir := range d.subdirs {
			if subdir.size < delta {
				if i == 0 {
					return d
				} else {
					return subdir.find_smallest2delete(delta)
				}

			}
		}
		return d
	}
	return nil
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

	//	pos := 0
	//	nchar := 4
	root := dir{name: "/", size: 0}

	dirp := &root

	lines, err := aoc_helpers.ReadLines(args[0])
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println("*** Part", *partFlag+1, "***")
	for _, line := range lines {
		w := strings.Split(line, " ")
		//fmt.Println(w)

		switch w[0] {
		case "$":
			//fmt.Println("Found command: ", w[1:])
			switch w[1] {
			case "ls":

			case "cd":
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
			//fmt.Println("Found dir: ", w[1])
			dirp.add_subdir(w[1])
		}
		if i, err := strconv.ParseUint(w[0], 10, 64); err == nil {
			//fmt.Println("Found file: ", w)
			dirp.add_size(i)
		}
	}

	// Print trer from root
	_ = root.tree("", 100000)
	fmt.Println("* Sum of Sizes: ", total_size, " *")

	fmt.Println("*** Part 2 ***")
	d = root.find_smallest2delete(root.size - 70000000 + 30000000)
	fmt.Println("* Smallest directory to delete: ", d.name)
	fmt.Println("* Size: ", d.size)
}
