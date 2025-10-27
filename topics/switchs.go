package main

import (
	"fmt"
	"runtime"
	"time"
)

func stdSwitch() {
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os { //allows both var declaration and calling at the same line to keep the local scope
	case "darwin":
			fmt.Println("mac0S.")
	case "linux":
			fmt.Println("Linux.")
	default:
			fmt.Println("%s.", os)
	}
}


func boolSwitch(){

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	default:
		fmt.Println("Your mom")
	}
}

func main() {
	boolSwitch()
}
