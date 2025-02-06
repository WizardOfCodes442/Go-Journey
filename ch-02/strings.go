// A Go string is a read-only byte slice
// that can hold any type of bytes
// and can have arbirary length.
package main

import (
	"fmt"
)

// The second portion of Go
func main() {

	//Each \xAB sequence representsn a 	single character\
	//Calling len(sLiteral) will return the number of charcters
	//of sLiteral.
	//Using %x in fmt.Printf() will return AB part of a \xAB sequence
	const sLiteral = "\x99\x42\x32\x55\x50\x29\x9c"
	fmt.Println(sLiteral)
	fmt.Printf("x: %x\n", sLiteral)

	fmt.Printf("sLiteral length: %d\n", len(sLiteral))

	//third part of code
	//you can access a string literal as if it is a slice
	//Using  %q in fmt.Printf() with a string argument
	//will print a double-quoted string that is safely escaped in go syntax
	//in order to print a string literal as a string you will need to
	//call fmt.Printf() with %s

	for i := 0; i < len(sLiteral); i++ {
		fmt.Printf("%x", sLiteral[i])
	}
	fmt.Println()

	fmt.Printf("q: %q\n", sLiteral)
	fmt.Printf("+q %+q\n", sLiteral)
	fmt.Printf("s: As a string: %s\n", sLiteral)

	//the fourth section of the code is as follow
	//Here a string is defined named s2 with three Unicode characters
	//Using fmt.Printf() with %#U will print the characters
	s2 := "pounds&euro symbol"
	for x, y := range s2 {
		fmt.Printf("%#U starts at byte position %d\n", y, x)
	}
	fmt.Printf("s2 length: %d\n", len(s2))

	//The last part of the code
	const s3 = "ab12AB"
	fmt.Println("s3:", s3)
	fmt.Println("x: % x\n", s3)

	fmt.Printf("s3 length: %d\n", len(s3))

	for i := 0; i < len(s3); i++ {
		fmt.Printf("%x ", s3[i])
	}
	fmt.Println()
}
