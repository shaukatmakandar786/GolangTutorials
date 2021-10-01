package main

import "fmt"

func main() {

	//declaring variable

	var mystr string = "Hello"
	var myInt int = 100
	var myFloat float64 = 4.5
	fmt.Println(mystr, myInt, myFloat)

	//Multiple Declarations
  
	var (
		empId               int    = 5
		firstName, lastName string = "shaukat", "Makandar"
	)

	fmt.Println(empId, firstName, lastName)

	//Short variable declaration syntax

	name := "Shaukat Makandar"
	age, salary, isProgrammer := 25, 25000, true

	fmt.Println(name, age, salary, isProgrammer)
}
