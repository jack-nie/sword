package algo

import "testing"

func TestRouteInMatrix(t *testing.T) {
	for _, v := range []struct {
		matrix [][]string
		strs   []string
		wanted bool
	}{
		{[][]string{{"a", "b", "t", "g"}, {"c", "f", "c", "s"}, {"j", "d", "e", "h"}}, []string{"b", "f", "c", "e"}, true},
		{[][]string{{"a", "b", "t", "g"}, {"c", "f", "c", "s"}, {"j", "d", "e", "h"}}, []string{"a", "b", "f", "b"}, false},
	} {
		if got := routeInMatrix(v.matrix, v.strs); got != v.wanted {
			t.Errorf("Failed, expected %t, got %t", v.wanted, got)
		}
	}
}
