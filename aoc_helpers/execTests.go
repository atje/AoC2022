package aoc_helpers

import (
	"fmt"
	"testing"
)

type File_resT struct {
	Args []string
	Res  int
}

func ExecTests(t *testing.T, f func([]string) int, tests []File_resT) {

	failed := false
	for i, test := range tests {
		fmt.Printf("Test #%d: Args = %v\tExpected = %d\t", i, test.Args, test.Res)
		output := f(test.Args)

		if output != test.Res {
			fmt.Printf("--> FAIL (Output %d)\n", output)
			failed = true
			//t.Errorf("\nTest %d: Output %v != expected %v", i, output, test.Res)
		} else {
			fmt.Printf("--> PASS\n")
		}
	}

	if failed {
		t.Fail()
	}

}
