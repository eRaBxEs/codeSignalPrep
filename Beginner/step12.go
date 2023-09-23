/****
Given an array of integers numbers, find and return the index i of the first integer within the array
that is less than its adjacent integers on both sides. Note that to satisfy these criteria, adjacent integers on both sides must exist.

Return -1 if none of the integers in the given array fit the criteria.

Assume that array indices are 0-based.

Note: You are not expected to provide the most optimal solution, but a solution with time complexity not worse than O(numbers.length3) will fit within the execution time limit.

***/
func findPeakElement(nums []int) int {
	if len(nums) < 3 {
		return -1
	}

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] < nums[i-1] && nums[i] < nums[i+1] {
			return i
		}
	}

	return -1
}


