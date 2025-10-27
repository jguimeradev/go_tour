package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func printAdd() {
	var x int = 10
	y := 5
	fmt.Println(add(x, y))
}

var (
	ToBe   bool    = true
	MaxInT uint64  = 1<<64 - 1
	fl32   float32 = 3.2
	fl64   float64 = 69.69
	num    int     = 5
)

func printVars() {
	const Pi = 3.14

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInT, MaxInT)
	fmt.Printf("Type: %T Value: %v\n", fl32, fl32)
	fmt.Printf("Type: %T Value: %v\n", fl64, fl64)
	fmt.Printf("Type: %T Value: %v\n", float64(num), float64(num)) //casting
	fmt.Printf("Type: %T Value: %v\n", Pi, Pi)
}

func main() {

	printVars()

}
