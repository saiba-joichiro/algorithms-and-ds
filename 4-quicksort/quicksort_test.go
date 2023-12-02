package quicksort

import "testing"

func TestQuicksort(t *testing.T) {
	var tests = []struct {
		in   []int
		want []int
	}{
		// [2 -5 0 0]      3   [4 77 32]     [-5 0 0 2 3 4 32 77]
		// [-5 0 0] 2 []   +   [] 4 [77 32]  [-5 0 0 2] + [4 32 77]
		// [] -5 [0 0]     +   [32] 77 []    [-5 0 0] + [32 77]
		{[]int{3, 2, 4, 77, 0, -5, 32, 0}, []int{-5, 0, 0, 2, 3, 4, 32, 77}},
		{[]int{}, []int{}},
		{[]int{-31, -222}, []int{-222, -31}},
		{[]int{3}, []int{3}},
	}

	for _, test := range tests {
		got := Quicksort(test.in)
		for i, e := range got {
			if e != test.want[i] {
				t.Errorf("Quicksort(%v) = %v, but got %v", test.in, test.want, got)
				break
			}
		}
	}
}
