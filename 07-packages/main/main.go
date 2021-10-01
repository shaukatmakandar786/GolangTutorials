//Package declaration
package main

//Importing Packages
import (
	"fmt"
	"math"
	"time"
	//"math/rand"
)

func main() {

	//Finding the Max of two numbers

	fmt.Println(math.Max(10, 20))

	//Calculate the Square root of a number

	fmt.Println(math.Sqrt(225))

	//Printing the value of PI

	fmt.Printf("%.2f\n", math.Pi)

	//Printing the current time

	epoch := time.Now()
	fmt.Println(epoch.Format("01-02-2006"))

}
/*

20
15
3.14
11-10-2009

*/
