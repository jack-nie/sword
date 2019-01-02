package algo

import (
	"reflect"
	"testing"
)

func TestMaxInQueue(t *testing.T) {
	for _, tt := range []struct {
		nums   []int
		size   int
		wanted []int
	}{
		{[]int{2, 3, 4, 2, 6, 2, 5, 1}, 3, []int{4, 4, 6, 6, 6, 5}},
	} {
		if got := maxInQueue(tt.nums, tt.size); !reflect.DeepEqual(tt.wanted, got) {
			t.Errorf("failed, expected %d, got %d", tt.wanted, got)
		}
	}
}
