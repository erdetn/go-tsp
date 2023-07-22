// Erdet Nasufi erdet.nasufi@gmail.com Copyrights 2022 (C) //

package tsp

type EchoParameters struct {
	delay       uint32
	attenuation T
}

func Echo(x []T, params []EchoParameters) []T {
	y := make([]T, 0)
	y0 := T(0)
	lx, lp := uint32(len(x)), len(params)

	for i := uint32(0); i < lx; i++ {
		y0 = x[i]
		for p := 0; p < lp; p++ {
			if i >= params[p].delay {
				y0 += T(params[p].attenuation * (x[i-params[p].delay]))
			}
		}
		y = append(y, y0)
	}
	return y
}
