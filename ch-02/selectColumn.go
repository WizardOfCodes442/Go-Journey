//selectColumn.go
//The utility needs at least two command-line arguments to operate;
//The first one is the required column number and the second one 
// is the path of the text file to process. However, you can use as many text
//files as you want -selectColumn.go will process all of them one by one 

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Printf("Usage: selectColumn column <file1> [<file2> [...
		  <fileN]]\n")
		os.Exit(1)
	}
	temp, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("Column value is not an integer:", temp)
		return
	}
	column := temp
	if column < 0 {
		fmt.Println("Invalid Column number!")
		os.Exit(1)
	}
	
	for _, filename := range arguments[2:] {
		fmt.Println("\t\t", filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file %s\n", err)
			continue	
		}
	
		defer f.Close()
		r := bufio.NewReader(f)
		for {
			line, err:= r.ReadString('\n')
		
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file %s", err)
			}
			data := strings.Fields(line)
			if len(data) >= column {
				fmt.Println((data[column-1]))
			}
		}
	}
}
