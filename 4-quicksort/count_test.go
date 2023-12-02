package quicksort

import "testing"

func TestCount(t *testing.T) {
	var tests = []struct {
		in   []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 7},
		{[]int{}, 0},
		{[]int{324}, 1},
		{[]int{-12, -23, -231}, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7, 12, 31, 321, 31, 3124, 523, 432}, 14},
	}

	for _, test := range tests {
		if got := Count(test.in); got != test.want {
			t.Errorf("Count(%v) = %d, but got %d", test.in, test.want, got)
		}
	}
}
