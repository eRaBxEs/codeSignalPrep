
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
	count := 0
	success := 0
	for sweepNum := 0; sweepNum < len(sequence); sweepNum++ {
		for i := 0; i < len(sequence)-1; i++ {
			if sequence[i] >= sequence[i+1] { // {1, 2, 5, 3, 5}
				count += 1
				if success > 0 {
					if sequence[i] >= sequence[i+1] {
						// Delete any element that is greater than it's successor
						sequence = append(sequence[:i], sequence[i+1:]...)
						// reset sweepNum back to 0
						sweepNum = 0
						break
					}
					// Delete any element that is greater than it's successor
					sequence = append(sequence[:i+1], sequence[i+1+1:]...)
					// reset sweepNum back to 0
					sweepNum = 0
					break
				}
				// Delete any element that is greater than it's successor
				sequence = append(sequence[:i], sequence[i+1:]...)
				// reset sweepNum back to 0
				sweepNum = 0
				break
			}
			success++
		}
	}
	if count > 1 {
		return false
	}
	return true
}
