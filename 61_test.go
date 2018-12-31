package algo

import "testing"

func TestIsContinues(t *testing.T) {
	for _, tt := range []struct {
		nums   []int
		wanted bool
	}{
		{[]int{}, false},
		{[]int{1, 2, 3, 0, 0}, true},
		{[]int{1, 2, 3, 6, 0}, false},
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{2, 3, 4, 5, 0}, true},
	} {
		if got := isContinues(tt.nums); got != tt.wanted {
			t.Errorf("failed, expected %t, got %t", tt.wanted, got)
		}
	}
}
