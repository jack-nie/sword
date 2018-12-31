package algo

import "testing"

func TestMaxProfit(t *testing.T) {
	for _, tt := range []struct {
		nums   []int
		wanted int
	}{
		{[]int{}, 0},
		{[]int{1}, 0},
		{[]int{9, 11, 8, 5, 7, 12, 16, 14}, 11},
	} {
		if got := maxProfit(tt.nums); tt.wanted != got {
			t.Errorf("failed, expected %d, got %d", tt.wanted, got)
		}
	}
}
