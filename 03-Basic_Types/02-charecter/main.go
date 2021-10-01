package main

import "fmt"

func main() {

	var myByte1 byte = 'a'
	var myRune1 rune = '♥'

	fmt.Printf("%c is of type of %T and unicode value is %d\n", myByte1, myByte1, myByte1)
	fmt.Printf("%c is of type of %T and unicode value is %d\n", myRune1, myRune1, myRune1)

	//byte alias for uint8
	//rune alias for int32

}

/*
a is of type of uint8 and unicode value is 97
♥ is of type of int32 and unicode value is 9829

*/
