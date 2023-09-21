/****
	converting a base 10 number to base 2 , inversing it and then converting the inversion to a number in base 10

***/

func changeAds(base10 int32) int32 {
	// Write your code here
	// 30 = 30/2 : 15 r 0
	// 15/2 : 7 r 1
	// 7/2 : 3 r 1
	// 3/2 : 1 r 1
	// 1/2 : 0 r 1
	var val, rem int32
	var strAds, newAds string

	for base10 >= 1 {
		val = base10 / 2
		rem = base10 % 2
		strAds = fmt.Sprintf("%d%s", rem, strAds)
		base10 = val
	}

	for _, str := range strAds {
		if string(str) == "1" {
			newAds = fmt.Sprintf("%s%s", newAds, "0")
		}
		if string(str) == "0" {
			newAds = fmt.Sprintf("%s%s", newAds, "1")
		}
	}
	// Now convert newAds to base10
	// 2pow4 * 1 + 2pow3 * 1 + 2pow2 * 1 + 2pow1 * 1 + 2pow0 * 0
	// 16 + 8 + 4 + 2 + 0 =  30
	// imageine newAds = 11110
	var multiplier float64
	for i := 0; i < len(newAds); i++ {
		pow := len(newAds) - i - 1
		convert, _ := strconv.Atoi(string(newAds[i]))
		multiplier += math.Pow(2, float64(pow)) * float64(convert)
	}

	return int32(multiplier)

}