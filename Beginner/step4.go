
func solution(inputArray []int) int {
	// initialize largest product to the smallest negative number by using shift left operator
	largestProduct := -1 << 63

	// Find the largest product of adjacent elements
	for i := 0; i < len(inputArray)-1; i++ {
		adjacentProduct := inputArray[i] * inputArray[i+1]

		if adjacentProduct > largestProduct {
			largestProduct = adjacentProduct
		}
	}

	return largestProduct
}

// https://go.dev/play/p/sLLR4uZksgu