package znusr

import "testing"

func TestPlog(t *testing.T) {
	cases := []struct {
		n, k, want int
	}{
		// k = 1: n^0 = 1 divides, n^1 does not
		{2, 1, 0},
		{3, 1, 0},
		{10, 1, 0},

		// k < n: no division
		{2, 3, 0},
		{3, 2, 0},
		{5, 7, 0},

		// k = n^p
		{2, 2, 1},
		{2, 4, 2},
		{2, 8, 3},
		{3, 9, 2},
		{3, 27, 3},
		{3, 81, 4},
		{5, 125, 3},

		// k not a pure power
		{2, 12, 2}, // 12 = 2^2 * 3
	}
	for _, c := range cases {
		got := Plog(c.n, c.k)
		if got != c.want {
			t.Errorf("Plog(%d, %d) = %d, want %d", c.n, c.k, got, c.want)
		}
	}
}
