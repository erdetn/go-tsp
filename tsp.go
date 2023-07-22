// Erdet Nasufi erdet.nasufi@gmail.com Copyrights 2022 (C) //

package tsp

import "fmt"

func Version() string {
	return "0.1.0"
}

func Stringify(x []T, name string) string {
	str := ""
	for i := 0; i < len(x); i++ {
		str += fmt.Sprintf("%s[%d] = %.4f\n", name, i, x[i])
	}
	return str
}
