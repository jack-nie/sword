package algo

import "testing"

func TestAdd(t *testing.T) {
	for _, tt := range []struct {
		num1   int
		num2   int
		wanted int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{1, -1, 0},
		{-1, -1, -2},
	} {
		if got := add(tt.num1, tt.num2); got != tt.wanted {
			t.Errorf("failed, expected %d, got %d", tt.wanted, got)
		}
	}
}
