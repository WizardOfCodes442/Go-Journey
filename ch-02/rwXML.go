//Go has support for XML, which is a markup language similar to HTML but much more
//advanced than HTML.
//The developed utility, which is called rwXML.go, will read a JSON record from disk, make
//a change to it, convert it to XML, and print it on screen. Then it will convert the XML data
//into JSON. The related Go code will be presented in four parts.

//The first part

package main

import (
	"encoding/json"
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

// The second part of the code
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

// The third part of the code
// After we read the input file and convert it into JSON, we put its data into a Go structure.
// Then, we make a change to the data of that structure (myRecord). After that, we convert
// that data into the XML format using the MarshalIndent() function and add a header
// using xml.Header.
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
		fmt.Println("JSON", myRecord)
	} else {
		fmt.Println(err)
	}

	myRecord.Name = "Dimitris"

	xmlData, _ := xml.MarshalIndent(myRecord, "", "		")
	xmlData = []byte(xml.Header + string(xmlData))
	fmt.Println("\nxmlData:", string(xmlData))

	//The last part of the program
	//In this part of the program, we convert the XML data into JSON using Marshal() and
	//Unmarshal() and print it on screen.
	data := &Record{}
	err = xml.Unmarshal(xmlData, data)
	if nil != err {
		fmt.Println("Error marshalling to JSON", err)
		return
	}

	result, err := json.Marshal(data)
	if nil != err {
		fmt.Println("Error marshalling to JSON", err)
		return
	}
	_ = json.Unmarshal([]byte(result), &myRecord)
}
