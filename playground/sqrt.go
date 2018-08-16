package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var z float64 = 1.0
	for {
		d := z*z - x
		if d < 0.00001 && d > -0.00001 {
			break
		}
		z += (x - z*z) / (z * 2)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(25))
}
