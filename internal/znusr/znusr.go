package znusr

import "slices"

func Z(_ []int) int {
	return 0
}

func N(x []int) int {
	return x[0] + 1
}

func U(_, k int) func([]int) int {
	return func(x []int) int {
		return x[k-1]
	}
}

func S(g func([]int) int, fs ...func([]int) int) func(x []int) int {
	return func(x []int) int {
		args := make([]int, len(fs))

		for i, f := range fs {
			args[i] = f(x)
		}

		return g(args)
	}
}

func R(f func([]int) int, g func([]int) int) func(x []int) int {
	return func(x []int) int {
		n := len(x)
		y := x[n-1]
		vec := x[:n-1]

		res := f(vec)

		for yi := range y {
			args := append(slices.Clone(vec), yi, res)
			res = g(args)
		}

		return res
	}
}
