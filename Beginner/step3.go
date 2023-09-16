/**
Given the string, check if it is a palindrome.

Example

For inputString = "aabaa", the output should be
solution(inputString) = true;
For inputString = "abac", the output should be
solution(inputString) = false;
For inputString = "a", the output should be
solution(inputString) = true.

**/

func solution(inputString string) bool {
	lengthOfString := len(inputString)
	count := 1
	for i := 0; i < len(inputString); i++ {
		if inputString[i] != inputString[lengthOfString-count] {
			return false
		}
		count += 1
	}
	return true
}

// https://go.dev/play/p/HY4RIRdJy02
