#Constants:

Constants are declared like variables, but with the const keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the := syntax.

Example:

package main

import "fmt"

const Pi= 3.14

func main() {
	const World= "Shaukat"
	const b int=20
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day",b)

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

output:

Hello Shaukat
Happy 3.14 Day 20
Go rules? true
