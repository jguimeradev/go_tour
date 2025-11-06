package main

import "fmt"

type Ops interface {
	Add() int
}

type Proc struct {
	S int
}

func (p Proc) Add() int {
	return p.S + 10
}

func main() {
	var s Ops = Proc{4}

	fmt.Println(s.Add())
}
