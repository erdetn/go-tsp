// Erdet Nasufi erdet.nasufi@gmail.com Copyrights 2022 (C) //

package tsp

type IirFilter struct {
	buffer       []T
	coefficients []T
	length       int
	sections     int
}

func NewIir(coefficients []T) *IirFilter {
	if len(coefficients)%2 == 0 {
		return nil
	}

	len := int((len(coefficients) - 1) / 2)
	sections := int(len / 2)

	new_buffer := make([]T, len)

	return &IirFilter{
		buffer:       new_buffer,
		coefficients: coefficients,
		length:       len,
		sections:     sections,
	}
}

// Implements cascaded direct form II second order sections.
func (f *IirFilter) Filter(input T) T {
	n := f.sections
	y := T(0)

	h1, h2 := T(0), T(0)

	y = input * (f.coefficients[0])
	for i := 0; i < n; i++ {
		h1 = f.buffer[2*i]
		h2 = f.buffer[2*i+1]

		y = y - h1*(f.coefficients[4*i+1])
		h3 := T(y - h2*(f.coefficients[4*i+2]))

		y = h3 + h1*(f.coefficients[4*i+3])
		y += h2 * (f.coefficients[4*i+4])

		f.buffer[2*i+1] = f.buffer[2*i]
		f.buffer[2*i] = h3
	}

	return y
}

func (f *IirFilter) FilterArray(samples []T) []T {
	var y []T = make([]T, len(samples))
	for i, sample := range samples {
		y[i] = f.Filter(sample)
	}
	return y
}

func (f *IirFilter) Coefficients() []T {
	return f.coefficients
}

func (f *IirFilter) History() []T {
	return f.buffer
}

func (f *IirFilter) Length() int {
	return f.length
}

func (f *IirFilter) Sections() int {
	return f.sections
}
