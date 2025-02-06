//A rune is an int32 value, and therefore a Go type 
//that is used for representing single Unicode characters 
//can also have alternative meaning 

//A rune literal is a character in single quotes 
//You may also consider a rune literal as a rune constant

package main 

import (
	"fmt"
)


 
func main() {
	
	//The second part of the code
	//First you define a rune literal named r1. (Please note that the Euro sign does not belong to
	//the ASCII table of characters.) Then, you print r1 using various statements. Next, you print
	//its int32 value and its hexadecimal value. After that, you try printing it as a string. Finally,
	//you print it as a character, which is what gives you the same output as the one used in the
	//definition of r1.
	const r1 = 'Euro'
	
	fmt.Println("(int32) r1:", r1)
	fmt.Printf("(HEX) r1: %x\n", r1)
	fmt.Printf("(as a String) r1: %s\n", r1)
	fmt.Printf("(as a character) r1: %c\n", r1)
	
	//The third part of the code 
	//Here, you see that a byte slice is a collection of runes and that printing a byte slice with
	//fmt.Println() might not return what you expected. In order to convert a rune into a
	//character, you should use %c in a fmt.Printf() statement. In order to print a byte slice as
	//a string, you will need to use fmt.Printf() with %s.
	//fmt.Println("A string is a collection of runes: ", []bytes("Mihalis"))
	//aString := []byte("Mihalis")
	for x, y := range aString {
		fmt.Println(x, y)
		fmt.Printf("Char: %c\n", aString[x])
	}
	fmt.Printf("%s\n", aString)
}