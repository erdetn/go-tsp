// Erdet Nasufi erdet.nasufi@gmail.com Copyrights 2022 (C) //

package tsp

import (
	"math"
)

type Z = complex128

func Dft(x []T) []Z {
	m := len(x)
	z := make([]Z, m)
	x_re, x_im, theta := T(0), T(0), T(0)
	for k := 0; k < m; k++ {
		x_re, x_im = T(0), T(0)
		for n := 0; n < m; n++ {
			theta = T(T(n) * T(k) * (2.0 * math.Pi / T(m)))
			x_re += T(x[n]) * math.Cos(theta)
			x_im += -1.0 * T(x[n]) * math.Sin(theta)
		}
		z[k] = Z(complex(x_re, x_im))
	}
	return z
}

func DftPower(z []Z) []T {
	l := len(z)
	pz := make([]T, l)

	for k := 0; k < l; k++ {
		pz[k] = T(math.Pow(real(z[k]), 2) + math.Pow(imag(z[k]), 2))
	}

	return pz
}

func Idft(z []Z) []T {
	nn := len(z)
	x, xk, theta := make([]T, nn), T(0.0), T(0.0)

	for k := 0; k < nn; k++ {
		xk = 0.0
		for n := 0; n < nn; n++ {
			theta = (2.0 * math.Pi * T(k) * T(n)) / T(nn)
			xk += T(real(z[n]))*math.Cos(theta) + T(imag(z[n])*math.Sin(theta))
		}
		xk = T(xk / T(nn))
		x[k] = xk
	}
	return x
}
