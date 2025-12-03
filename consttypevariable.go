package main

import "fmt"

func main() {
	//UnTyped Constant variable, as Datatype not declared
	//const <variable name> = <value>
	//Typed Constant variable systax
	//const <variable name> <data type> = <value>
	const age1 = 10
	fmt.Printf("%v, %T", age1, age1)

	//Constant can not re-assigned later
	const name = "Nishant"
	name = "Patel"
	fmt.Printf("%v, %T", name, name)

	//Const variable declare and assign at the same time, we can not assign later
	//example 
	/*
		const name
		name = "ABC"
	*/

	//Short hand assignment not possible in const
	const fname := "Nishant"
	fmt.Printf("%v, %T", fname, fname)
}
