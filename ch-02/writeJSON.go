// In this subsection, you will learn how to write JSON data. The utility that will be presented
// in three parts is called writeJSON.go and it will write to standard output (os.Stdout),
// which means that it will write on the terminal screen.
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

// The second part of the code
// The saveToJSON() function does all the work for us as it creates a JSON encoder variable
// named encodeJSON, which is associated with a filename, which is where the data is going
// to be put. The call to Encode() is what saves the data into the desired file after encoding it.

func saveToJSON(filename *os.File, key interface{}) {
	encodeJSON := json.NewEncoder(filename)
	err := encodeJSON.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// The last part of writeJSON.go is as follows
func main() {
	myRecord := Record{
		Name:    "Mihalis",
		Surname: "Tsoukalos",
		Tel: []Telephone{
			{Mobile: true, Number: "1234-567"},
			{Mobile: true, Number: "1234-abcd"},
			{Mobile: false, Number: "abcc-567"},
		},
	}
	saveToJSON(os.Stdout, myRecord)
}
