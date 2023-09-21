/***
getting the sum of concatenation for the a slice

***/

func solution(a []int) int64 {
	var sumConcat int64
	var stringConcat string
	var concat int64

	digitIndex := make(map[int]string)

	// convert all elements of a to strings and store them in digitIndex
	for i, v := range a {
		digitIndex[i] = strconv.Itoa(v)
	}

	// iterate over the keys of digitIndex
	for i := range digitIndex {
		// concatenate the current element with all other elements in a
		for _, v := range a {
			stringConcat = digitIndex[i] + strconv.Itoa(v)
			concat, _ = strconv.ParseInt(stringConcat, 10, 64)
			sumConcat += concat
		}
	}
	return sumConcat
	// https://go.dev/play/p/FxzpWBk1o7M
}