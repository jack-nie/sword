package algo

import (
	"reflect"
	"testing"
)

func TestMultiply(t *testing.T) {
	for _, tt := range []struct {
		nums1  []int
		nums2  []int
		wanted []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{0, 0, 0, 0, 0, 0}, []int{720, 360, 240, 180, 144, 120}},
		{[]int{0, 0, 2, 3, 4, 5}, []int{0, 0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0, 0}},
		{[]int{0, 1, 2, 3, 4, 5}, []int{0, 0, 0, 0, 0, 0}, []int{120, 0, 0, 0, 0, 0}},
		{[]int{-1, 2, 3, 4, 5, 6}, []int{0, 0, 0, 0, 0, 0}, []int{720, -360, -240, -180, -144, -120}},
		{[]int{-1, -2, 3, 4, 5, 6}, []int{0, 0, 0, 0, 0, 0}, []int{-720, -360, 240, 180, 144, 120}},
		{[]int{}, []int{}, []int{}},
	} {
		if got := multiply(tt.nums1, tt.nums2); !reflect.DeepEqual(got, tt.wanted) {
			t.Errorf("failed, expected %d, got %d", tt.wanted, got)
		}
	}
}
