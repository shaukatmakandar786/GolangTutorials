package main

import "fmt"

func main() {

	//If Statement
	var x = 25
	if x%5 == 0 {

		fmt.Printf("%d is multiple of 5\n", x)
	}

	// Parentheses are Optional
	var y = -1

	if y < 0 {
		fmt.Printf("%d is negative number\n", y)
	}

	// If with multiple condition
	var age = 21
	if age >= 17 && age <= 30 {
		fmt.Println("My Age is between 17 and 30")
	}

	// If with a short statement

	if no := 10; no%2 == 0 {

		fmt.Printf("%d  is even number\n", no)
	}

}


/*

25 is multiple of 5
-1 is negative number
My Age is between 17 and 30
10  is even number


*/
