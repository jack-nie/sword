package algo

import "testing"

func TestMoreThanHalfNums(t *testing.T) {
	for _, v := range []struct {
		nums   []int
		wanted int
	}{
		{[]int{1, 2, 2, 2, 2, 2, 3}, 2},
		{[]int{1}, 1},
		{[]int{}, -1},
		{nil, -1},
		{[]int{1, 2, 3, 3, 4, 5, 2}, -1},
	} {
		if got := moreThanHalfNums(v.nums); got != v.wanted {
			t.Errorf("Failed, expected %d, got %d", v.wanted, got)
		}
	}
}
