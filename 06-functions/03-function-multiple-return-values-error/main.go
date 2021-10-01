package main

import (
	"errors"
	"fmt"
)

func arithmeticOperations(x, y int) (int, int, error) {

	if x < 0 || y < 0 {

		err := errors.New("Negative numbers are not allowed")
		return 0, 0, err
	}

	add := x + y
	mul := x * y

	return add, mul, nil
}
func main() {
	x := -2 //you can put possitive value and check output
	y := 5

	add, mul, err := arithmeticOperations(x, y)

	if err != nil {

		fmt.Println(err)
	} else {

		fmt.Printf("Addition if %d and %d is %d\n", x, y, add)
		fmt.Printf("Multiplicatio if %d and %d is %d\n", x, y, mul)
	}

}


/*

Negative numbers are not allowed


*/

