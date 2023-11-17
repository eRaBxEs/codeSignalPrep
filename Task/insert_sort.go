package main

import (
	"fmt"
)

func main() {
	list := []int{8, 2, 5, 7, 0}
	fmt.Printf("%#v\n", list)
	InsertSortTask(list)
	fmt.Printf("%#v\n", list)

}

func InsertSortTask(list []int) {
	var sorted []int

	for _, item := range list {
		sorted = insertSort(sorted, item)
	}

	for i, val := range sorted {
		list[i] = val
	}
	// copy(list, sorted)

}

func insertSort(sortedList []int, item int) []int {
	for i, sortedItem := range sortedList {
		if item < sortedItem {
			return append(sortedList[:i], append([]int{item}, sortedList[i:]...)...)
		}
	}

	return append(sortedList, item)
}

func BubbleSortInt(list []int) {
	for sweepNum := 0; sweepNum < len(list); sweepNum++ {
		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
			}
		}
	}
}
