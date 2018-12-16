package algo

import "testing"

func TestFindGreatestSumOfSubArray(t *testing.T) {
	for _, v := range []struct {
		nums   []int
		wanted int
	}{
		{[]int{}, 0},
		{nil, 0},
		{[]int{1, -2, 3, 10, -4, 7, 2, -5}, 18},
	} {
		got := findGreatestSumOfSubArray(v.nums)
		if got != v.wanted {
			t.Errorf("Failed, expected %d, got %d", v.wanted, got)
		}
	}
}
