package main

import "fmt"


func pointers_basics(){
	i := 42 
	var p *int
	p = &i

	j := 44
	var n *int
	n = &j

	fmt.Println(*p)
	fmt.Println(p) //address 

	//modify value of the pointer
	*p = 21
	fmt.Println(*p)

	//addition
	*p++
	fmt.Println(*p)

	//it works
	fmt.Println(*p+*n)

	//fmt.Println(p++) //fails

}

func testingpointers(x *int){

	fmt.Printf("%d", *x)
	*x++
}

func references(j *int){
	*j++
}

func main() {

	i := 10;
	var p *int
	p = &i
	testingpointers(p)

	fmt.Printf("%d", *p) //11
	
	j := 5
	references(&j)
	fmt.Printf("%d", j) //6
	
}
