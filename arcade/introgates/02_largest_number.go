package introgates

/*
Given an integer n, return the largest number that contains exactly n digits.

Example

For n = 2, the output should be
largestNumber(n) = 99.

Input/Output

[time limit] 4000ms (go)
[input] integer n

Constraints:
1 ≤ n ≤ 7.

[output] integer

The largest integer of length n.
*/

func LargestNumber(n int) int {
	var num = 0;

	for i := 0; i < n; i++ {
		num = num * 10 + 9;
	}

	return num;
}

