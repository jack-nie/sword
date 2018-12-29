package algo

import "testing"

func TestFindNumsOf1Between1AndN(t *testing.T) {
	for _, v := range []struct {
		n      int
		wanted int
	}{
		{12, 5},
		{2, 1},
		{1000, 301},
	} {
		if got := findNumsOf1Between1AndN(v.n); got != v.wanted {
			t.Errorf("Failed, expeted %d, got %d", v.wanted, got)
		}
	}
}
