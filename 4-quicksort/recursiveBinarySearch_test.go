package quicksort

import "testing"

func TestRecursiveBinarySearch(t *testing.T) {
	var tests = []struct {
		array  []int
		target int
		want   int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 99, -1},
		{[]int{55}, 55, 0},
		{[]int{5, 38}, 38, 1},
	}

	for _, test := range tests {
		if got := BinarySearch(test.array, test.target); got != test.want {
			t.Errorf("BinarySearch(%v,%d) = %v", test.array, test.target, got)
		}
	}
}
