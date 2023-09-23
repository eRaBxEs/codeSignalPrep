/****

Solve this using golang with the solution function below as snippet :

func solution(numbers[]int, patterns[]int) int {
}

Given an array of integers numbers and an array pattern representing a comparison pattern,
find how many subarrays of numbers match the given pattern. pattern can only contain the following integers:

pattern[i] = 1, represents that the number corresponding to this element of the pattern is greater than the previous one.
pattern[i] = 0, represents that the number corresponding to this element of the pattern is equal to the previous one.
pattern[i] = -1, represents that the number corresponding to this element of the pattern is less than the previous one.
It is guaranteed that the numbers.length > pattern.length.
Note: You are not expected to provide the most optimal solution, but a solution with time complexity
not worse than O(numbers.length × pattern.length) will fit within the execution time limit.

Example

For numbers = [4, 1, 3, 4, 4, 5, 5, 1] and pattern = [1, 0, -1], the output should be solution(numbers, pattern) = 1.

Explanation:
Let's check all possible subarrays of length 3. Note, that the subarray [4, 1, 3], starting with numbers[0] = 4 does
not need to be checked, as there is nothing to compare the first element with.

Subarray [1, 3, 4] doesn't satisfy the pattern. pattern[0] = 1 means that the first element of the subarray
should be greater than the previous one, but numbers[1] = 1 < numbers[0] = 4.
Subarray [3, 4, 4] doesn't satisfy the pattern. pattern[1] = 0 means that the second element of the subarray
should be equal to the previous one, but numbers[3] = 4 != numbers[2] = 3.
Subarray [4, 4, 5] doesn't satisfy the pattern. pattern[2] = -1 means that the third element of the subarray
should be less than the previous one, but numbers[5] = 5 > numbers[4] = 4.
Following the same logic, subarray [4, 5, 5] doesn't satisfy the pattern.
Subarray [5, 5, 1] satisfies the pattern, because:
numbers[5] = 5 > numbers[4] = 4 and pattern[0] = 1;
numbers[6] = numbers[5] = 5 and pattern[1] = 0;
numbers[7] = 1 < numbers[6] = 5 and pattern[2] = -1;
Since there is a single subarray that satisfies the given pattern, the answer is 1.

****/

func solution(numbers []int, patterns []int) int {
	// initialize the count of matching subarrays to zero
	count := 0
	// loop through the numbers array from index 0 to length - pattern length
	for i := 0; i <= len(numbers)-len(patterns); i++ {
		// assume the current subarray matches the pattern until proven otherwise
		match := true
		// loop through the pattern array from index 1 to length - 1
		for j := 1; j < len(patterns); j++ {
			// compare the current and previous elements of the subarray according to the pattern
			if patterns[j] == 1 {
				// if the pattern is 1, the current element should be greater than the previous one
				if numbers[i+j] <= numbers[i+j-1] {
					// if not, the subarray does not match the pattern
					match = false
					break
				}
			} else if patterns[j] == 0 {
				// if the pattern is 0, the current element should be equal to the previous one
				if numbers[i+j] != numbers[i+j-1] {
					// if not, the subarray does not match the pattern
					match = false
					break
				}
			} else if patterns[j] == -1 {
				// if the pattern is -1, the current element should be less than the previous one
				if numbers[i+j] >= numbers[i+j-1] {
					// if not, the subarray does not match the pattern
					match = false
					break
				}
			}
		}
		// if the subarray matches the pattern, increment the count by one
		if match {
			count++
		}
	}
	// return the final count of matching subarrays
	return count
}
