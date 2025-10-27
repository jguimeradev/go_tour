package main

import (
	"fmt"
	"math"
)


func Sqrt(x float64) {

/* 
	execute the calculation 10 times and print each z along the way. See how close you get to the answer for various values of x (1, 2, 3, ...) and how quickly the guess improves.
	*/


    	for i := 0; i < 1000000; i++{
		z := float64(i)
		z -= (z*z - x) / (2*z)
		fmt.Println("i z:",i,z)
		if float64(z) == float64(i)	 {
			break
		} 
	}
	
}

func main() {

	
	const VAL = 640000
	Sqrt(VAL)
	fmt.Println("From math lib:", math.Sqrt(VAL))
}


