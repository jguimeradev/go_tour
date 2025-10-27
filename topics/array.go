package main

import "fmt"

func arrays(){
   //By default an array is zero-valued, which for ints means 0s.
	var a [10]int
	fmt.Println("var a [10]int:", a)

	b := [3]int{1,2,3}
	fmt.Println("b := [3]int{1,2,3}:", b)
	fmt.Println("length of b:", len(b))
	
	c := [...]int{1,2,3,4,5}
	fmt.Println("c = [...]int{1,2,3,4,5}", c)

	i := [...]int{100,3:400,500}
    fmt.Println("i := [...]int{100,3:400,500}", i) //index 1 and 2 are 0

	for i := range len(b){
		fmt.Println("i:", b[i])
	} 

	var matrix[2][3]int
	for i := range 2{
		for j:= range 3{
			matrix[i][j] = i + j
		}
	}
	fmt.Println("2d matrix with for:", matrix)

	matrix = [2][3]int{
		{1,2,3},
		{1,2,3},
	}
	fmt.Println("predefined 2d matrix:", matrix)
 
	

}


 


func main(){
	fmt.Println("Hello, Go")
	arrays()
}



