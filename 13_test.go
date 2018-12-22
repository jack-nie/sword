package algo

import "testing"

func TestMovingCount(t *testing.T) {
	for _, v := range []struct {
		rows      int
		cols      int
		threshold int
		wanted    int
	}{
		{7, 7, 11, 12},
		{12, 1, 12, 1},
		{1, 12, 12, 1},
		{11, 12, 0, 1},
		{11, 12, -1, 1},
	} {
		if got := movingCount(v.threshold, v.rows, v.cols); got != v.wanted {
			t.Errorf("Failed, expected %d, got %d", v.wanted, got)
		}
	}
}
