package main

import "fmt"

func main() {

	var name string
	fmt.Print("Enter your Name: ")
	fmt.Scanf("%s", &name) //Single Input
	fmt.Println("Hey there, ", name)

	//to run below code comment above code
	var namesecond string
	var is_muggle bool
	fmt.Print("Enter your name & are you a muggle:")
	fmt.Scanf("%s %t", &namesecond, &is_muggle) //Multiple Input
	fmt.Println(namesecond, is_muggle)
}
