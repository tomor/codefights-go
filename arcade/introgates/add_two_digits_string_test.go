package introgates

import (
	"testing"
)

func TestAddTwoDigitsString(t *testing.T) {
	var tests = []struct {
		n        int // input
		expected int // expected result
	}{
		{1, 1},
		{22, 4},
		{123, 6},
	}

	for _, tt := range tests {
		actual := AddTwoDigitsString(tt.n)
		if actual != tt.expected {
			t.Errorf("AddTwoDigitsString(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}