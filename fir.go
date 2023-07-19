// Erdet Nasufi erdet.nasufi@gmail.com Copyrights 2022 (C) //

package tsp

import "math"

type FirFilter struct {
	buffer       []float64
	coefficients []float64
	length       int
}

func NewFirFilter(coefficients []float64) FirFilter {
	var new_buffer []float64 = make([]float64, len(coefficients))

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

func (f *FirFilter) Filter(input float64) float64 {
	n := len(f.coefficients)
	y := float64(0)

	for i := 1; i < n; i++ {
		y += (f.buffer[n-i]) * (f.coefficients[n-i])
		f.buffer[n-i] = f.buffer[n-i-1]
	}
	y += input * (f.coefficients[0])
	f.buffer[1] = input

	return y
}

func (f *FirFilter) FilterArray(samples []float64) []float64 {
	var y []float64 = make([]float64, len(samples))

	for i, sample := range samples {
		y[i] = f.Filter(sample)
	}
	return y
}

func (f *FirFilter) Coefficients() []float64 {
	return f.coefficients
}

func (f *FirFilter) History() []float64 {
	return f.buffer
}

func (f *FirFilter) Length() int {
	return f.length
}

func HanningCoefficients(l int) []float64 {
	var coefficients []float64 = make([]float64, l)

	for i := 0; i < l; i++ {
		coefficients[i] = 0.5 * (1 - math.Cos(2.0*math.Pi*float64(i)/(float64(l-1))))
	}
	return coefficients
}

func HammingCoefficients(l int) []float64 {
	var coefficients []float64 = make([]float64, l)

	for i := 0; i < l; i++ {
		coefficients[i] = 0.54 - 0.46*math.Cos(2.0*math.Pi*float64(i)/float64((l-1)))
	}
	return coefficients
}
