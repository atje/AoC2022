package aoc_helpers

import (
	"testing"
)

type File_resT struct {
	Args []string
	Res  int
}

func ExecTests(t *testing.T, f func([]string) int, tests []File_resT) {

	for i, test := range tests {
		output := f(test.Args)

		if output != test.Res {
			t.Errorf("Test %d: Output %v != expected %v", i, output, test.Res)
		}
	}

}
