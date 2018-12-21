package algo

import "testing"

func TestFibnacci(t *testing.T) {
	for _, v := range []struct {
		n      int
		wanted int
	}{
		{0, 0},
		{-1, 0},
		{1, 1},
		{11, 55},
	} {
		if got := fibnacci(v.n); got != v.wanted {
			t.Errorf("Failed, expected %d, got %d", v.wanted, got)
		}
	}
}
