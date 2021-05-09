package main

import (
	"fmt"
//	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) sum() *float64 {
	a :=(v.X+v.Y)
	return &a
}

func (v *Vertex) sum2() {
	v.X=v.X+v.Y
}
func sum3(v Vertex) *float64 {
	a :=(v.X+v.Y)
	return &a
}
func main() {
	k1,k2,k3 := Vertex{3, 4},Vertex{3, 4},Vertex{3, 4}
	fmt.Println(*(k1.sum()))
	k2.sum2()
	fmt.Println(k2.X)
	fmt.Println(*(sum3(k3)))
}