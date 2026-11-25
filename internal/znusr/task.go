package znusr

var Add = R(U(1, 1), S(N, U(3, 3)))

func Add_(args []int) int {
	return args[0] + args[1]
}

// task 8.8a
func Mult_(args []int) int {
	return args[0] * args[1]
}

// task 8.8a
func Sub_(args []int) int {
	return max(args[0]-args[1], 0)
}

var One = S(N, Z)

// GE returns takes a and b and returns 1 if a >= b or 0 otherwise
var GE = S(Sub_, One, S(Sub_, U(2, 2), U(2, 1)))

// task 8.8b
func Div_(args []int) int {
	return args[0] / args[1]
}

// task 8.8b
func Mod_(args []int) int {
	return args[0] % args[1]
}

// Not returns 1, if x = 0 or 0 otherwise
var Not = S(Sub_, One, U(1, 1))

// IsMultipleOf returns 1, if a is multiple of b
// or 0, otherwise
var IsMultipleOf = S(Not, S(Mod_, U(2, 1), U(2, 2)))

func IsMultipleOf_(args []int) int {
	if args[0]%args[1] == 0 {
		return 1
	}
	return 0
}

// Eq returns 1, if a == b
// or 0, otherwise
var Eq = S(Mult_, S(GE, U(2, 1), U(2, 2)), S(GE, U(2, 2), U(2, 1)))

func Eq_(args []int) int {
	if args[0] == args[1] {
		return 1
	}
	return 0
}

// Gt returns 1, if a > b,
// and 0, otherwise
var Gt = S(Mult_, S(GE, U(2, 1), U(2, 2)), S(Not, S(Eq, U(2, 1), U(2, 2))))

func GT_(args []int) int {
	if args[0] > args[1] {
		return 1
	}
	return 0
}

// Le returns 1, if a <= b,
// and 0, otherwise
var Le = S(Not, S(Gt, U(2, 1), U(2, 2)))

func Le_(args []int) int {
	if args[0] <= args[1] {
		return 1
	}
	return 0
}

var Two = S(N, One)

// Neq returns 1 if a != b
// or 0, otherwise
var Neq = S(Not, S(Eq, U(2, 1), U(2, 2)))

var IsPrimeG = S(
	Mult_,
	U(3, 3),
	S(Not, S(IsMultipleOf, U(3, 1), S(Add, Two, U(3, 2)))),
)

var IsPrimeR = S(
	R(One, IsPrimeG),
	U(1, 1),
	S(Sub_, U(1, 1), Two),
)

// IsPrime returns 1, if n is prime
// and 0, otherwise
var IsPrime = S(
	Mult_,
	IsPrimeR,
	S(Mult_,
		S(Neq, U(1, 1), One),
		S(Neq, U(1, 1), Z),
	),
)

func IsPrime_(args []int) int {
	for d := 2; d*d <= args[0]; d++ {
		if args[0]%d == 0 {
			return 0
		}
	}
	return 1
}

// UpdIfZero returns a, if a > 0
// and b, if a == 0
var UpdIfZero = S(
	Add,

	S(Mult_,
		U(2, 1),
		S(Gt, U(2, 1), Z),
	),

	S(Mult_,
		U(2, 2),
		S(Eq, U(2, 1), Z),
	),
)

func UpdIfZero_(args []int) int {
	if args[0] > 0 {
		return args[0]
	}
	return args[1]
}

// UpdIfNonZero returns b, if b != 0
// and a, if b == 0
var UpdIfNonZero = S(
	Add,

	S(Mult_,
		S(Neq, U(2, 2), Z),
		U(2, 2),
	),

	S(Mult_,
		S(Eq, U(2, 2), Z),
		U(2, 1),
	),
)

func UpdIfNonZero_(args []int) int {
	if args[1] != 0 {
		return args[1]
	}
	return args[0]
}

var FirstPrimeG = S(
	UpdIfZero,
	U(3, 3),
	S(Mult_,
		S(Add, U(3, 2), U(3, 1)),
		S(IsPrime, S(Add, U(3, 2), U(3, 1)))),
)

var FirstPrimeR = R(Z, FirstPrimeG)

// FirstPrime returns the first prime number in [a, b]
var FirstPrime = S(
	FirstPrimeR,
	U(2, 1),
	S(Add, S(Sub_, U(2, 2), U(2, 1)), One),
)

var NthPrimeG = S(
	FirstPrime,
	S(Add, U(3, 3), One),
	S(Mult_, Two, S(Add, U(3, 3), One)),
)

// NthPrime returns n-th prime, assuming n >= 1
var NthPrime = S(
	R(One, NthPrimeG),
	U(1, 1),
	U(1, 1),
)

func NthPrime_(args []int) int {
	return first1000000Primes[args[0]-1]
}

var PowG = S(Mult_, U(3, 3), U(3, 1))

var Pow = S(R(One, PowG), U(2, 1), U(2, 2))

func Pow_(args []int) int {
	a := args[0]
	n := args[1]

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

var PlogG = S(UpdIfNonZero,
	U(4, 4),
	S(Mult_,
		U(4, 3),
		S(IsMultipleOf,
			U(4, 2),
			S(Pow,
				U(4, 1),
				U(4, 3),
			),
		),
	),
)

// Plog returns max{p | k ⫶ n^p}, assuming n > 1
var Plog = S(R(Z, PlogG),
	U(2, 1),
	U(2, 2),
	S(Add, U(2, 2), One),
)

func Plog_(args []int) int {
	n := args[0]
	k := args[1]

	p := 0
	for k != 0 && k%n == 0 {
		k /= n
		p++
	}
	return p
}

var SublistG = S(Mult_,
	U(4, 4),
	S(Pow_,
		S(NthPrime_, S(Add_, U(4, 3), One)),
		S(Plog_,
			S(NthPrime_, S(Add_, U(4, 2), S(Add_, U(4, 3), One))),
			U(4, 1),
		),
	),
)

var Sublist = S(R(One, SublistG),
	U(3, 1),
	U(3, 2),
	U(3, 3),
)

//func Sublist(L, I, d int) int {
//	f := constantly(1)
//
//	g := func(x ...int) int {
//		res, yi, args := unpack(x...)
//		l := U(0)(args...)
//		i := U(1)(args...)
//		return res * (pow_(nthPrime_(yi+1), plog_(nthPrime_(i+yi+1), l)))
//	}
//
//	return R(f, g)(L, I, d)
//}
