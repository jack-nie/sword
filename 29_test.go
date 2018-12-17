package algo

import (
	"reflect"
	"testing"
)

func TestPrintMatrixClockwisely(t *testing.T) {
	for _, v := range []struct {
		nums   [][]int
		wanted []int
	}{
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		{[][]int{{1, 2, 3, 4}}, []int{1, 2, 3, 4}},
		{[][]int{{1}, {2}, {3}, {4}}, []int{1, 2, 3, 4}},
		{[][]int{{1}}, []int{1}},
		{nil, nil},
	} {
		got := printMatrixClockwisely(v.nums)
		if !reflect.DeepEqual(v.wanted, got) {
			t.Errorf("Failed, expected %d, got %d", v.wanted, got)
		}
	}
}
