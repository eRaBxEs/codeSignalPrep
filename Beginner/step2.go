/**
Given a year, return the century it is in.
The first century spans from the year 1 up to and including the year 100,
the second - from the year 101 up to and including the year 200, etc.

Example
For year = 1905, the output should be
solution(year) = 20;
For year = 1700, the output should be
solution(year) = 17.

**/

func solution(year int) int {
	century := 100
	centuryCount := year / century

	remainderCount := year % century

	if remainderCount > 0 {
		centuryCount += 1
	}

	return centuryCount
}

// https://go.dev/play/p/aB-A7K0edJ7