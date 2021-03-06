
package main

import (
	"fmt"
	"math"
)

// Go does not have classes.
// A method is function with a special receiver argument.
// The receiver argument appears between the 'func' and the method name.

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	f := MyFloat(-215.55)
	fmt.Println(v.Abs())
	fmt.Println(f.Abs())
}
