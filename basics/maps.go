package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func simpleMaps() {
	m := make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		40.68, -74.30,
	}

	fmt.Println(m["Bell Labs"])
}

func mapLiterals() {
	m := map[string]Vertex{
		"Google": {
			40.68, -75.43,
		},
		"Apple": {
			40.68, -75.43,
		},
	}

	fmt.Println(m)
}

func mapOperations() {
	m := make(map[string]int)
	m["Answer"] = 54
	fmt.Println(m["Answer"])
	m["Answer"] = 55
	fmt.Println(m["Answer"])
	m["Answer"] = 56
	fmt.Println(m["Answer"])

	/*
	   If key is in m, ok is true. If not, ok is false.

	   If key is not in the map, then v is the zero value for the map's element type.
	*/

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	delete(m, "Answer")
	fmt.Println(m["Answer"]) // 0

	v, ok = m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func main() {
	mapOperations()
}
