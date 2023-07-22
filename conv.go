// Erdet Nasufi erdet.nasufi@gmail.com Copyrights 2022 (C) //

package tsp

type T = float64

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Conv(x []T, h []T) []T {
	len_y := len(x) + len(h) - 1
	var y []T = make([]T, 0)

	x0, x1, h0, yn := 0, 0, 0, T(0)

	for i := 0; i < len_y; i++ {
		x0 = Max(0, i-len(h)+1)
		x1 = Min(i+1, len(x))
		h0 = Min(i, len(h)-1)
		yn = T(0)

		for j := x0; j < x1; j++ {
			yn += h[h0] * x[j]
			h0--
		}
		y = append(y, yn)
	}
	return y
}

func Autocorrelate(x []T) []T {
	return Conv(x, x)
}
