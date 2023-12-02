package quicksort

func Quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var less = []int{}
	var greater = []int{}
	for _, num := range arr[1:] {
		if num < pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}

	less = append(Quicksort(less), pivot)
	greater = append(Quicksort(greater))
	return append(less, greater...)
}
