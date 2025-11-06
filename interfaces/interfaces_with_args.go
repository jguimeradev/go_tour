package main

import "fmt"

type Operations interface {
	Add() float64
	Sub() float64
	Mult(float64) float64
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Add() float64 {
	return float64(v.X + v.Y)
}

func (v *Vertex) Sub() float64 {
	return float64(v.X + v.Y)
}

func (v *Vertex) Mult(m float64) float64 {
	return float64((v.X + v.Y) * m)
}

func main() {
	var ops Operations

	v := Vertex{3, 5}

	ops = &v // v is a pointer

	fmt.Println(ops.Add(), ops.Sub(), ops.Mult(5))
}
