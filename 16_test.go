package algo

import "testing"

func TestPower(t *testing.T) {
	for _, v := range []struct {
		base   float64
		exp    int
		wanted float64
	}{
		{float64(1.2), 3, 1.728000},
		{float64(1.2), 4, 2.073600},
		{float64(1.2), -3, 0.578704},
		{float64(1.2), -4, 0.482253},
		{float64(1.2), 0, 1.000000},
		{float64(0), 0, 0.000000},
		{float64(0), -3, 0.000000},
		{float64(0), -4, 0.000000},
		{float64(0), 3, 0.000000},
		{float64(0), 4, 0.000000},
		{float64(-1.2), 3, -1.728000},
		{float64(-1.2), 4, 2.073600},
		{float64(-1.2), -3, -0.578704},
		{float64(-1.2), -4, 0.482253},
		{float64(-1.2), 0, 1.000000},
	} {
		got := power(v.base, v.exp)
		if !doubleEqual(got, v.wanted) {
			t.Errorf("Failed, expected %f, got %f", v.wanted, got)
		}
	}
}
