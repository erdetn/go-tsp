// Erdet Nasufi erdet.nasufi@gmail.com Copyrights 2022 (C) //

package tsp

import "math"

func NewSignal(t []int, signal_function func(int) T) []T {
	y := make([]T, 0)

	for _, it := range t {
		y = append(y, signal_function(it))
	}

	return y
}

func Range(args ...int) []int {
	min, max, step := 0, 0, 0

	if len(args) == 0 {
		return make([]int, 0)[:]
	}

	if len(args) == 1 {
		max = args[0]
		y := make([]int, max)
		for it := 0; it < max; it++ {
			y[it] = int(it)
		}
		return y[:]
	}

	if len(args) == 2 {
		min = args[0]
		max = args[1]
		if args[0] > args[1] {
			min = args[1]
			max = args[0]
		}
		l := max - min
		y := make([]int, l)
		for it := 0; it < l; it++ {
			y[it] = int(min + it)
		}

		return y[:]
	}

	if len(args) > 2 {
		min = args[0]
		step = args[1]
		max = args[2]

		if step < 0 {
			step = (-1) * step
		}
		if args[0] > args[2] {
			min = args[2]
			max = args[0]
		}
		l := int(max-min) / step
		y := make([]int, l)

		for it := 0; it < l; it++ {
			y[it] = int(min + it*step)
		}
		return y[:]
	}

	return make([]int, 0)[:]
}

func Operation(x1 []T, x2 []T, op func(T, T) T) []T {
	y := make([]T, 0)

	if len(x1) == len(x2) {
		for i := 0; i < len(x1); i++ {
			y = append(y, op(x1[i], x2[i]))
		}
		return y
	} else if len(x1) > len(x2) {
		for i := 0; i < len(x2); i++ {
			y = append(y, op(x1[i], x2[i]))
		}
		for j := len(x2); j < len(x1); j++ {
			y = append(y, op(T(0), x1[j]))
		}
		return y
	} else if len(x2) > len(x1) {
		for i := 0; i < len(x1); i++ {
			y = append(y, op(x1[i], x2[i]))
		}
		for j := len(x1); j < len(x2); j++ {
			y = append(y, op(0, x2[j]))
		}
		return y
	}
	return y
}

func Scale(samples []T, a T) []T {
	y := make([]T, len(samples))
	for i, x := range samples {
		y[i] = T(a * x)
	}
	return y
}

func Shift(samples []T, k int) []T {
	y := make([]T, 0)
	l := len(samples)

	if k > 0 {
		y = append(y, make([]T, k)...)
		y = append(y, samples...)
		return y
	} else {
		if k < 0 && k < l {
			k1 := (-1) * k
			y = samples[k1:]
			return y
		}
	}
	return y
}

func Mirror(samples []T) []T {
	y := make([]T, len(samples))

	for i := 0; i < len(samples); i++ {
		y[i] = samples[len(samples)-i-1]
	}
	return y
}

func Sum(samples []T) T {
	sum := T(0)
	for _, sample := range samples {
		sum += sample
	}
	return sum
}

func Product(samples []T) T {
	prod := T(1)
	for _, sample := range samples {
		prod *= T(sample)
	}
	return prod
}

func Energy(samples []T) T {
	se := T(0)
	for _, sample := range samples {
		se += math.Pow(T(sample), math.Abs(sample))
	}
	return se
}

func Power(samples []T, period int) T {
	pe := T(0)
	x := T(0)

	n := len(samples)
	if n > period {
		n = period
	}

	for i := 0; i < n; i++ {
		x = samples[i]
		pe += math.Pow(T(x), math.Abs(x))
	}

	pe = T(pe / T(n))

	return pe
}
