package algo

import "testing"

func TestFfndGreatestValue(t *testing.T) {
	for _, v := range []struct {
		nums   [][]int
		wanted int
	}{
		{[][]int{{1, 10, 3, 8}, {12, 2, 9, 6}, {5, 7, 4, 11}, {3, 7, 16, 5}}, 22},
	} {
		if got := findGreatestValue(v.nums); got != v.wanted {
			t.Errorf("Failed, expected %d, got %d", v.wanted, got)
		}
	}
}
