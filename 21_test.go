package algo

import (
	"reflect"
	"testing"
)

func TestReorderOddEven(t *testing.T) {
	for _, v := range []struct {
		nums   []int
		wanted []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 5, 3, 4, 2, 6}},
		{nil, nil},
	} {
		got := reorderOddEven(v.nums)
		if !reflect.DeepEqual(got, v.wanted) {
			t.Errorf("Failed, expected %d, got %d", v.wanted, got)
		}
	}
}

func TestReorderOddEvenTwo(t *testing.T) {
	for _, v := range []struct {
		nums   []int
		fn     func(n int) bool
		wanted []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, isEven, []int{1, 5, 3, 4, 2, 6}},
		{nil, isEven, nil},
	} {
		got := reorderOddEvenTwo(v.nums, isEven)
		if !reflect.DeepEqual(got, v.wanted) {
			t.Errorf("Failed, expected %d, got %d", v.wanted, got)
		}
	}
}
