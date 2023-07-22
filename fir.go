// Erdet Nasufi erdet.nasufi@gmail.com Copyrights 2022 (C) //

package tsp

import "math"

type FirFilter struct {
	buffer       []T
	coefficients []T
	length       int
}

func NewFirFilter(coefficients []T) FirFilter {
	var new_buffer []T = make([]T, len(coefficients))

	return FirFilter{
		buffer:       new_buffer,
		coefficients: coefficients,
		length:       len(coefficients),
	}
}

// ---------------------------------------------------------------- //
//                                                                  //
//         x[0]         x[1]         x[2]           x[N]            //
//           |            |            |              |             //
//  input ---+--->[z-1]---+--->[z-2]---+--->[z-N]-----+             //
//           |            |            |              |             //
//         c[0]         c[1]         c[2]           c[N]            //
//           |            |            |              |             //
//           |            |            |              +-->+         //
//           |            |            |                  |         //
//           |            |            +----------------->+         //
//           |            |                               |         //
//           |            +------------------------------>+         //
//           |                                            |         //
//           +------------------------------------------->+         //
//                                                        |         //
//                                                        +---> y   //
//                                                                  //
// ---------------------------------------------------------------- //

func (f *FirFilter) Filter(input T) T {
	n := len(f.coefficients)
	y := T(0)

	for i := 1; i < n; i++ {
		y += (f.buffer[n-i]) * (f.coefficients[n-i])
		f.buffer[n-i] = f.buffer[n-i-1]
	}
	y += input * (f.coefficients[0])
	f.buffer[1] = input

	return y
}

func (f *FirFilter) FilterArray(samples []T) []T {
	var y []T = make([]T, len(samples))

	for i, sample := range samples {
		y[i] = f.Filter(sample)
	}
	return y
}

func (f *FirFilter) Coefficients() []T {
	return f.coefficients
}

func (f *FirFilter) History() []T {
	return f.buffer
}

func (f *FirFilter) Length() int {
	return f.length
}

func HanningCoefficients(l int) []T {
	var coefficients []T = make([]T, l)

	for i := 0; i < l; i++ {
		coefficients[i] = 0.5 * (1 - math.Cos(2.0*math.Pi*T(i)/(T(l-1))))
	}
	return coefficients
}

func HammingCoefficients(l int) []T {
	var coefficients []T = make([]T, l)

	for i := 0; i < l; i++ {
		coefficients[i] = 0.54 - 0.46*math.Cos(2.0*math.Pi*T(i)/T((l-1)))
	}
	return coefficients
}
