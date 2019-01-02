package algo

import "testing"

func TestMinNumber(t *testing.T) {
	for _, tt := range []struct {
		nums   []int
		wanted int
	}{
		{[]int{3, 323, 32123}, 321233233},
	} {
		if got := minNumber(tt.nums); got != tt.wanted {
			t.Errorf("failed, expected %d, got %d", tt.wanted, got)
		}
	}
}
