// In this subsection, you will learn how to read an XML file from disk and store it into a Go
// structure. The name of the program is readXML.go and it will be presented in three parts.
// The first part of readXML.go is as follows:
package main

import (
	"encoding/xml"
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

// The second part of readXML.go is the following
// The presented process is very similar to the way that you read a JSON file from disk.
func loadfromXML(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	decodeXML := xml.NewDecoder(in)
	err = decodeXML.Decode(key)
	if err != nil {
		return err
	}
	in.Close()
	return nil
}

// The last part is as follow
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a filename!")
		return
	}
	filename := arguments[1]
	var myRecord Record
	err := loadfromXML(filename, &myRecord)

	if err == nil {
		fmt.Println("XML:", myRecord)
	} else {
		fmt.Println(err)
	}
}
