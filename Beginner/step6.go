/***

Ratiorg got statues of different sizes as a present from CodeMaster for his birthday,
each statue having an non-negative integer size. Since he likes to make things perfect,
he wants to arrange them from smallest to largest so that
each statue will be bigger than the previous one exactly by 1.
He may need some additional statues to be able to accomplish that.
Help him figure out the minimum number of additional statues needed.

Example

For statues = [6, 2, 3, 8], the output should be
solution(statues) = 3.

Ratiorg needs statues of sizes 4, 5 and 7.

***/

func solution(statues []int) int {
	sortedStatues := bubbleSortInt(statues)
	count := 0
	// 2, 4, 8, 11
	for i, statue := range sortedStatues {
		if i > 0 {
			// get the difference
			diff := statue - sortedStatues[i-1]
			if diff > 1 {
				count = count + (diff - 1)
			}
		}
	}
}

func bubbleSortInt(statues []int) []int {
	for sweepNum := 0; sweepNum < len(statues); sweepNum++ {
		for i := 0; i < len(statues)-1; i++ {
			if statues[i] > statues[i+1] {
				// Do the swap
				statues[i], statues[i+1] = statues[i+1], statues[i]
			}
		}
	}

	return statues
}
