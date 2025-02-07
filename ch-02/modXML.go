//In this subsection, you will learn how to modify and customize the generated XML output.
//The name of the utility, which will be presented in three parts, is modXML.go. Note that the
//data that is going to be converted into XML is hardcoded in the program for reasons of
//simplicity.
//This the first part of the code

package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func main() {

	//This is where the structure for the XML data is defined. However, there is additional
	//information regarding the name and the type of the XML elements. The XMLName field
	//provides the name of the XML record, which in this case will be employee.
	//A field with the tag ",comment" is a comment and it is formatted as such in the output. A
	//field with the tag ",attr" appears as an attribute to the provided field name (which is id

	//in this case) in the output. The "name>first" notation tells Go to embed the first tag
	//inside a tag called name.
	//Lastly, a field with the "omitempty" option is omitted from the output if it is empty. An
	//empty value is any of 0, false, nil pointer, or interface, and any array, slice, map, or string
	//with a length of zero.

	type Address struct {
		City, Country string
	}

	type Employee struct {
		XMLName   xml.Name `xml: "employee"`
		Id        int      `xml:"id,attr"`
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"`
		Initials  string   `xml:"name>initials"`
		Height    float32  `xml:"height,omitempty`
		Address
		Comment string `xml: ",comment"`
	}

	//The second part of modXML.go is as follows
	// Here we define and initialize an employee structure.
	r := &Employee{Id: 7, FirstName: "Mihalis", LastName: "Tsoukalos",
		Initials: "MIT"}
	r.Comment = "Technical Writer + DevOps"
	r.Address = Address{"SomeWhere 12", "12312, Greece"}

	//The last part of modXML.go is shown here
	output, err := xml.MarshalIndent(r, "  ", "		")
	if err != nil {
		fmt.Println("Error:", err)
	}
	output = []byte(xml.Header + string(output))
	os.Stdout.Write(output)
	os.Stdout.Write([]byte("\n"))
}
