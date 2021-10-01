package main

import "fmt"

func arithmeticOperations(x, y int) (add, mul int) {

	add = x + y
	mul = x * y

	return add, mul
}
func main() {
	x := 10
	y := 5

	add, mul := arithmeticOperations(x, y)

	fmt.Printf("Addition if %d and %d is %d\n", x, y, add)
	fmt.Printf("Multiplicatio if %d and %d is %d\n", x, y, mul)

}
/*

Addition if 10 and 5 is 15
Multiplicatio if 10 and 5 is 50

*/
