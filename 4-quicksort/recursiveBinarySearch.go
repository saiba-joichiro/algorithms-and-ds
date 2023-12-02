package quicksort

// Divide and Conquer version of binary search
func BinarySearch(arr []int, target int) int {
	return recursiveBinarySearch(arr, target, 0, len(arr)-1)
}

func recursiveBinarySearch(arr []int, target int, start int, end int) int {
	mid := (start + end) / 2
	if start > end {
		return -1
	}
	if arr[mid] == target {
		return mid
	}
	if arr[mid] > target {
		return recursiveBinarySearch(arr, target, start, mid-1)
	}
	if arr[mid] < target {
		return recursiveBinarySearch(arr, target, mid+1, end)
	}
	return -1
}
