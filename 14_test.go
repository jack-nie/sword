package algo

import "testing"

func TestMaxCuttingRobe(t *testing.T) {
	for _, v := range []struct {
		n      int
		wanted int
	}{
		{0, 0},
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 4},
		{5, 6},
	} {
		if got := maxCuttingRobe(v.n); got != v.wanted {
			t.Errorf("Failed, expected %d, got %d", v.wanted, got)
		}
	}
}
