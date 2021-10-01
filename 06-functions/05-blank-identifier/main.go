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

	add, _ := arithmeticOperations(x, y)

	fmt.Printf("Addition if %d and %d is %d\n", x, y, add)
	//fmt.Printf("Multiplicatio if %d and %d is %d\n", x, y, _)

}


/*

Addition if 10 and 5 is 15

Note: If I try to execute line number 19 then i will get ( ./prog.go:19:12: cannot use _ as value ) compile time error.

*/
