package algo

import "testing"

func TestVerifySequenceOfBST(t *testing.T) {
	for _, c := range []struct {
		nums   []int
		length int
		wanted bool
	}{
		{[]int{5, 7, 6, 9, 11, 10, 8}, 7, true},
		{[]int{}, 0, false},
		{[]int{7, 4, 6, 5}, 4, false},
	} {
		got := verifySequenceOfBST(c.nums, c.length)
		if got != c.wanted {
			t.Errorf("Failed, expected %t, got %t", c.wanted, got)
		}
	}
}
