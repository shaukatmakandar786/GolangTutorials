package main

import "fmt"

func main() {

	//Type inference

	var name = "shaukat Makandar" //Type declaration is optional here

	fmt.Printf("Variable 'name' is of type %T\n", name)

	//=================================================

	// Multiple variable declarations with inferred types

	var firstName, lastName, age, salary = "Shaukat", "Makandar", 25, 25000.00

	fmt.Printf("firstName:%T,LastName:%T,age:%T,salary:%T\n", firstName, lastName, age, salary)
}
