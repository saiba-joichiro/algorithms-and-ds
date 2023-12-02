package quicksort

import "testing"

func TestMaxNumber(t *testing.T) {
	var tests = []struct {
		in   []int
		want int
	}{
		{[]int{22, 44, 55, 66, 99}, 99},
		{[]int{-232, -131, -10, -1}, -1},
		{[]int{-232, 131, 10, -1, 322}, 322},
		{[]int{32}, 32},
		{[]int{-2}, -2},
		// TODO: add empty arr test case
	}

	for _, test := range tests {
		if got := MaxNumber(test.in); got != test.want {
			t.Errorf("MaxNumber(%v) = %d, but got %d", test.in, test.want, got)
		}
	}
}
