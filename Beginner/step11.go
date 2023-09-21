// insertion sort algo for slice of int
func InsertionSortInt(list []int) {
	var sorted []int

	// insert and sort
	for _, item := range list {
		sorted = insertSort(sorted, item)
	}

	for i, val := range sorted {
		list[i] = val
	}
}

func insertSort(sorted []int, item int) []int {
	for i, sortedItem := range sorted {
		// if item is smaller than the present sorted item in iteration, shift all the others to the right and append the item
		if item < sortedItem {
			return append(sorted[:i], append([]int{item}, sorted[i:]...)...)
		}
	}
	return append(sorted, item) // hence append the item at the very back
	// https://go.dev/play/p/hy5Gw7wzvTk
}