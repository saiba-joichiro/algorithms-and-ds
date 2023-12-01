package selection_sort

import "fmt"

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
// ATTENTION: modifies original array
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
