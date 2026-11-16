package znusr

import "testing"

func TestLen(t *testing.T) {
	cases := []struct {
		L, want int
	}{
		// empty list
		{1, 0},

		// pure powers of 2: length 1
		{pow_(2, 1), 1},
		{pow_(2, 2), 1},
		{pow_(2, 3), 1},
		{pow_(2, 4), 1},
		{pow_(2, 5), 1},
		{pow_(2, 10), 1},

		// pure powers of 3: length 2
		{pow_(3, 1), 2},
		{pow_(3, 2), 2},
		{pow_(3, 3), 2},
		{pow_(3, 4), 2},
		{pow_(3, 5), 2},
		{pow_(3, 6), 2},

		// pure powers of 5: length 3
		{pow_(5, 1), 3},
		{pow_(5, 2), 3},
		{pow_(5, 3), 3},
		{pow_(5, 4), 3},
		{pow_(5, 5), 3},

		// powers of 7: length 4
		{pow_(7, 1), 4},
		{pow_(7, 2), 4},
		{pow_(7, 3), 4},

		// powers of larger primes
		{pow_(11, 1), 5},
		{pow_(13, 1), 6},
		{pow_(13, 2), 6},
		{pow_(13, 4), 6},
		{pow_(17, 1), 7},
		{pow_(17, 4), 7},
		{pow_(47, 1), 15},
		{pow_(47, 3), 15},

		// two primes
		{pow_(2, 1) * pow_(3, 1), 2},
		{pow_(2, 2) * pow_(3, 1), 2},
		{pow_(2, 2) * pow_(3, 2), 2},
		{pow_(2, 3) * pow_(3, 3), 2},
		{pow_(2, 7) * pow_(3, 3), 2},
		{pow_(2, 5) * pow_(3, 4), 2},

		// three primes
		{pow_(2, 2) * pow_(3, 1) * pow_(5, 1), 3},
		{pow_(2, 3) * pow_(3, 3) * pow_(5, 3), 3},
		{pow_(2, 3) * pow_(3, 5) * pow_(5, 2), 3},

		// four primes
		{pow_(2, 1) * pow_(3, 1) * pow_(5, 1) * pow_(7, 1), 4},
		{pow_(2, 2) * pow_(3, 2) * pow_(5, 2) * pow_(7, 2), 4},
		{pow_(2, 4) * pow_(3, 3) * pow_(5, 2) * pow_(7, 2), 4},

		// five primes
		{pow_(2, 1) * pow_(3, 1) * pow_(5, 1) * pow_(7, 1) * pow_(11, 1), 5},
		{pow_(2, 3) * pow_(3, 2) * pow_(5, 1) * pow_(7, 1) * pow_(11, 1), 5},

		// six primes
		{pow_(2, 1) * pow_(3, 1) * pow_(5, 1) * pow_(7, 1) * pow_(11, 1) * pow_(13, 1), 6},
		{pow_(2, 4) * pow_(3, 2) * pow_(5, 1) * pow_(7, 1) * pow_(11, 1) * pow_(13, 1), 6},

		// seven primes
		{pow_(2, 1) * pow_(3, 1) * pow_(5, 1) * pow_(7, 1) * pow_(11, 1) * pow_(13, 1) * pow_(17, 1), 7},
	}
	for _, c := range cases {
		got := Len(c.L)
		if got != c.want {
			t.Errorf("Len(%d) = %d, want %d", c.L, got, c.want)
		}
	}
}
