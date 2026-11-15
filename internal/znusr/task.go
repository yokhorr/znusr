package znusr

func zeroVariadic(x ...int) int {
	return Z(x[0])
}

func addN(n int) func(int) int {
	f := U(0)

	g := func(x ...int) int {
		res, _, _ := unpack(x...)
		return N(res)
	}

	return func(x int) int {
		return R(f, g)(n, x)
	}
}

func multN(n int) func(int) int {
	f := zeroVariadic

	g := func(x ...int) int {
		res, _, _ := unpack(x...)
		return addN(n)(res)
	}

	return func(x int) int {
		return R(f, g)(n, x)
	}
}

// Mult takes two arguments: a, b; and returns a * b
func Mult(x ...int) int {
	return multN(U(0)(x...))(U(1)(x...))
}

func pred(x int) int {
	f := zeroVariadic

	g := func(x ...int) int {
		_, yi, _ := unpack(x...)
		return yi
	}

	return R(f, g)(x, x)
}

func subN(n int) func(int) int {
	f := U(0)

	g := func(x ...int) int {
		res, _, _ := unpack(x...)
		return pred(res)
	}

	return func(x int) int {
		return R(f, g)(x, n)
	}
}

// Sub takes two arguments: a, b; and returns max(a - b, 0)
func Sub(x ...int) int {
	return subN(U(1)(x...))(U(0)(x...))
}
