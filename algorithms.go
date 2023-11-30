package algorithms_and_ds

// Binary search function
// On success returns index of searched element
// On fail returns -1
func BinarySearch(array []int, target int) int {
	start := 0
	end := len(array) - 1
	for start <= end {
		mid := (start + end) / 2
		if array[mid] == target {
			return mid
		}
		if array[mid] > target {
			end = mid - 1
		} else if array[mid] < target {
			start = mid + 1
		}
	}

	return -1
}
