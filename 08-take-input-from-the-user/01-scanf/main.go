//example to take input from the user
package main

import "fmt"

func main() {

	var name string
	var a int
	var b int
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hello", name)

	fmt.Println("enter value of a and b\n")
	fmt.Scanf("%d%d",&a,&b)
	fmt.Println("Addition of two number is",(a+b))
}


