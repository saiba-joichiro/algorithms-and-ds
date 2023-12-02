package quicksort

import "testing"

func TestSum(t *testing.T) {
	var tests = []struct {
		in   []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{}, 0},
		{[]int{33}, 33},
		{[]int{-3}, -3},
		{[]int{-23, -44, 8, 36}, -23},
	}

	for _, test := range tests {
		if got := Sum(test.in); got != test.want {
			t.Errorf("Sum(%v) = %d, but got %d", test.in, test.want, got)
		}
	}
}
