//In this subsection, you will see how to use the Marshal() and Unmarshal() methods in
//order to implement the functionality of readJSON.go and writeJSON.go. The Go code
//that illustrates the Marshal() and Unmarshal() functions can be found in mUJSON.go,
//and it will be presented in three parts.

package main

import (
	"encoding/json"
	"fmt"
)

//In this part of the program, we are defining two structures named Record and Telephone,
//which will be used for storing the data that will be put into a JSON record.
type Record struct {
	Name string
	Surname string
	Tel [] Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

//the second part of mUJSON.go is as follows 
//In this part of the utility, we define the myRecord variable, which holds the desired data.
//You can also see the use of the json.Marshal() function, which accepts a reference to the
//myRecord variable. Note that json.Marshal() requires a pointer variable that converts
//into the JSON format.
func main() {
	myrecord := Record{
		Name: "Mihalis",
		Surname: "Tsoukalos",
		Tel: []Telephone {
			{Mobile: true, Number:"1234-567"},
			{Mobile: true, Number:"abcc-567"},
			{Mobile: false, Number: "abcc-567"},
			
		}
		
	}
	rec, err := json.Marshal(&myrecord)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println((strig(rec)))
	//The last part of mUJSON.go contains the folowing code:
	//The json.Unmarshal() function gets JSON input and converts it into a Go structure. As it
	//happened with json.Marshal(), json.Unmarshal() also requires a pointer argument.

	var unRec Record
	err1 := json.Unmarshal(rec, &unRec)
	if err1 != nil (err1) {
		fmt.Println(err1)
		return
	}
	fmt.Println(unRec)
	
}