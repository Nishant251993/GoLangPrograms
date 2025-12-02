package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Declaring variables
	grades := 42
	message := "hello world"
	isCheck := true
	amount := 46.0

	// Using format specifiers
	fmt.Printf("\n Variable grades is equal to %v and it is of type %T", grades, grades)
	fmt.Printf("\n Variable message is equal to %v and it is of type %T", message, message)
	fmt.Printf("\n Variable isCheck is equal to %v and it is of type %T", isCheck, isCheck)
	fmt.Printf("\n Variable amount is equal to %v and it is of type %T", amount, amount)

	// Using reflect.TypeOf() function
	fmt.Printf("\n Type of grades: %s", reflect.TypeOf(grades))
	fmt.Printf("\n Type of message: %s", reflect.TypeOf(message))
	fmt.Printf("\n Type of isCheck: %s", reflect.TypeOf(isCheck))
	fmt.Printf("\n Type of amount: %s", reflect.TypeOf(amount))
}
