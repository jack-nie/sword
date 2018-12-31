package algo

import "testing"

func TestLastRemainning(t *testing.T) {
	for _, tt := range []struct {
		n      int
		m      int
		wanted int
	}{
		{5, 3, 3},
		{6, 6, 3},
		{6, 7, 4},
		{-1, 0, -1},
	} {
		if got := lastRemainning(tt.n, tt.m); got != tt.wanted {
			t.Errorf("failed, expected %d, got %d", tt.wanted, got)
		}
	}
}
