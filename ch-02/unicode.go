//The unicode standard package contains various handy functions 
//One of them, which
//is called unicode.IsPrint(), can help you to identify the parts of a string that are
//printable using runes. This technique will be illustrated in the Go code of unicode.go,
//which will be presented in two parts. The first part of unicode.go is as follows:

package main 

import {
	"fmt"
	"unicode"
}

func main() {
	const sL = "\x99\x00ab\x50\x00\x23\x50\x29\x9c"
	
	//The second code segment of unicode.go is shown like below 
	for i := 0; i < len(sL); i++ {
		if unmicode.IsPrint(rune(sL[i])) {
			fmt.Printf("%c\n", sL[i])
		} else {
			fmt.Println("Not Printable!")
		}
	} 
}