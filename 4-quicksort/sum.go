package quicksort

// Recursive version of sum algorithm
// Using Divide and Conquer method
func Sum(arr []int) int {
	// base case
	if len(arr) == 0 {
		return 0
	}
	// dividing problem until it becomes the base case
	return arr[0] + Sum(arr[1:])
}
