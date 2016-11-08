package introgates

import (
	"testing"
)

func TestLargestNumber(t *testing.T) {
	var tests = []struct {
		n        int // input
		expected int // expected result
	}{
		{0, 0},
		{1, 9},
		{2, 99},
		{5, 99999},
	}

	for _, tt := range tests {
		actual := LargestNumber(tt.n)
		if actual != tt.expected {
			t.Errorf("LargestNumber(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}
