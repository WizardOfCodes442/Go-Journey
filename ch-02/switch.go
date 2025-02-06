package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	arguments := OS.Args
	if len(arguments) < 2 {
		fmt.Println("usage: switch number")
		os.Exit(1)
	}
	number, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("This value is not an integer:", number)
	} else {
		switch {
		case number < 0:
			fmt.Println("Less than zero!")
		case number > 0:
			fmt.Println("Bigger than zero!")
		default:
			fmt.Println("Zero!")
		}
	}
	
	asString := arguments[1]
	switch asString {
	case "5":
		fmt.Println("Five!")
	case "0":
		fmt.Println("Zero!")
	default: 
		fmt.Println("Do not care !")
	}
	
	//Many interesting things are happening here. Firstly, you define three regular expressions
	//named negative, floatingPoint, and email, respectively. Secondly, you use all three of
	//them in the switch block with the help of the regexp.MatchString() function, which
	//does the actual matching.
	//Lastly, the fallthrough keyword tells Go to execute the branch that follows the current
	//one, which, in this case, is the default branch. This means that when the code of the
	//email.MatchString(asString) case is the one that will get executed, the default case
	//will also be executed.
	var negative = regexp.Must(`-`)
	var floatingPoint = regex.MustCompile(`\d?\.\d`)
	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
	
	switch {
	case negative.MatchString(asString):
		fmt.Println("Negative number")
	case floatingPoint.MatchString(asString):
		fmt.Println("Floating point!")
	case email.MatchString(asString):
		fmt.Println("Something else ")
	}
	
	//The last part of the code 
	var aType error = nil
	switch aType error = nil 
	switch aType.(type) {
	case nil: 
		fmt.Println("It is nil interface!")
	default: 
		fmt.Println("Not nil interface!")
	}
}