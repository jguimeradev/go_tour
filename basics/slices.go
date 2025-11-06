package main

import "fmt"

func simpleSlices() {

	//A slice is formed by specifying two indices, a low and high bound, separated by a colon:
	// A slice does not store any data, it just describes a section of an underlying array.

	primes := [6]int{1, 2, 3, 4, 5, 6}

	var s []int = primes[1:4]

	s[0] = 5 //changes the array primes

	fmt.Println(primes) // [1 5 3 4 5 6]

	q := []int{2, 3, 4} //slice literal (without specific length as the arrays)

	fmt.Println(q)

}

func sliceDefaults() {

	a := []int{1, 2, 3, 4, 5}
	a = a[:2]
	fmt.Println(a) //[1 2]
	a = a[1:]      //[2]
	fmt.Println(a)

	var s []int //zero value of slice is nil

	printSlice(s) //len=0 cap=0 []

	if s == nil {
		fmt.Println("nil, madafaka!")
	}

}

func lengthCapacity() {

	/*
	   The length of a slice is the number of elements it contains.
	   The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	*/

	s := []int{1, 2, 3, 4, 5}
	printSlice(s) //len = 5, cap = 5

	s = s[1:4]
	printSlice(s) //len:3, cap:4 (from 1 to 5) -> points to 1st s

	s = s[2:]
	printSlice(s) //len:1, cap:2 (from 2 to 4) -> points to 2dn s
}

func makeSlices() {
	//Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
	//The make function allocates a zeroed array and returns a slice that refers to that array:

	a := make([]int, 5) //len=5 cap=5 [0 0 0 0 0]
	printSlice(a)

	b := make([]int, 0, 5) //len=0 cap=5 []
	printSlice(b)

	c := b[:2] //len=2 cap=5 [0 0] -> points to b
	printSlice(c)

	d := c[2:5] //len=3 cap=3 [0 0 0] -> points to b
	printSlice(d)

}

func appendSlices(){

	s := make([]int, 0, 5)  //len=0 cap=5 []
	printSlice(s)
	
	s = append(s,0) //len=1 cap=5 [0]
	printSlice(s)
	
	s = append(s, 1, 2, 3, 4) //len=5 cap=5 [0 1 2 3 4]
	printSlice(s)

	s = append(s, 5, 6, 7, 8) //len=9 cap=10 [0 1 2 3 4 5 6 7 8]
	printSlice(s)
}


func main() {
	appendSlices()
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
