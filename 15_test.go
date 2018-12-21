package algo

import "testing"

func TestNumsOf1(t *testing.T) {
	for _, v := range []struct {
		n      int
		wanted int
	}{
		{1, 1},
		{-1, 64},
		{0, 0},
		{11, 3},
		{9, 2},
	} {
		if got := numsOf1(v.n); got != v.wanted {
			t.Errorf("Failed, expected %d, got %d", v.wanted, got)
		}
	}
}

func TestNumsOf1Two(t *testing.T) {
	for _, v := range []struct {
		n      int
		wanted int
	}{
		{1, 1},
		{-1, 64},
		{0, 0},
		{11, 3},
		{9, 2},
	} {
		if got := numsOf1Two(v.n); got != v.wanted {
			t.Errorf("Failed, expected %d, got %d", v.wanted, got)
		}
	}
}
