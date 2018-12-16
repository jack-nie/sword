package algo

import "testing"

func TestFindInRotateArray(t *testing.T) {
	for _, c := range []struct {
		want int
		nums []int
	}{
		{1, []int{4, 5, 6, 1, 2, 3}},
		{1, []int{1, 2, 3}},
		{0, []int{1, 1, 1, 0, 1, 1}},
		{-1, []int{}},
	} {
		got := findInRotateArray(c.nums)
		if got != c.want {
			t.Errorf("Failed, expected %d, got %d", c.want, got)
		}
	}
}
