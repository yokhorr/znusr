package znusr

import "slices"

// unpack returns the last, the last but one, and the rest of arguments
func unpack(x ...int) (int, int, []int) {
	m := len(x)
	return x[m-1], x[m-2], x[:m-2]
}

func Z(_ int) int {
	return 0
}

func N(x int) int {
	return x + 1
}

func U(k int) func(...int) int {
	return func(x ...int) int {
		return x[k]
	}
}

func S(g func(...int) int, fs ...func(...int) int) func(x ...int) int {
	return func(x ...int) int {
		args := make([]int, len(fs))

		for i, f := range fs {
			args[i] = f(x...)
		}

		return g(args...)
	}
}

func R(f func(...int) int, g func(...int) int) func(x ...int) int {
	return func(x ...int) int {
		n := len(x)
		y := x[n-1]
		vec := x[:n-1]

		res := f(vec...)

		for yi := range y {
			args := append(slices.Clone(vec), yi, res)
			res = g(args...)
		}

		return res
	}
}
