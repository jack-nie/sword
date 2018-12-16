package algo

import "testing"

func TestFindMissingNums(t *testing.T) {
	for _, v := range []struct {
		nums   []int
		n      int
		wanted int
	}{
		{[]int{0, 1, 2, 4, 5, 6}, 6, 3},
		{[]int{1, 2, 3, 4, 5, 6}, 6, 0},
	} {
		got := findMissingNums(v.nums, v.n)
		if got != v.wanted {
			t.Fatalf("Failed, expected %d, got %d", v.wanted, got)
		}
	}
}
