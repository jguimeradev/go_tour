package main

import (
	"fmt"
	"math"
)

type User struct {
	name  string
	age   int
	admin bool
}

func (u *User) Create() {
	u.name = "James"
	u.age = 41
	u.admin = true
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Add(v Vertex) float64 {
	return v.X + v.Y
}

func method_interfaces() {
	u := User{}
	u.Create()
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Add(v))
	fmt.Println(u)
}

func main() {
	method_interfaces()
}
