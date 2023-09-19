
/***

Given a sequence of integers as an array,
determine whether it is possible to obtain a strictly increasing sequence
by removing no more than one element from the array.

Note: sequence a0, a1, ..., an is considered to be a strictly increasing
if a0 < a1 < ... < an. Sequence containing only one element is also considered to be strictly increasing.

Example

For sequence = [1, 3, 2, 1], the output should be
solution(sequence) = false.

There is no one element in this array that can be removed in order to get a strictly increasing sequence.

For sequence = [1, 3, 2], the output should be
solution(sequence) = true.

You can remove 3 from the array to get the strictly increasing sequence [1, 2].
Alternately, you can remove 2 to get the strictly increasing sequence [1, 3].



***/

func solution(sequence []int) bool {
	removed := false

	for i := 1; i < len(sequence); i++ {
		// check if the current element (sequence[i]) is less than or equal to the previous element (sequence[i-1])
		if sequence[i] <= sequence[i-1] {
			if removed {
				return false
			}
			removed = true

			// Check if removing the current element or the previous one results in a strictly increasing sequence
			if i == 1 || sequence[i] > sequence[i-2] {
				sequence[i-1] = sequence[i]
			} else {
				sequence[i] = sequence[i-1]
			}
		}
	}

	return true
}

