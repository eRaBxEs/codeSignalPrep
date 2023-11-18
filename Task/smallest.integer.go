package main

/**
Write a function:

func Solution(A []int) int

that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

Given A = [1, 2, 3], the function should return 4.

Given A = [−1, −3], the function should return 1.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000..1,000,000].

***/
func Solution(A []int) int {
	// Implement your solution here
	// create a map to store the elements of A
	m := make(map[int]bool)
	// loop through A and add each elements to the map
	for _, a := range A {
		m[a] = true
	}

	// start from 1 and check if it is in A
	// if not return it as the answer
	// otherwise increment it by 1 and repeat

	ans := 1
	for {
		if _, ok := m[ans]; !ok {
			return ans
		}
		// increment ans if not found
		ans++
	}
}
