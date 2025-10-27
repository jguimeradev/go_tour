package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	split := strings.Fields(s)
	words := make(map[string]int)

	var initial int = 1

	for i := range split {

		_, ok := words[split[i]]

		if ok == true {
			words[split[i]]++
		} else {
			words[split[i]] = initial
		}

	}

	return words
}

func WordCountImproved(s string) map[string]int {
	words := make(map[string]int)

	for w := range strings.FieldsSeq(s) {
		words[w]++
	}

	// The words[w]++ idiom works because missing keys default to zero in Go maps.

	return words
}

func main() {
	words := WordCountImproved("Hello Hello Bye")

	for i, v := range words {
		fmt.Println(i, v)
	}
}
