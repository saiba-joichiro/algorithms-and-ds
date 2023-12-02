package quicksort

// Recursive function for counting the number of items in the list
func Count(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return 1 + Count(arr[1:])
}
