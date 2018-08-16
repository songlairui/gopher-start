package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	var d Vertex
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	a = f
	b := f
	fmt.Println(f.Abs())
	fmt.Println(a.Abs(), b.Abs())
	a = &v
	c := &v
	d = v
	fmt.Println(a.Abs(), c.Abs(), d.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
