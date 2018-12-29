package algo

import "testing"

func TestReplaceSpace(t *testing.T) {
	for _, v := range []struct {
		str    string
		wanted string
	}{
		{"Hello world!", "Hello%20world!"},
		{"", ""},
		{" ", "%20"},
		{"   ", "%20%20%20"},
		{" Hello", "%20Hello"},
		{"Hello ", "Hello%20"},
		{"Hello world!I am jack.", "Hello%20world!I%20am%20jack."},
		{"Hello   world!", "Hello%20%20%20world!"},
	} {
		if got := replaceSpace(v.str); got != v.wanted {
			t.Errorf("Failed, expected %s, got %s", v.wanted, got)
		}
	}
}
