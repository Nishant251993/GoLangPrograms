package main

import "fmt"

func main() {
	city := "London"
	{
		country := "UK"
		fmt.Println(country)
		fmt.Println(city)
	}
	fmt.Println(city)
	//fmt.Println(country) // will throw error because variable is block scope
	//  and outer blocks can not access variables declared within inner block
}
