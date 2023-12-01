package selection_sort

import "testing"

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
