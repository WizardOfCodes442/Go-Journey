package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type myElement struct {
	Name    string
	Surname string
	Id      string
}

// The key-value store is stored in a native Go map because using a built-in Go structure is
// usually faster. The map variable is defined as a global variable, where its keys are string
// variables and its values are myElement variables.
var DATA = make(map[string]myElement)

// This code contains the implementation of two functions that support the functionality of
// the ADD and DELETE commands. Note that if the user tries to add a new element to the store
// without giving enough values to populate the myElement struct, the ADD function will
// fail. For this particular program, the missing fields of the myElement struct will be set to
// the empty string. However, if you try to add a key that already exists, you will get an error
// message instead of modifying the value of the existing key.
func ADD(k string, n myElement) bool {
	if k == "" {
		return false
	}
	if _, exists := DATA[k]; !exists {
		DATA[k] = n
		return true
	}
	return false
}

func DELETE(k string) bool {
	if _, exists := DATA[k]; exists {
		delete(DATA, k)
		return true
	}
	return false
}

// In this Go code segment, you can see the implementation of the functions that support the
// functionality of the LOOKUP and CHANGE commands. If you try to change a key that does not
// exist, the program will add that key to the store without generating any error messages. In
// this part, you can also see the implementation of the PRINT() function that prints the full
// contents of the key-value store.
func LOOKUP(k string) *myElement {
	if n, exists := DATA[k]; exists {
		return &n
	}
	return nil
}

func CHANGE(k string, n myElement) bool {
	DATA[k] = n
	return true
}

func PRINT() {
	for k, d := range DATA {
		fmt.Printf("Key: %s, Value: %+v\n", k, d)
	}
}

// In this part of keyValue.go, you read the input from the user. Firstly, the for loop makes
// sure that the program will keep running for as long as the user provides some input.
// Secondly, the program makes sure that the tokens slice has at least five elements, even
// though only the ADD command needs that number of elements. Thus, for an ADD operation
// to be complete and not have any missing values, you will need an input that looks like ADD
// aKey Field1 Field2 Field3.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Key-Value Store is running. Type 'STOP' to exit.")

	for {
		fmt.Print("> ")      // Prompt for user input
		if !scanner.Scan() { // Handle EOF
			break
		}

		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}

		tokens := strings.Fields(text)
		if len(tokens) == 0 {
			continue
		}

		// Extend tokens to at least 5 elements for ADD and CHANGE commands
		for len(tokens) < 5 {
			tokens = append(tokens, "")
		}

		// The last part of the keyValueStore program
		switch tokens[0] {
		case "PRINT":
			PRINT()
		case "STOP":
			fmt.Println("Exiting program...")
			return
		case "DELETE":
			if !DELETE(tokens[1]) {
				fmt.Println("Delete operation failed!")
			}
		case "ADD":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !ADD(tokens[1], n) {
				fmt.Println("Add operation failed!")
			}
		case "LOOKUP":
			n := LOOKUP(tokens[1])
			if n != nil {
				fmt.Printf("Found: %+v\n", *n)
			} else {
				fmt.Println("Key not found!")
			}
		case "CHANGE":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !CHANGE(tokens[1], n) {
				fmt.Println("Update operation failed!")
			}
		default:
			fmt.Println("Unknown command - please try again")
		}
	}
}
