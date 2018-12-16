package algo

import "testing"

func TestFindNumSameAsIndex(t *testing.T) {
	for _, v := range []struct {
		nums   []int
		wanted int
	}{
		{[]int{-1, 0, 1, 3, 5, 6}, 3},
		{[]int{-1, 0, 1, 2, 3, 4}, -1},
	} {
		got := findNumSameAsIndex(v.nums)
		if got != v.wanted {
			t.Fatalf("Failed, expected %d, got %d", v.wanted, got)
		}
	}
}
