package introgates

import (
	"testing"
)

func TestCandies(t *testing.T) {
	var tests = []struct {
		n        int // input
		m        int // input
		expected int // expected result
	}{
		{3, 10, 9},
		{3, 9, 9},
		{1, 1, 1},
		{2, 21, 20},
		{50, 49, 0},
	}

	for _, tt := range tests {
		actual := Candies(tt.n, tt.m)
		if actual != tt.expected {
			t.Errorf("Candies(%d, %d): expected %d, actual %d", tt.n, tt.m, tt.expected, actual)
		}
	}
}
