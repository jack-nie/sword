package algo

import (
	"reflect"
	"testing"
)

func TestFindListKNums(t *testing.T) {
	for _, v := range []struct {
		nums   []int
		k      int
		wanted []int
	}{
		{[]int{3, 4, 5, 7, 1, 2}, 3, []int{1, 2, 3}},
		{[]int{}, 1, nil},
	} {

		if got := findListKNums(v.nums, v.k); !reflect.DeepEqual(got, v.wanted) {
			t.Errorf("Failed, expected %d, got %d", v.wanted, got)
		}
	}
}
