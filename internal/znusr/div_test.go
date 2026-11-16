package znusr

import "testing"

func TestDivBasic(t *testing.T) {
	cases := []struct {
		a, b, want int
	}{
		{0, 1, 0},
		{0, 5, 0},
		{1, 1, 1},
		{5, 1, 5},
		{7, 2, 3},
		{6, 2, 3},
		{1, 2, 0},
		{2, 2, 1},
		{3, 2, 1},
		{10, 3, 3},
		{9, 3, 3},
		{11, 3, 3},
		{100, 10, 10},
		{99, 10, 9},
	}
	for _, c := range cases {
		got := Div(c.a, c.b)
		if got != c.want {
			t.Errorf("Div(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestDivEdges(t *testing.T) {
	// a < b must give 0
	for a := range 20 {
		for b := a + 1; b < 20; b++ {
			if got := Div(a, b); got != 0 {
				t.Errorf("Div(%d, %d) = %d, want 0", a, b, got)
			}
		}
	}
	// a == b must give 1 (for a > 0)
	for n := 1; n < 50; n++ {
		if got := Div(n, n); got != 1 {
			t.Errorf("Div(%d, %d) = %d, want 1", n, n, got)
		}
	}
	// Div(a, 1) == a
	for a := range 100 {
		if got := Div(a, 1); got != a {
			t.Errorf("Div(%d, 1) = %d, want %d", a, got, a)
		}
	}
}

// a = b * Div(a,b) + Mod(a,b), with 0 <= Mod(a,b) < b.
func TestDivModIdentity(t *testing.T) {
	for a := 10; a <= 30; a++ {
		for b := 1; b < 20; b++ {
			q := Div(a, b)
			r := Mod(a, b)
			if b*q+r != a {
				t.Errorf("a=%d b=%d: b*q+r = %d*%d+%d = %d, want %d",
					a, b, b, q, r, b*q+r, a)
			}
			if r < 0 || r >= b {
				t.Errorf("a=%d b=%d: r = %d out of [0,%d)", a, b, r, b)
			}
		}
	}
}

// Cross-check against Go's built-in division for b > 0.
func TestDivMatchesGoDiv(t *testing.T) {
	for a := 10; a <= 30; a++ {
		for b := 1; b < 20; b++ {
			if got, want := Div(a, b), a/b; got != want {
				t.Errorf("Div(%d, %d) = %d, want %d", a, b, got, want)
			}
		}
	}
}

func TestModBasic(t *testing.T) {
	cases := []struct {
		a, b, want int
	}{
		{0, 1, 0},
		{0, 5, 0},
		{1, 1, 0},
		{5, 1, 0},
		{7, 2, 1},
		{6, 2, 0},
		{1, 2, 1},
		{2, 2, 0},
		{3, 2, 1},
		{10, 3, 1},
		{9, 3, 0},
		{11, 3, 2},
		{100, 10, 0},
		{99, 10, 9},
	}
	for _, c := range cases {
		got := Mod(c.a, c.b)
		if got != c.want {
			t.Errorf("Mod(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestModEdges(t *testing.T) {
	// a < b must give a
	for a := range 20 {
		for b := a + 1; b < 20; b++ {
			if got := Mod(a, b); got != a {
				t.Errorf("Mod(%d, %d) = %d, want %d", a, b, got, a)
			}
		}
	}
	// Mod(a, 1) == 0
	for a := range 50 {
		if got := Mod(a, 1); got != 0 {
			t.Errorf("Mod(%d, 1) = %d, want 0", a, got)
		}
	}
	// Mod(a, a) == 0
	for a := 1; a < 50; a++ {
		if got := Mod(a, a); got != 0 {
			t.Errorf("Mod(%d, %d) = %d, want 0", a, a, got)
		}
	}
}

// Cross-check against Go's built-in division for b > 0.
func TestModMatchesGoMod(t *testing.T) {
	for a := 10; a <= 30; a++ {
		for b := 1; b < 20; b++ {
			if got, want := Mod(a, b), a%b; got != want {
				t.Errorf("Mod(%d, %d) = %d, want %d", a, b, got, want)
			}
		}
	}
}
