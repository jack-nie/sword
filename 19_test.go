package algo

import "testing"

func TestMatch(t *testing.T) {
	for _, v := range []struct {
		str     string
		pattern string
		wanted  bool
	}{
		{"aaa", "a.a", true},
		{"aaa", "ab*ac*a", true},
	} {
		if got := match(v.str, v.pattern); got != v.wanted {
			t.Errorf("failed, expected %t, got %t", v.wanted, got)
		}
	}
}
