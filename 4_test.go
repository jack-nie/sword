package algo

import "testing"

func TestFindInTwoDimensionArray(t *testing.T) {
	for _, c := range []struct {
		nums   [][]int
		target int
		want   bool
	}{
		{[][]int{{1, 2, 3, 4}, {2, 3, 4, 5}, {3, 4, 5, 6}, {5, 6, 7, 8}}, 7, true},
		{[][]int{{1, 2, 3, 4}, {2, 3, 4, 5}, {3, 4, 5, 6}, {5, 6, 7, 8}}, 100, false},
		{[][]int{{}}, 1, false},
	} {
		got := findInTwoDimensionArray(c.nums, c.target)
		if got != c.want {
			t.Errorf("Failed, expected %t, got %t", c.want, got)
		}
	}
}
