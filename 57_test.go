package algo

import (
	"reflect"
	"testing"
)

func TestFindNumsWithSum(t *testing.T) {
	for _, v := range []struct {
		nums   []int
		target int
		wanted []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 7, []int{1, 6}},
		{[]int{}, 8, nil},
	} {
		got := findNumsWithSum(v.nums, v.target)

		if !reflect.DeepEqual(got, v.wanted) {
			t.Errorf("Failed, expected  %d, got %d", v.wanted, got)
		}
	}
}
