package algo

import (
	"reflect"
	"testing"
)

func TestPermutation(t *testing.T) {
	for _, v := range []struct {
		str    string
		wanted []string
	}{
		{"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{"", nil},
	} {
		got := permutation(v.str)
		if !reflect.DeepEqual(got, v.wanted) {
			t.Errorf("Failed, expected %s, got %s", v.wanted, got)
		}
	}
}
