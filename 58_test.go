package algo

import "testing"

func TestReverseString(t *testing.T) {
	for _, tt := range []struct {
		str    string
		wanted string
	}{
		{"hello world", "world hello"},
		{"", ""},
	} {
		if got := reverseString(tt.str); got != tt.wanted {
			t.Errorf("failed, expected %s, got %s", tt.wanted, got)
		}
	}
}
