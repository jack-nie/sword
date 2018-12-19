package algo

import "testing"

func TestIsPopOrder(t *testing.T) {
	for _, v := range []struct {
		push   []int
		pop    []int
		length int
		wanted bool
	}{
		{[]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}, 5, true},
		{[]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}, 5, false},
		{nil, nil, 0, false},
		{[]int{1}, []int{1}, 1, true},
	} {
		if got := isPopOrder(v.push, v.pop, v.length); got != v.wanted {
			t.Errorf("Failed, expected %t, got %t", v.wanted, got)
		}
	}
}
