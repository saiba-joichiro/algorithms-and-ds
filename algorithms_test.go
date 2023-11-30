package algorithms_and_ds

import (
	"testing"
)

func TestBinarySearchk(t *testing.T) {
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

func TestSelectionSort(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{4, -2, 55, 0, 4, -78}, []int{-78, -2, 0, 4, 4, 55}},
		{[]int{}, []int{}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{32}, []int{32}},
		{[]int{-1, -2, -55, -22}, []int{-55, -22, -2, -1}},
	}

	for _, test := range tests {
		got := SelectionSort(test.input)
		for i, e := range got {
			if e != test.want[i] {
				t.Errorf("SelectionSort(%v) = %v", test.input, test.want)
				break
			}
		}
	}
}
