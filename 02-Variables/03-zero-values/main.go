package main

import "fmt"

func main() {

	var (
		firstName, lastName string
		age                 int
		salary              float64
		isProgrammer        bool
	)

	fmt.Printf("FirstName: %s,lastName:%s,age:%d,salary:%f,isProgrammer:%t\n", firstName, lastName, age, salary, isProgrammer)
}

/*
output:

FirstName: ,lastName:,age:0,salary:0.000000,isProgrammer:false

*/
