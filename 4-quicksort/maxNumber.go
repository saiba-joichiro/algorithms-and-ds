package quicksort

// Recurive function to find max number in array
func MaxNumber(arr []int) int {
	// TODO: we should return error on len == 0 because there is no valid value for empty arr
	if len(arr) == 1 {
		return arr[0]
	}
	return max(arr[0], MaxNumber(arr[1:]))
}
