package main

import "fmt"

type Ops interface {
	Add() float64
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func (f MyFloat) Add() float64 {
	return float64(f + 10)
}

func (v *Vertex) Add() float64 {
	return float64(v.X + v.Y)
}

func main() {
	// var a Ops
	v := Vertex{2, 3}
	f := MyFloat(100)

	var a Ops

	a = f
	fmt.Println(a.Add()) // 110

	a = &v
	fmt.Println(a.Add()) // 5
}
