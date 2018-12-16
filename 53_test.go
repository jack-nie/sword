package algo

import "testing"

func TestFindNumsOfK(t *testing.T) {
	for _, v := range []struct {
		nums   []int
		k      int
		wanted int
	}{
		{[]int{1, 2, 3, 3, 3, 3, 5}, 3, 4},
		{[]int{}, 2, 0},
		{[]int{1, 1, 1, 1, 1}, 3, 0},
	} {
		got := findNumsOfK(v.nums, v.k)
		if got != v.wanted {
			t.Fatalf("Failed, expected %d, got %d", v.wanted, got)
		}
	}
}
