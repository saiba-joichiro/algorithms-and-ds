package recursion

import "testing"

func TestFactorial(t *testing.T) {
	var tests = []struct {
		in   int
		want int
	}{
		{3, 6},
		{0, 0},
		{1, 1},
		{5, 120},
		{10, 3628800},
	}

	for _, test := range tests {
		if got := Factorial(test.in); got != test.want {
			t.Errorf("Factorial(%v) = %v, but got %v", test.in, test.want, got)
		}
	}

}
