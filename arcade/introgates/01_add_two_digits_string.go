package introgates

import (
	"strconv"
	"strings"
)

/*
You are given a two-digit integer n. Return the sum of its digits.

Example

For n = 29, the output should be
addTwoDigits(n) = 11.

Input/Output

[time limit] 4000ms (go)
[input] integer n

A positive two-digit integer.

Constraints:
10 ≤ n ≤ 99.

[output] integer

The sum of the first and second digits of the input number.
*/

func AddTwoDigitsString(n int) int {
	var result = 0
	numAsArray := strings.Split(strconv.Itoa(n), "")

	for _, element := range numAsArray {
		digit, _ := strconv.Atoi(element)
		result += digit
	}

	return result
}
