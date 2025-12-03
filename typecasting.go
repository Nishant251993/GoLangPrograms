package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Declaring an integer
	var num int = 42

	// Converting integer to float
	var floatNum float64 = float64(num)

	// Converting integer to string
	var strNum string = fmt.Sprintf("%d", num)

	// Printing results
	fmt.Printf("Integer: %d", num)
	fmt.Printf("\nConverted to Float: %f", floatNum)
	fmt.Printf("\nConverted to String: %s", strNum)
	fmt.Println()
	// Example of converting float to integer
	var anotherFloat float64 = 3.14
	var intFromFloat int = int(anotherFloat)

	fmt.Printf("Float: %f", anotherFloat)
	fmt.Printf("\nConverted to Integer: %d", intFromFloat)

	//convert int to sting Itoa
	var intnum int = 34
	var stringnum string = strconv.Itoa(intnum)
	fmt.Printf("\nInt to String convert : %q", stringnum)

	//convert string to int
	var svalue string = "200"
	i, err := strconv.Atoi(svalue)
	fmt.Println()
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", err, err)
}
