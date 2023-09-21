/***
	Check if concatenation pair is tiny

***/
func solution(a []int, b []int, k int) int {
	countTiny := 0 // count tiny pairs
	for i := 0; i < len(a); i++ {

		stringConcat := fmt.Sprintf("%d%d", a[i], b[len(b)-i-1])
		concat, _ := strconv.Atoi(stringConcat)

		if concat < k { // pair is tiny
			countTiny++
		}

	}
	return countTiny

	// https://go.dev/play/p/gwxyi79nfi1
}