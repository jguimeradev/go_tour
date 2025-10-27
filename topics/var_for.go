package main

import "fmt"


func read_vars(){
	var a int = 1
	c := "hello"
	d := "world" //_ only works when declaring variables with 'var'
	fmt.Println(a,c+d)
}

func read_consts(){
	const s string = "constant"
	const d = 50
	fmt.Println(s+" hello " , d+3)
}

func for_loops(){
	i := 1
	for i <=3 {
		fmt.Println("i:", i)
		i++
	}

	for j:= 0; j < 3; j++{
		fmt.Println("j:", j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for{
		fmt.Println("loop")
		break
	}

	for n := range 6 {

		if n%2 == 0{
			continue
		} 
		fmt.Println(n)
	}

}


func main(){
	fmt.Println("Hello, Go")
	for_loops()
}



