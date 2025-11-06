package main

import "fmt"

func ifstmt(b int) bool{
	
	
	if a := 10; a < b{ //declare and use within the same statement (local scope)
		return true
	}

	return false
}

func main(){


	fmt.Printf("Value: %v", ifstmt(20))
	fmt.Println("Value: ", ifstmt(20))
	
}