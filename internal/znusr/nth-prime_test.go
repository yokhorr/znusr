package znusr

import "testing"

func TestNthPrimeBasic(t *testing.T) {
	cases := []struct {
		n, want int
	}{
		{1, 2},
		{2, 3},
		{3, 5},
		{4, 7},
		{5, 11},
		{6, 13},
		{7, 17},
	}
	for _, c := range cases {
		got := NthPrime(c.n)
		if got != c.want {
			t.Errorf("NthPrime(%d) = %d, want %d", c.n, got, c.want)
		}
	}
}
