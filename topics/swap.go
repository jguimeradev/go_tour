package main

import "fmt"

func swap(x int, y int) (int, int) {
	return y, x
}

func main() {

	x, y := 10, 5

	
	fmt.Println(swap(x, y))

}
