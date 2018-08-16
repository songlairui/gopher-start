package main

import (
	"fmt"
)

type ErrNegativeSqrt struct {
	S string
}

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprint(e.S)
}

func Sqrt(x float64) (float64, error) {
	var z float64
	if x == 0 {
		return 0, nil
	}
	if x < 0 {
		return z, &ErrNegativeSqrt{"cannot Sqrt negative number: " + fmt.Sprint(x) + "\n"}
	}
	z = 1.0
	for {
		d := z*z - x
		if d < 0.00001 && d > -0.00001 {
			break
		}
		z += (x - z*z) / (z * 2)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
