package znusr

import "testing"

func TestZ(t *testing.T) {
	if Z(0) != 0 {
		t.Errorf("Z(0) = %d; want 0", Z(0))
	}
	if Z(100) != 0 {
		t.Errorf("Z(100) = %d; want 0", Z(100))
	}
}

func TestN(t *testing.T) {
	if N(0) != 1 {
		t.Errorf("N(0) = %d; want 1", N(0))
	}
	if N(9) != 10 {
		t.Errorf("N(9) = %d; want 10", N(9))
	}
}

func TestU(t *testing.T) {
	proj0 := U(0)
	proj1 := U(1)

	if proj0(10, 20) != 10 {
		t.Errorf("U(0)(10, 20) = %d; want 10", proj0(10, 20))
	}
	if proj1(10, 20) != 20 {
		t.Errorf("U(1)(10, 20) = %d; want 20", proj1(10, 20))
	}
}

func TestS(t *testing.T) {
	g := func(x ...int) int {
		return x[0] + x[1]
	}

	f1 := U(0)
	f2 := U(1)

	sub := S(g, f1, f2)

	got := sub(10, 20)
	expected := 30

	if got != expected {
		t.Errorf("S(g, f1, f2)(10, 20) = %d; want %d", got, expected)
	}
}

func TestR(t *testing.T) {
	f := U(0)

	g := func(args ...int) int {
		res, _, _ := unpack(args...)
		return N(res)
	}

	add := R(f, g)

	got := add(5, 3)
	expected := 8

	if got != expected {
		t.Errorf("R addition (5 + 3) = %d; want %d", got, expected)
	}
}
