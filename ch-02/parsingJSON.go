// This subsection will tell you how to read and store unstructured JSON data. The critical
// thing to remember is that unstructured JSON data is put into Go maps instead of Go
// structures – this will be illustrated in parsingJSON.go, which will be presented in four
// parts.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//The second part
//The presented code reads the command-line arguments of the program and gets the first
//one, which is the JSON file that is going to be read.

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a filename")
		return
	}
	filename := arguments[1]

	//The third part of parsJSON.go is as follows
	//The ioutil.ReadFile() function allows you to read a file all at once, which is what we
	//want here.
	//In this part, there is also a definition of a map named parsedData that will hold the
	//contents of the JSON file that was read. Each map key, which is a string, corresponds to a
	//JSON property. The value of each map key is of the type interface{}, which can be of
	//any type – this means that the value of a map key can also be another map.
	//The json.Unmarshal() function is used for putting the contents of the file into the
	//parsedData map.
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	var parsedData map[string]interface{}
	json.Unmarshal([]byte(fileData), &parsedData)

	//The last part of parsingJSON.go contains the following code
	for key, value := range parsedData {
		fmt.Println("key:", key, "value:", value)
	}
}
