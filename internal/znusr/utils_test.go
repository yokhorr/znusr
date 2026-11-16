package znusr

import "testing"

func TestAnd(t *testing.T) {
	cases := []struct {
		a, b, want int
	}{
		{0, 0, 0},
		{0, 1, 0},
		{1, 0, 0},
		{1, 1, 1},
	}
	for _, c := range cases {
		got := and(c.a, c.b)
		if got != c.want {
			t.Errorf("and(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestGe(t *testing.T) {
	cases := []struct {
		a, b, want int
	}{
		{3, 2, 1},
		{2, 3, 0},
		{3, 3, 1},
	}
	for _, c := range cases {
		got := ge(c.a, c.b)
		if got != c.want {
			t.Errorf("ge(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestEq(t *testing.T) {
	cases := []struct {
		a, b, want int
	}{
		{3, 2, 0},
		{2, 3, 0},
		{3, 3, 1},
	}

	for _, c := range cases {
		got := eq(c.a, c.b)
		if got != c.want {
			t.Errorf("eq(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestGt(t *testing.T) {
	cases := []struct {
		a, b, want int
	}{
		{3, 2, 1},
		{2, 3, 0},
		{3, 3, 0},
	}

	for _, c := range cases {
		got := gt(c.a, c.b)
		if got != c.want {
			t.Errorf("gt(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestUpdIfZero(t *testing.T) {
	cases := []struct {
		a, b, want int
	}{
		{0, 2, 2},
		{2, 3, 2},
		{3, 5, 3},
	}
	for _, c := range cases {
		got := updIfZero(c.a, c.b)
		if got != c.want {
			t.Errorf("updIfZero(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestUpdIfNonZero(t *testing.T) {
	cases := []struct {
		a, b, want int
	}{
		{0, 2, 2},
		{2, 3, 3},
		{3, 5, 5},
		{3, 0, 3},
		{5, 0, 5},
	}
	for _, c := range cases {
		got := updIfNonZero(c.a, c.b)
		if got != c.want {
			t.Errorf("updIfNonZero(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestIsMultipleOf(t *testing.T) {
	cases := []struct {
		a, b, want int
	}{
		{4, 2, 1},
		{5, 2, 0},
		{6, 3, 1},
		{7, 3, 0},
	}
	for _, c := range cases {
		got := isMultipleOf(c.a, c.b)
		if got != c.want {
			t.Errorf("isMultipleOf(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestIsPrime(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 0},
		{5, 1},
		{6, 0},
		{7, 1},
		{8, 0},
		{9, 0},
		{10, 0},
	}
	for _, c := range cases {
		got := isPrime(c.n)
		if got != c.want {
			t.Errorf("isPrime(%d) = %d, want %d", c.n, got, c.want)
		}
	}
}

func TestFirstPrime(t *testing.T) {
	cases := []struct {
		l, r, want int
	}{
		{1, 10, 2},
		{10, 20, 11},
		{20, 30, 23},
	}
	for _, c := range cases {
		got := firstPrime(c.l, c.r)
		if got != c.want {
			t.Errorf("firstPrime(%d, %d) = %d, want %d", c.l, c.r, got, c.want)
		}
	}
}

func TestPow(t *testing.T) {
	cases := []struct {
		a, n, want int
	}{
		// n = 0: a^0 = 1 (including 0^0 = 1 by convention)
		{0, 0, 1},
		{1, 0, 1},
		{2, 0, 1},
		{5, 0, 1},
		{10, 0, 1},

		// a = 0, n > 0
		{0, 1, 0},
		{0, 2, 0},
		{0, 5, 0},

		// a = 1
		{1, 1, 1},
		{1, 10, 1},
		{1, 100, 1},

		// small powers
		{2, 1, 2},
		{2, 2, 4},
		{2, 3, 8},
		{2, 4, 16},
		{2, 5, 32},
		{2, 10, 1024},
		{2, 16, 65536},
		{3, 1, 3},
		{3, 2, 9},
		{3, 3, 27},
		{3, 4, 81},
		{3, 5, 243},
		{5, 1, 5},
		{5, 2, 25},
		{5, 3, 125},
		{10, 1, 10},
		{10, 2, 100},
		{10, 3, 1000},
		{7, 3, 343},
	}
	for _, c := range cases {
		got := pow(c.a, c.n)
		if got != c.want {
			t.Errorf("pow(%d, %d) = %d, want %d", c.a, c.n, got, c.want)
		}
	}
}
