package introgates

/*
You are given a two-digit integer n. Return the sum of its digits.

Example

For n = 29, the output should be
addTwoDigits(n) = 11.

Input/Output

[time limit] 4000ms (php)
[input] integer n

A positive two-digit integer.

Constraints:
10 ≤ n ≤ 99.

[output] integer

The sum of the first and second digits of the input number.
*/

func AddTwoDigits(n int) int {
	var result, dig int = 0, 0
	var cur int = n

	for cur >= 10 {
		dig = cur % 10 		// get current digit
		cur = (cur - dig) / 10 	// create new number without the last digit
		result += dig		// add result of sm of digits
	}

	// add the last digit
	result += cur

	return result
}
