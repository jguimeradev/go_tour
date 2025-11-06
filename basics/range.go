package main

import "fmt"

func iterateRange() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	fmt.Println(pow)

	for i, v := range pow {
		fmt.Printf("%d^2 = %d\n", i, v)
	}
}

func iterateRangeContinued() {

	pow := make([]int, 10)

	for i := range pow { //index
		pow[i] = 1 << uint(i) 
		fmt.Printf("%08b\n", pow[i])
	}

	for _, v := range pow{ //skip index, print value
		fmt.Printf("%d\n", v)
	}  

}

func main() {
	fmt.Println("shit")
	iterateRangeContinued()

}
