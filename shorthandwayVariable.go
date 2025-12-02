package main

import "fmt"

func main() {

	//for same data type
	var firstname, lastanme string = "Nishant", "Patel"
	fmt.Println(firstname)
	fmt.Println(lastanme)

	//for different data type
	var (
		s string = "Nishant"
		t int    = 50000
	)

	fmt.Println(s)
	fmt.Println(t)

	//short variable declaration
	shortv := "Short variable declaration"

	fmt.Println(shortv)
}
