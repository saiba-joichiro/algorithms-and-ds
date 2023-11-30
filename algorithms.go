package algorithms_and_ds

import "fmt"

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

// Return index of min element
func findMinElement(arr []int) int {
	if len(arr) == 0 {
		panic("Empty array")
	}
	smallest := arr[0]
	smallestIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}

	return smallestIndex
}

// Selection sort algorithm
// Return new sorted array
func SelectionSort(arr []int) []int {
	size := len(arr)
	sorted := make([]int, size)
	for i := 0; i < size; i++ {
		smallest := findMinElement(arr)
		fmt.Println(smallest)
		sorted[i] = arr[smallest]
		// TODO: Read about this. Do we modify input array because of this?
		arr = append(arr[:smallest], arr[smallest+1:]...)
	}
	return sorted
}
