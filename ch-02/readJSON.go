//Reading JSON data
//In this section, you will learn how to read a JSON record from disk using the code of
//readJSON.go, which will be presented in three parts.

package main

import (
	"encoding/json"
	"fmt"
	"os"
) //we define the structure variables that are going to keep the JSON data.
type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

// The second part is as follows
// Here, we define a new function named loadFromJSON() that is used for decoding the data
// of a JSON file according to a data structure that is given as the second argument to it. We
// first call the json.NewDecoder() function to create a new JSON decode variable that is
// associated with a file, and then we call the Decode() function for actually decoding the
// contents of the file and putting them into the desired variable.
func loadFromJSON(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}
	in.Close()
	return nil
}

// The last part of readJSON.go is as follow
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a filename!")
		return
	}

	filename := arguments[1]
	var myRecord Record
	err := loadFromJSON(filename, &myRecord)
	if err == nil {
		fmt.Println(myRecord)
	} else {
		fmt.Println(err)
	}
}
