package algo

import "testing"

func TestFindDup(t *testing.T) {
	for _, c := range []struct {
		want   int
		nums   []int
		length int
	}{
		{1, []int{0, 1, 1, 2}, 3},
		{2, []int{2, 2, 2, 2}, 3},
		{-1, []int{}, 0},
		{1, []int{1, 1, 1, 1}, 3},
	} {
		got := findDup(c.nums, c.length)
		if got != c.want {
			t.Errorf("Failed, expect %d, got %d", c.want, got)
		}
	}
}
