package znusr

import "slices"

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

// Add takes two arguments: a, b;
// and returns a + b
func Add(x ...int) int {
	a := U(0)(x...)
	b := U(1)(x...)
	return addN(a)(b)
}

func multN(n int) func(int) int {
	f := zeroVariadic

	g := func(x ...int) int {
		res, _, args := unpack(x...)
		delta := U(0)(args...)
		return addN(delta)(res)
	}

	return func(x int) int {
		return R(f, g)(n, x)
	}
}

// Mult takes two arguments: a, b;
// and returns a * b
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

// Sub takes two arguments: a, b;
// and returns max(a - b, 0)
func Sub(x ...int) int {
	return subN(U(1)(x...))(U(0)(x...))
}

func one(x int) int {
	return N(Z(x))
}

// ge takes two arguments: a, b;
// and returns 1, if a >= b,
//
//	or 0, otherwise
func ge(x ...int) int {
	a := U(0)(x...)
	b := U(1)(x...)
	return Sub(one(a), Sub(b, a))
}

// Div takes two arguments: a, b;
// and return floor(a / b)
func Div(x ...int) int {
	f := zeroVariadic

	g := func(x ...int) int {
		res, yi, args := unpack(x...)
		a := U(0)(args...)
		b := U(1)(args...)

		return res + ge(a, b*N(yi))
	}

	return R(f, g)(append(slices.Clone(x), U(0)(x...))...)
}

// Mod takes two arguments: a, b;
// and return a % b
func Mod(x ...int) int {
	a := U(0)(x...)
	b := U(1)(x...)

	return a - b*Div(a, b)
}

// not returns 1, if x = 0
//
//	or 0, otherwise
func not(x int) int {
	return Sub(1, x)
}

// boolean returns 1, if x = 1,
//
//	or 0, otherwise
func boolean(x int) int {
	return not(not(x))
}

// isMultipleOf returns 1, if a is multiple of b
//
//	or 0, otherwise
func isMultipleOf(a, b int) int {
	return not(Mod(a, b))
}

func isMultipleOf_(a, b int) int {
	if a%b == 0 {
		return 1
	}
	return 0
}

func constantly(a int) func(...int) int {
	return func(...int) int {
		return a
	}
}

// eq returns 1, if a == b
// or 0, otherwise
func eq(a, b int) int {
	return and(ge(a, b), ge(b, a))
}

// gt returns 1, if a > b,
// and 0, otherwise
func gt(a, b int) int {
	A := ge(a, b)
	B := not(eq(a, b))
	return and(A, B)
}

// le returns 1, if a <= b,
// and 0, otherwise
func le(a, b int) int {
	return not(gt(b, a))
}

// and returns 1, if a >= 1 and b >= 1
// or 0, otherwise
func and(a, b int) int {
	return Mult(boolean(a), boolean(b))
}

// isPrime returns 1, if n is prime
// and 0, otherwise
func isPrime(n int) int {
	f := constantly(one(n))

	g := func(x ...int) int {
		res, yi, args := unpack(x...)
		arg := U(0)(args...)

		return Mult(res, not(isMultipleOf(arg, addN(2)(yi))))
	}

	return R(f, g)(n, n-2) * not(eq(n, 1)) * not(eq(n, 0))
}

func neq(a, b int) int {
	return not(eq(a, b))
}

// updIfZero returns a, if a > 0
// and b, if a == 0
func updIfZero(a, b int) int {
	A := a * gt(a, Z(a))
	B := b * eq(a, Z(a))
	return A + B
}

func updIfNonZero(a, b int) int {
	return neq(b, 0)*b + eq(b, 0)*a
}

func updIfNonZero_(a, b int) int {
	if b != 0 {
		return b
	}
	return a
}

// firstPrime returns the first prime number in [l, r]
func firstPrime(a, b int) int {
	f := zeroVariadic

	g := func(x ...int) int {
		res, yi, args := unpack(x...)
		l := U(0)(args...)

		return updIfZero(res, (yi+l)*isPrime(yi+l))
	}

	return R(f, g)(a, b, b-a+one(b))
}

func two(x int) int {
	return N(one(x))
}

// NthPrime returns n-th prime, assuming n >= 1
func NthPrime(n int) int {
	f := constantly(1)

	g := func(x ...int) int {
		res, _, _ := unpack(x...)
		return firstPrime(res+1, 2*(res+1))
	}

	return R(f, g)(n, n)
}

func pow(a, n int) int {
	f := constantly(1)

	g := func(x ...int) int {
		res, _, args := unpack(x...)
		delta := U(0)(args...)
		return res * delta
	}

	return R(f, g)(a, n)
}

func pow_(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res *= a
		}
		a *= a
		n /= 2
	}
	return res
}

// Plog returns max{p | k ⫶ n^p}, assuming n > 1
func Plog(n, k int) int {
	f := constantly(0)

	g := func(x ...int) int {
		res, yi, args := unpack(x...)
		n, k := U(0)(args...), U(1)(args...)
		return updIfNonZero(res, yi*isMultipleOf_(k, pow_(n, yi)))
	}

	return R(f, g)(n, k, k+1)
}

func plog_(n, k int) int {
	p := 0
	for k != 0 && k%n == 0 {
		k /= n
		p++
	}
	return p
}

// nthPrime_ returns the n-th prime for 1 <= n <= 1_000_000.
func nthPrime_(n int) int {
	if n < 1 || n > len(first1000000Primes) {
		panic("NthPrimeTable: n must be in [1, 1_000_000]")
	}
	return first1000000Primes[n-1]
}

func Len(L int) int {
	f := constantly(0)

	g := func(x ...int) int {
		res, yi, args := unpack(x...)
		l := U(0)(args...)
		return updIfNonZero_(res, (yi+1)*isMultipleOf_(l, nthPrime_(yi+1)))
	}

	return R(f, g)(L, L+1)
}

func Sublist(L, I, d int) int {
	f := constantly(1)

	g := func(x ...int) int {
		res, yi, args := unpack(x...)
		l := U(0)(args...)
		i := U(1)(args...)
		return res * (pow_(nthPrime_(yi+1), plog_(nthPrime_(i+yi+1), l)))
	}

	return R(f, g)(L, I, d)
}
